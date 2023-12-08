// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractDelegationManager

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

// IDelegationManagerOperatorDetails is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerOperatorDetails struct {
	EarningsReceiver         common.Address
	DelegationApprover       common.Address
	StakerOptOutWindowBlocks uint32
}

// IDelegationManagerQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerQueuedWithdrawalParams struct {
	Strategies []common.Address
	Shares     []*big.Int
	Withdrawer common.Address
}

// IDelegationManagerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// IStrategyManagerDeprecatedStructQueuedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructQueuedWithdrawal struct {
	Strategies           []common.Address
	Shares               []*big.Int
	Staker               common.Address
	WithdrawerAndNonce   IStrategyManagerDeprecatedStructWithdrawerAndNonce
	WithdrawalStartBlock uint32
	DelegatedAddress     common.Address
}

// IStrategyManagerDeprecatedStructWithdrawerAndNonce is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructWithdrawerAndNonce struct {
	Withdrawer common.Address
	Nonce      *big.Int
}

// ContractDelegationManagerMetaData contains all meta data concerning the ContractDelegationManager contract.
var ContractDelegationManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"_strategyManager\",\"type\":\"address\"},{\"internalType\":\"contractISlasher\",\"name\":\"_slasher\",\"type\":\"address\"},{\"internalType\":\"contractIEigenPodManager\",\"name\":\"_eigenPodManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"OperatorDetailsModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"OperatorMetadataURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"operatorDetails\",\"type\":\"tuple\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"OperatorSharesDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"OperatorSharesIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIStakeRegistryStub\",\"name\":\"stakeRegistry\",\"type\":\"address\"}],\"name\":\"StakeRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"StakerDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"StakerForceUndelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"StakerUndelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"WithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"WithdrawalDelayBlocksSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldWithdrawalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newWithdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"WithdrawalMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"WithdrawalQueued\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STAKER_OPT_OUT_WINDOW_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKER_DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconChainETHStrategy\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateCurrentStakerDelegationDigestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegationApprover\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"approverSalt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateDelegationApprovalDigestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakerNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateStakerDelegationDigestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"internalType\":\"structIDelegationManager.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"calculateWithdrawalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"internalType\":\"structIDelegationManager.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAsTokens\",\"type\":\"bool\"}],\"name\":\"completeQueuedWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"internalType\":\"structIDelegationManager.Withdrawal[]\",\"name\":\"withdrawals\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20[][]\",\"name\":\"tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"receiveAsTokens\",\"type\":\"bool[]\"}],\"name\":\"completeQueuedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeWithdrawalsQueued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"decreaseDelegatedShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"approverSalt\",\"type\":\"bytes32\"}],\"name\":\"delegateTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"name\":\"stakerSignatureAndExpiry\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"approverSalt\",\"type\":\"bytes32\"}],\"name\":\"delegateToBySignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegatedTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"delegationApprover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delegationApproverSaltIsSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"earningsReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getDelegatableShares\",\"outputs\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"increaseDelegatedShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialPausedStatus\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"}],\"internalType\":\"structIStrategyManager.DeprecatedStruct_WithdrawerAndNonce\",\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"}],\"internalType\":\"structIStrategyManager.DeprecatedStruct_QueuedWithdrawal[]\",\"name\":\"withdrawalsToMigrate\",\"type\":\"tuple[]\"}],\"name\":\"migrateQueuedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"modifyOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"operatorDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"}],\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\"}],\"name\":\"queueWithdrawals\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"registeringOperatorDetails\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakeRegistryStub\",\"name\":\"_stakeRegistry\",\"type\":\"address\"}],\"name\":\"setStakeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newWithdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalDelayBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistryStub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"stakerOptOutWindowBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyManager\",\"outputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateOperatorMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200599f3803806200599f833981016040819052620000359162000140565b6001600160a01b0380841660805280821660c052821660a0526200005862000065565b50504660e0525062000194565b600054610100900460ff1615620000d25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000125576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013d57600080fd5b50565b6000806000606084860312156200015657600080fd5b8351620001638162000127565b6020850151909350620001768162000127565b6040850151909250620001898162000127565b809150509250925092565b60805160a05160c05160e05161577662000229600039600061266c01526000818161058b01528181610fc30152818161131901528181611ce8015281816129c501528181613ac00152613fd0015260006107770152600081816104d301528181610f91015281816112e7015281816116e601528181611d7c01528181612a7701528181613bf7015261407601526157766000f3fe608060405234801561001057600080fd5b50600436106103425760003560e01c80635f966f14116101b8578063b134427111610104578063da8be864116100a2578063f16172b01161007c578063f16172b014610906578063f2fde38b14610919578063f698da251461092c578063fabc1cbc1461093457600080fd5b8063da8be864146108cd578063e3b05f2f146108e0578063eea9064b146108f357600080fd5b8063c5e480db116100de578063c5e480db146107ea578063c94b511114610890578063ca661c04146108a3578063cf80873e146108ac57600080fd5b8063b134427114610772578063b7f06ebe14610799578063bb45fef2146107bc57600080fd5b8063778e55f3116101715780638da5cb5b1161014b5780638da5cb5b146107135780639104c3191461072457806399be81c81461073f578063a17884841461075257600080fd5b8063778e55f3146106c25780637f548071146106ed578063886f11951461070057600080fd5b80635f966f141461062c57806360d7faed1461065857806365da12641461066b57806368304835146106945780636d70f7ae146106a7578063715018a6146106ba57600080fd5b806333404396116102925780634fc40b6111610230578063597b36da1161020a578063597b36da146105db5780635ac86ab7146105ee5780635c975abb146106115780635cfe8d2c1461061957600080fd5b80634fc40b61146105c057806350f73e7c146105ca578063595c6a67146105d357600080fd5b80633e28391d1161026c5780633e28391d1461053c578063433773821461055f5780634665bcda146105865780634d50f9a4146105ad57600080fd5b806333404396146104bb57806339b70e38146104ce5780633cdeb5e01461050d57600080fd5b8063136439dd116102ff5780631bbce091116102d95780631bbce0911461044e57806320606b701461046157806328a573ae1461048857806329c77d4f1461049b57600080fd5b8063136439dd146103ef57806316928365146104025780631794bb3c1461043b57600080fd5b806304a4f979146103475780630b9f487a146103815780630dd8dd02146103945780630f589e59146103b457806310d67a2f146103c9578063132d4967146103dc575b600080fd5b61036e7f3b89fca151cbe5122d58acee86cf184413d751d585779bdd19d3bbfa3a306dce81565b6040519081526020015b60405180910390f35b61036e61038f366004614474565b610947565b6103a76103a2366004614513565b610a09565b6040516103789190614554565b6103c76103c23660046145eb565b610d83565b005b6103c76103d736600461463e565b610ed3565b6103c76103ea366004614662565b610f86565b6103c76103fd3660046146a3565b611046565b61036e61041036600461463e565b6001600160a01b0316600090815260996020526040902060010154600160a01b900463ffffffff1690565b6103c7610449366004614662565b611185565b61036e61045c366004614662565b6112ae565b61036e7f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6103c7610496366004614662565b6112dc565b61036e6104a936600461463e565b609b6020526000908152604090205481565b6103c76104c93660046146bc565b61138c565b6104f57f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610378565b6104f561051b36600461463e565b6001600160a01b039081166000908152609960205260409020600101541690565b61054f61054a36600461463e565b6114c9565b6040519015158152602001610378565b61036e7f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b81565b6104f57f000000000000000000000000000000000000000000000000000000000000000081565b6103c76105bb3660046146a3565b6114e9565b61036e6213c68081565b61036e609d5481565b6103c76115c0565b61036e6105e93660046149fc565b611687565b61054f6105fc366004614a38565b606654600160ff9092169190911b9081161490565b60665461036e565b6103c7610627366004614aae565b6116b7565b6104f561063a36600461463e565b6001600160a01b039081166000908152609960205260409020541690565b6103c7610666366004614c0e565b611962565b6104f561067936600461463e565b609a602052600090815260409020546001600160a01b031681565b60a0546104f5906001600160a01b031681565b61054f6106b536600461463e565b6119fd565b6103c7611a1d565b61036e6106d0366004614c9d565b609860209081526000928352604080842090915290825290205481565b6103c76106fb366004614d7e565b611a31565b6065546104f5906001600160a01b031681565b6033546001600160a01b03166104f5565b6104f573beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b6103c761074d366004614e0e565b611b36565b61036e61076036600461463e565b609f6020526000908152604090205481565b6104f57f000000000000000000000000000000000000000000000000000000000000000081565b61054f6107a73660046146a3565b609e6020526000908152604090205460ff1681565b61054f6107ca366004614e43565b609c60209081526000928352604080842090915290825290205460ff1681565b61085a6107f836600461463e565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03948516815260998352839020835191820184528054851682526001015493841691810191909152600160a01b90920463ffffffff169082015290565b6040805182516001600160a01b039081168252602080850151909116908201529181015163ffffffff1690820152606001610378565b61036e61089e366004614e6f565b611c08565b61036e61c4e081565b6108bf6108ba36600461463e565b611cc1565b604051610378929190614f2b565b61036e6108db36600461463e565b612079565b6103c76108ee36600461463e565b6123ed565b6103c7610901366004614f50565b612555565b6103c7610914366004614fa8565b612561565b6103c761092736600461463e565b6125f2565b61036e612668565b6103c76109423660046146a3565b6126a6565b604080517f3b89fca151cbe5122d58acee86cf184413d751d585779bdd19d3bbfa3a306dce6020808301919091526001600160a01b038681168385015288811660608401528716608083015260a0820185905260c08083018590528351808403909101815260e09092019092528051910120600090816109c5612668565b60405161190160f01b602082015260228101919091526042810183905260620160408051808303601f19018152919052805160209091012098975050505050505050565b60665460609060019060029081161415610a3e5760405162461bcd60e51b8152600401610a3590614fc4565b60405180910390fd5b6000836001600160401b03811115610a5857610a5861477f565b604051908082528060200260200182016040528015610a81578160200160208202803683370190505b50905060005b84811015610d7a57858582818110610aa157610aa1614ffb565b9050602002810190610ab39190615011565b610ac1906020810190615031565b9050868683818110610ad557610ad5614ffb565b9050602002810190610ae79190615011565b610af19080615031565b905014610b665760405162461bcd60e51b815260206004820152603860248201527f44656c65676174696f6e4d616e616765722e717565756557697468647261776160448201527f6c3a20696e707574206c656e677468206d69736d6174636800000000000000006064820152608401610a35565b6000868683818110610b7a57610b7a614ffb565b9050602002810190610b8c9190615011565b610b9d90606081019060400161463e565b6001600160a01b03161415610c2b5760405162461bcd60e51b815260206004820152604860248201527f44656c65676174696f6e4d616e616765722e717565756557697468647261776160448201527f6c3a206d7573742070726f766964652076616c6964207769746864726177616c606482015267206164647265737360c01b608482015260a401610a35565b336000818152609a60205260409020546001600160a01b031690610d4a9082898986818110610c5c57610c5c614ffb565b9050602002810190610c6e9190615011565b610c7f90606081019060400161463e565b8a8a87818110610c9157610c91614ffb565b9050602002810190610ca39190615011565b610cad9080615031565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508d9150899050818110610cf357610cf3614ffb565b9050602002810190610d059190615011565b610d13906020810190615031565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061280292505050565b838381518110610d5c57610d5c614ffb565b60209081029190910101525080610d7281615090565b915050610a87565b50949350505050565b336000908152609960205260409020546001600160a01b031615610e1d5760405162461bcd60e51b815260206004820152604560248201527f44656c65676174696f6e4d616e616765722e726567697374657241734f70657260448201527f61746f723a206f70657261746f722068617320616c72656164792072656769736064820152641d195c995960da1b608482015260a401610a35565b610e273384612c3c565b604080518082019091526060815260006020820152610e493380836000612ed8565b336001600160a01b03167f8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e285604051610e8291906150ab565b60405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908484604051610ec59291906150fd565b60405180910390a250505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4a919061512c565b6001600160a01b0316336001600160a01b031614610f7a5760405162461bcd60e51b8152600401610a3590615149565b610f8381613291565b50565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610fe55750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b6110015760405162461bcd60e51b8152600401610a3590615193565b61100a836114c9565b15611041576001600160a01b038084166000908152609a60205260409020541661103681858585613388565b61103f81613403565b505b505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561108e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b291906151f0565b6110ce5760405162461bcd60e51b8152600401610a359061520d565b606654818116146111475760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610a35565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156111a55750600054600160ff909116105b806111bf5750303b1580156111bf575060005460ff166001145b6112225760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610a35565b6000805460ff191660011790558015611245576000805461ff0019166101001790555b61124f83836134bd565b6112576135a3565b6097556112638461363a565b801561103f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b6001600160a01b0383166000908152609b60205260408120546112d385828686611c08565b95945050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061133b5750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b6113575760405162461bcd60e51b8152600401610a3590615193565b611360836114c9565b15611041576001600160a01b038084166000908152609a6020526040902054166110368185858561368c565b606654600290600490811614156113b55760405162461bcd60e51b8152600401610a3590614fc4565b600260c95414156114085760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610a35565b600260c95560005b888110156114b8576114a88a8a8381811061142d5761142d614ffb565b905060200281019061143f9190615255565b89898481811061145157611451614ffb565b90506020028101906114639190615031565b89898681811061147557611475614ffb565b9050602002013588888781811061148e5761148e614ffb565b90506020020160208101906114a3919061526b565b613707565b6114b181615090565b9050611410565b5050600160c9555050505050505050565b6001600160a01b039081166000908152609a602052604090205416151590565b6114f1613d81565b61c4e081111561157f5760405162461bcd60e51b815260206004820152604d60248201527f44656c65676174696f6e4d616e616765722e7365745769746864726177616c4460448201527f656c6179426c6f636b733a206e65775769746864726177616c44656c6179426c60648201526c0dec6d6e640e8dede40d0d2ced609b1b608482015260a401610a35565b609d5460408051918252602082018390527f4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e910160405180910390a1609d55565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611608573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162c91906151f0565b6116485760405162461bcd60e51b8152600401610a359061520d565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60008160405160200161169a91906152fc565b604051602081830303815290604052805190602001209050919050565b60005b815181101561195e5760008282815181106116d7576116d7614ffb565b602002602001015190506000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cd293f6f846040518263ffffffff1660e01b8152600401611730919061530f565b60408051808303816000875af115801561174e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061177291906153be565b915091508115611950576040808401516001600160a01b0381166000908152609f602052918220805491928291906117a983615090565b919050555060006040518060e00160405280846001600160a01b031681526020018760a001516001600160a01b031681526020018760600151600001516001600160a01b03168152602001838152602001876080015163ffffffff1681526020018760000151815260200187602001518152509050600061182982611687565b6000818152609e602052604090205490915060ff16156118bf5760405162461bcd60e51b815260206004820152604560248201527f44656c65676174696f6e4d616e616765722e6d6967726174655175657565645760448201527f69746864726177616c733a207769746864726177616c20616c72656164792065606482015264786973747360d81b608482015260a401610a35565b6000818152609e602052604090819020805460ff19166001179055517f9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f99061190a90839085906153ec565b60405180910390a160408051868152602081018390527fdc00758b65eef71dc3780c04ebe36cab6bdb266c3a698187e29e0f0dca012630910160405180910390a1505050505b8360010193505050506116ba565b5050565b6066546002906004908116141561198b5760405162461bcd60e51b8152600401610a3590614fc4565b600260c95414156119de5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610a35565b600260c9556119f08686868686613707565b5050600160c95550505050565b6001600160a01b0390811660009081526099602052604090205416151590565b611a25613d81565b611a2f600061363a565b565b4283602001511015611ab55760405162461bcd60e51b815260206004820152604160248201527f44656c65676174696f6e4d616e616765722e64656c6567617465546f4279536960448201527f676e61747572653a207374616b6572207369676e6174757265206578706972656064820152601960fa1b608482015260a401610a35565b6000609b6000876001600160a01b03166001600160a01b031681526020019081526020016000205490506000611af18783888860200151611c08565b6001600160a01b0388166000908152609b602052604090206001840190558551909150611b219088908390613ddb565b611b2d87878686612ed8565b50505050505050565b611b3f336119fd565b611bc15760405162461bcd60e51b815260206004820152604760248201527f44656c65676174696f6e4d616e616765722e7570646174654f70657261746f7260448201527f4d657461646174615552493a2063616c6c6572206d75737420626520616e206f6064820152663832b930ba37b960c91b608482015260a401610a35565b336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908383604051611bfc9291906150fd565b60405180910390a25050565b604080517f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b6020808301919091526001600160a01b0387811683850152851660608301526080820186905260a08083018590528351808403909101815260c0909201909252805191012060009081611c7e612668565b60405161190160f01b602082015260228101919091526042810183905260620160408051808303601f190181529190528051602090910120979650505050505050565b6040516360f4062b60e01b81526001600160a01b03828116600483015260609182916000917f0000000000000000000000000000000000000000000000000000000000000000909116906360f4062b90602401602060405180830381865afa158015611d31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d559190615405565b6040516394f649dd60e01b81526001600160a01b03868116600483015291925060009182917f0000000000000000000000000000000000000000000000000000000000000000909116906394f649dd90602401600060405180830381865afa158015611dc5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ded9190810190615479565b9150915060008313611e0457909590945092505050565b606080835160001415611ebe576040805160018082528183019092529060208083019080368337505060408051600180825281830190925292945090506020808301908036833701905050905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac082600081518110611e7957611e79614ffb565b60200260200101906001600160a01b031690816001600160a01b0316815250508481600081518110611ead57611ead614ffb565b60200260200101818152505061206c565b8351611ecb90600161553d565b6001600160401b03811115611ee257611ee261477f565b604051908082528060200260200182016040528015611f0b578160200160208202803683370190505b50915081516001600160401b03811115611f2757611f2761477f565b604051908082528060200260200182016040528015611f50578160200160208202803683370190505b50905060005b8451811015611fea57848181518110611f7157611f71614ffb565b6020026020010151838281518110611f8b57611f8b614ffb565b60200260200101906001600160a01b031690816001600160a01b031681525050838181518110611fbd57611fbd614ffb565b6020026020010151828281518110611fd757611fd7614ffb565b6020908102919091010152600101611f56565b5073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0826001845161200f9190615555565b8151811061201f5761201f614ffb565b60200260200101906001600160a01b031690816001600160a01b03168152505084816001845161204f9190615555565b8151811061205f5761205f614ffb565b6020026020010181815250505b9097909650945050505050565b606654600090600190600290811614156120a55760405162461bcd60e51b8152600401610a3590614fc4565b6120ae836114c9565b61212e5760405162461bcd60e51b8152602060048201526044602482018190527f44656c65676174696f6e4d616e616765722e756e64656c65676174653a207374908201527f616b6572206d7573742062652064656c65676174656420746f20756e64656c656064820152636761746560e01b608482015260a401610a35565b6001600160a01b038084166000908152609a602052604090205416612152846119fd565b156121c55760405162461bcd60e51b815260206004820152603d60248201527f44656c65676174696f6e4d616e616765722e756e64656c65676174653a206f7060448201527f657261746f72732063616e6e6f7420626520756e64656c6567617465640000006064820152608401610a35565b6001600160a01b0384166122415760405162461bcd60e51b815260206004820152603c60248201527f44656c65676174696f6e4d616e616765722e756e64656c65676174653a20636160448201527f6e6e6f7420756e64656c6567617465207a65726f2061646472657373000000006064820152608401610a35565b336001600160a01b03851614806122605750336001600160a01b038216145b8061228757506001600160a01b038181166000908152609960205260409020600101541633145b6122f95760405162461bcd60e51b815260206004820152603d60248201527f44656c65676174696f6e4d616e616765722e756e64656c65676174653a20636160448201527f6c6c65722063616e6e6f7420756e64656c6567617465207374616b65720000006064820152608401610a35565b60008061230586611cc1565b9092509050336001600160a01b0387161461235b57826001600160a01b0316866001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a35b826001600160a01b0316866001600160a01b03167ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467660405160405180910390a36001600160a01b0386166000908152609a6020526040902080546001600160a01b031916905581516123d45750600093506123e7915050565b6123e18684888585612802565b94505050505b50919050565b6123f5613d81565b60a0546001600160a01b0316156124745760405162461bcd60e51b815260206004820152603d60248201527f44656c65676174696f6e4d616e616765722e7365745374616b6552656769737460448201527f72793a207374616b65526567697374727920616c7265616479207365740000006064820152608401610a35565b6001600160a01b0381166125015760405162461bcd60e51b815260206004820152604860248201527f44656c65676174696f6e4d616e616765722e7365745374616b6552656769737460448201527f72793a207374616b6552656769737472792063616e6e6f74206265207a65726f606482015267206164647265737360c01b608482015260a401610a35565b60a080546001600160a01b0319166001600160a01b0383169081179091556040519081527fce6d874069bceda1867eca9c60636bbf262e15213041658273d803a2b609a51f9060200160405180910390a150565b61104133848484612ed8565b61256a336119fd565b6125e85760405162461bcd60e51b815260206004820152604360248201527f44656c65676174696f6e4d616e616765722e6d6f646966794f70657261746f7260448201527f44657461696c733a2063616c6c6572206d75737420626520616e206f706572616064820152623a37b960e91b608482015260a401610a35565b610f833382612c3c565b6125fa613d81565b6001600160a01b03811661265f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a35565b610f838161363a565b60007f0000000000000000000000000000000000000000000000000000000000000000461415612699575060975490565b6126a16135a3565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156126f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061271d919061512c565b6001600160a01b0316336001600160a01b03161461274d5760405162461bcd60e51b8152600401610a3590615149565b6066541981196066541916146127cb5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610a35565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161117a565b60006001600160a01b0386166128995760405162461bcd60e51b815260206004820152605060248201527f44656c65676174696f6e4d616e616765722e5f72656d6f76655368617265734160448201527f6e6451756575655769746864726177616c3a207374616b65722063616e6e6f7460648201526f206265207a65726f206164647265737360801b608482015260a401610a35565b82516129235760405162461bcd60e51b815260206004820152604d60248201527f44656c65676174696f6e4d616e616765722e5f72656d6f76655368617265734160448201527f6e6451756575655769746864726177616c3a207374726174656769657320636160648201526c6e6e6f7420626520656d70747960981b608482015260a401610a35565b60005b8351811015612b32576001600160a01b0386161561297c5761297c868886848151811061295557612955614ffb565b602002602001015186858151811061296f5761296f614ffb565b6020026020010151613388565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06001600160a01b03168482815181106129ac576129ac614ffb565b60200260200101516001600160a01b03161415612a75577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663beffbb8988858481518110612a0557612a05614ffb565b60200260200101516040518363ffffffff1660e01b8152600401612a3e9291906001600160a01b03929092168252602082015260400190565b600060405180830381600087803b158015612a5857600080fd5b505af1158015612a6c573d6000803e3d6000fd5b50505050612b2a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638c80d4e588868481518110612ab757612ab7614ffb565b6020026020010151868581518110612ad157612ad1614ffb565b60200260200101516040518463ffffffff1660e01b8152600401612af79392919061556c565b600060405180830381600087803b158015612b1157600080fd5b505af1158015612b25573d6000803e3d6000fd5b505050505b600101612926565b506001600160a01b03851615612b4b57612b4b85613403565b6001600160a01b0386166000908152609f60205260408120805491829190612b7283615090565b919050555060006040518060e00160405280896001600160a01b03168152602001886001600160a01b03168152602001876001600160a01b031681526020018381526020014363ffffffff1681526020018681526020018581525090506000612bda82611687565b6000818152609e602052604090819020805460ff19166001179055519091507f9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f990612c2890839085906153ec565b60405180910390a198975050505050505050565b6000612c4b602083018361463e565b6001600160a01b03161415612ce55760405162461bcd60e51b815260206004820152605460248201527f44656c65676174696f6e4d616e616765722e5f7365744f70657261746f72446560448201527f7461696c733a2063616e6e6f742073657420606561726e696e677352656365696064820152737665726020746f207a65726f206164647265737360601b608482015260a401610a35565b6213c680612cf96060830160408401615590565b63ffffffff161115612dae5760405162461bcd60e51b815260206004820152606c60248201527f44656c65676174696f6e4d616e616765722e5f7365744f70657261746f72446560448201527f7461696c733a207374616b65724f70744f757457696e646f77426c6f636b732060648201527f63616e6e6f74206265203e204d41585f5354414b45525f4f50545f4f55545f5760848201526b494e444f575f424c4f434b5360a01b60a482015260c401610a35565b6001600160a01b0382166000908152609960205260409081902060010154600160a01b900463ffffffff1690612dea9060608401908401615590565b63ffffffff161015612e805760405162461bcd60e51b815260206004820152605360248201527f44656c65676174696f6e4d616e616765722e5f7365744f70657261746f72446560448201527f7461696c733a207374616b65724f70744f757457696e646f77426c6f636b732060648201527218d85b9b9bdd08189948191958dc99585cd959606a1b608482015260a401610a35565b6001600160a01b03821660009081526099602052604090208190612ea482826155cd565b505060405133907ffebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac90611bfc9084906150ab565b60665460009060019081161415612f015760405162461bcd60e51b8152600401610a3590614fc4565b612f0a856114c9565b15612f875760405162461bcd60e51b815260206004820152604160248201527f44656c65676174696f6e4d616e616765722e5f64656c65676174653a2073746160448201527f6b657220697320616c7265616479206163746976656c792064656c65676174656064820152601960fa1b608482015260a401610a35565b612f90846119fd565b6130105760405162461bcd60e51b815260206004820152604560248201527f44656c65676174696f6e4d616e616765722e5f64656c65676174653a206f706560448201527f7261746f72206973206e6f74207265676973746572656420696e20456967656e6064820152642630bcb2b960d91b608482015260a401610a35565b6001600160a01b038085166000908152609960205260409020600101541680158015906130465750336001600160a01b03821614155b801561305b5750336001600160a01b03861614155b156131c85742846020015110156130da5760405162461bcd60e51b815260206004820152603760248201527f44656c65676174696f6e4d616e616765722e5f64656c65676174653a2061707060448201527f726f766572207369676e617475726520657870697265640000000000000000006064820152608401610a35565b6001600160a01b0381166000908152609c6020908152604080832086845290915290205460ff16156131745760405162461bcd60e51b815260206004820152603760248201527f44656c65676174696f6e4d616e616765722e5f64656c65676174653a2061707060448201527f726f76657253616c7420616c7265616479207370656e740000000000000000006064820152608401610a35565b6001600160a01b0381166000908152609c6020908152604080832086845282528220805460ff191660011790558501516131b5908890889085908890610947565b90506131c682828760000151613ddb565b505b6001600160a01b038681166000818152609a602052604080822080546001600160a01b031916948a169485179055517fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d87433049190a360008061322788611cc1565b9150915060005b825181101561327d57613275888a85848151811061324e5761324e614ffb565b602002602001015185858151811061326857613268614ffb565b602002602001015161368c565b60010161322e565b5061328787613403565b5050505050505050565b6001600160a01b03811661331f5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610a35565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b038085166000908152609860209081526040808320938616835292905290812080548392906133bf908490615555565b92505081905550836001600160a01b03167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd848484604051610ec59392919061556c565b60a0546001600160a01b031615610f835760408051600180825281830190925260009160208083019080368337019050509050818160008151811061344a5761344a614ffb565b6001600160a01b03928316602091820292909201015260a05460405163ce977ec360e01b815291169063ce977ec390613487908490600401615630565b600060405180830381600087803b1580156134a157600080fd5b505af11580156134b5573d6000803e3d6000fd5b505050505050565b6065546001600160a01b03161580156134de57506001600160a01b03821615155b6135605760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610a35565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261195e82613291565b604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038085166000908152609860209081526040808320938616835292905290812080548392906136c390849061553d565b92505081905550836001600160a01b03167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c848484604051610ec59392919061556c565b60006137156105e987615671565b6000818152609e602052604090205490915060ff1661378a5760405162461bcd60e51b815260206004820152603e602482015260008051602061572183398151915260448201527f416374696f6e3a20616374696f6e206973206e6f7420696e20717565756500006064820152608401610a35565b609d54439061379f60a0890160808a01615590565b63ffffffff166137af919061553d565b11156138375760405162461bcd60e51b8152602060048201526057602482015260008051602061572183398151915260448201527f416374696f6e3a207769746864726177616c44656c6179426c6f636b7320706560648201527f72696f6420686173206e6f742079657420706173736564000000000000000000608482015260a401610a35565b613847606087016040880161463e565b6001600160a01b0316336001600160a01b0316146138cf5760405162461bcd60e51b815260206004820152604b602482015260008051602061572183398151915260448201527f416374696f6e3a206f6e6c7920776974686472617765722063616e20636f6d7060648201526a3632ba329030b1ba34b7b760a91b608482015260a401610a35565b8115613946576138e260a0870187615031565b851490506139465760405162461bcd60e51b815260206004820152603d602482015260008051602061572183398151915260448201527f416374696f6e3a20696e707574206c656e677468206d69736d617463680000006064820152608401610a35565b6000818152609e60205260409020805460ff191690558115613a1c5760005b61397260a0880188615031565b9050811015613a1657613a0e61398b602089018961463e565b3361399960a08b018b615031565b858181106139a9576139a9614ffb565b90506020020160208101906139be919061463e565b6139cb60c08c018c615031565b868181106139db576139db614ffb565b905060200201358a8a878181106139f4576139f4614ffb565b9050602002016020810190613a09919061463e565b613f95565b600101613965565b50613d46565b336000908152609a60205260408120546001600160a01b0316905b613a4460a0890189615031565b9050811015613d3a5773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0613a6f60a08a018a615031565b83818110613a7f57613a7f614ffb565b9050602002016020810190613a94919061463e565b6001600160a01b03161415613bed576000613ab260208a018a61463e565b905060006001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016630e81073c83613af360c08e018e615031565b87818110613b0357613b03614ffb565b6040516001600160e01b031960e087901b1681526001600160a01b03909416600485015260200291909101356024830152506044016020604051808303816000875af1158015613b57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b7b9190615405565b6001600160a01b038084166000908152609a6020526040902054919250168015613be557613bdc8184613bb160a08f018f615031565b88818110613bc157613bc1614ffb565b9050602002016020810190613bd6919061463e565b8561368c565b613be581613403565b505050613d32565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166350ff722533613c2a60a08c018c615031565b85818110613c3a57613c3a614ffb565b9050602002016020810190613c4f919061463e565b613c5c60c08d018d615031565b86818110613c6c57613c6c614ffb565b905060200201356040518463ffffffff1660e01b8152600401613c919392919061556c565b600060405180830381600087803b158015613cab57600080fd5b505af1158015613cbf573d6000803e3d6000fd5b505050506001600160a01b03821615613d3257613d328233613ce460a08c018c615031565b85818110613cf457613cf4614ffb565b9050602002016020810190613d09919061463e565b613d1660c08d018d615031565b86818110613d2657613d26614ffb565b9050602002013561368c565b600101613a37565b50613d4481613403565b505b6040518181527fc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d9060200160405180910390a1505050505050565b6033546001600160a01b03163314611a2f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a35565b6001600160a01b0383163b15613ef557604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e90613e1b9086908690600401615683565b602060405180830381865afa158015613e38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e5c91906156e0565b6001600160e01b031916146110415760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a401610a35565b826001600160a01b0316613f0983836140da565b6001600160a01b0316146110415760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a401610a35565b6001600160a01b03831673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014156140405760405162387b1360e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063387b1300906140099088908890879060040161556c565b600060405180830381600087803b15801561402357600080fd5b505af1158015614037573d6000803e3d6000fd5b505050506140d3565b60405163c608c7f360e01b81526001600160a01b03858116600483015284811660248301526044820184905282811660648301527f0000000000000000000000000000000000000000000000000000000000000000169063c608c7f390608401600060405180830381600087803b1580156140ba57600080fd5b505af11580156140ce573d6000803e3d6000fd5b505050505b5050505050565b60008060006140e985856140fe565b915091506140f68161416e565b509392505050565b6000808251604114156141355760208301516040840151606085015160001a61412987828585614329565b94509450505050614167565b82516040141561415f5760208301516040840151614154868383614416565b935093505050614167565b506000905060025b9250929050565b60008160048111156141825761418261570a565b141561418b5750565b600181600481111561419f5761419f61570a565b14156141ed5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610a35565b60028160048111156142015761420161570a565b141561424f5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610a35565b60038160048111156142635761426361570a565b14156142bc5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610a35565b60048160048111156142d0576142d061570a565b1415610f835760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610a35565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614360575060009050600361440d565b8460ff16601b1415801561437857508460ff16601c14155b15614389575060009050600461440d565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156143dd573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166144065760006001925092505061440d565b9150600090505b94509492505050565b6000806001600160ff1b0383168161443360ff86901c601b61553d565b905061444187828885614329565b935093505050935093915050565b6001600160a01b0381168114610f8357600080fd5b803561446f8161444f565b919050565b600080600080600060a0868803121561448c57600080fd5b85356144978161444f565b945060208601356144a78161444f565b935060408601356144b78161444f565b94979396509394606081013594506080013592915050565b60008083601f8401126144e157600080fd5b5081356001600160401b038111156144f857600080fd5b6020830191508360208260051b850101111561416757600080fd5b6000806020838503121561452657600080fd5b82356001600160401b0381111561453c57600080fd5b614548858286016144cf565b90969095509350505050565b6020808252825182820181905260009190848201906040850190845b8181101561458c57835183529284019291840191600101614570565b50909695505050505050565b6000606082840312156123e757600080fd5b60008083601f8401126145bc57600080fd5b5081356001600160401b038111156145d357600080fd5b60208301915083602082850101111561416757600080fd5b60008060006080848603121561460057600080fd5b61460a8585614598565b925060608401356001600160401b0381111561462557600080fd5b614631868287016145aa565b9497909650939450505050565b60006020828403121561465057600080fd5b813561465b8161444f565b9392505050565b60008060006060848603121561467757600080fd5b83356146828161444f565b925060208401356146928161444f565b929592945050506040919091013590565b6000602082840312156146b557600080fd5b5035919050565b6000806000806000806000806080898b0312156146d857600080fd5b88356001600160401b03808211156146ef57600080fd5b6146fb8c838d016144cf565b909a50985060208b013591508082111561471457600080fd5b6147208c838d016144cf565b909850965060408b013591508082111561473957600080fd5b6147458c838d016144cf565b909650945060608b013591508082111561475e57600080fd5b5061476b8b828c016144cf565b999c989b5096995094979396929594505050565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156147b7576147b761477f565b60405290565b604080519081016001600160401b03811182821017156147b7576147b761477f565b60405160c081016001600160401b03811182821017156147b7576147b761477f565b604051601f8201601f191681016001600160401b03811182821017156148295761482961477f565b604052919050565b63ffffffff81168114610f8357600080fd5b803561446f81614831565b60006001600160401b038211156148675761486761477f565b5060051b60200190565b600082601f83011261488257600080fd5b813560206148976148928361484e565b614801565b82815260059290921b840181019181810190868411156148b657600080fd5b8286015b848110156148da5780356148cd8161444f565b83529183019183016148ba565b509695505050505050565b600082601f8301126148f657600080fd5b813560206149066148928361484e565b82815260059290921b8401810191818101908684111561492557600080fd5b8286015b848110156148da5780358352918301918301614929565b600060e0828403121561495257600080fd5b61495a614795565b905061496582614464565b815261497360208301614464565b602082015261498460408301614464565b60408201526060820135606082015261499f60808301614843565b608082015260a08201356001600160401b03808211156149be57600080fd5b6149ca85838601614871565b60a084015260c08401359150808211156149e357600080fd5b506149f0848285016148e5565b60c08301525092915050565b600060208284031215614a0e57600080fd5b81356001600160401b03811115614a2457600080fd5b614a3084828501614940565b949350505050565b600060208284031215614a4a57600080fd5b813560ff8116811461465b57600080fd5b600060408284031215614a6d57600080fd5b614a756147bd565b90508135614a828161444f565b815260208201356bffffffffffffffffffffffff81168114614aa357600080fd5b602082015292915050565b60006020808385031215614ac157600080fd5b82356001600160401b0380821115614ad857600080fd5b818501915085601f830112614aec57600080fd5b8135614afa6148928261484e565b81815260059190911b83018401908481019088831115614b1957600080fd5b8585015b83811015614bf357803585811115614b355760008081fd5b860160e0818c03601f1901811315614b4d5760008081fd5b614b556147df565b8983013588811115614b675760008081fd5b614b758e8c83870101614871565b82525060408084013589811115614b8c5760008081fd5b614b9a8f8d838801016148e5565b8c840152506060614bac818601614464565b8284015260809150614bc08f838701614a5b565b90830152614bd060c08501614843565b90820152614bdf838301614464565b60a082015285525050918601918601614b1d565b5098975050505050505050565b8015158114610f8357600080fd5b600080600080600060808688031215614c2657600080fd5b85356001600160401b0380821115614c3d57600080fd5b9087019060e0828a031215614c5157600080fd5b90955060208701359080821115614c6757600080fd5b50614c74888289016144cf565b909550935050604086013591506060860135614c8f81614c00565b809150509295509295909350565b60008060408385031215614cb057600080fd5b8235614cbb8161444f565b91506020830135614ccb8161444f565b809150509250929050565b600060408284031215614ce857600080fd5b614cf06147bd565b905081356001600160401b0380821115614d0957600080fd5b818401915084601f830112614d1d57600080fd5b8135602082821115614d3157614d3161477f565b614d43601f8301601f19168201614801565b92508183528681838601011115614d5957600080fd5b8181850182850137600081838501015282855280860135818601525050505092915050565b600080600080600060a08688031215614d9657600080fd5b8535614da18161444f565b94506020860135614db18161444f565b935060408601356001600160401b0380821115614dcd57600080fd5b614dd989838a01614cd6565b94506060880135915080821115614def57600080fd5b50614dfc88828901614cd6565b95989497509295608001359392505050565b60008060208385031215614e2157600080fd5b82356001600160401b03811115614e3757600080fd5b614548858286016145aa565b60008060408385031215614e5657600080fd5b8235614e618161444f565b946020939093013593505050565b60008060008060808587031215614e8557600080fd5b8435614e908161444f565b9350602085013592506040850135614ea78161444f565b9396929550929360600135925050565b600081518084526020808501945080840160005b83811015614ef05781516001600160a01b031687529582019590820190600101614ecb565b509495945050505050565b600081518084526020808501945080840160005b83811015614ef057815187529582019590820190600101614f0f565b604081526000614f3e6040830185614eb7565b82810360208401526112d38185614efb565b600080600060608486031215614f6557600080fd5b8335614f708161444f565b925060208401356001600160401b03811115614f8b57600080fd5b614f9786828701614cd6565b925050604084013590509250925092565b600060608284031215614fba57600080fd5b61465b8383614598565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b60008235605e1983360301811261502757600080fd5b9190910192915050565b6000808335601e1984360301811261504857600080fd5b8301803591506001600160401b0382111561506257600080fd5b6020019150600581901b360382131561416757600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156150a4576150a461507a565b5060010190565b6060810182356150ba8161444f565b6001600160a01b0390811683526020840135906150d68261444f565b16602083015260408301356150ea81614831565b63ffffffff811660408401525092915050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561513e57600080fd5b815161465b8161444f565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60208082526037908201527f44656c65676174696f6e4d616e616765723a206f6e6c7953747261746567794d60408201527f616e616765724f72456967656e506f644d616e61676572000000000000000000606082015260800190565b60006020828403121561520257600080fd5b815161465b81614c00565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000823560de1983360301811261502757600080fd5b60006020828403121561527d57600080fd5b813561465b81614c00565b600060018060a01b03808351168452806020840151166020850152806040840151166040850152506060820151606084015263ffffffff608083015116608084015260a082015160e060a08501526152e360e0850182614eb7565b905060c083015184820360c08601526112d38282614efb565b60208152600061465b6020830184615288565b602081526000825160e0602084015261532c610100840182614eb7565b90506020840151601f198483030160408501526153498282614efb565b915050604084015160018060a01b03808216606086015260608601519150808251166080860152506bffffffffffffffffffffffff60208201511660a08501525060808401516153a160c085018263ffffffff169052565b5060a08401516001600160a01b03811660e0850152509392505050565b600080604083850312156153d157600080fd5b82516153dc81614c00565b6020939093015192949293505050565b828152604060208201526000614a306040830184615288565b60006020828403121561541757600080fd5b5051919050565b600082601f83011261542f57600080fd5b8151602061543f6148928361484e565b82815260059290921b8401810191818101908684111561545e57600080fd5b8286015b848110156148da5780518352918301918301615462565b6000806040838503121561548c57600080fd5b82516001600160401b03808211156154a357600080fd5b818501915085601f8301126154b757600080fd5b815160206154c76148928361484e565b82815260059290921b840181019181810190898411156154e657600080fd5b948201945b8386101561550d5785516154fe8161444f565b825294820194908201906154eb565b9188015191965090935050508082111561552657600080fd5b506155338582860161541e565b9150509250929050565b600082198211156155505761555061507a565b500190565b6000828210156155675761556761507a565b500390565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6000602082840312156155a257600080fd5b813561465b81614831565b80546001600160a01b0319166001600160a01b0392909216919091179055565b81356155d88161444f565b6155e281836155ad565b506001810160208301356155f58161444f565b6155ff81836155ad565b50604083013561560e81614831565b815463ffffffff60a01b191660a09190911b63ffffffff60a01b161790555050565b6020808252825182820181905260009190848201906040850190845b8181101561458c5783516001600160a01b03168352928401929184019160010161564c565b600061567d3683614940565b92915050565b82815260006020604081840152835180604085015260005b818110156156b75785810183015185820160600152820161569b565b818111156156c9576000606083870101525b50601f01601f191692909201606001949350505050565b6000602082840312156156f257600080fd5b81516001600160e01b03198116811461465b57600080fd5b634e487b7160e01b600052602160045260246000fdfe44656c65676174696f6e4d616e616765722e636f6d706c657465517565756564a26469706673582212205ea0131cf5b5fb99da1fff7b870c94af10f99c971aedbf34ac1853af7e1af52464736f6c634300080c0033",
}

// ContractDelegationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractDelegationManagerMetaData.ABI instead.
var ContractDelegationManagerABI = ContractDelegationManagerMetaData.ABI

// ContractDelegationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractDelegationManagerMetaData.Bin instead.
var ContractDelegationManagerBin = ContractDelegationManagerMetaData.Bin

// DeployContractDelegationManager deploys a new Ethereum contract, binding an instance of ContractDelegationManager to it.
func DeployContractDelegationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _slasher common.Address, _eigenPodManager common.Address) (common.Address, *types.Transaction, *ContractDelegationManager, error) {
	parsed, err := ContractDelegationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractDelegationManagerBin), backend, _strategyManager, _slasher, _eigenPodManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractDelegationManager{ContractDelegationManagerCaller: ContractDelegationManagerCaller{contract: contract}, ContractDelegationManagerTransactor: ContractDelegationManagerTransactor{contract: contract}, ContractDelegationManagerFilterer: ContractDelegationManagerFilterer{contract: contract}}, nil
}

// ContractDelegationManagerMethods is an auto generated interface around an Ethereum contract.
type ContractDelegationManagerMethods interface {
	ContractDelegationManagerCalls
	ContractDelegationManagerTransacts
	ContractDelegationManagerFilters
}

// ContractDelegationManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractDelegationManagerCalls interface {
	DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	MAXSTAKEROPTOUTWINDOWBLOCKS(opts *bind.CallOpts) (*big.Int, error)

	MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error)

	STAKERDELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error)

	CalculateCurrentStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error)

	CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error)

	CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerWithdrawal) ([32]byte, error)

	CumulativeWithdrawalsQueued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	DelegatedTo(opts *bind.CallOpts, arg0 common.Address) (common.Address, error)

	DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error)

	DelegationApproverSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	EarningsReceiver(opts *bind.CallOpts, operator common.Address) (common.Address, error)

	EigenPodManager(opts *bind.CallOpts) (common.Address, error)

	GetDelegatableShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error)

	IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error)

	IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error)

	OperatorDetails(opts *bind.CallOpts, operator common.Address) (IDelegationManagerOperatorDetails, error)

	OperatorShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PendingWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	Slasher(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)

	StakerNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	StakerOptOutWindowBlocks(opts *bind.CallOpts, operator common.Address) (*big.Int, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)

	WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error)
}

// ContractDelegationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractDelegationManagerTransacts interface {
	CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error)

	CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error)

	DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error)

	DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error)

	DelegateToBySignature(opts *bind.TransactOpts, staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error)

	IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	MigrateQueuedWithdrawals(opts *bind.TransactOpts, withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error)

	ModifyOperatorDetails(opts *bind.TransactOpts, newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	QueueWithdrawals(opts *bind.TransactOpts, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error)

	RegisterAsOperator(opts *bind.TransactOpts, registeringOperatorDetails IDelegationManagerOperatorDetails, metadataURI string) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error)

	SetStakeRegistry(opts *bind.TransactOpts, _stakeRegistry common.Address) (*types.Transaction, error)

	SetWithdrawalDelayBlocks(opts *bind.TransactOpts, newWithdrawalDelayBlocks *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperatorMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)
}

// ContractDelegationManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractDelegationManagerFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractDelegationManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractDelegationManagerInitialized, error)

	FilterOperatorDetailsModified(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorDetailsModifiedIterator, error)
	WatchOperatorDetailsModified(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorDetailsModified, operator []common.Address) (event.Subscription, error)
	ParseOperatorDetailsModified(log types.Log) (*ContractDelegationManagerOperatorDetailsModified, error)

	FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorMetadataURIUpdatedIterator, error)
	WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error)
	ParseOperatorMetadataURIUpdated(log types.Log) (*ContractDelegationManagerOperatorMetadataURIUpdated, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorRegistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractDelegationManagerOperatorRegistered, error)

	FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesDecreasedIterator, error)
	WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesDecreased(log types.Log) (*ContractDelegationManagerOperatorSharesDecreased, error)

	FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesIncreasedIterator, error)
	WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesIncreased(log types.Log) (*ContractDelegationManagerOperatorSharesIncreased, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractDelegationManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractDelegationManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractDelegationManagerPaused, error)

	FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractDelegationManagerPauserRegistrySetIterator, error)
	WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPauserRegistrySet) (event.Subscription, error)
	ParsePauserRegistrySet(log types.Log) (*ContractDelegationManagerPauserRegistrySet, error)

	FilterStakeRegistrySet(opts *bind.FilterOpts) (*ContractDelegationManagerStakeRegistrySetIterator, error)
	WatchStakeRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakeRegistrySet) (event.Subscription, error)
	ParseStakeRegistrySet(log types.Log) (*ContractDelegationManagerStakeRegistrySet, error)

	FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerDelegatedIterator, error)
	WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerDelegated(log types.Log) (*ContractDelegationManagerStakerDelegated, error)

	FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerForceUndelegatedIterator, error)
	WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerForceUndelegated(log types.Log) (*ContractDelegationManagerStakerForceUndelegated, error)

	FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerUndelegatedIterator, error)
	WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerUndelegated(log types.Log) (*ContractDelegationManagerStakerUndelegated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractDelegationManagerUnpaused, error)

	FilterWithdrawalCompleted(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalCompletedIterator, error)
	WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalCompleted) (event.Subscription, error)
	ParseWithdrawalCompleted(log types.Log) (*ContractDelegationManagerWithdrawalCompleted, error)

	FilterWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalDelayBlocksSetIterator, error)
	WatchWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalDelayBlocksSet) (event.Subscription, error)
	ParseWithdrawalDelayBlocksSet(log types.Log) (*ContractDelegationManagerWithdrawalDelayBlocksSet, error)

	FilterWithdrawalMigrated(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalMigratedIterator, error)
	WatchWithdrawalMigrated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalMigrated) (event.Subscription, error)
	ParseWithdrawalMigrated(log types.Log) (*ContractDelegationManagerWithdrawalMigrated, error)

	FilterWithdrawalQueued(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalQueuedIterator, error)
	WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalQueued) (event.Subscription, error)
	ParseWithdrawalQueued(log types.Log) (*ContractDelegationManagerWithdrawalQueued, error)
}

// ContractDelegationManager is an auto generated Go binding around an Ethereum contract.
type ContractDelegationManager struct {
	ContractDelegationManagerCaller     // Read-only binding to the contract
	ContractDelegationManagerTransactor // Write-only binding to the contract
	ContractDelegationManagerFilterer   // Log filterer for contract events
}

// ContractDelegationManager implements the ContractDelegationManagerMethods interface.
var _ ContractDelegationManagerMethods = (*ContractDelegationManager)(nil)

// ContractDelegationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractDelegationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerCaller implements the ContractDelegationManagerCalls interface.
var _ ContractDelegationManagerCalls = (*ContractDelegationManagerCaller)(nil)

// ContractDelegationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractDelegationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerTransactor implements the ContractDelegationManagerTransacts interface.
var _ ContractDelegationManagerTransacts = (*ContractDelegationManagerTransactor)(nil)

// ContractDelegationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractDelegationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerFilterer implements the ContractDelegationManagerFilters interface.
var _ ContractDelegationManagerFilters = (*ContractDelegationManagerFilterer)(nil)

// ContractDelegationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractDelegationManagerSession struct {
	Contract     *ContractDelegationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractDelegationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractDelegationManagerCallerSession struct {
	Contract *ContractDelegationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractDelegationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractDelegationManagerTransactorSession struct {
	Contract     *ContractDelegationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractDelegationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractDelegationManagerRaw struct {
	Contract *ContractDelegationManager // Generic contract binding to access the raw methods on
}

// ContractDelegationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractDelegationManagerCallerRaw struct {
	Contract *ContractDelegationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractDelegationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractDelegationManagerTransactorRaw struct {
	Contract *ContractDelegationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractDelegationManager creates a new instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManager(address common.Address, backend bind.ContractBackend) (*ContractDelegationManager, error) {
	contract, err := bindContractDelegationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManager{ContractDelegationManagerCaller: ContractDelegationManagerCaller{contract: contract}, ContractDelegationManagerTransactor: ContractDelegationManagerTransactor{contract: contract}, ContractDelegationManagerFilterer: ContractDelegationManagerFilterer{contract: contract}}, nil
}

// NewContractDelegationManagerCaller creates a new read-only instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractDelegationManagerCaller, error) {
	contract, err := bindContractDelegationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerCaller{contract: contract}, nil
}

// NewContractDelegationManagerTransactor creates a new write-only instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractDelegationManagerTransactor, error) {
	contract, err := bindContractDelegationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerTransactor{contract: contract}, nil
}

// NewContractDelegationManagerFilterer creates a new log filterer instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractDelegationManagerFilterer, error) {
	contract, err := bindContractDelegationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerFilterer{contract: contract}, nil
}

// bindContractDelegationManager binds a generic wrapper to an already deployed contract.
func bindContractDelegationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractDelegationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDelegationManager.Contract.ContractDelegationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ContractDelegationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ContractDelegationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDelegationManager *ContractDelegationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDelegationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDelegationManager *ContractDelegationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDelegationManager *ContractDelegationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "DELEGATION_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DOMAINTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DOMAINTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// MAXSTAKEROPTOUTWINDOWBLOCKS is a free data retrieval call binding the contract method 0x4fc40b61.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW_BLOCKS() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) MAXSTAKEROPTOUTWINDOWBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "MAX_STAKER_OPT_OUT_WINDOW_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKEROPTOUTWINDOWBLOCKS is a free data retrieval call binding the contract method 0x4fc40b61.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW_BLOCKS() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) MAXSTAKEROPTOUTWINDOWBLOCKS() (*big.Int, error) {
	return _ContractDelegationManager.Contract.MAXSTAKEROPTOUTWINDOWBLOCKS(&_ContractDelegationManager.CallOpts)
}

// MAXSTAKEROPTOUTWINDOWBLOCKS is a free data retrieval call binding the contract method 0x4fc40b61.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW_BLOCKS() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) MAXSTAKEROPTOUTWINDOWBLOCKS() (*big.Int, error) {
	return _ContractDelegationManager.Contract.MAXSTAKEROPTOUTWINDOWBLOCKS(&_ContractDelegationManager.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _ContractDelegationManager.Contract.MAXWITHDRAWALDELAYBLOCKS(&_ContractDelegationManager.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _ContractDelegationManager.Contract.MAXWITHDRAWALDELAYBLOCKS(&_ContractDelegationManager.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) STAKERDELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "STAKER_DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.STAKERDELEGATIONTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.STAKERDELEGATIONTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractDelegationManager.Contract.BeaconChainETHStrategy(&_ContractDelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractDelegationManager.Contract.BeaconChainETHStrategy(&_ContractDelegationManager.CallOpts)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateCurrentStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateCurrentStakerDelegationDigestHash", staker, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateCurrentStakerDelegationDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, expiry)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateCurrentStakerDelegationDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, _delegationApprover, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateStakerDelegationDigestHash", staker, _stakerNonce, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateStakerDelegationDigestHash(&_ContractDelegationManager.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateStakerDelegationDigestHash(&_ContractDelegationManager.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateWithdrawalRoot", withdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerWithdrawal) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateWithdrawalRoot(&_ContractDelegationManager.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerWithdrawal) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateWithdrawalRoot(&_ContractDelegationManager.CallOpts, withdrawal)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.CumulativeWithdrawalsQueued(&_ContractDelegationManager.CallOpts, arg0)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.CumulativeWithdrawalsQueued(&_ContractDelegationManager.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegatedTo(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegatedTo", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegatedTo(&_ContractDelegationManager.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegatedTo(&_ContractDelegationManager.CallOpts, arg0)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegationApprover", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegationApprover(&_ContractDelegationManager.CallOpts, operator)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegationApprover(&_ContractDelegationManager.CallOpts, operator)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegationApproverSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.DelegationApproverSaltIsSpent(&_ContractDelegationManager.CallOpts, arg0, arg1)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.DelegationApproverSaltIsSpent(&_ContractDelegationManager.CallOpts, arg0, arg1)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) DomainSeparator() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DomainSeparator(&_ContractDelegationManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DomainSeparator(&_ContractDelegationManager.CallOpts)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) EarningsReceiver(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "earningsReceiver", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.EarningsReceiver(&_ContractDelegationManager.CallOpts, operator)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.EarningsReceiver(&_ContractDelegationManager.CallOpts, operator)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) EigenPodManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.EigenPodManager(&_ContractDelegationManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.EigenPodManager(&_ContractDelegationManager.CallOpts)
}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetDelegatableShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getDelegatableShares", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetDelegatableShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractDelegationManager.Contract.GetDelegatableShares(&_ContractDelegationManager.CallOpts, staker)
}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetDelegatableShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractDelegationManager.Contract.GetDelegatableShares(&_ContractDelegationManager.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "isDelegated", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) IsDelegated(staker common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsDelegated(&_ContractDelegationManager.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) IsDelegated(staker common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsDelegated(&_ContractDelegationManager.CallOpts, staker)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) IsOperator(operator common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsOperator(&_ContractDelegationManager.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsOperator(&_ContractDelegationManager.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_ContractDelegationManager *ContractDelegationManagerCaller) OperatorDetails(opts *bind.CallOpts, operator common.Address) (IDelegationManagerOperatorDetails, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "operatorDetails", operator)

	if err != nil {
		return *new(IDelegationManagerOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationManagerOperatorDetails)).(*IDelegationManagerOperatorDetails)

	return out0, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_ContractDelegationManager *ContractDelegationManagerSession) OperatorDetails(operator common.Address) (IDelegationManagerOperatorDetails, error) {
	return _ContractDelegationManager.Contract.OperatorDetails(&_ContractDelegationManager.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) OperatorDetails(operator common.Address) (IDelegationManagerOperatorDetails, error) {
	return _ContractDelegationManager.Contract.OperatorDetails(&_ContractDelegationManager.CallOpts, operator)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) OperatorShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "operatorShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.OperatorShares(&_ContractDelegationManager.CallOpts, arg0, arg1)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.OperatorShares(&_ContractDelegationManager.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) Owner() (common.Address, error) {
	return _ContractDelegationManager.Contract.Owner(&_ContractDelegationManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Owner() (common.Address, error) {
	return _ContractDelegationManager.Contract.Owner(&_ContractDelegationManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) Paused(index uint8) (bool, error) {
	return _ContractDelegationManager.Contract.Paused(&_ContractDelegationManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractDelegationManager.Contract.Paused(&_ContractDelegationManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) Paused0() (*big.Int, error) {
	return _ContractDelegationManager.Contract.Paused0(&_ContractDelegationManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractDelegationManager.Contract.Paused0(&_ContractDelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.PauserRegistry(&_ContractDelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.PauserRegistry(&_ContractDelegationManager.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.PendingWithdrawals(&_ContractDelegationManager.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.PendingWithdrawals(&_ContractDelegationManager.CallOpts, arg0)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) Slasher() (common.Address, error) {
	return _ContractDelegationManager.Contract.Slasher(&_ContractDelegationManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Slasher() (common.Address, error) {
	return _ContractDelegationManager.Contract.Slasher(&_ContractDelegationManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.StakeRegistry(&_ContractDelegationManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.StakeRegistry(&_ContractDelegationManager.CallOpts)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) StakerNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "stakerNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.StakerNonce(&_ContractDelegationManager.CallOpts, arg0)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.StakerNonce(&_ContractDelegationManager.CallOpts, arg0)
}

// StakerOptOutWindowBlocks is a free data retrieval call binding the contract method 0x16928365.
//
// Solidity: function stakerOptOutWindowBlocks(address operator) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) StakerOptOutWindowBlocks(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "stakerOptOutWindowBlocks", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerOptOutWindowBlocks is a free data retrieval call binding the contract method 0x16928365.
//
// Solidity: function stakerOptOutWindowBlocks(address operator) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) StakerOptOutWindowBlocks(operator common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.StakerOptOutWindowBlocks(&_ContractDelegationManager.CallOpts, operator)
}

// StakerOptOutWindowBlocks is a free data retrieval call binding the contract method 0x16928365.
//
// Solidity: function stakerOptOutWindowBlocks(address operator) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) StakerOptOutWindowBlocks(operator common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.StakerOptOutWindowBlocks(&_ContractDelegationManager.CallOpts, operator)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) StrategyManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.StrategyManager(&_ContractDelegationManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) StrategyManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.StrategyManager(&_ContractDelegationManager.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "withdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _ContractDelegationManager.Contract.WithdrawalDelayBlocks(&_ContractDelegationManager.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _ContractDelegationManager.Contract.WithdrawalDelayBlocks(&_ContractDelegationManager.CallOpts)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x60d7faed.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "completeQueuedWithdrawal", withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x60d7faed.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawal(&_ContractDelegationManager.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x60d7faed.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawal(&_ContractDelegationManager.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "completeQueuedWithdrawals", withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "decreaseDelegatedShares", staker, strategy, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) DecreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DecreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) DecreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DecreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, shares)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateTo(&_ContractDelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateTo(&_ContractDelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) DelegateToBySignature(opts *bind.TransactOpts, staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "delegateToBySignature", staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateToBySignature(&_ContractDelegationManager.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateToBySignature(&_ContractDelegationManager.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "increaseDelegatedShares", staker, strategy, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.IncreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.IncreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, shares)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Initialize(&_ContractDelegationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Initialize(&_ContractDelegationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// MigrateQueuedWithdrawals is a paid mutator transaction binding the contract method 0x5cfe8d2c.
//
// Solidity: function migrateQueuedWithdrawals((address[],uint256[],address,(address,uint96),uint32,address)[] withdrawalsToMigrate) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) MigrateQueuedWithdrawals(opts *bind.TransactOpts, withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "migrateQueuedWithdrawals", withdrawalsToMigrate)
}

// MigrateQueuedWithdrawals is a paid mutator transaction binding the contract method 0x5cfe8d2c.
//
// Solidity: function migrateQueuedWithdrawals((address[],uint256[],address,(address,uint96),uint32,address)[] withdrawalsToMigrate) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) MigrateQueuedWithdrawals(withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.MigrateQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, withdrawalsToMigrate)
}

// MigrateQueuedWithdrawals is a paid mutator transaction binding the contract method 0x5cfe8d2c.
//
// Solidity: function migrateQueuedWithdrawals((address[],uint256[],address,(address,uint96),uint32,address)[] withdrawalsToMigrate) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) MigrateQueuedWithdrawals(withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.MigrateQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, withdrawalsToMigrate)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "modifyOperatorDetails", newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) ModifyOperatorDetails(newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ModifyOperatorDetails(&_ContractDelegationManager.TransactOpts, newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) ModifyOperatorDetails(newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ModifyOperatorDetails(&_ContractDelegationManager.TransactOpts, newOperatorDetails)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Pause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Pause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.PauseAll(&_ContractDelegationManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.PauseAll(&_ContractDelegationManager.TransactOpts)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerTransactor) QueueWithdrawals(opts *bind.TransactOpts, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "queueWithdrawals", queuedWithdrawalParams)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerSession) QueueWithdrawals(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.QueueWithdrawals(&_ContractDelegationManager.TransactOpts, queuedWithdrawalParams)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) QueueWithdrawals(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.QueueWithdrawals(&_ContractDelegationManager.TransactOpts, queuedWithdrawalParams)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, registeringOperatorDetails IDelegationManagerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "registerAsOperator", registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) RegisterAsOperator(registeringOperatorDetails IDelegationManagerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RegisterAsOperator(&_ContractDelegationManager.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) RegisterAsOperator(registeringOperatorDetails IDelegationManagerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RegisterAsOperator(&_ContractDelegationManager.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RenounceOwnership(&_ContractDelegationManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RenounceOwnership(&_ContractDelegationManager.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SetPauserRegistry(&_ContractDelegationManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SetPauserRegistry(&_ContractDelegationManager.TransactOpts, newPauserRegistry)
}

// SetStakeRegistry is a paid mutator transaction binding the contract method 0xe3b05f2f.
//
// Solidity: function setStakeRegistry(address _stakeRegistry) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) SetStakeRegistry(opts *bind.TransactOpts, _stakeRegistry common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "setStakeRegistry", _stakeRegistry)
}

// SetStakeRegistry is a paid mutator transaction binding the contract method 0xe3b05f2f.
//
// Solidity: function setStakeRegistry(address _stakeRegistry) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) SetStakeRegistry(_stakeRegistry common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SetStakeRegistry(&_ContractDelegationManager.TransactOpts, _stakeRegistry)
}

// SetStakeRegistry is a paid mutator transaction binding the contract method 0xe3b05f2f.
//
// Solidity: function setStakeRegistry(address _stakeRegistry) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) SetStakeRegistry(_stakeRegistry common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SetStakeRegistry(&_ContractDelegationManager.TransactOpts, _stakeRegistry)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newWithdrawalDelayBlocks) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) SetWithdrawalDelayBlocks(opts *bind.TransactOpts, newWithdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "setWithdrawalDelayBlocks", newWithdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newWithdrawalDelayBlocks) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) SetWithdrawalDelayBlocks(newWithdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SetWithdrawalDelayBlocks(&_ContractDelegationManager.TransactOpts, newWithdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newWithdrawalDelayBlocks) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) SetWithdrawalDelayBlocks(newWithdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SetWithdrawalDelayBlocks(&_ContractDelegationManager.TransactOpts, newWithdrawalDelayBlocks)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.TransferOwnership(&_ContractDelegationManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.TransferOwnership(&_ContractDelegationManager.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "undelegate", staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Undelegate(&_ContractDelegationManager.TransactOpts, staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Undelegate(&_ContractDelegationManager.TransactOpts, staker)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Unpause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Unpause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "updateOperatorMetadataURI", metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.UpdateOperatorMetadataURI(&_ContractDelegationManager.TransactOpts, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.UpdateOperatorMetadataURI(&_ContractDelegationManager.TransactOpts, metadataURI)
}

// ContractDelegationManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractDelegationManager contract.
type ContractDelegationManagerInitializedIterator struct {
	Event *ContractDelegationManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerInitialized)
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
		it.Event = new(ContractDelegationManagerInitialized)
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
func (it *ContractDelegationManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerInitialized represents a Initialized event raised by the ContractDelegationManager contract.
type ContractDelegationManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractDelegationManagerInitializedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerInitializedIterator{contract: _ContractDelegationManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerInitialized)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseInitialized(log types.Log) (*ContractDelegationManagerInitialized, error) {
	event := new(ContractDelegationManagerInitialized)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorDetailsModifiedIterator is returned from FilterOperatorDetailsModified and is used to iterate over the raw logs and unpacked data for OperatorDetailsModified events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorDetailsModifiedIterator struct {
	Event *ContractDelegationManagerOperatorDetailsModified // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorDetailsModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorDetailsModified)
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
		it.Event = new(ContractDelegationManagerOperatorDetailsModified)
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
func (it *ContractDelegationManagerOperatorDetailsModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorDetailsModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorDetailsModified represents a OperatorDetailsModified event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorDetailsModified struct {
	Operator           common.Address
	NewOperatorDetails IDelegationManagerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorDetailsModified is a free log retrieval operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorDetailsModified(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorDetailsModifiedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorDetailsModifiedIterator{contract: _ContractDelegationManager.contract, event: "OperatorDetailsModified", logs: logs, sub: sub}, nil
}

// WatchOperatorDetailsModified is a free log subscription operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorDetailsModified(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorDetailsModified, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorDetailsModified)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
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

// ParseOperatorDetailsModified is a log parse operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorDetailsModified(log types.Log) (*ContractDelegationManagerOperatorDetailsModified, error) {
	event := new(ContractDelegationManagerOperatorDetailsModified)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorMetadataURIUpdatedIterator is returned from FilterOperatorMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for OperatorMetadataURIUpdated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorMetadataURIUpdatedIterator struct {
	Event *ContractDelegationManagerOperatorMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorMetadataURIUpdated)
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
		it.Event = new(ContractDelegationManagerOperatorMetadataURIUpdated)
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
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorMetadataURIUpdated represents a OperatorMetadataURIUpdated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorMetadataURIUpdated struct {
	Operator    common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataURIUpdated is a free log retrieval operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorMetadataURIUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorMetadataURIUpdatedIterator{contract: _ContractDelegationManager.contract, event: "OperatorMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataURIUpdated is a free log subscription operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorMetadataURIUpdated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
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

// ParseOperatorMetadataURIUpdated is a log parse operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorMetadataURIUpdated(log types.Log) (*ContractDelegationManagerOperatorMetadataURIUpdated, error) {
	event := new(ContractDelegationManagerOperatorMetadataURIUpdated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorRegisteredIterator struct {
	Event *ContractDelegationManagerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorRegistered)
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
		it.Event = new(ContractDelegationManagerOperatorRegistered)
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
func (it *ContractDelegationManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorRegistered represents a OperatorRegistered event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorRegistered struct {
	Operator        common.Address
	OperatorDetails IDelegationManagerOperatorDetails
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorRegisteredIterator{contract: _ContractDelegationManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorRegistered)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorRegistered(log types.Log) (*ContractDelegationManagerOperatorRegistered, error) {
	event := new(ContractDelegationManagerOperatorRegistered)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesDecreasedIterator is returned from FilterOperatorSharesDecreased and is used to iterate over the raw logs and unpacked data for OperatorSharesDecreased events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesDecreasedIterator struct {
	Event *ContractDelegationManagerOperatorSharesDecreased // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesDecreased)
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
		it.Event = new(ContractDelegationManagerOperatorSharesDecreased)
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
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesDecreased represents a OperatorSharesDecreased event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesDecreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesDecreased is a free log retrieval operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesDecreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesDecreasedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesDecreased is a free log subscription operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesDecreased)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
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

// ParseOperatorSharesDecreased is a log parse operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesDecreased(log types.Log) (*ContractDelegationManagerOperatorSharesDecreased, error) {
	event := new(ContractDelegationManagerOperatorSharesDecreased)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesIncreasedIterator is returned from FilterOperatorSharesIncreased and is used to iterate over the raw logs and unpacked data for OperatorSharesIncreased events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesIncreasedIterator struct {
	Event *ContractDelegationManagerOperatorSharesIncreased // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesIncreased)
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
		it.Event = new(ContractDelegationManagerOperatorSharesIncreased)
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
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesIncreased represents a OperatorSharesIncreased event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesIncreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesIncreased is a free log retrieval operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesIncreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesIncreasedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesIncreased is a free log subscription operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesIncreased)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
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

// ParseOperatorSharesIncreased is a log parse operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesIncreased(log types.Log) (*ContractDelegationManagerOperatorSharesIncreased, error) {
	event := new(ContractDelegationManagerOperatorSharesIncreased)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOwnershipTransferredIterator struct {
	Event *ContractDelegationManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOwnershipTransferred)
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
		it.Event = new(ContractDelegationManagerOwnershipTransferred)
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
func (it *ContractDelegationManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractDelegationManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOwnershipTransferredIterator{contract: _ContractDelegationManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOwnershipTransferred)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractDelegationManagerOwnershipTransferred, error) {
	event := new(ContractDelegationManagerOwnershipTransferred)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractDelegationManager contract.
type ContractDelegationManagerPausedIterator struct {
	Event *ContractDelegationManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerPaused)
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
		it.Event = new(ContractDelegationManagerPaused)
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
func (it *ContractDelegationManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerPaused represents a Paused event raised by the ContractDelegationManager contract.
type ContractDelegationManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerPausedIterator{contract: _ContractDelegationManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerPaused)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParsePaused(log types.Log) (*ContractDelegationManagerPaused, error) {
	event := new(ContractDelegationManagerPaused)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractDelegationManager contract.
type ContractDelegationManagerPauserRegistrySetIterator struct {
	Event *ContractDelegationManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerPauserRegistrySet)
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
		it.Event = new(ContractDelegationManagerPauserRegistrySet)
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
func (it *ContractDelegationManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractDelegationManager contract.
type ContractDelegationManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractDelegationManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerPauserRegistrySetIterator{contract: _ContractDelegationManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerPauserRegistrySet)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractDelegationManagerPauserRegistrySet, error) {
	event := new(ContractDelegationManagerPauserRegistrySet)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakeRegistrySetIterator is returned from FilterStakeRegistrySet and is used to iterate over the raw logs and unpacked data for StakeRegistrySet events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakeRegistrySetIterator struct {
	Event *ContractDelegationManagerStakeRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerStakeRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakeRegistrySet)
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
		it.Event = new(ContractDelegationManagerStakeRegistrySet)
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
func (it *ContractDelegationManagerStakeRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakeRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakeRegistrySet represents a StakeRegistrySet event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakeRegistrySet struct {
	StakeRegistry common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeRegistrySet is a free log retrieval operation binding the contract event 0xce6d874069bceda1867eca9c60636bbf262e15213041658273d803a2b609a51f.
//
// Solidity: event StakeRegistrySet(address stakeRegistry)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakeRegistrySet(opts *bind.FilterOpts) (*ContractDelegationManagerStakeRegistrySetIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakeRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakeRegistrySetIterator{contract: _ContractDelegationManager.contract, event: "StakeRegistrySet", logs: logs, sub: sub}, nil
}

// WatchStakeRegistrySet is a free log subscription operation binding the contract event 0xce6d874069bceda1867eca9c60636bbf262e15213041658273d803a2b609a51f.
//
// Solidity: event StakeRegistrySet(address stakeRegistry)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakeRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakeRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakeRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakeRegistrySet)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakeRegistrySet", log); err != nil {
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

// ParseStakeRegistrySet is a log parse operation binding the contract event 0xce6d874069bceda1867eca9c60636bbf262e15213041658273d803a2b609a51f.
//
// Solidity: event StakeRegistrySet(address stakeRegistry)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakeRegistrySet(log types.Log) (*ContractDelegationManagerStakeRegistrySet, error) {
	event := new(ContractDelegationManagerStakeRegistrySet)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakeRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerDelegatedIterator is returned from FilterStakerDelegated and is used to iterate over the raw logs and unpacked data for StakerDelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerDelegatedIterator struct {
	Event *ContractDelegationManagerStakerDelegated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerStakerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerDelegated)
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
		it.Event = new(ContractDelegationManagerStakerDelegated)
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
func (it *ContractDelegationManagerStakerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerDelegated represents a StakerDelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerDelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerDelegated is a free log retrieval operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerDelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerDelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerDelegated", logs: logs, sub: sub}, nil
}

// WatchStakerDelegated is a free log subscription operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerDelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
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

// ParseStakerDelegated is a log parse operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerDelegated(log types.Log) (*ContractDelegationManagerStakerDelegated, error) {
	event := new(ContractDelegationManagerStakerDelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerForceUndelegatedIterator is returned from FilterStakerForceUndelegated and is used to iterate over the raw logs and unpacked data for StakerForceUndelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerForceUndelegatedIterator struct {
	Event *ContractDelegationManagerStakerForceUndelegated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerForceUndelegated)
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
		it.Event = new(ContractDelegationManagerStakerForceUndelegated)
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
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerForceUndelegated represents a StakerForceUndelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerForceUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerForceUndelegated is a free log retrieval operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerForceUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerForceUndelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerForceUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerForceUndelegated is a free log subscription operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerForceUndelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
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

// ParseStakerForceUndelegated is a log parse operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerForceUndelegated(log types.Log) (*ContractDelegationManagerStakerForceUndelegated, error) {
	event := new(ContractDelegationManagerStakerForceUndelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerUndelegatedIterator is returned from FilterStakerUndelegated and is used to iterate over the raw logs and unpacked data for StakerUndelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerUndelegatedIterator struct {
	Event *ContractDelegationManagerStakerUndelegated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerStakerUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerUndelegated)
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
		it.Event = new(ContractDelegationManagerStakerUndelegated)
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
func (it *ContractDelegationManagerStakerUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerUndelegated represents a StakerUndelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerUndelegated is a free log retrieval operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerUndelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerUndelegated is a free log subscription operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerUndelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
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

// ParseStakerUndelegated is a log parse operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerUndelegated(log types.Log) (*ContractDelegationManagerStakerUndelegated, error) {
	event := new(ContractDelegationManagerStakerUndelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractDelegationManager contract.
type ContractDelegationManagerUnpausedIterator struct {
	Event *ContractDelegationManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerUnpaused)
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
		it.Event = new(ContractDelegationManagerUnpaused)
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
func (it *ContractDelegationManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerUnpaused represents a Unpaused event raised by the ContractDelegationManager contract.
type ContractDelegationManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerUnpausedIterator{contract: _ContractDelegationManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerUnpaused)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseUnpaused(log types.Log) (*ContractDelegationManagerUnpaused, error) {
	event := new(ContractDelegationManagerUnpaused)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalCompletedIterator struct {
	Event *ContractDelegationManagerWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerWithdrawalCompleted)
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
		it.Event = new(ContractDelegationManagerWithdrawalCompleted)
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
func (it *ContractDelegationManagerWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerWithdrawalCompleted represents a WithdrawalCompleted event raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalCompletedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerWithdrawalCompletedIterator{contract: _ContractDelegationManager.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerWithdrawalCompleted)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseWithdrawalCompleted(log types.Log) (*ContractDelegationManagerWithdrawalCompleted, error) {
	event := new(ContractDelegationManagerWithdrawalCompleted)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerWithdrawalDelayBlocksSetIterator is returned from FilterWithdrawalDelayBlocksSet and is used to iterate over the raw logs and unpacked data for WithdrawalDelayBlocksSet events raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalDelayBlocksSetIterator struct {
	Event *ContractDelegationManagerWithdrawalDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerWithdrawalDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerWithdrawalDelayBlocksSet)
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
		it.Event = new(ContractDelegationManagerWithdrawalDelayBlocksSet)
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
func (it *ContractDelegationManagerWithdrawalDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerWithdrawalDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerWithdrawalDelayBlocksSet represents a WithdrawalDelayBlocksSet event raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalDelayBlocksSet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayBlocksSet is a free log retrieval operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalDelayBlocksSetIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerWithdrawalDelayBlocksSetIterator{contract: _ContractDelegationManager.contract, event: "WithdrawalDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayBlocksSet is a free log subscription operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerWithdrawalDelayBlocksSet)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
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

// ParseWithdrawalDelayBlocksSet is a log parse operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseWithdrawalDelayBlocksSet(log types.Log) (*ContractDelegationManagerWithdrawalDelayBlocksSet, error) {
	event := new(ContractDelegationManagerWithdrawalDelayBlocksSet)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerWithdrawalMigratedIterator is returned from FilterWithdrawalMigrated and is used to iterate over the raw logs and unpacked data for WithdrawalMigrated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalMigratedIterator struct {
	Event *ContractDelegationManagerWithdrawalMigrated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerWithdrawalMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerWithdrawalMigrated)
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
		it.Event = new(ContractDelegationManagerWithdrawalMigrated)
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
func (it *ContractDelegationManagerWithdrawalMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerWithdrawalMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerWithdrawalMigrated represents a WithdrawalMigrated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalMigrated struct {
	OldWithdrawalRoot [32]byte
	NewWithdrawalRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalMigrated is a free log retrieval operation binding the contract event 0xdc00758b65eef71dc3780c04ebe36cab6bdb266c3a698187e29e0f0dca012630.
//
// Solidity: event WithdrawalMigrated(bytes32 oldWithdrawalRoot, bytes32 newWithdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterWithdrawalMigrated(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalMigratedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "WithdrawalMigrated")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerWithdrawalMigratedIterator{contract: _ContractDelegationManager.contract, event: "WithdrawalMigrated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalMigrated is a free log subscription operation binding the contract event 0xdc00758b65eef71dc3780c04ebe36cab6bdb266c3a698187e29e0f0dca012630.
//
// Solidity: event WithdrawalMigrated(bytes32 oldWithdrawalRoot, bytes32 newWithdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchWithdrawalMigrated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalMigrated) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "WithdrawalMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerWithdrawalMigrated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalMigrated", log); err != nil {
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

// ParseWithdrawalMigrated is a log parse operation binding the contract event 0xdc00758b65eef71dc3780c04ebe36cab6bdb266c3a698187e29e0f0dca012630.
//
// Solidity: event WithdrawalMigrated(bytes32 oldWithdrawalRoot, bytes32 newWithdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseWithdrawalMigrated(log types.Log) (*ContractDelegationManagerWithdrawalMigrated, error) {
	event := new(ContractDelegationManagerWithdrawalMigrated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalQueuedIterator struct {
	Event *ContractDelegationManagerWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerWithdrawalQueued)
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
		it.Event = new(ContractDelegationManagerWithdrawalQueued)
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
func (it *ContractDelegationManagerWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerWithdrawalQueued represents a WithdrawalQueued event raised by the ContractDelegationManager contract.
type ContractDelegationManagerWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     IDelegationManagerWithdrawal
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*ContractDelegationManagerWithdrawalQueuedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerWithdrawalQueuedIterator{contract: _ContractDelegationManager.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerWithdrawalQueued)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
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

// ParseWithdrawalQueued is a log parse operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseWithdrawalQueued(log types.Log) (*ContractDelegationManagerWithdrawalQueued, error) {
	event := new(ContractDelegationManagerWithdrawalQueued)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
