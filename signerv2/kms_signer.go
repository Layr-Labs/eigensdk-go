package signerv2

import (
	"context"
	"crypto/ecdsa"
	"encoding/asn1"
	"fmt"
	"math/big"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func KMSSignerFn(address common.Address, chainID *big.Int) (bind.SignerFn, error) {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return nil, errors.New("unimplemented")
	}, errors.New("unimplemented")
}

func getPubKey(ctx context.Context, svc *kms.Client, keyId string) (*ecdsa.PublicKey, error) {
	getPubKeyOutput, err := svc.GetPublicKey(ctx, &kms.GetPublicKeyInput{
		KeyId: aws.String(keyId),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get public key for KeyId=%s: %w", keyId, err)
	}

	var asn1pubk asn1EcPublicKey
	_, err = asn1.Unmarshal(getPubKeyOutput.PublicKey, &asn1pubk)
	if err != nil {
		return nil, errors.Wrapf(err, "can not parse asn1 public key for KeyId=%s", keyId)
	}

	pubkey, err := crypto.UnmarshalPubkey(asn1pubk.PublicKey.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "can not construct secp256k1 public key from key bytes")
	}

	return pubkey, nil
}
