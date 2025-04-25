package challenger

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChallengeVerifier[Input any, Output any] interface {
	VerifyChallenge(uint32, GenericInputTask[Input], TaskResponseData[Output]) error
}

type GenericInputTask[Input any] struct {
	InputValue                Input
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

type GenericOutputTaskResponse[Output any] struct {
	ReferenceTaskIndex uint32
	OutputValue        Output
}

type GenericTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

type TaskResponseData[Output any] struct {
	TaskResponse              GenericOutputTaskResponse[Output]
	TaskResponseMetadata      GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

type NewTaskCreatedEvent[Input any] struct {
	TaskIndex uint32
	Task      GenericInputTask[Input]
	Raw       types.Log
}

type TaskRespondedEvent[Output any] struct {
	TaskResponse              GenericOutputTaskResponse[Output]
	TaskResponseMetadata      GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

func (taskRespEvent TaskRespondedEvent[Output]) TaskIndex() uint32 {
	return taskRespEvent.TaskResponse.ReferenceTaskIndex
}

func (taskRespEvent TaskRespondedEvent[Output]) GetTaskResponse() GenericOutputTaskResponse[Output] {
	return taskRespEvent.TaskResponse
}

func (taskRespEvent TaskRespondedEvent[Output]) GetTaskResponseMetadata() GenericTaskResponseMetadata {
	return taskRespEvent.TaskResponseMetadata
}

type ChallengerConfig struct {
	EthWsUrl       string
	Logger         logging.Logger
	TaskManagerAbi *abi.ABI
	EthClient      *ethclient.Client
}

type Challenger[Input any, Output any] struct {
	logger             logging.Logger
	challengeVerifier  ChallengeVerifier[Input, Output]
	taskResponseChan   chan types.Log
	newTaskCreatedChan chan types.Log

	taskManagerAbi *abi.ABI
	tasks          map[uint32]GenericInputTask[Input]
	taskResponses  map[uint32]TaskResponseData[Output]

	ethClient *ethclient.Client
}

func NewChallenger[Input any, Output any](
	c ChallengerConfig,
	challengeVerifier ChallengeVerifier[Input, Output],
) (*Challenger[Input, Output], error) {
	client, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		c.Logger.Fatalf("error connecting to web socket: %v", err)
	}

	newTaskEventHash := c.TaskManagerAbi.Events["NewTaskCreated"].ID
	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{newTaskEventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		c.Logger.Fatalf("error subscribing to newTaskCreated events: %v", err)
	}

	taskRespondedEventHash := c.TaskManagerAbi.Events["TaskResponded"].ID
	query.Topics[0][0] = taskRespondedEventHash

	taskRespondedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, taskRespondedLogs)
	if err != nil {
		c.Logger.Fatalf("error subscribing to taskResponded events: %v", err)
	}

	return &Challenger[Input, Output]{
		logger:             c.Logger,
		challengeVerifier:  challengeVerifier,
		newTaskCreatedChan: newTaskCreatedLogs,
		taskResponseChan:   taskRespondedLogs,
		taskManagerAbi:     c.TaskManagerAbi,
		tasks:              make(map[uint32]GenericInputTask[Input]),
		taskResponses:      make(map[uint32]TaskResponseData[Output]),
		ethClient:          c.EthClient,
	}, nil
}

func (c *Challenger[Input, Output]) Start(ctx context.Context) error {
	c.logger.Info("Starting Challenger.")

	for {
		select {
		case newTaskCreatedLog := <-c.newTaskCreatedChan:
			c.logger.Info("New task created log received")
			err := c.processNewTaskCreatedLog(newTaskCreatedLog)
			if err != nil {
				c.logger.Fatalf("Error processing NewTaskCreated log: %v", err)
			}
		case taskResponseLog := <-c.taskResponseChan:
			c.logger.Info("Task response log received")
			err := c.processTaskResponseLog(taskResponseLog)
			if err != nil {
				c.logger.Fatalf("Error processing TaskResponded log: %v", err)
			}
		}
	}

}

func (c *Challenger[Input, Output]) processNewTaskCreatedLog(log types.Log) error {
	var newTaskCreatedLog NewTaskCreatedEvent[Input]

	err := c.taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())
	c.tasks[newTaskIndex] = newTaskCreatedLog.Task

	return nil
}

func (c *Challenger[Input, Output]) processTaskResponseLog(
	log types.Log,
) error {
	var taskRespondedLog TaskRespondedEvent[Output]

	err := c.taskManagerAbi.UnpackIntoInterface(&taskRespondedLog, "TaskResponded", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	taskIndex := taskRespondedLog.TaskIndex()

	// get the inputs necessary for raising a challenge
	nonSigningOperatorPubKeys := c.getNonSigningOperatorPubKeys(log.TxHash)
	taskResponseData := TaskResponseData[Output]{
		TaskResponse:              taskRespondedLog.GetTaskResponse(),
		TaskResponseMetadata:      taskRespondedLog.GetTaskResponseMetadata(),
		NonSigningOperatorPubKeys: nonSigningOperatorPubKeys,
	}

	c.taskResponses[taskIndex] = taskResponseData

	if task, found := c.tasks[taskIndex]; found {
		err = c.challengeVerifier.VerifyChallenge(taskIndex, task, taskResponseData)
		if err != nil {
			return fmt.Errorf("error verifying the challenge: %w", err)
		}
	}

	return nil
}

func (c *Challenger[Input, Output]) getNonSigningOperatorPubKeys(
	transactionHash common.Hash,
) []BN254G1Point {
	// get the nonSignerStakesAndSignature
	tx, _, err := c.ethClient.TransactionByHash(context.Background(), transactionHash)
	if err != nil {
		c.logger.Error("Error getting transaction by hash",
			"txHash", transactionHash,
			"err", err,
		)
	}

	calldata := tx.Data()
	methodSig := calldata[:4]
	method, err := c.taskManagerAbi.MethodById(methodSig)
	if err != nil {
		c.logger.Error("Error getting method", "err", err)
	}

	inputs, err := method.Inputs.Unpack(calldata[4:])
	if err != nil {
		c.logger.Error("Error unpacking calldata", "err", err)
	}

	// Note: this implies the abi of the Task Manager contract implemented by the AVS should respect this values, or it
	// wont work. Other solution is to receive this as parameter, but it looks more difficult than only replace this.
	nonSignerStakesAndSignatureInput := inputs[2].(struct {
		NonSignerQuorumBitmapIndices []uint32 "json:\"nonSignerQuorumBitmapIndices\""
		NonSignerPubkeys             []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"nonSignerPubkeys\""
		QuorumApks []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"quorumApks\""
		ApkG2 struct {
			X [2]*big.Int "json:\"X\""
			Y [2]*big.Int "json:\"Y\""
		} "json:\"apkG2\""
		Sigma struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"sigma\""
		QuorumApkIndices      []uint32   "json:\"quorumApkIndices\""
		TotalStakeIndices     []uint32   "json:\"totalStakeIndices\""
		NonSignerStakeIndices [][]uint32 "json:\"nonSignerStakeIndices\""
	})

	// get pubkeys of non-signing operators and submit them to the contract
	nonSigningOperatorPubKeys := make(
		[]BN254G1Point,
		len(nonSignerStakesAndSignatureInput.NonSignerPubkeys),
	)
	for i, pubkey := range nonSignerStakesAndSignatureInput.NonSignerPubkeys {
		nonSigningOperatorPubKeys[i] = BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	return nonSigningOperatorPubKeys
}
