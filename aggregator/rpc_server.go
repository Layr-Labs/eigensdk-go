package aggregator

import (
	"context"
	"net"
	"net/http"
	"net/rpc"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// The aggregator RPC server is responsible for receiving signed responses from the operators, and sending
// those responses to the BLS aggregation service.
// Warning: Public methods added to this struct may be exposed via RPC and accessible over the network.
// Ensure that any exported method has proper validation as needed.
type AggregatorRpcServer[Input any, Output any] struct {
	logger logging.Logger

	// IP address and port where the aggregator will listen to operator task responses
	serverIpPortAddr string

	// BLS aggregation service where the operator responses will be sent
	blsAggregationService blsagg.BlsAggregationService
}

// NewAggregatorRpcServer creates a new AggregatorRpcServer with a logger, an IP address and port,
// and the BLS aggregation service.
func NewAggregatorRpcServer[Input any, Output any](
	logger logging.Logger,
	serverIpPortAddr string,
	blsAggregationService blsagg.BlsAggregationService,
) *AggregatorRpcServer[Input, Output] {
	return &AggregatorRpcServer[Input, Output]{
		logger:                logger,
		serverIpPortAddr:      serverIpPortAddr,
		blsAggregationService: blsAggregationService,
	}
}

// When starting the server, the aggregator RPC server starts listening at the address specified by the
// config for calls to the `ProcessSignedTaskResponse` method
func (aggServ *AggregatorRpcServer[Input, Output]) StartServer() error {
	server := rpc.NewServer()
	err := server.RegisterName("Aggregator", aggServ)
	if err != nil {
		return utils.WrapError("Error registering aggregator service (maybe the format of service task manager isn't correct)", err)
	}

	// TODO: Replace with http.ListenAndServe()
	err = aggServ.listenAndServe(server)
	if err != nil {
		return utils.WrapError("Failed to listen and serve", err)
	}

	return nil
}

// This function should be replaced by http.ListenAndServe() in a future. It listens for new connections
// from operators, and will respond to calls to the ProcessSignedTaskResponse method
func (aggServ *AggregatorRpcServer[Input, Output]) listenAndServe(server *rpc.Server) error {
	listener, err := net.Listen("tcp", aggServ.serverIpPortAddr)
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

// RPC endpoint which is called by operators.
// Reply doesn't need to be checked. If there are no errors, the task response is accepted
// RPC framework forces a reply type to exist, so we put bool as a placeholder.
func (aggServ *AggregatorRpcServer[Input, Output]) ProcessSignedTaskResponse(signedTaskResponse *SignedTaskResponse[Output], reply *bool) error {
	aggServ.logger.Infof("Received signed task response: %#v", signedTaskResponse)
	taskIndex := signedTaskResponse.TaskResponse.ReferenceTaskIndex

	taskSignature := blsagg.NewTaskSignature(
		taskIndex,
		signedTaskResponse.TaskResponse,
		&signedTaskResponse.BlsSignature,
		signedTaskResponse.OperatorId,
	)

	err := aggServ.blsAggregationService.ProcessNewSignature(context.Background(), taskSignature)

	return err
}
