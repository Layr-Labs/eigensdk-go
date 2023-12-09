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

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	eigenabi "github.com/Layr-Labs/eigensdk-go/types/abi"

	blspkcompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/Slasher"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

type ELReader interface {
	IsOperatorRegistered(opts *bind.CallOpts, operator types.Operator) (bool, error)

	GetOperatorDetails(opts *bind.CallOpts, operator types.Operator) (types.Operator, error)

	// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
	// use GetStrategyAndUnderlyingERC20Token if the contract address confirms with ERC20 standard
	GetStrategyAndUnderlyingToken(
		opts *bind.CallOpts, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, gethcommon.Address, error)

	// GetStrategyAndUnderlyingERC20Token returns the strategy contract and the underlying ERC20 token address
	GetStrategyAndUnderlyingERC20Token(
		opts *bind.CallOpts, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error)

	QueryExistingRegisteredOperatorPubKeys(
		ctx context.Context, startBlock *big.Int, stopBlock *big.Int,
	) ([]types.OperatorAddr, []types.OperatorPubkeys, error)

	GetOperatorPubkeyHash(opts *bind.CallOpts, operator types.Operator) ([32]byte, error)

	GetOperatorAddressFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (gethcommon.Address, error)

	ServiceManagerCanSlashOperatorUntilBlock(
		opts *bind.CallOpts,
		operatorAddr gethcommon.Address,
		serviceManagerAddr gethcommon.Address,
	) (uint32, error)

	OperatorIsFrozen(opts *bind.CallOpts, operatorAddr gethcommon.Address) (bool, error)

	GetOperatorSharesInStrategy(
		opts *bind.CallOpts,
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
) *ELChainReader {
	return &ELChainReader{
		slasher:                 slasher,
		delegationManager:       delegationManager,
		strategyManager:         strategyManager,
		blsPubkeyCompendium:     blsPubkeyCompendium,
		blsPubKeyCompendiumAddr: blsPubKeyCompendiumAddr,
		logger:                  logger,
		ethClient:               ethClient,
	}
}

func BuildELChainReader(
	slasherAddr gethcommon.Address,
	blsPubKeyCompendiumAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*ELChainReader, error) {
	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		slasherAddr,
		blsPubKeyCompendiumAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.BlsPubkeyCompendium,
		blsPubKeyCompendiumAddr,
		logger,
		ethClient,
	), nil
}

// TODO(samlaf): should we just pass the CallOpts directly as argument instead of the context?
func (r *ELChainReader) IsOperatorRegistered(opts *bind.CallOpts, operator types.Operator) (bool, error) {
	isOperator, err := r.delegationManager.IsOperator(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return false, err
	}

	return isOperator, nil
}

func (r *ELChainReader) GetOperatorPubkeyHash(opts *bind.CallOpts, operator types.Operator) ([32]byte, error) {
	operatorPubkeyHash, err := r.blsPubkeyCompendium.OperatorToPubkeyHash(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return [32]byte{}, err
	}

	return operatorPubkeyHash, nil
}

func (r *ELChainReader) GetOperatorAddressFromPubkeyHash(
	opts *bind.CallOpts,
	pubkeyHash [32]byte,
) (gethcommon.Address, error) {
	return r.blsPubkeyCompendium.PubkeyHashToOperator(
		opts,
		pubkeyHash,
	)
}

func (r *ELChainReader) GetOperatorDetails(opts *bind.CallOpts, operator types.Operator) (types.Operator, error) {
	operatorDetails, err := r.delegationManager.OperatorDetails(
		opts,
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
	opts *bind.CallOpts, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		r.logger.Error("Failed to fetch strategy contract", "err", err)
		return nil, common.Address{}, err
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		r.logger.Error("Failed to fetch token contract", "err", err)
		return nil, common.Address{}, err
	}
	return contractStrategy, underlyingTokenAddr, nil
}

// GetStrategyAndUnderlyingERC20Token returns the strategy contract, the erc20 bindings for the underlying token
// and the underlying token address
func (r *ELChainReader) GetStrategyAndUnderlyingERC20Token(
	opts *bind.CallOpts, strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		r.logger.Error("Failed to fetch strategy contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(opts)
	if err != nil {
		r.logger.Error("Failed to fetch token contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	if err != nil {
		r.logger.Error("Failed to fetch token contract", "err", err)
		return nil, nil, common.Address{}, err
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (r *ELChainReader) ServiceManagerCanSlashOperatorUntilBlock(
	opts *bind.CallOpts,
	operatorAddr gethcommon.Address,
	serviceManagerAddr gethcommon.Address,
) (uint32, error) {
	serviceManagerCanSlashOperatorUntilBlock, err := r.slasher.ContractCanSlashOperatorUntilBlock(
		opts, operatorAddr, serviceManagerAddr,
	)
	if err != nil {
		return 0, err
	}
	return serviceManagerCanSlashOperatorUntilBlock, nil
}

func (r *ELChainReader) OperatorIsFrozen(opts *bind.CallOpts, operatorAddr gethcommon.Address) (bool, error) {
	operatorIsFrozen, err := r.slasher.IsFrozen(opts, operatorAddr)
	if err != nil {
		return false, err
	}
	return operatorIsFrozen, nil
}

func (r *ELChainReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
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

	logs, err := r.ethClient.FilterLogs(ctx, query)
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
	opts *bind.CallOpts,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	operatorSharesInStrategy, err := r.delegationManager.OperatorShares(
		opts,
		operatorAddr,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}
	return operatorSharesInStrategy, nil
}
