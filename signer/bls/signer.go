package bls

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/signer/bls/cerberus"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/local"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/privatekey"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/types"
)

type Signer interface {
	// Sign signs the message using the BLS signature scheme
	Sign(ctx context.Context, msg []byte) ([]byte, error)

	// SignG1 signs the message using the BLS signature scheme
	// This message is assumed to be already mapped to G1 point
	SignG1(ctx context.Context, msg []byte) ([]byte, error)

	// GetOperatorId returns the operator ID of the signer.
	// This is hash of the G1 public key of the signer
	GetOperatorId() (string, error)

	GetPublicKeyG1() string

	GetPublicKeyG2() string
}

// NewSigner creates a new Signer instance based on the provided configuration.
// It supports different types of signers, such as Local and Cerberus.
//
// Parameters:
//   - cfg: A SignerConfig struct that contains the configuration for the signer.
//
// Returns:
//   - Signer: An instance of the Signer interface.
//   - error: An error if the signer type is invalid or if there is an issue creating the signer.
func NewSigner(cfg types.SignerConfig) (Signer, error) {
	switch cfg.SignerType {
	case types.Local:
		return local.New(local.Config{
			Path:     cfg.Path,
			Password: cfg.Password,
		})
	case types.Cerberus:
		return cerberus.New(cerberus.Config{
			URL:             cfg.CerberusUrl,
			PublicKeyHex:    cfg.PublicKeyHex,
			Password:        cfg.CerberusPassword,
			EnableTLS:       cfg.EnableTLS,
			TLSCertFilePath: cfg.TLSCertFilePath,
			SignerAPIKey:    cfg.CerberusAPIKey,
		})
	case types.PrivateKey:
		return privatekey.New(privatekey.Config{
			PrivateKey: cfg.PrivateKey,
		})
	default:
		return nil, types.ErrInvalidSignerType
	}
}
