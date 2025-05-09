package types

type SignerType string

const (
	// Local signer type
	Local SignerType = "local"
	// Cerberus signer type
	Cerberus SignerType = "cerberus"
	// PrivateKey signer type
	PrivateKey SignerType = "privatekey"
)

type SignerConfig struct {
	// PrivateKey is the private key of the signer
	PrivateKey string

	// Type of the signer
	SignerType SignerType

	// Params for local signer
	// Path to the key file
	Path string
	// Password to decrypt the key file
	Password string

	// Params for cerberus signer
	// CerberusUrl of the cerberus signer
	CerberusUrl string
	// PublicKeyHex is the public key of the cerberus signer
	PublicKeyHex string
	// CerberusPassword is the password to encrypt the key if cerberus using local keystore
	CerberusPassword string
	// EnableTLS enables TLS for the cerberus signer
	EnableTLS bool
	// TLSCertFilePath is the path to the TLS cert file
	TLSCertFilePath string
	// CerberusAPIKey is the API key for the cerberus signer
	CerberusAPIKey string
}
