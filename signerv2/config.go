package signerv2

import "crypto/ecdsa"

type Config struct {
	PrivateKey   *ecdsa.PrivateKey
	KeystorePath string
	Password     string
	Endpoint     string
	Address      string
}

func (c Config) IsPrivateKeySigner() bool {
	return c.PrivateKey != nil
}

func (c Config) IsLocalKeystoreSigner() bool {
	return c.KeystorePath != ""
}

func (c Config) IsRemoteSigner() bool {
	if c.Endpoint == "" || c.Address == "" {
		return false
	}
	return true
}
