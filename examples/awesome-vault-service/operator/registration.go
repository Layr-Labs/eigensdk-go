package main

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterOperatorOnStartup(logger logging.Logger) error {
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)

	registrationConfig := sdkoperator.RegistrationConfig{
		Logger: logger,

		OperatorAddr:            common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		AllocationManagerAddr:   common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"),
		AvsAddress:              common.HexToAddress("0xcd8a1c3ba11cf5ecfa6267617243239504a98d90"),
		RegistryCoordinatorAddr: common.HexToAddress("0xfd471836031dc5108809d173a067e8486b9047a3"),
		StrategyAddrs:           []common.Address{common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5")},

		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),

		EthRpcUrl: "http://localhost:8545",

		EcdsaKeyStorePath: "keys/test.ecdsa.key.json",
		BlsKeyStorePath:   "keys/test.bls.key.json",

		AmountToMint:          amount,
		AllocatableMagnitudes: []uint64{1000000000000000},

		OperatorSetIds: []uint32{0},
	}

	err := sdkoperator.RegisterOperatorOnStartup(registrationConfig)
	if err != nil {
		logger.Errorf("Failed to register operator on startup: %w", err)
		return err
	}

	return nil
}
