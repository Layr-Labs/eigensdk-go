package aggregator_example

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/task-processor"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/bindings/taskManager"
	taskspammerexample "github.com/Layr-Labs/eigensdk-go/examples/task-spammer"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	ethHttpUrl := "http://localhost:8545"
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		return
	}

	txMgr, err := taskspammerexample.GetTxManager(logger, ethHttpClient, testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		return
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	cfg := aggregator.AggregatorConfig{
		RegistryCoordinatorAddress:    common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650"),
		OperatorStateRetrieverAddress: common.HexToAddress("0x4c5859f0f772848b2d91f1d83e2fe57935348029"),
		ServiceManagerAddress:         common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		EthHttpClient:                 ethHttpClient,
		Logger:                        logger,
		EthHttpUrl:                    ethHttpUrl,
		EthWsUrl:                      "ws://localhost:8545",
		EcdsaPrivateKey:               ecdsaPrivateKey,
		AggregatorServerIpPortAddr:    "localhost:8090",
		TaskManagerAbi:                taskManagerAbi,
	}

	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(cfg.ServiceManagerAddress, ethHttpClient)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})
	if err != nil {
		logger.Fatalf(err.Error())
	}

	contractTaskManager, err := cstaskmanager.NewContractIncredibleSquaringTaskManager(taskManagerAddr, ethHttpClient)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	taskResponder := NewTaskResponder(taskManagerAbi, contractTaskManager, txMgr)

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(taskManagerAbi, logger, taskResponder)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	agg, err := aggregator.NewAggregator(cfg, taskProcessor)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = agg.Start(context.Background())
	if err != nil {
		logger.Fatalf(err.Error())
	}
}

type TaskResponder struct {
	taskManagerAbi *abi.ABI

	taskManagerContract *cstaskmanager.ContractIncredibleSquaringTaskManager
	txMgr               txmgr.TxManager
}

func NewTaskResponder(taskManagerAbi *abi.ABI, taskManagerContract *cstaskmanager.ContractIncredibleSquaringTaskManager, txMgr txmgr.TxManager) TaskResponder {
	return TaskResponder{
		taskManagerAbi:      taskManagerAbi,
		taskManagerContract: taskManagerContract,
		txMgr:               txMgr,
	}
}

func (tr TaskResponder) RespondToTask(task types.GenericInputTask[*big.Int], taskResponse types.GenericOutputTaskResponse[*big.Int], nonSignersStakesAndSig types.NonSignerStakesAndSignature) error {

	incredibleTask := cstaskmanager.IIncredibleSquaringTaskManagerTask{
		NumberToBeSquared:         task.InputValue,
		TaskCreatedBlock:          task.TaskCreatedBlock,
		QuorumNumbers:             task.QuorumNumbers,
		QuorumThresholdPercentage: task.QuorumThresholdPercentage,
	}

	incredibleTaskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{}

	incredibleNonSignersStakesAndSig := cstaskmanager.IBLSSignatureCheckerTypesNonSignerStakesAndSignature{}

	txOpts, err := tr.txMgr.GetNoSendTxOpts()
	if err != nil {
		return err
	}

	tx, err := tr.taskManagerContract.RespondToTask(txOpts, incredibleTask, incredibleTaskResponse, incredibleNonSignersStakesAndSig)
	if err != nil {
		return err
	}

	_, err = tr.txMgr.Send(context.Background(), tx, true)
	if err != nil {
		return err
	}

	return nil
}

func (tr TaskResponder) PackTaskResponse(taskResponse types.GenericOutputTaskResponse[*big.Int]) (types.Bytes32, error) {
	abiType, err := extractTypeFromAbi(tr.taskManagerAbi)
	if err != nil {
		return [32]byte{}, err
	}
	hashFn := getDefaultHashFunction(abiType)

	return hashFn(taskResponse)
}

func extractTypeFromAbi(taskManagerAbi *abi.ABI) (abi.Type, error) {
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "OutputValue", // Left because abi does not support purely anonymous or underscored fields
			Type: taskManagerAbi.Events["TaskResponded"].Inputs[0].Type.TupleElems[1].String(),
		},
	})
	if err != nil {
		return abi.Type{}, fmt.Errorf("error creating abi task response type: %w", err)
	}

	return taskResponseType, nil
}

func getDefaultHashFunction(taskResponseType abi.Type) types.TaskResponseHashFunction {
	return func(taskResponse types.TaskResponse) (types.TaskResponseDigest, error) {
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		encodeTaskResponseByte, err := arguments.Pack(taskResponse)
		if err != nil {
			return types.Bytes32{}, fmt.Errorf("error encoding task response: %w", err)
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

		return taskResponseDigest, nil
	}
}
