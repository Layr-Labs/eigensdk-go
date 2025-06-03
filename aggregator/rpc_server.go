package aggregator

import (
	"context"
	"net"
	"net/http"
	"net/rpc"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// When starting the server, the aggregator start listening at the address specified by config the calls to
// the ProcessSignedTaskResponse method
func (agg *Aggregator[Input, Output]) startServer(ctx context.Context) error {
	server := rpc.NewServer()
	err := server.RegisterName("Aggregator", agg)
	if err != nil {
		return utils.WrapError("Error registering aggregator service (maybe the format of service task manager isn't correct)", err)
	}

	// TODO: Replace with http.ListenAndServe()
	err = agg.listenAndServe(server)
	if err != nil {
		return utils.WrapError("Failed to listen and serve", err)
	}

	return nil
}

func (agg *Aggregator[Input, Output]) listenAndServe(server *rpc.Server) error {
	listener, err := net.Listen("tcp", agg.serverIpPortAddr)
	if err != nil {
		return utils.WrapError("Err wile listening", err)
	}

	err = http.Serve(listener, server)
	if err != nil {
		return utils.WrapError("Err while serving", err)
	}

	return nil
}

type SignedTaskResponse[Output any] struct {
	TaskResponse taskmanager.TaskResponse[Output]
	BlsSignature bls.Signature
	OperatorId   sdktypes.OperatorId
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator[Input, Output]) ProcessSignedTaskResponse(signedTaskResponse *SignedTaskResponse[Output], reply *bool) error {
	agg.logger.Infof("Received signed task response: %#v", signedTaskResponse)
	taskIndex := signedTaskResponse.TaskResponse.ReferenceTaskIndex

	taskSignature := blsagg.NewTaskSignature(
		taskIndex,
		signedTaskResponse.TaskResponse,
		&signedTaskResponse.BlsSignature,
		signedTaskResponse.OperatorId,
	)

	err := agg.blsAggregationService.ProcessNewSignature(context.Background(), taskSignature)

	return err
}
