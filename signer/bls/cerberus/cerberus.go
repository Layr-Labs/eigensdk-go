package cerberus

import (
	"context"
	"encoding/hex"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	sdkBls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/consensys/gnark-crypto/ecc/bn254"

	v1 "github.com/Layr-Labs/cerberus-api/pkg/api/v1"
)

type Config struct {
	URL          string
	PublicKeyHex string

	SignerAPIKey string

	// Optional: in case if your signer uses local keystore
	Password string

	EnableTLS       bool
	TLSCertFilePath string

	SigningTimeout time.Duration
}

type Signer struct {
	signerClient v1.SignerClient
	kmsClient    v1.KeyManagerClient
	pubKeyHex    string
	password     string
	signerAPIKey string
}

func New(cfg Config) (Signer, error) {
	opts := make([]grpc.DialOption, 0)
	if cfg.EnableTLS {
		creds, err := credentials.NewClientTLSFromFile(cfg.TLSCertFilePath, "")
		if err != nil {
			return Signer{}, utils.WrapError("could not load tls cert", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(cfg.URL, opts...)
	if err != nil {
		return Signer{}, utils.WrapError("did not connect", err)
	}

	signerClient := v1.NewSignerClient(conn)
	kmsClient := v1.NewKeyManagerClient(conn)
	return Signer{
		signerClient: signerClient,
		kmsClient:    kmsClient,
		pubKeyHex:    cfg.PublicKeyHex,
		password:     cfg.Password,
		signerAPIKey: cfg.SignerAPIKey,
	}, nil
}

func (s Signer) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 32 {
		return nil, types.ErrInvalidMessageLength
	}

	// Pass the API key to the signer client
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", s.signerAPIKey)

	resp, err := s.signerClient.SignGeneric(ctx, &v1.SignGenericRequest{
		Data:        msg,
		PublicKeyG1: s.pubKeyHex,
		Password:    s.password,
	})
	if err != nil {
		return nil, err
	}

	return resp.Signature, nil
}

func (s Signer) SignG1(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 64 {
		return nil, types.ErrInvalidMessageLength
	}

	// Pass the API key to the signer client
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", s.signerAPIKey)

	resp, err := s.signerClient.SignG1(ctx, &v1.SignG1Request{
		Data:        msg,
		PublicKeyG1: s.pubKeyHex,
		Password:    s.password,
	})
	if err != nil {
		return nil, err
	}

	return resp.Signature, nil
}

func (s Signer) GetOperatorId() (string, error) {
	pkBytes, err := hex.DecodeString(s.pubKeyHex)
	if err != nil {
		return "", utils.WrapError("failed to decode BLS public key", err)
	}
	var point bn254.G1Affine
	_, err = point.SetBytes(pkBytes)
	if err != nil {
		return "", utils.WrapError("failed to set BLS public key", err)
	}
	pubkey := &sdkBls.G1Point{G1Affine: &point}
	return pubkey.GetOperatorID(), nil
}

func (s Signer) GetPublicKeyG1() string {
	return s.pubKeyHex
}

func (s Signer) GetPublicKeyG2() string {
	resp, err := s.kmsClient.GetKeyMetadata(context.Background(), &v1.GetKeyMetadataRequest{
		PublicKeyG1: s.pubKeyHex,
	})
	if err != nil {
		return ""
	}

	return resp.PublicKeyG2
}
