package aggregator

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

// When starting the server, the aggregator start listening at the address specified by config the calls to
// the ProcessSignedTaskResponse method
func (agg *Aggregator[Input, Output]) startServer(ctx context.Context) {
	server := rpc.NewServer()
	err := server.RegisterName("Aggregator", agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}

	listener, err := net.Listen("tcp", agg.serverIpPortAddr)
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(listener, server)
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
