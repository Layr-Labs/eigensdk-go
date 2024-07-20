package fakes

import (
	"context"
	"math/big"

	apkregistrybindings "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type TestOperator struct {
	OperatorAddr     common.Address
	OperatorInfo     types.OperatorInfo
	ContractG1Pubkey apkregistrybindings.BN254G1Point
	ContractG2Pubkey apkregistrybindings.BN254G2Point
	OperatorId       types.OperatorId
}

type FakeAVSRegistryReader struct {
	opAddress  []types.OperatorAddr
	opPubKeys  []types.OperatorPubkeys
	operatorId types.OperatorId
	socket     types.Socket
	err        error
}

func NewFakeAVSRegistryReader(
	opr *TestOperator,
	err error,
) *FakeAVSRegistryReader {
	if opr == nil {
		return &FakeAVSRegistryReader{}
	}
	return &FakeAVSRegistryReader{
		opAddress:  []common.Address{opr.OperatorAddr},
		opPubKeys:  []types.OperatorPubkeys{opr.OperatorInfo.Pubkeys},
		socket:     opr.OperatorInfo.Socket,
		operatorId: opr.OperatorId,
		err:        err,
	}
}

func (f *FakeAVSRegistryReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {
	return f.opAddress, f.opPubKeys, f.err
}

func (f *FakeAVSRegistryReader) QueryExistingRegisteredOperatorSockets(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) (map[types.OperatorId]types.Socket, error) {
	if len(f.opPubKeys) == 0 {
		return nil, nil
	}

	return map[types.OperatorId]types.Socket{
		types.OperatorIdFromG1Pubkey(f.opPubKeys[0].G1Pubkey): f.socket,
	}, nil
}

func (f *FakeAVSRegistryReader) GetOperatorFromId(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (common.Address, error) {
	return f.opAddress[0], f.err
}

func (f *FakeAVSRegistryReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
	blockNumber types.BlockNum,
) ([][]opstateretrievar.OperatorStateRetrieverOperator, error) {
	return [][]opstateretrievar.OperatorStateRetrieverOperator{
		{
			{
				OperatorId: f.operatorId,
				Stake:      big.NewInt(123),
			},
		},
	}, nil
}

func (f *FakeAVSRegistryReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers types.QuorumNums,
	nonSignerOperatorIds []types.OperatorId,
) (opstateretrievar.OperatorStateRetrieverCheckSignaturesIndices, error) {
	return opstateretrievar.OperatorStateRetrieverCheckSignaturesIndices{}, nil
}
