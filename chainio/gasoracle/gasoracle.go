package gasoracle

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Params struct {
	FallbackGasTipCap          uint64
	GasMultiplierPercentage    uint64
	GasTipMultiplierPercentage uint64
}

var DefaultParams = Params{
	FallbackGasTipCap:          uint64(5_000_000_000), // 5 gwei
	GasMultiplierPercentage:    uint64(120),           // add an extra 20% gas buffer to the gas limit
	GasTipMultiplierPercentage: uint64(125),           // add an extra 25% to the gas tip
}

type GasOracle struct {
	params Params
	client eth.Client
	logger logging.Logger
}

// params are optional gas parameters any of which will be filled with default values if not provided
func New(client eth.Client, logger logging.Logger, params Params) *GasOracle {
	if params.FallbackGasTipCap == 0 {
		params.FallbackGasTipCap = DefaultParams.FallbackGasTipCap
	}
	if params.GasMultiplierPercentage == 0 {
		params.GasMultiplierPercentage = DefaultParams.GasMultiplierPercentage
	}
	if params.GasTipMultiplierPercentage == 0 {
		params.GasTipMultiplierPercentage = DefaultParams.GasTipMultiplierPercentage
	}
	return &GasOracle{
		params: params,
		client: client,
		logger: logger,
	}
}

func (o *GasOracle) GetLatestGasCaps(ctx context.Context) (gasTipCap, gasFeeCap *big.Int, err error) {
	gasTipCap, err = o.client.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		// Currently Alchemy is the only backend provider that exposes this
		// method, so in the event their API is unreachable we can fallback to a
		// degraded mode of operation. This also applies to our test
		// environments, as hardhat doesn't support the query either.
		o.logger.Info("eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		gasTipCap = big.NewInt(0).SetUint64(o.params.FallbackGasTipCap)
	}

	gasTipCap.Mul(gasTipCap, big.NewInt(int64(o.params.GasTipMultiplierPercentage))).Div(gasTipCap, big.NewInt(100))

	header, err := o.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	gasFeeCap = getGasFeeCap(gasTipCap, header.BaseFee)
	return
}

func (o *GasOracle) UpdateGas(
	ctx context.Context,
	tx *types.Transaction,
	value, gasTipCap, gasFeeCap *big.Int,
	from common.Address,
) (*types.Transaction, error) {
	gasLimit, err := o.client.EstimateGas(ctx, ethereum.CallMsg{
		From:      from,
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     value,
		Data:      tx.Data(),
	})

	if err != nil {
		return nil, err
	}

	noopSigner := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return tx, nil
	}
	opts := &bind.TransactOpts{
		From:   from,
		Signer: noopSigner,
		NoSend: true,
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = o.addGasBuffer(gasLimit)
	opts.Value = value

	contract := bind.NewBoundContract(*tx.To(), abi.ABI{}, o.client, o.client, o.client)
	return contract.RawTransact(opts, tx.Data())
}

func (o *GasOracle) addGasBuffer(gasLimit uint64) uint64 {
	return o.params.GasMultiplierPercentage * gasLimit / 100
}

// getGasFeeCap returns the gas fee cap for a transaction, calculated as:
// gasFeeCap = 2 * baseFee + gasTipCap
// Rationale: https://www.blocknative.com/blog/eip-1559-fees
func getGasFeeCap(gasTipCap *big.Int, baseFee *big.Int) *big.Int {
	return new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), gasTipCap)
}
