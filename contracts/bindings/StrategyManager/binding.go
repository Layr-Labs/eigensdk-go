// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractStrategyManager

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

// IStrategyManagerQueuedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerQueuedWithdrawal struct {
	Strategies           []common.Address
	Shares               []*big.Int
	Depositor            common.Address
	WithdrawerAndNonce   IStrategyManagerWithdrawerAndNonce
	WithdrawalStartBlock uint32
	DelegatedAddress     common.Address
}

// IStrategyManagerWithdrawerAndNonce is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerWithdrawerAndNonce struct {
	Withdrawer common.Address
	Nonce      *big.Int
}

// ContractStrategyManagerMetaData contains all meta data concerning the ContractStrategyManager contract.
var ContractStrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegation\",\"type\":\"address\"},{\"internalType\":\"contractIEigenPodManager\",\"name\":\"_eigenPodManager\",\"type\":\"address\"},{\"internalType\":\"contractISlasher\",\"name\":\"_slasher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"ShareWithdrawalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyAddedToDepositWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyRemovedFromDepositWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"StrategyWhitelisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"WithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"WithdrawalDelayBlocksSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"WithdrawalQueued\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\"}],\"name\":\"addStrategiesToDepositWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconChainETHStrategy\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"}],\"internalType\":\"structIStrategyManager.WithdrawerAndNonce\",\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"}],\"internalType\":\"structIStrategyManager.QueuedWithdrawal\",\"name\":\"queuedWithdrawal\",\"type\":\"tuple\"}],\"name\":\"calculateWithdrawalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"}],\"internalType\":\"structIStrategyManager.WithdrawerAndNonce\",\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"}],\"internalType\":\"structIStrategyManager.QueuedWithdrawal\",\"name\":\"queuedWithdrawal\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAsTokens\",\"type\":\"bool\"}],\"name\":\"completeQueuedWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"}],\"internalType\":\"structIStrategyManager.WithdrawerAndNonce\",\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"}],\"internalType\":\"structIStrategyManager.QueuedWithdrawal[]\",\"name\":\"queuedWithdrawals\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20[][]\",\"name\":\"tokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"receiveAsTokens\",\"type\":\"bool[]\"}],\"name\":\"completeQueuedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositBeaconChainETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositIntoStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"depositIntoStrategyWithSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"forceTotalWithdrawal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"getDeposits\",\"outputs\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialStrategyWhitelister\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialPausedStatus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numWithdrawalsQueued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"strategyIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"undelegateIfPossible\",\"type\":\"bool\"}],\"name\":\"queueWithdrawal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"beaconChainETHStrategyIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"sharesDelta\",\"type\":\"int256\"}],\"name\":\"recordBeaconChainETHBalanceUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\"}],\"name\":\"removeStrategiesFromDepositWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStrategyWhitelister\",\"type\":\"address\"}],\"name\":\"setStrategyWhitelister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalDelayBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"}],\"internalType\":\"structIStrategyManager.WithdrawerAndNonce\",\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"}],\"internalType\":\"structIStrategyManager.QueuedWithdrawal\",\"name\":\"queuedWithdrawal\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indicesToSkip\",\"type\":\"uint256[]\"}],\"name\":\"slashQueuedWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slashedAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"strategyIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shareAmounts\",\"type\":\"uint256[]\"}],\"name\":\"slashShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerStrategyList\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakerStrategyListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerStrategyShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"strategyIsWhitelistedForDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyWhitelister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawalRootPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005f4c38038062005f4c833981016040819052620000359162000140565b6001600160a01b0380841660805280831660a052811660c0526200005862000065565b50504660e0525062000194565b600054610100900460ff1615620000d25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000125576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013d57600080fd5b50565b6000806000606084860312156200015657600080fd5b8351620001638162000127565b6020850151909350620001768162000127565b6040850151909250620001898162000127565b809150509250925092565b60805160a05160c05160e051615d07620002456000396000611fce015260008181610595015281816121dd015281816125a30152613c1b015260008181610324015281816108310152818161102e01528181612f0c015281816133fd0152818161401c0152614277015260008181610634015281816109b5015281816113dc01528181612c29015281816130e30152818161339c015281816136e50152818161385201526138fb0152615d076000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c80638da5cb5b1161015c578063b5d8b5b8116100ce578063e7a050aa11610087578063e7a050aa14610656578063f123991e14610669578063f2fde38b1461067c578063f3be65d31461068f578063f698da25146106a2578063fabc1cbc146106aa57600080fd5b8063b5d8b5b8146105ca578063c3c6b3a9146105dd578063c665670214610600578063ca661c0414610613578063cbc2bd621461061c578063df5cf7231461062f57600080fd5b80639f00fa24116101205780639f00fa2414610544578063a1ca780b14610557578063a6b63eb81461056a578063a782d9451461057d578063b134427114610590578063b43b514b146105b757600080fd5b80638da5cb5b146104dc5780639104c319146104ed57806392ab89bb1461050857806394f649dd14610510578063967fc0d21461053157600080fd5b8063595c6a67116101f5578063663c1de4116101b9578063663c1de41461042a578063715018a61461044d5780637a7e0d92146104555780637ecebe0014610480578063886f1195146104a05780638b8aac3c146104b357600080fd5b8063595c6a67146103c15780635ac86ab7146103c95780635c975abb146103fc5780635de08ff2146104045780635e97f8bd1461041757600080fd5b806343c090611161024757806343c090611461030c5780634665bcda1461031f57806348825e941461035e5780634d50f9a41461038557806350f73e7c1461039857806356631028146103a157600080fd5b806306f1f6841461028457806310d67a2f14610299578063136439dd146102ac57806320606b70146102bf57806332e89ace146102f9575b600080fd5b610297610292366004614d8b565b6106bd565b005b6102976102a7366004614e72565b610a37565b6102976102ba366004614e8f565b610aea565b6102e67f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6040519081526020015b60405180910390f35b6102e6610307366004614f16565b610c29565b61029761031a36600461500c565b610e0c565b6103467f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102f0565b6102e67f0a564d4cfe5cb0d4ee082aab2ca54b8c48e129485a8f7c77766ab5ef0c3566f181565b610297610393366004614e8f565b6111d0565b6102e660cc5481565b6102e66103af366004614e72565b60d06020526000908152604090205481565b6102976111e1565b6103ec6103d73660046150b2565b609854600160ff9092169190911b9081161490565b60405190151581526020016102f0565b6098546102e6565b6102976104123660046150d5565b6112a8565b6102e6610425366004614e72565b6113cf565b6103ec610438366004614e72565b60d16020526000908152604090205460ff1681565b61029761171d565b6102e6610463366004615116565b60cd60209081526000928352604080842090915290825290205481565b6102e661048e366004614e72565b60ca6020526000908152604090205481565b609754610346906001600160a01b031681565b6102e66104c1366004614e72565b6001600160a01b0316600090815260ce602052604090205490565b6033546001600160a01b0316610346565b61034673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b610297611731565b61052361051e366004614e72565b61173a565b6040516102f09291906151c3565b60cb54610346906001600160a01b031681565b6102976105523660046151f1565b6118b9565b61029761056536600461521d565b611944565b610297610578366004615252565b61197f565b61029761058b3660046152ad565b611abd565b6103467f000000000000000000000000000000000000000000000000000000000000000081565b6102e66105c536600461559f565b611bbb565b6102976105d83660046150d5565b611c08565b6103ec6105eb366004614e8f565b60cf6020526000908152604090205460ff1681565b61029761060e366004614e72565b611d2a565b6102e661c4e081565b61034661062a3660046151f1565b611d3b565b6103467f000000000000000000000000000000000000000000000000000000000000000081565b6102e66106643660046155d3565b611d73565b6102e6610677366004615622565b611de9565b61029761068a366004614e72565b611ee9565b61029761069d3660046156e1565b611f5f565b6102e6611fca565b6102976106b8366004614e8f565b612008565b6106c5612164565b600260655414156106f15760405162461bcd60e51b81526004016106e890615768565b60405180910390fd5b60026065556106ff8a6121be565b8487146107695760405162461bcd60e51b815260206004820152603260248201527f53747261746567794d616e616765722e736c6173685368617265733a20696e706044820152710eae840d8cadccee8d040dad2e6dac2e8c6d60731b60648201526084016106e8565b600087815b8181101561099d576107d88d88888681811061078c5761078c61579f565b905060200201358d8d858181106107a5576107a561579f565b90506020020160208101906107ba9190614e72565b8888868181106107cc576107cc61579f565b905060200201356122b3565b156107e4578260010192505b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac08b8b8381811061080b5761080b61579f565b90506020020160208101906108209190614e72565b6001600160a01b031614156108ce577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631739ec9e8e8e8888868181106108725761087261579f565b905060200201356040518463ffffffff1660e01b8152600401610897939291906157b5565b600060405180830381600087803b1580156108b157600080fd5b505af11580156108c5573d6000803e3d6000fd5b50505050610995565b8a8a828181106108e0576108e061579f565b90506020020160208101906108f59190614e72565b6001600160a01b031663d9caed128d8b8b858181106109165761091661579f565b905060200201602081019061092b9190614e72565b88888681811061093d5761093d61579f565b905060200201356040518463ffffffff1660e01b8152600401610962939291906157b5565b600060405180830381600087803b15801561097c57600080fd5b505af1158015610990573d6000803e3d6000fd5b505050505b60010161076e565b50604051631608124760e21b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635820491c906109f2908f908e908e908a908a906004016157d9565b600060405180830381600087803b158015610a0c57600080fd5b505af1158015610a20573d6000803e3d6000fd5b505060016065555050505050505050505050505050565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aae9190615868565b6001600160a01b0316336001600160a01b031614610ade5760405162461bcd60e51b81526004016106e890615885565b610ae78161248d565b50565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610b32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5691906158cf565b610b725760405162461bcd60e51b81526004016106e8906158ec565b60985481811614610beb5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016106e8565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b609854600090819060019081161415610c545760405162461bcd60e51b81526004016106e890615934565b60026065541415610c775760405162461bcd60e51b81526004016106e890615768565b6002606555610c8585612584565b42841015610d075760405162461bcd60e51b815260206004820152604360248201527f53747261746567794d616e616765722e6465706f736974496e746f537472617460448201527f656779576974685369676e61747572653a207369676e617475726520657870696064820152621c995960ea1b608482015260a4016106e8565b6001600160a01b038516600090815260ca6020526040812080546001810190915590610d31611fca565b604080517f0a564d4cfe5cb0d4ee082aab2ca54b8c48e129485a8f7c77766ab5ef0c3566f160208201526001600160a01b03808e1692820192909252908b166060820152608081018a905260a0810184905260c0810188905260e00160405160208183030381529060405280519060200120604051602001610dca92919061190160f01b81526002810192909252602282015260420190565b604051602081830303815290604052805190602001209050610ded87828761269e565b610df9878b8b8b61285d565b60016065559a9950505050505050505050565b610e14612164565b60026065541415610e375760405162461bcd60e51b81526004016106e890615768565b6002606555610e54610e4f60e0870160c08801614e72565b6121be565b610e5e858061596b565b84149050610ed45760405162461bcd60e51b815260206004820152603c60248201527f53747261746567794d616e616765722e736c617368517565756564576974686460448201527f726177616c3a20696e707574206c656e677468206d69736d617463680000000060648201526084016106e8565b6000610ee26105c5876159b4565b600081815260cf602052604090205490915060ff16610f6b576040805162461bcd60e51b81526020600482015260248101919091527f53747261746567794d616e616765722e736c617368517565756564576974686460448201527f726177616c3a207769746864726177616c206973206e6f742070656e64696e6760648201526084016106e8565b600081815260cf60205260408120805460ff1916905580610f8c888061596b565b9050905060005b818110156111bf578483108015610fc1575080868685818110610fb857610fb861579f565b90506020020135145b15610fd1578260010192506111b7565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0610ff08a8061596b565b838181106110005761100061579f565b90506020020160208101906110159190614e72565b6001600160a01b031614156110dd576001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016631739ec9e61106360608c0160408d01614e72565b8c61107160208e018e61596b565b868181106110815761108161579f565b905060200201356040518463ffffffff1660e01b81526004016110a6939291906157b5565b600060405180830381600087803b1580156110c057600080fd5b505af11580156110d4573d6000803e3d6000fd5b505050506111b7565b6110e7898061596b565b828181106110f7576110f761579f565b905060200201602081019061110c9190614e72565b6001600160a01b031663d9caed128b8a8a8581811061112d5761112d61579f565b90506020020160208101906111429190614e72565b61114f60208e018e61596b565b8681811061115f5761115f61579f565b905060200201356040518463ffffffff1660e01b8152600401611184939291906157b5565b600060405180830381600087803b15801561119e57600080fd5b505af11580156111b2573d6000803e3d6000fd5b505050505b600101610f93565b505060016065555050505050505050565b6111d8612164565b610ae7816129f3565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124d91906158cf565b6112695760405162461bcd60e51b81526004016106e8906158ec565b600019609881905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6112b0612ab8565b8060005b818110156113c95760d160008585848181106112d2576112d261579f565b90506020020160208101906112e79190614e72565b6001600160a01b0316815260208101919091526040016000205460ff166113c157600160d160008686858181106113205761132061579f565b90506020020160208101906113359190614e72565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe8484838181106113905761139061579f565b90506020020160208101906113a59190614e72565b6040516001600160a01b03909116815260200160405180910390a15b6001016112b4565b50505050565b6000336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611471576040805162461bcd60e51b81526020600482015260248101919091527f53747261746567794d616e616765722e6f6e6c7944656c65676174696f6e4d6160448201527f6e616765723a206e6f74207468652044656c65676174696f6e4d616e6167657260648201526084016106e8565b6098546001906002908116141561149a5760405162461bcd60e51b81526004016106e890615934565b600260655414156114bd5760405162461bcd60e51b81526004016106e890615768565b60026065556114cb83612584565b6001600160a01b038316600090815260ce602052604081205490816001600160401b038111156114fd576114fd614ea8565b604051908082528060200260200182016040528015611526578160200160208202803683370190505b5090506000826001600160401b0381111561154357611543614ea8565b60405190808252806020026020018201604052801561156c578160200160208202803683370190505b5090506000836001600160401b0381111561158957611589614ea8565b6040519080825280602002602001820160405280156115b2578160200160208202803683370190505b50905060005b848110156116fd576000816115ce6001886159dc565b6115d891906159dc565b6001600160a01b038a16600090815260ce60205260409020805491925090829081106116065761160661579f565b9060005260206000200160009054906101000a90046001600160a01b03168583815181106116365761163661579f565b60200260200101906001600160a01b031690816001600160a01b03168152505060cd60008a6001600160a01b03166001600160a01b03168152602001908152602001600020600086848151811061168f5761168f61579f565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548483815181106116ca576116ca61579f565b602002602001018181525050808383815181106116e9576116e961579f565b6020908102919091010152506001016115b8565b5061170d878285858b6001612b46565b6001606555979650505050505050565b611725612164565b61172f6000613296565b565b61172f336132e8565b6001600160a01b038116600090815260ce6020526040812054606091829190816001600160401b0381111561177157611771614ea8565b60405190808252806020026020018201604052801561179a578160200160208202803683370190505b50905060005b8281101561182b576001600160a01b038616600090815260cd6020908152604080832060ce90925282208054919291849081106117df576117df61579f565b60009182526020808320909101546001600160a01b0316835282019290925260400190205482518390839081106118185761181861579f565b60209081029190910101526001016117a0565b5060ce6000866001600160a01b03166001600160a01b0316815260200190815260200160002081818054806020026020016040519081016040528092919081815260200182805480156118a757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611889575b50505050509150935093505050915091565b609854600090600190811614156118e25760405162461bcd60e51b81526004016106e890615934565b600260655414156119055760405162461bcd60e51b81526004016106e890615768565b60026065556119126133fb565b61191b83612584565b61193a8373beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac084613499565b5050600160655550565b600260655414156119675760405162461bcd60e51b81526004016106e890615768565b60026065556119746133fb565b61193a838383613755565b600054610100900460ff161580801561199f5750600054600160ff909116105b806119b95750303b1580156119b9575060005460ff166001145b611a1c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106e8565b6000805460ff191660011790558015611a3f576000805461ff0019166101001790555b611a47613980565b60c955611a548484613a17565b611a5d86613296565b611a6685613b01565b611a6f826129f3565b8015611ab5576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60985460019060029081161415611ae65760405162461bcd60e51b81526004016106e890615934565b60026065541415611b095760405162461bcd60e51b81526004016106e890615768565b600260655560005b888110156111bf57611ba98a8a83818110611b2e57611b2e61579f565b9050602002810190611b4091906159f3565b898984818110611b5257611b5261579f565b9050602002810190611b64919061596b565b898986818110611b7657611b7661579f565b90506020020135888887818110611b8f57611b8f61579f565b9050602002016020810190611ba49190615a13565b613b6a565b80611bb381615a30565b915050611b11565b80516020808301516040808501516060860151608087015160a08801519351600097611beb979096959101615a4b565b604051602081830303815290604052805190602001209050919050565b611c10612ab8565b8060005b818110156113c95760d16000858584818110611c3257611c3261579f565b9050602002016020810190611c479190614e72565b6001600160a01b0316815260208101919091526040016000205460ff1615611d2257600060d16000868685818110611c8157611c8161579f565b9050602002016020810190611c969190614e72565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030848483818110611cf157611cf161579f565b9050602002016020810190611d069190614e72565b6040516001600160a01b03909116815260200160405180910390a15b600101611c14565b611d32612164565b610ae781613b01565b60ce6020528160005260406000208181548110611d5757600080fd5b6000918252602090912001546001600160a01b03169150829050565b609854600090819060019081161415611d9e5760405162461bcd60e51b81526004016106e890615934565b60026065541415611dc15760405162461bcd60e51b81526004016106e890615768565b6002606555611dcf33612584565b611ddb3386868661285d565b600160655595945050505050565b60985460009060019060029081161415611e155760405162461bcd60e51b81526004016106e890615934565b60026065541415611e385760405162461bcd60e51b81526004016106e890615768565b6002606555611e4633612584565b610df9338b8b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808f0282810182019093528e82529093508e92508d91829185019084908082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152508b92508a9150612b469050565b611ef1612164565b6001600160a01b038116611f565760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106e8565b610ae781613296565b60985460019060029081161415611f885760405162461bcd60e51b81526004016106e890615934565b60026065541415611fab5760405162461bcd60e51b81526004016106e890615768565b6002606555611fbd8686868686613b6a565b5050600160655550505050565b60007f0000000000000000000000000000000000000000000000000000000000000000461415611ffb575060c95490565b612003613980565b905090565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561205b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061207f9190615868565b6001600160a01b0316336001600160a01b0316146120af5760405162461bcd60e51b81526004016106e890615885565b60985419811960985419161461212d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016106e8565b609881905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610c1e565b6033546001600160a01b0316331461172f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106e8565b6040516372c1cc1b60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063e583983690602401602060405180830381865afa158015612224573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061224891906158cf565b610ae75760405162461bcd60e51b815260206004820152603660248201527f53747261746567794d616e616765722e6f6e6c7946726f7a656e3a207374616b60448201527532b9103430b9903737ba103132b2b710333937bd32b760511b60648201526084016106e8565b60006001600160a01b0385166123315760405162461bcd60e51b815260206004820152603f60248201527f53747261746567794d616e616765722e5f72656d6f76655368617265733a206460448201527f65706f7369746f722063616e6e6f74206265207a65726f20616464726573730060648201526084016106e8565b816123a45760405162461bcd60e51b815260206004820152603e60248201527f53747261746567794d616e616765722e5f72656d6f76655368617265733a207360448201527f68617265416d6f756e742073686f756c64206e6f74206265207a65726f21000060648201526084016106e8565b6001600160a01b03808616600090815260cd6020908152604080832093871683529290522054808311156124365760405162461bcd60e51b815260206004820152603360248201527f53747261746567794d616e616765722e5f72656d6f76655368617265733a20736044820152720d0c2e4ca82dadeeadce840e8dede40d0d2ced606b1b60648201526084016106e8565b6001600160a01b03808716600090815260cd60209081526040808320938816835292905220838203908190559083141561247f576124758686866143b0565b6001915050612485565b60009150505b949350505050565b6001600160a01b03811661251b5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106e8565b609754604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1609780546001600160a01b0319166001600160a01b0392909216919091179055565b6040516372c1cc1b60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063e583983690602401602060405180830381865afa1580156125ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061260e91906158cf565b15610ae75760405162461bcd60e51b815260206004820152605460248201527f53747261746567794d616e616765722e6f6e6c794e6f7446726f7a656e3a207360448201527f74616b657220686173206265656e2066726f7a656e20616e64206d6179206265606482015273207375626a65637420746f20736c617368696e6760601b608482015260a4016106e8565b6001600160a01b0383163b156127bd57604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906126de9086908690600401615b17565b602060405180830381865afa1580156126fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061271f9190615b30565b6001600160e01b031916146127b85760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a4016106e8565b505050565b826001600160a01b03166127d18383614694565b6001600160a01b0316146127b85760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a4016106e8565b6001600160a01b038316600090815260d16020526040812054849060ff166129035760405162461bcd60e51b815260206004820152604d60248201527f53747261746567794d616e616765722e6f6e6c7953747261746567696573576860448201527f6974656c6973746564466f724465706f7369743a207374726174656779206e6f60648201526c1d081dda1a5d195b1a5cdd1959609a1b608482015260a4016106e8565b6129186001600160a01b0385163387866146b8565b6040516311f9fbc960e21b81526001600160a01b038581166004830152602482018590528616906347e7ef24906044016020604051808303816000875af1158015612967573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061298b9190615b5a565b9150612998868684613499565b604080516001600160a01b03888116825286811660208301528716818301526060810184905290517f7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a969181900360800190a150949350505050565b61c4e0811115612a775760405162461bcd60e51b815260206004820152604360248201527f53747261746567794d616e616765722e7365745769746864726177616c44656c60448201527f61793a205f7769746864726177616c44656c6179426c6f636b7320746f6f20686064820152620d2ced60eb1b608482015260a4016106e8565b60cc5460408051918252602082018390527f4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e910160405180910390a160cc55565b60cb546001600160a01b0316331461172f5760405162461bcd60e51b8152602060048201526044602482018190527f53747261746567794d616e616765722e6f6e6c79537472617465677957686974908201527f656c69737465723a206e6f742074686520737472617465677957686974656c6960648201526339ba32b960e11b608482015260a4016106e8565b60008351855114612ba65760405162461bcd60e51b81526020600482015260366024820152600080516020615cb2833981519152604482015275040d2dce0eae840d8cadccee8d040dad2e6dac2e8c6d60531b60648201526084016106e8565b6001600160a01b038316612c12576040805162461bcd60e51b8152602060048201526024810191909152600080516020615cb283398151915260448201527f2063616e6e6f7420776974686472617720746f207a65726f206164647265737360648201526084016106e8565b604051631608124760e21b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635820491c90612c62908a9089908990600401615b73565b600060405180830381600087803b158015612c7c57600080fd5b505af1158015612c90573d6000803e3d6000fd5b5050506001600160a01b038816600090815260d060205260408120549150805b87518110156130c05773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06001600160a01b0316888281518110612ce957612ce961579f565b60200260200101516001600160a01b03161415612fb857896001600160a01b0316866001600160a01b031614612da95760405162461bcd60e51b81526020600482015260656024820152600080516020615cb283398151915260448201527f2063616e6e6f742071756575652061207769746864726177616c206f6620426560648201527f61636f6e20436861696e2045544820746f206120646966666572656e74206164608482015264647265737360d81b60a482015260c4016106e8565b8751600114612e435760405162461bcd60e51b81526020600482015260666024820152600080516020615cb283398151915260448201527f2063616e6e6f742071756575652061207769746864726177616c20696e636c7560648201527f64696e6720426561636f6e20436861696e2045544820616e64206f7468657220608482015265746f6b656e7360d01b60a482015260c4016106e8565b633b9aca00878281518110612e5a57612e5a61579f565b6020026020010151612e6c9190615bb3565b15612f0a5760405162461bcd60e51b815260206004820152606e6024820152600080516020615cb283398151915260448201527f2063616e6e6f742071756575652061207769746864726177616c206f6620426560648201527f61636f6e20436861696e2045544820666f7220616e206e6f6e2d77686f6c652060848201526d616d6f756e74206f66206777656960901b60a482015260c4016106e8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637723cc1a33898481518110612f4c57612f4c61579f565b60200260200101516040518363ffffffff1660e01b8152600401612f859291906001600160a01b03929092168252602082015260400190565b600060405180830381600087803b158015612f9f57600080fd5b505af1158015612fb3573d6000803e3d6000fd5b505050505b61300f8a8a8481518110612fce57612fce61579f565b60200260200101518a8481518110612fe857612fe861579f565b60200260200101518a85815181106130025761300261579f565b60200260200101516122b3565b1561301b578160010191505b7fcf1c2370141bbd0a6d971beb0e3a2455f24d6e773ddc20ccc1c4e32f3dd9f9f78a848a84815181106130505761305061579f565b60200260200101518a858151811061306a5761306a61579f565b60200260200101516040516130b094939291906001600160a01b0394851681526001600160601b0393909316602084015292166040820152606081019190915260800190565b60405180910390a1600101612cb0565b50604051631976849960e21b81526001600160a01b038a811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906365da126490602401602060405180830381865afa15801561312c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131509190615868565b905061315a614cb6565b506040805180820182526001600160a01b0380891682526001600160601b038087166020808501919091528e8316600081815260d0835286812060018b01909416909355855160c0810187528e81529182018d905294810194909452606084019290925263ffffffff43166080840152831660a08301526131da82611bbb565b600081815260cf60205260409020805460ff19166001179055905086801561321857506001600160a01b038c16600090815260ce6020526040902054155b15613226576132268c6132e8565b604080516001600160a01b038e811682526001600160601b03881660208301528a811682840152851660608201526080810183905290517f32cf9fc97155f52860a59a99879a2e89c1e53f28126a9ab6a2ff29344299e6749181900360a00190a19b9a5050505050505050505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6132f181612584565b6001600160a01b038116600090815260ce60205260409020541561337d5760405162461bcd60e51b815260206004820152603a60248201527f53747261746567794d616e616765722e5f756e64656c65676174653a2064657060448201527f6f7369746f722068617320616374697665206465706f7369747300000000000060648201526084016106e8565b6040516336a2fa1960e21b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063da8be86490602401600060405180830381600087803b1580156133e057600080fd5b505af11580156133f4573d6000803e3d6000fd5b5050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316331461172f5760405162461bcd60e51b815260206004820152603c60248201527f53747261746567794d616e616765722e6f6e6c79456967656e506f644d616e6160448201527f6765723a206e6f742074686520656967656e506f644d616e616765720000000060648201526084016106e8565b6001600160a01b0383166135155760405162461bcd60e51b815260206004820152603c60248201527f53747261746567794d616e616765722e5f6164645368617265733a206465706f60448201527f7369746f722063616e6e6f74206265207a65726f20616464726573730000000060648201526084016106e8565b806135815760405162461bcd60e51b815260206004820152603660248201527f53747261746567794d616e616765722e5f6164645368617265733a207368617260448201527565732073686f756c64206e6f74206265207a65726f2160501b60648201526084016106e8565b6001600160a01b03808416600090815260cd6020908152604080832093861683529290522054613692576001600160a01b038316600090815260ce6020908152604090912054106136535760405162461bcd60e51b815260206004820152605060248201527f53747261746567794d616e616765722e5f6164645368617265733a206465706f60448201527f73697420776f756c6420657863656564204d41585f5354414b45525f5354524160648201526f0a88a8eb2be9892a6a8be988a9c8ea8960831b608482015260a4016106e8565b6001600160a01b03838116600090815260ce602090815260408220805460018101825590835291200180546001600160a01b0319169184169190911790555b6001600160a01b03808416600090815260cd60209081526040808320938616835292905290812080548392906136c9908490615bd5565b9091555050604051631452b9d760e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906328a573ae9061371e908690869086906004016157b5565b600060405180830381600087803b15801561373857600080fd5b505af115801561374c573d6000803e3d6000fd5b50505050505050565b60008112156138c4576040805160018082528183019092526000916020808301908036833701905050905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0816000815181106137a8576137a861579f565b6001600160a01b0392909216602092830291909101909101526040805160018082528183019092526000918160200160208202803683370190505090506137ee83615bed565b816000815181106138015761380161579f565b60200260200101818152505061383a858573beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0846000815181106130025761300261579f565b50604051631608124760e21b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635820491c9061388b90889086908690600401615b73565b600060405180830381600087803b1580156138a557600080fd5b505af11580156138b9573d6000803e3d6000fd5b505050505050505050565b806138e48473beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac083613499565b604051631452b9d760e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906328a573ae9061394890879073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac09086906004016157b5565b600060405180830381600087803b15801561396257600080fd5b505af1158015613976573d6000803e3d6000fd5b5050505050505050565b604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b6097546001600160a01b0316158015613a3857506001600160a01b03821615155b613aba5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106e8565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613afd8261248d565b5050565b60cb54604080516001600160a01b03928316815291831660208301527f4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29910160405180910390a160cb80546001600160a01b0319166001600160a01b0392909216919091179055565b613b82613b7d60e0870160c08801614e72565b612584565b6000613b906105c5876159b4565b600081815260cf602052604090205490915060ff16613c115760405162461bcd60e51b81526020600482015260436024820152600080516020615c9283398151915260448201527f746864726177616c3a207769746864726177616c206973206e6f742070656e64606482015262696e6760e81b608482015260a4016106e8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016638105e043613c5060e0890160c08a01614e72565b613c6060c08a0160a08b01615c0a565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015263ffffffff166024820152604481018690526064016020604051808303816000875af1158015613cb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cdc91906158cf565b613d625760405162461bcd60e51b81526020600482015260576024820152600080516020615c9283398151915260448201527f746864726177616c3a207368617265732070656e64696e67207769746864726160648201527f77616c20617265207374696c6c20736c61736861626c65000000000000000000608482015260a4016106e8565b60cc544390613d7760c0890160a08a01615c0a565b63ffffffff16613d879190615bd5565b111580613ddf575073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0613dae878061596b565b6000818110613dbf57613dbf61579f565b9050602002016020810190613dd49190614e72565b6001600160a01b0316145b613e655760405162461bcd60e51b81526020600482015260596024820152600080516020615c9283398151915260448201527f746864726177616c3a207769746864726177616c44656c6179426c6f636b732060648201527f706572696f6420686173206e6f74207965742070617373656400000000000000608482015260a4016106e8565b613e756080870160608801614e72565b6001600160a01b0316336001600160a01b031614613f1d5760405162461bcd60e51b815260206004820152606460248201819052600080516020615c9283398151915260448301527f746864726177616c3a206f6e6c79207370656369666965642077697468647261908201527f7765722063616e20636f6d706c65746520612071756575656420776974686472608482015263185dd85b60e21b60a482015260c4016106e8565b600081815260cf60205260408120805460ff19169055613f3d878061596b565b91505082156141b357613f50878061596b565b86149050613fb45760405162461bcd60e51b815260206004820152603f6024820152600080516020615c9283398151915260448201527f746864726177616c3a20696e707574206c656e677468206d69736d617463680060648201526084016106e8565b60005b818110156141ad5773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0613fde898061596b565b83818110613fee57613fee61579f565b90506020020160208101906140039190614e72565b6001600160a01b031614156140cb576001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016631739ec9e61405160608b0160408c01614e72565b3361405f60208d018d61596b565b8681811061406f5761406f61579f565b905060200201356040518463ffffffff1660e01b8152600401614094939291906157b5565b600060405180830381600087803b1580156140ae57600080fd5b505af11580156140c2573d6000803e3d6000fd5b505050506141a5565b6140d5888061596b565b828181106140e5576140e561579f565b90506020020160208101906140fa9190614e72565b6001600160a01b031663d9caed123389898581811061411b5761411b61579f565b90506020020160208101906141309190614e72565b61413d60208d018d61596b565b8681811061414d5761414d61579f565b905060200201356040518463ffffffff1660e01b8152600401614172939291906157b5565b600060405180830381600087803b15801561418c57600080fd5b505af11580156141a0573d6000803e3d6000fd5b505050505b600101613fb7565b5061433b565b60005b818110156143395761421a336141cc8a8061596b565b848181106141dc576141dc61579f565b90506020020160208101906141f19190614e72565b6141fe60208c018c61596b565b8581811061420e5761420e61579f565b90506020020135613499565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0614239898061596b565b838181106142495761424961579f565b905060200201602081019061425e9190614e72565b6001600160a01b03161415614331576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663d34dedd06142ac60608b0160408c01614e72565b6142b960208c018c61596b565b858181106142c9576142c961579f565b6040516001600160e01b031960e087901b1681526001600160a01b0390941660048501526020029190910135602483015250604401600060405180830381600087803b15801561431857600080fd5b505af115801561432c573d6000803e3d6000fd5b505050505b6001016141b6565b505b3361434c6060890160408a01614e72565b6001600160a01b03167fe7eb0ca11b83744ece3d78e9be01b913425fbae70c32ce27726d0ecde92ef8d261438660a08b0160808c01615c25565b604080516001600160601b039092168252602082018790520160405180910390a350505050505050565b6001600160a01b03838116600090815260ce60205260409020805491831691849081106143df576143df61579f565b6000918252602090912001546001600160a01b031614156144a0576001600160a01b038316600090815260ce602052604090208054614420906001906159dc565b815481106144305761443061579f565b60009182526020808320909101546001600160a01b03868116845260ce909252604090922080549190921691908490811061446d5761446d61579f565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550614646565b6001600160a01b038316600090815260ce6020526040812054905b818110156145bb576001600160a01b03858116600090815260ce60205260409020805491851691839081106144f2576144f261579f565b6000918252602090912001546001600160a01b031614156145b3576001600160a01b038516600090815260ce602052604090208054614533906001906159dc565b815481106145435761454361579f565b60009182526020808320909101546001600160a01b03888116845260ce90925260409092208054919092169190839081106145805761458061579f565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506145bb565b6001016144bb565b818114156146435760405162461bcd60e51b815260206004820152604960248201527f53747261746567794d616e616765722e5f72656d6f766553747261746567794660448201527f726f6d5374616b657253747261746567794c6973743a207374726174656779206064820152681b9bdd08199bdd5b9960ba1b608482015260a4016106e8565b50505b6001600160a01b038316600090815260ce6020526040902080548061466d5761466d615c40565b600082815260209020810160001990810180546001600160a01b0319169055019055505050565b60008060006146a38585614710565b915091506146b081614780565b509392505050565b6113c9846323b872dd60e01b8585856040516024016146d9939291906157b5565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261493b565b6000808251604114156147475760208301516040840151606085015160001a61473b87828585614a0d565b94509450505050614779565b8251604014156147715760208301516040840151614766868383614afa565b935093505050614779565b506000905060025b9250929050565b600081600481111561479457614794615c56565b141561479d5750565b60018160048111156147b1576147b1615c56565b14156147ff5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016106e8565b600281600481111561481357614813615c56565b14156148615760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016106e8565b600381600481111561487557614875615c56565b14156148ce5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016106e8565b60048160048111156148e2576148e2615c56565b1415610ae75760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016106e8565b6000614990826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614b339092919063ffffffff16565b8051909150156127b857808060200190518101906149ae91906158cf565b6127b85760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016106e8565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614a445750600090506003614af1565b8460ff16601b14158015614a5c57508460ff16601c14155b15614a6d5750600090506004614af1565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614ac1573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614aea57600060019250925050614af1565b9150600090505b94509492505050565b6000806001600160ff1b03831681614b1760ff86901c601b615bd5565b9050614b2587828885614a0d565b935093505050935093915050565b6060614b428484600085614b4c565b90505b9392505050565b606082471015614bad5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016106e8565b6001600160a01b0385163b614c045760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016106e8565b600080866001600160a01b03168587604051614c209190615c6c565b60006040518083038185875af1925050503d8060008114614c5d576040519150601f19603f3d011682016040523d82523d6000602084013e614c62565b606091505b5091509150614c72828286614c7d565b979650505050505050565b60608315614c8c575081614b45565b825115614c9c5782518084602001fd5b8160405162461bcd60e51b81526004016106e89190615c7e565b6040518060c00160405280606081526020016060815260200160006001600160a01b03168152602001614d0e604051806040016040528060006001600160a01b0316815260200160006001600160601b031681525090565b815260006020820181905260409091015290565b6001600160a01b0381168114610ae757600080fd5b8035614d4281614d22565b919050565b60008083601f840112614d5957600080fd5b5081356001600160401b03811115614d7057600080fd5b6020830191508360208260051b850101111561477957600080fd5b60008060008060008060008060008060c08b8d031215614daa57600080fd5b614db38b614d37565b9950614dc160208c01614d37565b985060408b01356001600160401b0380821115614ddd57600080fd5b614de98e838f01614d47565b909a50985060608d0135915080821115614e0257600080fd5b614e0e8e838f01614d47565b909850965060808d0135915080821115614e2757600080fd5b614e338e838f01614d47565b909650945060a08d0135915080821115614e4c57600080fd5b50614e598d828e01614d47565b915080935050809150509295989b9194979a5092959850565b600060208284031215614e8457600080fd5b8135614b4581614d22565b600060208284031215614ea157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b0381118282101715614ee057614ee0614ea8565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614f0e57614f0e614ea8565b604052919050565b60008060008060008060c08789031215614f2f57600080fd5b8635614f3a81614d22565b9550602087810135614f4b81614d22565b9550604088013594506060880135614f6281614d22565b93506080880135925060a08801356001600160401b0380821115614f8557600080fd5b818a0191508a601f830112614f9957600080fd5b813581811115614fab57614fab614ea8565b614fbd601f8201601f19168501614ee6565b91508082528b84828501011115614fd357600080fd5b80848401858401376000848284010152508093505050509295509295509295565b600060e0828403121561500657600080fd5b50919050565b6000806000806000806080878903121561502557600080fd5b863561503081614d22565b955060208701356001600160401b038082111561504c57600080fd5b6150588a838b01614ff4565b9650604089013591508082111561506e57600080fd5b61507a8a838b01614d47565b9096509450606089013591508082111561509357600080fd5b506150a089828a01614d47565b979a9699509497509295939492505050565b6000602082840312156150c457600080fd5b813560ff81168114614b4557600080fd5b600080602083850312156150e857600080fd5b82356001600160401b038111156150fe57600080fd5b61510a85828601614d47565b90969095509350505050565b6000806040838503121561512957600080fd5b823561513481614d22565b9150602083013561514481614d22565b809150509250929050565b600081518084526020808501945080840160005b838110156151885781516001600160a01b031687529582019590820190600101615163565b509495945050505050565b600081518084526020808501945080840160005b83811015615188578151875295820195908201906001016151a7565b6040815260006151d6604083018561514f565b82810360208401526151e88185615193565b95945050505050565b6000806040838503121561520457600080fd5b823561520f81614d22565b946020939093013593505050565b60008060006060848603121561523257600080fd5b833561523d81614d22565b95602085013595506040909401359392505050565b600080600080600060a0868803121561526a57600080fd5b853561527581614d22565b9450602086013561528581614d22565b9350604086013561529581614d22565b94979396509394606081013594506080013592915050565b6000806000806000806000806080898b0312156152c957600080fd5b88356001600160401b03808211156152e057600080fd5b6152ec8c838d01614d47565b909a50985060208b013591508082111561530557600080fd5b6153118c838d01614d47565b909850965060408b013591508082111561532a57600080fd5b6153368c838d01614d47565b909650945060608b013591508082111561534f57600080fd5b5061535c8b828c01614d47565b999c989b5096995094979396929594505050565b60006001600160401b0382111561538957615389614ea8565b5060051b60200190565b600082601f8301126153a457600080fd5b813560206153b96153b483615370565b614ee6565b82815260059290921b840181019181810190868411156153d857600080fd5b8286015b848110156153fc5780356153ef81614d22565b83529183019183016153dc565b509695505050505050565b600082601f83011261541857600080fd5b813560206154286153b483615370565b82815260059290921b8401810191818101908684111561544757600080fd5b8286015b848110156153fc578035835291830191830161544b565b80356001600160601b0381168114614d4257600080fd5b60006040828403121561548b57600080fd5b604051604081018181106001600160401b03821117156154ad576154ad614ea8565b60405290508082356154be81614d22565b81526154cc60208401615462565b60208201525092915050565b803563ffffffff81168114614d4257600080fd5b600060e082840312156154fe57600080fd5b615506614ebe565b905081356001600160401b038082111561551f57600080fd5b61552b85838601615393565b8352602084013591508082111561554157600080fd5b5061554e84828501615407565b60208301525061556060408301614d37565b60408201526155728360608401615479565b606082015261558360a083016154d8565b608082015261559460c08301614d37565b60a082015292915050565b6000602082840312156155b157600080fd5b81356001600160401b038111156155c757600080fd5b612485848285016154ec565b6000806000606084860312156155e857600080fd5b83356155f381614d22565b9250602084013561560381614d22565b929592945050506040919091013590565b8015158114610ae757600080fd5b60008060008060008060008060a0898b03121561563e57600080fd5b88356001600160401b038082111561565557600080fd5b6156618c838d01614d47565b909a50985060208b013591508082111561567a57600080fd5b6156868c838d01614d47565b909850965060408b013591508082111561569f57600080fd5b506156ac8b828c01614d47565b90955093505060608901356156c081614d22565b915060808901356156d081615614565b809150509295985092959890939650565b6000806000806000608086880312156156f957600080fd5b85356001600160401b038082111561571057600080fd5b61571c89838a01614ff4565b9650602088013591508082111561573257600080fd5b5061573f88828901614d47565b90955093505060408601359150606086013561575a81615614565b809150509295509295909350565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03868116825260606020808401829052908301869052600091879160808501845b8981101561582857843561581481614d22565b841682529382019390820190600101615801565b5085810360408701528681526001600160fb1b0387111561584857600080fd5b8660051b9350838883830137600093010191825250979650505050505050565b60006020828403121561587a57600080fd5b8151614b4581614d22565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156158e157600080fd5b8151614b4581615614565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b6000808335601e1984360301811261598257600080fd5b8301803591506001600160401b0382111561599c57600080fd5b6020019150600581901b360382131561477957600080fd5b60006159c036836154ec565b92915050565b634e487b7160e01b600052601160045260246000fd5b6000828210156159ee576159ee6159c6565b500390565b6000823560de19833603018112615a0957600080fd5b9190910192915050565b600060208284031215615a2557600080fd5b8135614b4581615614565b6000600019821415615a4457615a446159c6565b5060010190565b60e081526000615a5e60e083018961514f565b8281036020840152615a708189615193565b6001600160a01b0397881660408501528651881660608501526020909601516001600160601b03166080840152505063ffffffff9290921660a083015290921660c09092019190915292915050565b60005b83811015615ada578181015183820152602001615ac2565b838111156113c95750506000910152565b60008151808452615b03816020860160208601615abf565b601f01601f19169290920160200192915050565b828152604060208201526000614b426040830184615aeb565b600060208284031215615b4257600080fd5b81516001600160e01b031981168114614b4557600080fd5b600060208284031215615b6c57600080fd5b5051919050565b6001600160a01b0384168152606060208201819052600090615b979083018561514f565b8281036040840152615ba98185615193565b9695505050505050565b600082615bd057634e487b7160e01b600052601260045260246000fd5b500690565b60008219821115615be857615be86159c6565b500190565b6000600160ff1b821415615c0357615c036159c6565b5060000390565b600060208284031215615c1c57600080fd5b614b45826154d8565b600060208284031215615c3757600080fd5b614b4582615462565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b60008251615a09818460208701615abf565b602081526000614b456020830184615aeb56fe53747261746567794d616e616765722e636f6d706c657465517565756564576953747261746567794d616e616765722e71756575655769746864726177616c3aa2646970667358221220479cd7cb468a365be1dc0612804c0b8c3907723b55d226c5868bd74b4e9815b964736f6c634300080c0033",
}

// ContractStrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractStrategyManagerMetaData.ABI instead.
var ContractStrategyManagerABI = ContractStrategyManagerMetaData.ABI

// ContractStrategyManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractStrategyManagerMetaData.Bin instead.
var ContractStrategyManagerBin = ContractStrategyManagerMetaData.Bin

// DeployContractStrategyManager deploys a new Ethereum contract, binding an instance of ContractStrategyManager to it.
func DeployContractStrategyManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _eigenPodManager common.Address, _slasher common.Address) (common.Address, *types.Transaction, *ContractStrategyManager, error) {
	parsed, err := ContractStrategyManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractStrategyManagerBin), backend, _delegation, _eigenPodManager, _slasher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractStrategyManager{ContractStrategyManagerCaller: ContractStrategyManagerCaller{contract: contract}, ContractStrategyManagerTransactor: ContractStrategyManagerTransactor{contract: contract}, ContractStrategyManagerFilterer: ContractStrategyManagerFilterer{contract: contract}}, nil
}

// ContractStrategyManager is an auto generated Go binding around an Ethereum contract.
type ContractStrategyManager struct {
	ContractStrategyManagerCaller     // Read-only binding to the contract
	ContractStrategyManagerTransactor // Write-only binding to the contract
	ContractStrategyManagerFilterer   // Log filterer for contract events
}

// ContractStrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractStrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractStrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractStrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractStrategyManagerSession struct {
	Contract     *ContractStrategyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractStrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractStrategyManagerCallerSession struct {
	Contract *ContractStrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractStrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractStrategyManagerTransactorSession struct {
	Contract     *ContractStrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractStrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractStrategyManagerRaw struct {
	Contract *ContractStrategyManager // Generic contract binding to access the raw methods on
}

// ContractStrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractStrategyManagerCallerRaw struct {
	Contract *ContractStrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractStrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractStrategyManagerTransactorRaw struct {
	Contract *ContractStrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractStrategyManager creates a new instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManager(address common.Address, backend bind.ContractBackend) (*ContractStrategyManager, error) {
	contract, err := bindContractStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManager{ContractStrategyManagerCaller: ContractStrategyManagerCaller{contract: contract}, ContractStrategyManagerTransactor: ContractStrategyManagerTransactor{contract: contract}, ContractStrategyManagerFilterer: ContractStrategyManagerFilterer{contract: contract}}, nil
}

// NewContractStrategyManagerCaller creates a new read-only instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractStrategyManagerCaller, error) {
	contract, err := bindContractStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerCaller{contract: contract}, nil
}

// NewContractStrategyManagerTransactor creates a new write-only instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractStrategyManagerTransactor, error) {
	contract, err := bindContractStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerTransactor{contract: contract}, nil
}

// NewContractStrategyManagerFilterer creates a new log filterer instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractStrategyManagerFilterer, error) {
	contract, err := bindContractStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerFilterer{contract: contract}, nil
}

// bindContractStrategyManager binds a generic wrapper to an already deployed contract.
func bindContractStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractStrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyManager.Contract.ContractStrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ContractStrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ContractStrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyManager *ContractStrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyManager *ContractStrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyManager *ContractStrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DEPOSITTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DEPOSITTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DOMAINTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DOMAINTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _ContractStrategyManager.Contract.MAXWITHDRAWALDELAYBLOCKS(&_ContractStrategyManager.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _ContractStrategyManager.Contract.MAXWITHDRAWALDELAYBLOCKS(&_ContractStrategyManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractStrategyManager.Contract.BeaconChainETHStrategy(&_ContractStrategyManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractStrategyManager.Contract.BeaconChainETHStrategy(&_ContractStrategyManager.CallOpts)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, queuedWithdrawal IStrategyManagerQueuedWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "calculateWithdrawalRoot", queuedWithdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerQueuedWithdrawal) ([32]byte, error) {
	return _ContractStrategyManager.Contract.CalculateWithdrawalRoot(&_ContractStrategyManager.CallOpts, queuedWithdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerQueuedWithdrawal) ([32]byte, error) {
	return _ContractStrategyManager.Contract.CalculateWithdrawalRoot(&_ContractStrategyManager.CallOpts, queuedWithdrawal)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) Delegation() (common.Address, error) {
	return _ContractStrategyManager.Contract.Delegation(&_ContractStrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractStrategyManager.Contract.Delegation(&_ContractStrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) DomainSeparator() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DomainSeparator(&_ContractStrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DomainSeparator(&_ContractStrategyManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) EigenPodManager() (common.Address, error) {
	return _ContractStrategyManager.Contract.EigenPodManager(&_ContractStrategyManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _ContractStrategyManager.Contract.EigenPodManager(&_ContractStrategyManager.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address depositor) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerCaller) GetDeposits(opts *bind.CallOpts, depositor common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "getDeposits", depositor)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address depositor) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerSession) GetDeposits(depositor common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractStrategyManager.Contract.GetDeposits(&_ContractStrategyManager.CallOpts, depositor)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address depositor) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) GetDeposits(depositor common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractStrategyManager.Contract.GetDeposits(&_ContractStrategyManager.CallOpts, depositor)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.Nonces(&_ContractStrategyManager.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.Nonces(&_ContractStrategyManager.CallOpts, arg0)
}

// NumWithdrawalsQueued is a free data retrieval call binding the contract method 0x56631028.
//
// Solidity: function numWithdrawalsQueued(address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) NumWithdrawalsQueued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "numWithdrawalsQueued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumWithdrawalsQueued is a free data retrieval call binding the contract method 0x56631028.
//
// Solidity: function numWithdrawalsQueued(address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) NumWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.NumWithdrawalsQueued(&_ContractStrategyManager.CallOpts, arg0)
}

// NumWithdrawalsQueued is a free data retrieval call binding the contract method 0x56631028.
//
// Solidity: function numWithdrawalsQueued(address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) NumWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.NumWithdrawalsQueued(&_ContractStrategyManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) Owner() (common.Address, error) {
	return _ContractStrategyManager.Contract.Owner(&_ContractStrategyManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Owner() (common.Address, error) {
	return _ContractStrategyManager.Contract.Owner(&_ContractStrategyManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyManager.Contract.Paused(&_ContractStrategyManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyManager.Contract.Paused(&_ContractStrategyManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) Paused0() (*big.Int, error) {
	return _ContractStrategyManager.Contract.Paused0(&_ContractStrategyManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractStrategyManager.Contract.Paused0(&_ContractStrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyManager.Contract.PauserRegistry(&_ContractStrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyManager.Contract.PauserRegistry(&_ContractStrategyManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) Slasher() (common.Address, error) {
	return _ContractStrategyManager.Contract.Slasher(&_ContractStrategyManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Slasher() (common.Address, error) {
	return _ContractStrategyManager.Contract.Slasher(&_ContractStrategyManager.CallOpts)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerStrategyList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerStrategyList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerStrategyList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _ContractStrategyManager.Contract.StakerStrategyList(&_ContractStrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerStrategyList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _ContractStrategyManager.Contract.StakerStrategyList(&_ContractStrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyListLength(&_ContractStrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyListLength(&_ContractStrategyManager.CallOpts, staker)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerStrategyShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerStrategyShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerStrategyShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyShares(&_ContractStrategyManager.CallOpts, arg0, arg1)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerStrategyShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyShares(&_ContractStrategyManager.CallOpts, arg0, arg1)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerSession) StrategyIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _ContractStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_ContractStrategyManager.CallOpts, arg0)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StrategyIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _ContractStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_ContractStrategyManager.CallOpts, arg0)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractStrategyManager.Contract.StrategyWhitelister(&_ContractStrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractStrategyManager.Contract.StrategyWhitelister(&_ContractStrategyManager.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "withdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _ContractStrategyManager.Contract.WithdrawalDelayBlocks(&_ContractStrategyManager.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _ContractStrategyManager.Contract.WithdrawalDelayBlocks(&_ContractStrategyManager.CallOpts)
}

// WithdrawalRootPending is a free data retrieval call binding the contract method 0xc3c6b3a9.
//
// Solidity: function withdrawalRootPending(bytes32 ) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCaller) WithdrawalRootPending(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "withdrawalRootPending", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalRootPending is a free data retrieval call binding the contract method 0xc3c6b3a9.
//
// Solidity: function withdrawalRootPending(bytes32 ) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerSession) WithdrawalRootPending(arg0 [32]byte) (bool, error) {
	return _ContractStrategyManager.Contract.WithdrawalRootPending(&_ContractStrategyManager.CallOpts, arg0)
}

// WithdrawalRootPending is a free data retrieval call binding the contract method 0xc3c6b3a9.
//
// Solidity: function withdrawalRootPending(bytes32 ) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) WithdrawalRootPending(arg0 [32]byte) (bool, error) {
	return _ContractStrategyManager.Contract.WithdrawalRootPending(&_ContractStrategyManager.CallOpts, arg0)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToWhitelist)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xf3be65d3.
//
// Solidity: function completeQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, queuedWithdrawal IStrategyManagerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "completeQueuedWithdrawal", queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xf3be65d3.
//
// Solidity: function completeQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) CompleteQueuedWithdrawal(queuedWithdrawal IStrategyManagerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.CompleteQueuedWithdrawal(&_ContractStrategyManager.TransactOpts, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xf3be65d3.
//
// Solidity: function completeQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) CompleteQueuedWithdrawal(queuedWithdrawal IStrategyManagerQueuedWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.CompleteQueuedWithdrawal(&_ContractStrategyManager.TransactOpts, queuedWithdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xa782d945.
//
// Solidity: function completeQueuedWithdrawals((address[],uint256[],address,(address,uint96),uint32,address)[] queuedWithdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, queuedWithdrawals []IStrategyManagerQueuedWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "completeQueuedWithdrawals", queuedWithdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xa782d945.
//
// Solidity: function completeQueuedWithdrawals((address[],uint256[],address,(address,uint96),uint32,address)[] queuedWithdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) CompleteQueuedWithdrawals(queuedWithdrawals []IStrategyManagerQueuedWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.CompleteQueuedWithdrawals(&_ContractStrategyManager.TransactOpts, queuedWithdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xa782d945.
//
// Solidity: function completeQueuedWithdrawals((address[],uint256[],address,(address,uint96),uint32,address)[] queuedWithdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) CompleteQueuedWithdrawals(queuedWithdrawals []IStrategyManagerQueuedWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.CompleteQueuedWithdrawals(&_ContractStrategyManager.TransactOpts, queuedWithdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// DepositBeaconChainETH is a paid mutator transaction binding the contract method 0x9f00fa24.
//
// Solidity: function depositBeaconChainETH(address staker, uint256 amount) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) DepositBeaconChainETH(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "depositBeaconChainETH", staker, amount)
}

// DepositBeaconChainETH is a paid mutator transaction binding the contract method 0x9f00fa24.
//
// Solidity: function depositBeaconChainETH(address staker, uint256 amount) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) DepositBeaconChainETH(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositBeaconChainETH(&_ContractStrategyManager.TransactOpts, staker, amount)
}

// DepositBeaconChainETH is a paid mutator transaction binding the contract method 0x9f00fa24.
//
// Solidity: function depositBeaconChainETH(address staker, uint256 amount) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) DepositBeaconChainETH(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositBeaconChainETH(&_ContractStrategyManager.TransactOpts, staker, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategy(&_ContractStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategy(&_ContractStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// ForceTotalWithdrawal is a paid mutator transaction binding the contract method 0x5e97f8bd.
//
// Solidity: function forceTotalWithdrawal(address staker) returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) ForceTotalWithdrawal(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "forceTotalWithdrawal", staker)
}

// ForceTotalWithdrawal is a paid mutator transaction binding the contract method 0x5e97f8bd.
//
// Solidity: function forceTotalWithdrawal(address staker) returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) ForceTotalWithdrawal(staker common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ForceTotalWithdrawal(&_ContractStrategyManager.TransactOpts, staker)
}

// ForceTotalWithdrawal is a paid mutator transaction binding the contract method 0x5e97f8bd.
//
// Solidity: function forceTotalWithdrawal(address staker) returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) ForceTotalWithdrawal(staker common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ForceTotalWithdrawal(&_ContractStrategyManager.TransactOpts, staker)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, address _pauserRegistry, uint256 initialPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "initialize", initialOwner, initialStrategyWhitelister, _pauserRegistry, initialPausedStatus, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, address _pauserRegistry, uint256 initialPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Initialize(&_ContractStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, _pauserRegistry, initialPausedStatus, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, address _pauserRegistry, uint256 initialPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Initialize(&_ContractStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, _pauserRegistry, initialPausedStatus, _withdrawalDelayBlocks)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Pause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Pause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.PauseAll(&_ContractStrategyManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.PauseAll(&_ContractStrategyManager.TransactOpts)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xf123991e.
//
// Solidity: function queueWithdrawal(uint256[] strategyIndexes, address[] strategies, uint256[] shares, address withdrawer, bool undelegateIfPossible) returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) QueueWithdrawal(opts *bind.TransactOpts, strategyIndexes []*big.Int, strategies []common.Address, shares []*big.Int, withdrawer common.Address, undelegateIfPossible bool) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "queueWithdrawal", strategyIndexes, strategies, shares, withdrawer, undelegateIfPossible)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xf123991e.
//
// Solidity: function queueWithdrawal(uint256[] strategyIndexes, address[] strategies, uint256[] shares, address withdrawer, bool undelegateIfPossible) returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) QueueWithdrawal(strategyIndexes []*big.Int, strategies []common.Address, shares []*big.Int, withdrawer common.Address, undelegateIfPossible bool) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.QueueWithdrawal(&_ContractStrategyManager.TransactOpts, strategyIndexes, strategies, shares, withdrawer, undelegateIfPossible)
}

// QueueWithdrawal is a paid mutator transaction binding the contract method 0xf123991e.
//
// Solidity: function queueWithdrawal(uint256[] strategyIndexes, address[] strategies, uint256[] shares, address withdrawer, bool undelegateIfPossible) returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) QueueWithdrawal(strategyIndexes []*big.Int, strategies []common.Address, shares []*big.Int, withdrawer common.Address, undelegateIfPossible bool) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.QueueWithdrawal(&_ContractStrategyManager.TransactOpts, strategyIndexes, strategies, shares, withdrawer, undelegateIfPossible)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 beaconChainETHStrategyIndex, int256 sharesDelta) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, beaconChainETHStrategyIndex *big.Int, sharesDelta *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "recordBeaconChainETHBalanceUpdate", podOwner, beaconChainETHStrategyIndex, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 beaconChainETHStrategyIndex, int256 sharesDelta) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, beaconChainETHStrategyIndex *big.Int, sharesDelta *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RecordBeaconChainETHBalanceUpdate(&_ContractStrategyManager.TransactOpts, podOwner, beaconChainETHStrategyIndex, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xa1ca780b.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 beaconChainETHStrategyIndex, int256 sharesDelta) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, beaconChainETHStrategyIndex *big.Int, sharesDelta *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RecordBeaconChainETHBalanceUpdate(&_ContractStrategyManager.TransactOpts, podOwner, beaconChainETHStrategyIndex, sharesDelta)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RenounceOwnership(&_ContractStrategyManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RenounceOwnership(&_ContractStrategyManager.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetPauserRegistry(&_ContractStrategyManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetPauserRegistry(&_ContractStrategyManager.TransactOpts, newPauserRegistry)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "setStrategyWhitelister", newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetStrategyWhitelister(&_ContractStrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetStrategyWhitelister(&_ContractStrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) SetWithdrawalDelayBlocks(opts *bind.TransactOpts, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "setWithdrawalDelayBlocks", _withdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) SetWithdrawalDelayBlocks(_withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetWithdrawalDelayBlocks(&_ContractStrategyManager.TransactOpts, _withdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) SetWithdrawalDelayBlocks(_withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetWithdrawalDelayBlocks(&_ContractStrategyManager.TransactOpts, _withdrawalDelayBlocks)
}

// SlashQueuedWithdrawal is a paid mutator transaction binding the contract method 0x43c09061.
//
// Solidity: function slashQueuedWithdrawal(address recipient, (address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256[] indicesToSkip) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) SlashQueuedWithdrawal(opts *bind.TransactOpts, recipient common.Address, queuedWithdrawal IStrategyManagerQueuedWithdrawal, tokens []common.Address, indicesToSkip []*big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "slashQueuedWithdrawal", recipient, queuedWithdrawal, tokens, indicesToSkip)
}

// SlashQueuedWithdrawal is a paid mutator transaction binding the contract method 0x43c09061.
//
// Solidity: function slashQueuedWithdrawal(address recipient, (address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256[] indicesToSkip) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) SlashQueuedWithdrawal(recipient common.Address, queuedWithdrawal IStrategyManagerQueuedWithdrawal, tokens []common.Address, indicesToSkip []*big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SlashQueuedWithdrawal(&_ContractStrategyManager.TransactOpts, recipient, queuedWithdrawal, tokens, indicesToSkip)
}

// SlashQueuedWithdrawal is a paid mutator transaction binding the contract method 0x43c09061.
//
// Solidity: function slashQueuedWithdrawal(address recipient, (address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal, address[] tokens, uint256[] indicesToSkip) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) SlashQueuedWithdrawal(recipient common.Address, queuedWithdrawal IStrategyManagerQueuedWithdrawal, tokens []common.Address, indicesToSkip []*big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SlashQueuedWithdrawal(&_ContractStrategyManager.TransactOpts, recipient, queuedWithdrawal, tokens, indicesToSkip)
}

// SlashShares is a paid mutator transaction binding the contract method 0x06f1f684.
//
// Solidity: function slashShares(address slashedAddress, address recipient, address[] strategies, address[] tokens, uint256[] strategyIndexes, uint256[] shareAmounts) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) SlashShares(opts *bind.TransactOpts, slashedAddress common.Address, recipient common.Address, strategies []common.Address, tokens []common.Address, strategyIndexes []*big.Int, shareAmounts []*big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "slashShares", slashedAddress, recipient, strategies, tokens, strategyIndexes, shareAmounts)
}

// SlashShares is a paid mutator transaction binding the contract method 0x06f1f684.
//
// Solidity: function slashShares(address slashedAddress, address recipient, address[] strategies, address[] tokens, uint256[] strategyIndexes, uint256[] shareAmounts) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) SlashShares(slashedAddress common.Address, recipient common.Address, strategies []common.Address, tokens []common.Address, strategyIndexes []*big.Int, shareAmounts []*big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SlashShares(&_ContractStrategyManager.TransactOpts, slashedAddress, recipient, strategies, tokens, strategyIndexes, shareAmounts)
}

// SlashShares is a paid mutator transaction binding the contract method 0x06f1f684.
//
// Solidity: function slashShares(address slashedAddress, address recipient, address[] strategies, address[] tokens, uint256[] strategyIndexes, uint256[] shareAmounts) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) SlashShares(slashedAddress common.Address, recipient common.Address, strategies []common.Address, tokens []common.Address, strategyIndexes []*big.Int, shareAmounts []*big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SlashShares(&_ContractStrategyManager.TransactOpts, slashedAddress, recipient, strategies, tokens, strategyIndexes, shareAmounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.TransferOwnership(&_ContractStrategyManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.TransferOwnership(&_ContractStrategyManager.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Undelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "undelegate")
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Undelegate() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Undelegate(&_ContractStrategyManager.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Undelegate() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Undelegate(&_ContractStrategyManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Unpause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Unpause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// ContractStrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ContractStrategyManager contract.
type ContractStrategyManagerDepositIterator struct {
	Event *ContractStrategyManagerDeposit // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerDeposit)
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
		it.Event = new(ContractStrategyManagerDeposit)
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
func (it *ContractStrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerDeposit represents a Deposit event raised by the ContractStrategyManager contract.
type ContractStrategyManagerDeposit struct {
	Depositor common.Address
	Token     common.Address
	Strategy  common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address depositor, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*ContractStrategyManagerDepositIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerDepositIterator{contract: _ContractStrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address depositor, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerDeposit)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address depositor, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseDeposit(log types.Log) (*ContractStrategyManagerDeposit, error) {
	event := new(ContractStrategyManagerDeposit)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractStrategyManager contract.
type ContractStrategyManagerInitializedIterator struct {
	Event *ContractStrategyManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerInitialized)
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
		it.Event = new(ContractStrategyManagerInitialized)
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
func (it *ContractStrategyManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerInitialized represents a Initialized event raised by the ContractStrategyManager contract.
type ContractStrategyManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractStrategyManagerInitializedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerInitializedIterator{contract: _ContractStrategyManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerInitialized)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseInitialized(log types.Log) (*ContractStrategyManagerInitialized, error) {
	event := new(ContractStrategyManagerInitialized)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractStrategyManager contract.
type ContractStrategyManagerOwnershipTransferredIterator struct {
	Event *ContractStrategyManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerOwnershipTransferred)
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
		it.Event = new(ContractStrategyManagerOwnershipTransferred)
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
func (it *ContractStrategyManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractStrategyManager contract.
type ContractStrategyManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractStrategyManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerOwnershipTransferredIterator{contract: _ContractStrategyManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerOwnershipTransferred)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractStrategyManagerOwnershipTransferred, error) {
	event := new(ContractStrategyManagerOwnershipTransferred)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractStrategyManager contract.
type ContractStrategyManagerPausedIterator struct {
	Event *ContractStrategyManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerPaused)
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
		it.Event = new(ContractStrategyManagerPaused)
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
func (it *ContractStrategyManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerPaused represents a Paused event raised by the ContractStrategyManager contract.
type ContractStrategyManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerPausedIterator{contract: _ContractStrategyManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerPaused)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParsePaused(log types.Log) (*ContractStrategyManagerPaused, error) {
	event := new(ContractStrategyManagerPaused)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractStrategyManager contract.
type ContractStrategyManagerPauserRegistrySetIterator struct {
	Event *ContractStrategyManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerPauserRegistrySet)
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
		it.Event = new(ContractStrategyManagerPauserRegistrySet)
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
func (it *ContractStrategyManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractStrategyManager contract.
type ContractStrategyManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractStrategyManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerPauserRegistrySetIterator{contract: _ContractStrategyManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerPauserRegistrySet)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractStrategyManagerPauserRegistrySet, error) {
	event := new(ContractStrategyManagerPauserRegistrySet)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerShareWithdrawalQueuedIterator is returned from FilterShareWithdrawalQueued and is used to iterate over the raw logs and unpacked data for ShareWithdrawalQueued events raised by the ContractStrategyManager contract.
type ContractStrategyManagerShareWithdrawalQueuedIterator struct {
	Event *ContractStrategyManagerShareWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerShareWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerShareWithdrawalQueued)
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
		it.Event = new(ContractStrategyManagerShareWithdrawalQueued)
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
func (it *ContractStrategyManagerShareWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerShareWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerShareWithdrawalQueued represents a ShareWithdrawalQueued event raised by the ContractStrategyManager contract.
type ContractStrategyManagerShareWithdrawalQueued struct {
	Depositor common.Address
	Nonce     *big.Int
	Strategy  common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterShareWithdrawalQueued is a free log retrieval operation binding the contract event 0xcf1c2370141bbd0a6d971beb0e3a2455f24d6e773ddc20ccc1c4e32f3dd9f9f7.
//
// Solidity: event ShareWithdrawalQueued(address depositor, uint96 nonce, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterShareWithdrawalQueued(opts *bind.FilterOpts) (*ContractStrategyManagerShareWithdrawalQueuedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "ShareWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerShareWithdrawalQueuedIterator{contract: _ContractStrategyManager.contract, event: "ShareWithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchShareWithdrawalQueued is a free log subscription operation binding the contract event 0xcf1c2370141bbd0a6d971beb0e3a2455f24d6e773ddc20ccc1c4e32f3dd9f9f7.
//
// Solidity: event ShareWithdrawalQueued(address depositor, uint96 nonce, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchShareWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerShareWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "ShareWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerShareWithdrawalQueued)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "ShareWithdrawalQueued", log); err != nil {
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

// ParseShareWithdrawalQueued is a log parse operation binding the contract event 0xcf1c2370141bbd0a6d971beb0e3a2455f24d6e773ddc20ccc1c4e32f3dd9f9f7.
//
// Solidity: event ShareWithdrawalQueued(address depositor, uint96 nonce, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseShareWithdrawalQueued(log types.Log) (*ContractStrategyManagerShareWithdrawalQueued, error) {
	event := new(ContractStrategyManagerShareWithdrawalQueued)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "ShareWithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *ContractStrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
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
		it.Event = new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
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
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _ContractStrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *ContractStrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
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
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _ContractStrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *ContractStrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyWhitelisterChanged)
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
		it.Event = new(ContractStrategyManagerStrategyWhitelisterChanged)
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
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyWhitelisterChangedIterator{contract: _ContractStrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyWhitelisterChanged)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*ContractStrategyManagerStrategyWhitelisterChanged, error) {
	event := new(ContractStrategyManagerStrategyWhitelisterChanged)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractStrategyManager contract.
type ContractStrategyManagerUnpausedIterator struct {
	Event *ContractStrategyManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerUnpaused)
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
		it.Event = new(ContractStrategyManagerUnpaused)
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
func (it *ContractStrategyManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerUnpaused represents a Unpaused event raised by the ContractStrategyManager contract.
type ContractStrategyManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerUnpausedIterator{contract: _ContractStrategyManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerUnpaused)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseUnpaused(log types.Log) (*ContractStrategyManagerUnpaused, error) {
	event := new(ContractStrategyManagerUnpaused)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the ContractStrategyManager contract.
type ContractStrategyManagerWithdrawalCompletedIterator struct {
	Event *ContractStrategyManagerWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerWithdrawalCompleted)
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
		it.Event = new(ContractStrategyManagerWithdrawalCompleted)
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
func (it *ContractStrategyManagerWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerWithdrawalCompleted represents a WithdrawalCompleted event raised by the ContractStrategyManager contract.
type ContractStrategyManagerWithdrawalCompleted struct {
	Depositor      common.Address
	Nonce          *big.Int
	Withdrawer     common.Address
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xe7eb0ca11b83744ece3d78e9be01b913425fbae70c32ce27726d0ecde92ef8d2.
//
// Solidity: event WithdrawalCompleted(address indexed depositor, uint96 nonce, address indexed withdrawer, bytes32 withdrawalRoot)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts, depositor []common.Address, withdrawer []common.Address) (*ContractStrategyManagerWithdrawalCompletedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "WithdrawalCompleted", depositorRule, withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerWithdrawalCompletedIterator{contract: _ContractStrategyManager.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xe7eb0ca11b83744ece3d78e9be01b913425fbae70c32ce27726d0ecde92ef8d2.
//
// Solidity: event WithdrawalCompleted(address indexed depositor, uint96 nonce, address indexed withdrawer, bytes32 withdrawalRoot)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerWithdrawalCompleted, depositor []common.Address, withdrawer []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "WithdrawalCompleted", depositorRule, withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerWithdrawalCompleted)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xe7eb0ca11b83744ece3d78e9be01b913425fbae70c32ce27726d0ecde92ef8d2.
//
// Solidity: event WithdrawalCompleted(address indexed depositor, uint96 nonce, address indexed withdrawer, bytes32 withdrawalRoot)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseWithdrawalCompleted(log types.Log) (*ContractStrategyManagerWithdrawalCompleted, error) {
	event := new(ContractStrategyManagerWithdrawalCompleted)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerWithdrawalDelayBlocksSetIterator is returned from FilterWithdrawalDelayBlocksSet and is used to iterate over the raw logs and unpacked data for WithdrawalDelayBlocksSet events raised by the ContractStrategyManager contract.
type ContractStrategyManagerWithdrawalDelayBlocksSetIterator struct {
	Event *ContractStrategyManagerWithdrawalDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerWithdrawalDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerWithdrawalDelayBlocksSet)
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
		it.Event = new(ContractStrategyManagerWithdrawalDelayBlocksSet)
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
func (it *ContractStrategyManagerWithdrawalDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerWithdrawalDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerWithdrawalDelayBlocksSet represents a WithdrawalDelayBlocksSet event raised by the ContractStrategyManager contract.
type ContractStrategyManagerWithdrawalDelayBlocksSet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayBlocksSet is a free log retrieval operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*ContractStrategyManagerWithdrawalDelayBlocksSetIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerWithdrawalDelayBlocksSetIterator{contract: _ContractStrategyManager.contract, event: "WithdrawalDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayBlocksSet is a free log subscription operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerWithdrawalDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerWithdrawalDelayBlocksSet)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
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
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseWithdrawalDelayBlocksSet(log types.Log) (*ContractStrategyManagerWithdrawalDelayBlocksSet, error) {
	event := new(ContractStrategyManagerWithdrawalDelayBlocksSet)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the ContractStrategyManager contract.
type ContractStrategyManagerWithdrawalQueuedIterator struct {
	Event *ContractStrategyManagerWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerWithdrawalQueued)
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
		it.Event = new(ContractStrategyManagerWithdrawalQueued)
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
func (it *ContractStrategyManagerWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerWithdrawalQueued represents a WithdrawalQueued event raised by the ContractStrategyManager contract.
type ContractStrategyManagerWithdrawalQueued struct {
	Depositor        common.Address
	Nonce            *big.Int
	Withdrawer       common.Address
	DelegatedAddress common.Address
	WithdrawalRoot   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x32cf9fc97155f52860a59a99879a2e89c1e53f28126a9ab6a2ff29344299e674.
//
// Solidity: event WithdrawalQueued(address depositor, uint96 nonce, address withdrawer, address delegatedAddress, bytes32 withdrawalRoot)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*ContractStrategyManagerWithdrawalQueuedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerWithdrawalQueuedIterator{contract: _ContractStrategyManager.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x32cf9fc97155f52860a59a99879a2e89c1e53f28126a9ab6a2ff29344299e674.
//
// Solidity: event WithdrawalQueued(address depositor, uint96 nonce, address withdrawer, address delegatedAddress, bytes32 withdrawalRoot)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerWithdrawalQueued)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
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

// ParseWithdrawalQueued is a log parse operation binding the contract event 0x32cf9fc97155f52860a59a99879a2e89c1e53f28126a9ab6a2ff29344299e674.
//
// Solidity: event WithdrawalQueued(address depositor, uint96 nonce, address withdrawer, address delegatedAddress, bytes32 withdrawalRoot)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseWithdrawalQueued(log types.Log) (*ContractStrategyManagerWithdrawalQueued, error) {
	event := new(ContractStrategyManagerWithdrawalQueued)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
