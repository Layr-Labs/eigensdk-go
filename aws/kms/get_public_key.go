package kms

import (
	"context"
	"crypto/ecdsa"
	"encoding/asn1"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/ethereum/go-ethereum/crypto"
)

type asn1EcPublicKey struct {
	EcPublicKeyInfo asn1EcPublicKeyInfo
	PublicKey       asn1.BitString
}

type asn1EcPublicKeyInfo struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.ObjectIdentifier
}

// GetECDSAPublicKey retrieves the ECDSA public key for a KMS key
// It assumes the key is set up with `ECC_SECG_P256K1` key spec and `SIGN_VERIFY` key usage
func GetECDSAPublicKey(ctx context.Context, svc *kms.Client, keyId string) (*ecdsa.PublicKey, error) {
	getPubKeyOutput, err := svc.GetPublicKey(ctx, &kms.GetPublicKeyInput{
		KeyId: aws.String(keyId),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get public key for KeyId=%s: %w", keyId, err)
	}

	var asn1pubk asn1EcPublicKey
	_, err = asn1.Unmarshal(getPubKeyOutput.PublicKey, &asn1pubk)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key for KeyId=%s: %w", keyId, err)
	}

	pubkey, err := crypto.UnmarshalPubkey(asn1pubk.PublicKey.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key for KeyId=%s: %w", keyId, err)
	}

	return pubkey, nil
}
