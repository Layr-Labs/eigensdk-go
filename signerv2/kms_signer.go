package signerv2

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"

	eigenkms "github.com/Layr-Labs/eigensdk-go/aws/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var secp256k1N = crypto.S256().Params().N
var secp256k1HalfN = new(big.Int).Div(secp256k1N, big.NewInt(2))

func NewKMSSigner(ctx context.Context, svc *kms.Client, pk *ecdsa.PublicKey, keyId string, chainID *big.Int) SignerFn {
	return func(ctx context.Context, address common.Address) (bind.SignerFn, error) {
		return KMSSignerFn(ctx, svc, pk, keyId, chainID)
	}
}

// KMSSignerFn returns a SignerFn that uses a KMS key to sign transactions
// Heavily taken from https://github.com/welthee/go-ethereum-aws-kms-tx-signer
// It constructs R and S values from KMS, and constructs the recovery id (V) by trying to recover with both 0 and 1 values:
// ref: https://github.com/aws-samples/aws-kms-ethereum-accounts?tab=readme-ov-file#the-recovery-identifier-v
//
// Its V value is 0/1 instead of 27/28 because `types.LatestSignerForChainID` expects 0/1 which turns it into 27/28
func KMSSignerFn(ctx context.Context, svc *kms.Client, pk *ecdsa.PublicKey, keyId string, chainID *big.Int) (bind.SignerFn, error) {
	if chainID == nil {
		return nil, errors.New("chainID is required")
	}
	if svc == nil {
		return nil, errors.New("kms client is required")
	}
	if pk == nil {
		return nil, errors.New("public key is required")
	}

	pubKeyBytes := secp256k1.S256().Marshal(pk.X, pk.Y)
	keyAddr := crypto.PubkeyToAddress(*pk)
	signer := types.LatestSignerForChainID(chainID)
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != keyAddr {
			return nil, bind.ErrNotAuthorized
		}

		txHashBytes := signer.Hash(tx).Bytes()

		rBytes, sBytes, err := eigenkms.GetECDSASignature(ctx, svc, keyId, txHashBytes)
		if err != nil {
			return nil, err
		}

		// Adjust S value from signature according to Ethereum standard
		sBigInt := new(big.Int).SetBytes(sBytes)
		if sBigInt.Cmp(secp256k1HalfN) > 0 {
			sBytes = new(big.Int).Sub(secp256k1N, sBigInt).Bytes()
		}

		signature, err := getEthereumSignature(pubKeyBytes, txHashBytes, rBytes, sBytes)
		if err != nil {
			return nil, err
		}

		return tx.WithSignature(signer, signature)
	}, nil
}

func getEthereumSignature(expectedPublicKeyBytes []byte, txHash []byte, r []byte, s []byte) ([]byte, error) {
	rsSignature := append(adjustSignatureLength(r), adjustSignatureLength(s)...)
	signature := append(rsSignature, []byte{0}...)

	recoveredPublicKeyBytes, err := crypto.Ecrecover(txHash, signature)
	if err != nil {
		return nil, err
	}

	if hex.EncodeToString(recoveredPublicKeyBytes) != hex.EncodeToString(expectedPublicKeyBytes) {
		signature = append(rsSignature, []byte{1}...)
		recoveredPublicKeyBytes, err = crypto.Ecrecover(txHash, signature)
		if err != nil {
			return nil, err
		}

		if hex.EncodeToString(recoveredPublicKeyBytes) != hex.EncodeToString(expectedPublicKeyBytes) {
			return nil, errors.New("can not reconstruct public key from sig")
		}
	}

	return signature, nil
}

func adjustSignatureLength(buffer []byte) []byte {
	buffer = bytes.TrimLeft(buffer, "\x00")
	for len(buffer) < 32 {
		zeroBuf := []byte{0}
		buffer = append(zeroBuf, buffer...)
	}
	return buffer
}
