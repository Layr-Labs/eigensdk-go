// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIEigenPod

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// BeaconChainProofsWithdrawalProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsWithdrawalProof struct {
	WithdrawalProof                 []byte
	SlotProof                       []byte
	ExecutionPayloadProof           []byte
	TimestampProof                  []byte
	HistoricalSummaryBlockRootProof []byte
	BlockRootIndex                  uint64
	HistoricalSummaryIndex          uint64
	WithdrawalIndex                 uint64
	BlockRoot                       [32]byte
	SlotRoot                        [32]byte
	TimestampRoot                   [32]byte
	ExecutionPayloadRoot            [32]byte
}

// IEigenPodValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodValidatorInfo struct {
	ValidatorIndex                   uint64
	RestakedBalanceGwei              uint64
	MostRecentBalanceUpdateTimestamp uint64
	Status                           uint8
}

// ContractIEigenPodMetaData contains all meta data concerning the ContractIEigenPod contract.
var ContractIEigenPodMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activateRestaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRestaked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mostRecentWithdrawalTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonBeaconChainETHBalanceWei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provenWithdrawal\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validatorPubkeyHashToInfo\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mostRecentBalanceUpdateTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyToInfo\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mostRecentBalanceUpdateTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyAndProcessWithdrawals\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"withdrawalProofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.WithdrawalProof[]\",\"components\":[{\"name\":\"withdrawalProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"slotProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executionPayloadProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestampProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"historicalSummaryBlockRootProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockRootIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"historicalSummaryIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"withdrawalIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"slotRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestampRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executionPayloadRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"},{\"name\":\"withdrawalFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyBalanceUpdates\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"withdrawalCredentialProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBeforeRestaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNonBeaconChainETHBalanceWei\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountToWithdraw\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRestakedBeaconChainETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EigenPodStaked\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FullWithdrawalRedeemed\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"withdrawalTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawalAmountGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHReceived\",\"inputs\":[{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountWithdrawn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialWithdrawalRedeemed\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"withdrawalTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"partialWithdrawalAmountGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakedBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakingActivated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorBalanceUpdated\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"balanceTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRestaked\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false}]",
}

// ContractIEigenPodABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIEigenPodMetaData.ABI instead.
var ContractIEigenPodABI = ContractIEigenPodMetaData.ABI

// ContractIEigenPodMethods is an auto generated interface around an Ethereum contract.
type ContractIEigenPodMethods interface {
	ContractIEigenPodCalls
	ContractIEigenPodTransacts
	ContractIEigenPodFilters
}

// ContractIEigenPodCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIEigenPodCalls interface {
	MAXRESTAKEDBALANCEGWEIPERVALIDATOR(opts *bind.CallOpts) (uint64, error)

	EigenPodManager(opts *bind.CallOpts) (common.Address, error)

	HasRestaked(opts *bind.CallOpts) (bool, error)

	MostRecentWithdrawalTimestamp(opts *bind.CallOpts) (uint64, error)

	NonBeaconChainETHBalanceWei(opts *bind.CallOpts) (*big.Int, error)

	PodOwner(opts *bind.CallOpts) (common.Address, error)

	ProvenWithdrawal(opts *bind.CallOpts, validatorPubkeyHash [32]byte, slot uint64) (bool, error)

	ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error)

	ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodValidatorInfo, error)

	ValidatorStatus(opts *bind.CallOpts, validatorPubkey []byte) (uint8, error)

	ValidatorStatus0(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error)

	WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error)
}

// ContractIEigenPodTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIEigenPodTransacts interface {
	ActivateRestaking(opts *bind.TransactOpts) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error)

	RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error)

	Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error)

	VerifyAndProcessWithdrawals(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error)

	VerifyBalanceUpdates(opts *bind.TransactOpts, oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error)

	VerifyWithdrawalCredentials(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error)

	WithdrawBeforeRestaking(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawNonBeaconChainETHBalanceWei(opts *bind.TransactOpts, recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error)

	WithdrawRestakedBeaconChainETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)
}

// ContractIEigenPodFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIEigenPodFilters interface {
	FilterEigenPodStaked(opts *bind.FilterOpts) (*ContractIEigenPodEigenPodStakedIterator, error)
	WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodEigenPodStaked) (event.Subscription, error)
	ParseEigenPodStaked(log types.Log) (*ContractIEigenPodEigenPodStaked, error)

	FilterFullWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodFullWithdrawalRedeemedIterator, error)
	WatchFullWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodFullWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error)
	ParseFullWithdrawalRedeemed(log types.Log) (*ContractIEigenPodFullWithdrawalRedeemed, error)

	FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*ContractIEigenPodNonBeaconChainETHReceivedIterator, error)
	WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodNonBeaconChainETHReceived) (event.Subscription, error)
	ParseNonBeaconChainETHReceived(log types.Log) (*ContractIEigenPodNonBeaconChainETHReceived, error)

	FilterNonBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodNonBeaconChainETHWithdrawnIterator, error)
	WatchNonBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodNonBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error)
	ParseNonBeaconChainETHWithdrawn(log types.Log) (*ContractIEigenPodNonBeaconChainETHWithdrawn, error)

	FilterPartialWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodPartialWithdrawalRedeemedIterator, error)
	WatchPartialWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodPartialWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error)
	ParsePartialWithdrawalRedeemed(log types.Log) (*ContractIEigenPodPartialWithdrawalRedeemed, error)

	FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator, error)
	WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error)
	ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*ContractIEigenPodRestakedBeaconChainETHWithdrawn, error)

	FilterRestakingActivated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodRestakingActivatedIterator, error)
	WatchRestakingActivated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodRestakingActivated, podOwner []common.Address) (event.Subscription, error)
	ParseRestakingActivated(log types.Log) (*ContractIEigenPodRestakingActivated, error)

	FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*ContractIEigenPodValidatorBalanceUpdatedIterator, error)
	WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorBalanceUpdated) (event.Subscription, error)
	ParseValidatorBalanceUpdated(log types.Log) (*ContractIEigenPodValidatorBalanceUpdated, error)

	FilterValidatorRestaked(opts *bind.FilterOpts) (*ContractIEigenPodValidatorRestakedIterator, error)
	WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorRestaked) (event.Subscription, error)
	ParseValidatorRestaked(log types.Log) (*ContractIEigenPodValidatorRestaked, error)
}

// ContractIEigenPod is an auto generated Go binding around an Ethereum contract.
type ContractIEigenPod struct {
	ContractIEigenPodCaller     // Read-only binding to the contract
	ContractIEigenPodTransactor // Write-only binding to the contract
	ContractIEigenPodFilterer   // Log filterer for contract events
}

// ContractIEigenPod implements the ContractIEigenPodMethods interface.
var _ ContractIEigenPodMethods = (*ContractIEigenPod)(nil)

// ContractIEigenPodCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIEigenPodCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodCaller implements the ContractIEigenPodCalls interface.
var _ ContractIEigenPodCalls = (*ContractIEigenPodCaller)(nil)

// ContractIEigenPodTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIEigenPodTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodTransactor implements the ContractIEigenPodTransacts interface.
var _ ContractIEigenPodTransacts = (*ContractIEigenPodTransactor)(nil)

// ContractIEigenPodFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIEigenPodFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodFilterer implements the ContractIEigenPodFilters interface.
var _ ContractIEigenPodFilters = (*ContractIEigenPodFilterer)(nil)

// ContractIEigenPodSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIEigenPodSession struct {
	Contract     *ContractIEigenPod // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractIEigenPodCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIEigenPodCallerSession struct {
	Contract *ContractIEigenPodCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractIEigenPodTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIEigenPodTransactorSession struct {
	Contract     *ContractIEigenPodTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractIEigenPodRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIEigenPodRaw struct {
	Contract *ContractIEigenPod // Generic contract binding to access the raw methods on
}

// ContractIEigenPodCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIEigenPodCallerRaw struct {
	Contract *ContractIEigenPodCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIEigenPodTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIEigenPodTransactorRaw struct {
	Contract *ContractIEigenPodTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIEigenPod creates a new instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPod(address common.Address, backend bind.ContractBackend) (*ContractIEigenPod, error) {
	contract, err := bindContractIEigenPod(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPod{ContractIEigenPodCaller: ContractIEigenPodCaller{contract: contract}, ContractIEigenPodTransactor: ContractIEigenPodTransactor{contract: contract}, ContractIEigenPodFilterer: ContractIEigenPodFilterer{contract: contract}}, nil
}

// NewContractIEigenPodCaller creates a new read-only instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPodCaller(address common.Address, caller bind.ContractCaller) (*ContractIEigenPodCaller, error) {
	contract, err := bindContractIEigenPod(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodCaller{contract: contract}, nil
}

// NewContractIEigenPodTransactor creates a new write-only instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPodTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIEigenPodTransactor, error) {
	contract, err := bindContractIEigenPod(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodTransactor{contract: contract}, nil
}

// NewContractIEigenPodFilterer creates a new log filterer instance of ContractIEigenPod, bound to a specific deployed contract.
func NewContractIEigenPodFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIEigenPodFilterer, error) {
	contract, err := bindContractIEigenPod(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodFilterer{contract: contract}, nil
}

// bindContractIEigenPod binds a generic wrapper to an already deployed contract.
func bindContractIEigenPod(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIEigenPodMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIEigenPod *ContractIEigenPodRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIEigenPod.Contract.ContractIEigenPodCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIEigenPod *ContractIEigenPodRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.ContractIEigenPodTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIEigenPod *ContractIEigenPodRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.ContractIEigenPodTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIEigenPod *ContractIEigenPodCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIEigenPod.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIEigenPod *ContractIEigenPodTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIEigenPod *ContractIEigenPodTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.contract.Transact(opts, method, params...)
}

// MAXRESTAKEDBALANCEGWEIPERVALIDATOR is a free data retrieval call binding the contract method 0x1d905d5c.
//
// Solidity: function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCaller) MAXRESTAKEDBALANCEGWEIPERVALIDATOR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXRESTAKEDBALANCEGWEIPERVALIDATOR is a free data retrieval call binding the contract method 0x1d905d5c.
//
// Solidity: function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodSession) MAXRESTAKEDBALANCEGWEIPERVALIDATOR() (uint64, error) {
	return _ContractIEigenPod.Contract.MAXRESTAKEDBALANCEGWEIPERVALIDATOR(&_ContractIEigenPod.CallOpts)
}

// MAXRESTAKEDBALANCEGWEIPERVALIDATOR is a free data retrieval call binding the contract method 0x1d905d5c.
//
// Solidity: function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) MAXRESTAKEDBALANCEGWEIPERVALIDATOR() (uint64, error) {
	return _ContractIEigenPod.Contract.MAXRESTAKEDBALANCEGWEIPERVALIDATOR(&_ContractIEigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodSession) EigenPodManager() (common.Address, error) {
	return _ContractIEigenPod.Contract.EigenPodManager(&_ContractIEigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) EigenPodManager() (common.Address, error) {
	return _ContractIEigenPod.Contract.EigenPodManager(&_ContractIEigenPod.CallOpts)
}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_ContractIEigenPod *ContractIEigenPodCaller) HasRestaked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "hasRestaked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_ContractIEigenPod *ContractIEigenPodSession) HasRestaked() (bool, error) {
	return _ContractIEigenPod.Contract.HasRestaked(&_ContractIEigenPod.CallOpts)
}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) HasRestaked() (bool, error) {
	return _ContractIEigenPod.Contract.HasRestaked(&_ContractIEigenPod.CallOpts)
}

// MostRecentWithdrawalTimestamp is a free data retrieval call binding the contract method 0x87e0d289.
//
// Solidity: function mostRecentWithdrawalTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCaller) MostRecentWithdrawalTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "mostRecentWithdrawalTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MostRecentWithdrawalTimestamp is a free data retrieval call binding the contract method 0x87e0d289.
//
// Solidity: function mostRecentWithdrawalTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodSession) MostRecentWithdrawalTimestamp() (uint64, error) {
	return _ContractIEigenPod.Contract.MostRecentWithdrawalTimestamp(&_ContractIEigenPod.CallOpts)
}

// MostRecentWithdrawalTimestamp is a free data retrieval call binding the contract method 0x87e0d289.
//
// Solidity: function mostRecentWithdrawalTimestamp() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) MostRecentWithdrawalTimestamp() (uint64, error) {
	return _ContractIEigenPod.Contract.MostRecentWithdrawalTimestamp(&_ContractIEigenPod.CallOpts)
}

// NonBeaconChainETHBalanceWei is a free data retrieval call binding the contract method 0xfe80b087.
//
// Solidity: function nonBeaconChainETHBalanceWei() view returns(uint256)
func (_ContractIEigenPod *ContractIEigenPodCaller) NonBeaconChainETHBalanceWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "nonBeaconChainETHBalanceWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonBeaconChainETHBalanceWei is a free data retrieval call binding the contract method 0xfe80b087.
//
// Solidity: function nonBeaconChainETHBalanceWei() view returns(uint256)
func (_ContractIEigenPod *ContractIEigenPodSession) NonBeaconChainETHBalanceWei() (*big.Int, error) {
	return _ContractIEigenPod.Contract.NonBeaconChainETHBalanceWei(&_ContractIEigenPod.CallOpts)
}

// NonBeaconChainETHBalanceWei is a free data retrieval call binding the contract method 0xfe80b087.
//
// Solidity: function nonBeaconChainETHBalanceWei() view returns(uint256)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) NonBeaconChainETHBalanceWei() (*big.Int, error) {
	return _ContractIEigenPod.Contract.NonBeaconChainETHBalanceWei(&_ContractIEigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCaller) PodOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "podOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodSession) PodOwner() (common.Address, error) {
	return _ContractIEigenPod.Contract.PodOwner(&_ContractIEigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) PodOwner() (common.Address, error) {
	return _ContractIEigenPod.Contract.PodOwner(&_ContractIEigenPod.CallOpts)
}

// ProvenWithdrawal is a free data retrieval call binding the contract method 0x34bea20a.
//
// Solidity: function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) view returns(bool)
func (_ContractIEigenPod *ContractIEigenPodCaller) ProvenWithdrawal(opts *bind.CallOpts, validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "provenWithdrawal", validatorPubkeyHash, slot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProvenWithdrawal is a free data retrieval call binding the contract method 0x34bea20a.
//
// Solidity: function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) view returns(bool)
func (_ContractIEigenPod *ContractIEigenPodSession) ProvenWithdrawal(validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	return _ContractIEigenPod.Contract.ProvenWithdrawal(&_ContractIEigenPod.CallOpts, validatorPubkeyHash, slot)
}

// ProvenWithdrawal is a free data retrieval call binding the contract method 0x34bea20a.
//
// Solidity: function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) view returns(bool)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ProvenWithdrawal(validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	return _ContractIEigenPod.Contract.ProvenWithdrawal(&_ContractIEigenPod.CallOpts, validatorPubkeyHash, slot)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyHashToInfo(&_ContractIEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyHashToInfo(&_ContractIEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorPubkeyToInfo", validatorPubkey)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyToInfo(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _ContractIEigenPod.Contract.ValidatorPubkeyToInfo(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorStatus(opts *bind.CallOpts, validatorPubkey []byte) (uint8, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorStatus", validatorPubkey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus(&_ContractIEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCaller) ValidatorStatus0(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "validatorStatus0", pubkeyHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus0(&_ContractIEigenPod.CallOpts, pubkeyHash)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _ContractIEigenPod.Contract.ValidatorStatus0(&_ContractIEigenPod.CallOpts, pubkeyHash)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCaller) WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPod.contract.Call(opts, &out, "withdrawableRestakedExecutionLayerGwei")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _ContractIEigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_ContractIEigenPod.CallOpts)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_ContractIEigenPod *ContractIEigenPodCallerSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _ContractIEigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_ContractIEigenPod.CallOpts)
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) ActivateRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "activateRestaking")
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_ContractIEigenPod *ContractIEigenPodSession) ActivateRestaking() (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.ActivateRestaking(&_ContractIEigenPod.TransactOpts)
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) ActivateRestaking() (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.ActivateRestaking(&_ContractIEigenPod.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Initialize(&_ContractIEigenPod.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Initialize(&_ContractIEigenPod.TransactOpts, owner)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "recoverTokens", tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.RecoverTokens(&_ContractIEigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.RecoverTokens(&_ContractIEigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPod *ContractIEigenPodSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Stake(&_ContractIEigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.Stake(&_ContractIEigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) VerifyAndProcessWithdrawals(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "verifyAndProcessWithdrawals", oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyAndProcessWithdrawals(&_ContractIEigenPod.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyAndProcessWithdrawals(&_ContractIEigenPod.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) VerifyBalanceUpdates(opts *bind.TransactOpts, oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "verifyBalanceUpdates", oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyBalanceUpdates(&_ContractIEigenPod.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyBalanceUpdates(&_ContractIEigenPod.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] withdrawalCredentialProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "verifyWithdrawalCredentials", oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] withdrawalCredentialProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyWithdrawalCredentials(&_ContractIEigenPod.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] withdrawalCredentialProofs, bytes32[][] validatorFields) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.VerifyWithdrawalCredentials(&_ContractIEigenPod.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) WithdrawBeforeRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "withdrawBeforeRestaking")
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_ContractIEigenPod *ContractIEigenPodSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawBeforeRestaking(&_ContractIEigenPod.TransactOpts)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawBeforeRestaking(&_ContractIEigenPod.TransactOpts)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) WithdrawNonBeaconChainETHBalanceWei(opts *bind.TransactOpts, recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "withdrawNonBeaconChainETHBalanceWei", recipient, amountToWithdraw)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawNonBeaconChainETHBalanceWei(&_ContractIEigenPod.TransactOpts, recipient, amountToWithdraw)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawNonBeaconChainETHBalanceWei(&_ContractIEigenPod.TransactOpts, recipient, amountToWithdraw)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactor) WithdrawRestakedBeaconChainETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.contract.Transact(opts, "withdrawRestakedBeaconChainETH", recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_ContractIEigenPod *ContractIEigenPodSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawRestakedBeaconChainETH(&_ContractIEigenPod.TransactOpts, recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_ContractIEigenPod *ContractIEigenPodTransactorSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPod.Contract.WithdrawRestakedBeaconChainETH(&_ContractIEigenPod.TransactOpts, recipient, amount)
}

// ContractIEigenPodEigenPodStakedIterator is returned from FilterEigenPodStaked and is used to iterate over the raw logs and unpacked data for EigenPodStaked events raised by the ContractIEigenPod contract.
type ContractIEigenPodEigenPodStakedIterator struct {
	Event *ContractIEigenPodEigenPodStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodEigenPodStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodEigenPodStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodEigenPodStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodEigenPodStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodEigenPodStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodEigenPodStaked represents a EigenPodStaked event raised by the ContractIEigenPod contract.
type ContractIEigenPodEigenPodStaked struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEigenPodStaked is a free log retrieval operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterEigenPodStaked(opts *bind.FilterOpts) (*ContractIEigenPodEigenPodStakedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodEigenPodStakedIterator{contract: _ContractIEigenPod.contract, event: "EigenPodStaked", logs: logs, sub: sub}, nil
}

// WatchEigenPodStaked is a free log subscription operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodEigenPodStaked) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodEigenPodStaked)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEigenPodStaked is a log parse operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseEigenPodStaked(log types.Log) (*ContractIEigenPodEigenPodStaked, error) {
	event := new(ContractIEigenPodEigenPodStaked)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodFullWithdrawalRedeemedIterator is returned from FilterFullWithdrawalRedeemed and is used to iterate over the raw logs and unpacked data for FullWithdrawalRedeemed events raised by the ContractIEigenPod contract.
type ContractIEigenPodFullWithdrawalRedeemedIterator struct {
	Event *ContractIEigenPodFullWithdrawalRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodFullWithdrawalRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodFullWithdrawalRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodFullWithdrawalRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodFullWithdrawalRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodFullWithdrawalRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodFullWithdrawalRedeemed represents a FullWithdrawalRedeemed event raised by the ContractIEigenPod contract.
type ContractIEigenPodFullWithdrawalRedeemed struct {
	ValidatorIndex       *big.Int
	WithdrawalTimestamp  uint64
	Recipient            common.Address
	WithdrawalAmountGwei uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFullWithdrawalRedeemed is a free log retrieval operation binding the contract event 0xb76a93bb649ece524688f1a01d184e0bbebcda58eae80c28a898bec3fb5a0963.
//
// Solidity: event FullWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 withdrawalAmountGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterFullWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodFullWithdrawalRedeemedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "FullWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodFullWithdrawalRedeemedIterator{contract: _ContractIEigenPod.contract, event: "FullWithdrawalRedeemed", logs: logs, sub: sub}, nil
}

// WatchFullWithdrawalRedeemed is a free log subscription operation binding the contract event 0xb76a93bb649ece524688f1a01d184e0bbebcda58eae80c28a898bec3fb5a0963.
//
// Solidity: event FullWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 withdrawalAmountGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchFullWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodFullWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "FullWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodFullWithdrawalRedeemed)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "FullWithdrawalRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFullWithdrawalRedeemed is a log parse operation binding the contract event 0xb76a93bb649ece524688f1a01d184e0bbebcda58eae80c28a898bec3fb5a0963.
//
// Solidity: event FullWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 withdrawalAmountGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseFullWithdrawalRedeemed(log types.Log) (*ContractIEigenPodFullWithdrawalRedeemed, error) {
	event := new(ContractIEigenPodFullWithdrawalRedeemed)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "FullWithdrawalRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodNonBeaconChainETHReceivedIterator is returned from FilterNonBeaconChainETHReceived and is used to iterate over the raw logs and unpacked data for NonBeaconChainETHReceived events raised by the ContractIEigenPod contract.
type ContractIEigenPodNonBeaconChainETHReceivedIterator struct {
	Event *ContractIEigenPodNonBeaconChainETHReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodNonBeaconChainETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodNonBeaconChainETHReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodNonBeaconChainETHReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodNonBeaconChainETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodNonBeaconChainETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodNonBeaconChainETHReceived represents a NonBeaconChainETHReceived event raised by the ContractIEigenPod contract.
type ContractIEigenPodNonBeaconChainETHReceived struct {
	AmountReceived *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNonBeaconChainETHReceived is a free log retrieval operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*ContractIEigenPodNonBeaconChainETHReceivedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodNonBeaconChainETHReceivedIterator{contract: _ContractIEigenPod.contract, event: "NonBeaconChainETHReceived", logs: logs, sub: sub}, nil
}

// WatchNonBeaconChainETHReceived is a free log subscription operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodNonBeaconChainETHReceived) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodNonBeaconChainETHReceived)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNonBeaconChainETHReceived is a log parse operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseNonBeaconChainETHReceived(log types.Log) (*ContractIEigenPodNonBeaconChainETHReceived, error) {
	event := new(ContractIEigenPodNonBeaconChainETHReceived)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodNonBeaconChainETHWithdrawnIterator is returned from FilterNonBeaconChainETHWithdrawn and is used to iterate over the raw logs and unpacked data for NonBeaconChainETHWithdrawn events raised by the ContractIEigenPod contract.
type ContractIEigenPodNonBeaconChainETHWithdrawnIterator struct {
	Event *ContractIEigenPodNonBeaconChainETHWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodNonBeaconChainETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodNonBeaconChainETHWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodNonBeaconChainETHWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodNonBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodNonBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodNonBeaconChainETHWithdrawn represents a NonBeaconChainETHWithdrawn event raised by the ContractIEigenPod contract.
type ContractIEigenPodNonBeaconChainETHWithdrawn struct {
	Recipient       common.Address
	AmountWithdrawn *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNonBeaconChainETHWithdrawn is a free log retrieval operation binding the contract event 0x30420aacd028abb3c1fd03aba253ae725d6ddd52d16c9ac4cb5742cd43f53096.
//
// Solidity: event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterNonBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodNonBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "NonBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodNonBeaconChainETHWithdrawnIterator{contract: _ContractIEigenPod.contract, event: "NonBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNonBeaconChainETHWithdrawn is a free log subscription operation binding the contract event 0x30420aacd028abb3c1fd03aba253ae725d6ddd52d16c9ac4cb5742cd43f53096.
//
// Solidity: event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchNonBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodNonBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "NonBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodNonBeaconChainETHWithdrawn)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "NonBeaconChainETHWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNonBeaconChainETHWithdrawn is a log parse operation binding the contract event 0x30420aacd028abb3c1fd03aba253ae725d6ddd52d16c9ac4cb5742cd43f53096.
//
// Solidity: event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseNonBeaconChainETHWithdrawn(log types.Log) (*ContractIEigenPodNonBeaconChainETHWithdrawn, error) {
	event := new(ContractIEigenPodNonBeaconChainETHWithdrawn)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "NonBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodPartialWithdrawalRedeemedIterator is returned from FilterPartialWithdrawalRedeemed and is used to iterate over the raw logs and unpacked data for PartialWithdrawalRedeemed events raised by the ContractIEigenPod contract.
type ContractIEigenPodPartialWithdrawalRedeemedIterator struct {
	Event *ContractIEigenPodPartialWithdrawalRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodPartialWithdrawalRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodPartialWithdrawalRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodPartialWithdrawalRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodPartialWithdrawalRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodPartialWithdrawalRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodPartialWithdrawalRedeemed represents a PartialWithdrawalRedeemed event raised by the ContractIEigenPod contract.
type ContractIEigenPodPartialWithdrawalRedeemed struct {
	ValidatorIndex              *big.Int
	WithdrawalTimestamp         uint64
	Recipient                   common.Address
	PartialWithdrawalAmountGwei uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterPartialWithdrawalRedeemed is a free log retrieval operation binding the contract event 0x8a7335714231dbd551aaba6314f4a97a14c201e53a3e25e1140325cdf67d7a4e.
//
// Solidity: event PartialWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 partialWithdrawalAmountGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterPartialWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodPartialWithdrawalRedeemedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "PartialWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodPartialWithdrawalRedeemedIterator{contract: _ContractIEigenPod.contract, event: "PartialWithdrawalRedeemed", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawalRedeemed is a free log subscription operation binding the contract event 0x8a7335714231dbd551aaba6314f4a97a14c201e53a3e25e1140325cdf67d7a4e.
//
// Solidity: event PartialWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 partialWithdrawalAmountGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchPartialWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodPartialWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "PartialWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodPartialWithdrawalRedeemed)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "PartialWithdrawalRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePartialWithdrawalRedeemed is a log parse operation binding the contract event 0x8a7335714231dbd551aaba6314f4a97a14c201e53a3e25e1140325cdf67d7a4e.
//
// Solidity: event PartialWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 partialWithdrawalAmountGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParsePartialWithdrawalRedeemed(log types.Log) (*ContractIEigenPodPartialWithdrawalRedeemed, error) {
	event := new(ContractIEigenPodPartialWithdrawalRedeemed)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "PartialWithdrawalRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator is returned from FilterRestakedBeaconChainETHWithdrawn and is used to iterate over the raw logs and unpacked data for RestakedBeaconChainETHWithdrawn events raised by the ContractIEigenPod contract.
type ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator struct {
	Event *ContractIEigenPodRestakedBeaconChainETHWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodRestakedBeaconChainETHWithdrawn represents a RestakedBeaconChainETHWithdrawn event raised by the ContractIEigenPod contract.
type ContractIEigenPodRestakedBeaconChainETHWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRestakedBeaconChainETHWithdrawn is a free log retrieval operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodRestakedBeaconChainETHWithdrawnIterator{contract: _ContractIEigenPod.contract, event: "RestakedBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRestakedBeaconChainETHWithdrawn is a free log subscription operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestakedBeaconChainETHWithdrawn is a log parse operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*ContractIEigenPodRestakedBeaconChainETHWithdrawn, error) {
	event := new(ContractIEigenPodRestakedBeaconChainETHWithdrawn)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodRestakingActivatedIterator is returned from FilterRestakingActivated and is used to iterate over the raw logs and unpacked data for RestakingActivated events raised by the ContractIEigenPod contract.
type ContractIEigenPodRestakingActivatedIterator struct {
	Event *ContractIEigenPodRestakingActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodRestakingActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodRestakingActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodRestakingActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodRestakingActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodRestakingActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodRestakingActivated represents a RestakingActivated event raised by the ContractIEigenPod contract.
type ContractIEigenPodRestakingActivated struct {
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRestakingActivated is a free log retrieval operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterRestakingActivated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodRestakingActivatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodRestakingActivatedIterator{contract: _ContractIEigenPod.contract, event: "RestakingActivated", logs: logs, sub: sub}, nil
}

// WatchRestakingActivated is a free log subscription operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchRestakingActivated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodRestakingActivated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodRestakingActivated)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestakingActivated is a log parse operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseRestakingActivated(log types.Log) (*ContractIEigenPodRestakingActivated, error) {
	event := new(ContractIEigenPodRestakingActivated)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodValidatorBalanceUpdatedIterator is returned from FilterValidatorBalanceUpdated and is used to iterate over the raw logs and unpacked data for ValidatorBalanceUpdated events raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorBalanceUpdatedIterator struct {
	Event *ContractIEigenPodValidatorBalanceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodValidatorBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodValidatorBalanceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodValidatorBalanceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodValidatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodValidatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodValidatorBalanceUpdated represents a ValidatorBalanceUpdated event raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorBalanceUpdated struct {
	ValidatorIndex          *big.Int
	BalanceTimestamp        uint64
	NewValidatorBalanceGwei uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterValidatorBalanceUpdated is a free log retrieval operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*ContractIEigenPodValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodValidatorBalanceUpdatedIterator{contract: _ContractIEigenPod.contract, event: "ValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBalanceUpdated is a free log subscription operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodValidatorBalanceUpdated)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorBalanceUpdated is a log parse operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseValidatorBalanceUpdated(log types.Log) (*ContractIEigenPodValidatorBalanceUpdated, error) {
	event := new(ContractIEigenPodValidatorBalanceUpdated)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodValidatorRestakedIterator is returned from FilterValidatorRestaked and is used to iterate over the raw logs and unpacked data for ValidatorRestaked events raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorRestakedIterator struct {
	Event *ContractIEigenPodValidatorRestaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIEigenPodValidatorRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodValidatorRestaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIEigenPodValidatorRestaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIEigenPodValidatorRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodValidatorRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodValidatorRestaked represents a ValidatorRestaked event raised by the ContractIEigenPod contract.
type ContractIEigenPodValidatorRestaked struct {
	ValidatorIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorRestaked is a free log retrieval operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) FilterValidatorRestaked(opts *bind.FilterOpts) (*ContractIEigenPodValidatorRestakedIterator, error) {

	logs, sub, err := _ContractIEigenPod.contract.FilterLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodValidatorRestakedIterator{contract: _ContractIEigenPod.contract, event: "ValidatorRestaked", logs: logs, sub: sub}, nil
}

// WatchValidatorRestaked is a free log subscription operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodValidatorRestaked) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPod.contract.WatchLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodValidatorRestaked)
				if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRestaked is a log parse operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_ContractIEigenPod *ContractIEigenPodFilterer) ParseValidatorRestaked(log types.Log) (*ContractIEigenPodValidatorRestaked, error) {
	event := new(ContractIEigenPodValidatorRestaked)
	if err := _ContractIEigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
