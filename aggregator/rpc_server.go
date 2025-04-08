package aggregator

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
)

type ProcessSignedTaskResponse interface {
	processSignedTaskResponse(ctx context.Context, event any) (bool, error)
}

type ProcessSignedTaskResponseServer struct {
	taskProcessor TaskProcessor
	serviceHandle blsagg.ServiceHandler
	logger        logging.Logger
}

func NewProcessSignedTaskResponseServer(taskProcessor TaskProcessor, serviceHandle blsagg.ServiceHandler, logger logging.Logger) ProcessSignedTaskResponseServer {
	return ProcessSignedTaskResponseServer{
		taskProcessor: taskProcessor,
		serviceHandle: serviceHandle,
		logger:        logger,
	}
}
