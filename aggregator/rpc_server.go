package aggregator

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
)

type ProcessSignedTaskResponse interface {
	processSignedTaskResponse(ctx context.Context, event any) (bool, error)
}

type ProcessSignedTaskResponseServer struct{
    taskProcessor TaskProcessor
    serviceHandle blsagg.ServiceHandler
	logger		logging.Logger
}

func NewProcessSignedTaskResponseServer(taskProcessor TaskProcessor, serviceHandle blsagg.ServiceHandler, logger logging.Logger) ProcessSignedTaskResponseServer {
    return ProcessSignedTaskResponseServer{
        taskProcessor: taskProcessor,
        serviceHandle: serviceHandle,
		logger: logger,
    }
}

func (serv *ProcessSignedTaskResponseServer) processSignedTaskResponse(
        signedTaskResponse SignedTaskResponse,
    ) error {
    taskIndex := signedTaskResponse.TaskResponse.ReferenceTaskIndex;

    taskResponseDigest, err := serv.taskProcessor.ProcessTaskResponse(context.Background(), signedTaskResponse.TaskResponse)
	if err != nil {
		serv.logger.Errorf("Failure processing task response: %v", err)
		return err
	}
    taskSignature := blsagg.NewTaskSignature(taskIndex, taskResponseDigest, &signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId);

    serv.serviceHandle.ProcessNewSignature(context.Background(), taskSignature);
        serv.logger.Infof("processed signature for index %v", taskIndex);
	return nil
}
