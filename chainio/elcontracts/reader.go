package elcontracts

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	eigenabi "github.com/Layr-Labs/eigensdk-go/types/abi"

	blspkcompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/Slasher"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

type ELReader interface {
	IsOperatorRegistered(ctx context.Context, operator types.Operator) (bool, error)

	GetOperatorDetails(ctx context.Context, operator types.Operator) (types.Operator, error)

	// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
	// use GetStrategyAndUnderlyingERC20Token if the contract address confirms with ERC20 standard
	GetStrategyAndUnderlyingToken(
		ctx context.Context, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, gethcommon.Address, error)

	// GetStrategyAndUnderlyingERC20Token returns the strategy contract and the underlying ERC20 token address
	GetStrategyAndUnderlyingERC20Token(
		ctx context.Context, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, clients.ERC20ContractClient, gethcommon.Address, error)

	QueryExistingRegisteredOperatorPubKeys(
		startBlock *big.Int,
		stopBlock *big.Int,
	) ([]types.OperatorAddr, []types.OperatorPubkeys, error)

	GetOperatorPubkeyHash(ctx context.Context, operator types.Operator) ([32]byte, error)

	GetOperatorAddressFromPubkeyHash(ctx context.Context, pubkeyHash [32]byte) (gethcommon.Address, error)

	ServiceManagerCanSlashOperatorUntilBlock(
		ctx context.Context,
		operatorAddr gethcommon.Address,
		serviceManagerAddr gethcommon.Address,
	) (uint32, error)

	OperatorIsFrozen(ctx context.Context, operatorAddr gethcommon.Address) (bool, error)

	GetOperatorSharesInStrategy(
		ctx context.Context,
		operatorAddr gethcommon.Address,
		strategyAddr gethcommon.Address,
	) (*big.Int, error)
}

type ELChainReader struct {
	logger                  logging.Logger
	slasher                 slasher.ContractSlasherCalls
	delegationManager       delegationmanager.ContractDelegationManagerCalls
	strategyManager         strategymanager.ContractStrategyManagerCalls
	blsPubkeyCompendium     blspkcompendium.ContractBLSPublicKeyCompendiumCalls
	blsPubKeyCompendiumAddr common.Address
	ethClient               eth.EthClient
}

// forces EthReader to implement the chainio.Reader interface
var _ ELReader = (*ELChainReader)(nil)

func NewELChainReader(
	slasher slasher.ContractSlasherCalls,
	delegationManager delegationmanager.ContractDelegationManagerCalls,
	strategyManager strategymanager.ContractStrategyManagerCalls,
	blsPubkeyCompendium blspkcompendium.ContractBLSPublicKeyCompendiumCalls,
	blsPubKeyCompendiumAddr common.Address,
	logger logging.Logger,
	ethClient eth.EthClient,
) (*ELChainReader, error) {
	return &ELChainReader{
		slasher:                 slasher,
		delegationManager:       delegationManager,
		strategyManager:         strategyManager,
		blsPubkeyCompendium:     blsPubkeyCompendium,
		blsPubKeyCompendiumAddr: blsPubKeyCompendiumAddr,
		logger:                  logger,
		ethClient:               ethClient,
	}, nil
}

// TODO(samlaf): should we just pass the CallOpts directly as argument instead of the context?
func (r *ELChainReader) IsOperatorRegistered(ctx context.Context, operator types.Operator) (bool, error) {
	isOperator, err := r.delegationManager.IsOperator(
		&bind.CallOpts{Context: ctx},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return false, err
	}

	return isOperator, nil
}

func (r *ELChainReader) GetOperatorPubkeyHash(ctx context.Context, operator types.Operator) ([32]byte, error) {
	operatorPubkeyHash, err := r.blsPubkeyCompendium.OperatorToPubkeyHash(
		&bind.CallOpts{},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return [32]byte{}, err
	}

	return operatorPubkeyHash, nil
}

func (r *ELChainReader) GetOperatorAddressFromPubkeyHash(
	ctx context.Context,
	pubkeyHash [32]byte,
) (gethcommon.Address, error) {
	return r.blsPubkeyCompendium.PubkeyHashToOperator(
		&bind.CallOpts{},
		pubkeyHash,
	)
}

func (r *ELChainReader) GetOperatorDetails(ctx context.Context, operator types.Operator) (types.Operator, error) {
	operatorDetails, err := r.delegationManager.OperatorDetails(
		&bind.CallOpts{},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return types.Operator{}, err
	}

	return types.Operator{
		Address:                   operator.Address,
		EarningsReceiverAddress:   operatorDetails.EarningsReceiver.Hex(),
		StakerOptOutWindowBlocks:  operatorDetails.StakerOptOutWindowBlocks,
		DelegationApproverAddress: operatorDetails.DelegationApprover.Hex(),
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ELChainReader) GetStrategyAndUnderlyingToken(
	ctx context.Context, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		r.logger.Error("Failed to fetch strategy contract", "err", err)
		return nil, common.Address{}, err
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{})
	if err != nil {
		r.logger.Error("Failed to fetch token contract", "err", err)
		return nil, common.Address{}, err
	}
	return contractStrategy, underlyingTokenAddr, nil
}

func (r *ELChainReader) GetStrategyAndUnderlyingERC20Token(
	ctx context.Context, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, clients.ERC20ContractClient, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		r.logger.Error("Failed to fetch strategy contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{})
	if err != nil {
		r.logger.Error("Failed to fetch token contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	contractUnderlyingToken, err := clients.NewERC20ContractChainClient(underlyingTokenAddr, r.ethClient)
	if err != nil {
		r.logger.Error("Failed to fetch token contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (r *ELChainReader) ServiceManagerCanSlashOperatorUntilBlock(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	serviceManagerAddr gethcommon.Address,
) (uint32, error) {
	serviceManagerCanSlashOperatorUntilBlock, err := r.slasher.ContractCanSlashOperatorUntilBlock(
		&bind.CallOpts{Context: ctx}, operatorAddr, serviceManagerAddr,
	)
	if err != nil {
		return 0, err
	}
	return serviceManagerCanSlashOperatorUntilBlock, nil
}

func (r *ELChainReader) OperatorIsFrozen(ctx context.Context, operatorAddr gethcommon.Address) (bool, error) {
	operatorIsFrozen, err := r.slasher.IsFrozen(&bind.CallOpts{Context: ctx}, operatorAddr)
	if err != nil {
		return false, err
	}
	return operatorIsFrozen, nil
}

func (r *ELChainReader) QueryExistingRegisteredOperatorPubKeys(
	startBlock *big.Int,
	stopBlock *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   stopBlock,
		Addresses: []gethcommon.Address{
			r.blsPubKeyCompendiumAddr,
		},
	}

	logs, err := r.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		r.logger.Error("Error filtering logs", "err", err)
		return nil, nil, err
	}
	r.logger.Info("logs:", "logs", logs)

	pubkeyCompendiumAbi, err := abi.JSON(bytes.NewReader(eigenabi.BLSPublicKeyCompendiumAbi))
	if err != nil {
		r.logger.Error("Error getting Abi", "err", err)
		return nil, nil, err
	}

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)

	for _, vLog := range logs {

		// get the operator address
		operatorAddr := gethcommon.HexToAddress(vLog.Topics[1].Hex())
		operatorAddresses = append(operatorAddresses, operatorAddr)

		event, err := pubkeyCompendiumAbi.Unpack("NewPubkeyRegistration", vLog.Data)
		if err != nil {
			r.logger.Error("Error unpacking event data", "err", err)
			return nil, nil, err
		}

		G1Pubkey := event[0].(struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		})

		G2Pubkey := event[1].(struct {
			X [2]*big.Int "json:\"X\""
			Y [2]*big.Int "json:\"Y\""
		})

		operatorPubkey := types.OperatorPubkeys{
			G1Pubkey: bls.NewG1Point(
				G1Pubkey.X,
				G1Pubkey.Y,
			),
			G2Pubkey: bls.NewG2Point(
				G2Pubkey.X,
				G2Pubkey.Y,
			),
		}

		operatorPubkeys = append(operatorPubkeys, operatorPubkey)

	}

	return operatorAddresses, operatorPubkeys, nil
}

func (r *ELChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	operatorSharesInStrategy, err := r.delegationManager.OperatorShares(
		&bind.CallOpts{Context: ctx},
		operatorAddr,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}
	return operatorSharesInStrategy, nil
}
