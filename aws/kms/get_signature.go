package kms

import (
	"context"
	"encoding/asn1"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

type asn1EcSig struct {
	R asn1.RawValue
	S asn1.RawValue
}

// GetSignature retrieves the ECDSA signature for a message using a KMS key
func GetSignature(
	ctx context.Context, svc *kms.Client, keyId string, msg []byte,
) ([]byte, []byte, error) {
	signInput := &kms.SignInput{
		KeyId:            aws.String(keyId),
		SigningAlgorithm: types.SigningAlgorithmSpecEcdsaSha256,
		MessageType:      types.MessageTypeDigest,
		Message:          msg,
	}

	signOutput, err := svc.Sign(ctx, signInput)
	if err != nil {
		return nil, nil, err
	}

	var sigAsn1 asn1EcSig
	_, err = asn1.Unmarshal(signOutput.Signature, &sigAsn1)
	if err != nil {
		return nil, nil, err
	}

	return sigAsn1.R.Bytes, sigAsn1.S.Bytes, nil
}
