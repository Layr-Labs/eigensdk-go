package bls

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	bn254utils "github.com/Layr-Labs/eigensdk-go/crypto/bn254"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// We are using similar structure for saving bls keys as ethereum keystore
// https://github.com/ethereum/go-ethereum/blob/master/accounts/keystore/key.go
//
// We are storing PubKey sepearately so that we can list the pubkey without
// needing password to decrypt the private key
type encryptedBLSKeyJSONV3 struct {
	PubKey string              `json:"pubKey"`
	Crypto keystore.CryptoJSON `json:"crypto"`
}

func newFpElement(x *big.Int) fp.Element {
	var p fp.Element
	p.SetBigInt(x)
	return p
}

type G1Point struct {
	*bn254.G1Affine
}

func NewG1Point(x, y *big.Int) *G1Point {
	return &G1Point{
		&bn254.G1Affine{
			X: newFpElement(x),
			Y: newFpElement(y),
		},
	}
}

func NewZeroG1Point() *G1Point {
	return NewG1Point(big.NewInt(0), big.NewInt(0))
}

// Add another G1 point to this one
func (p *G1Point) Add(p2 *G1Point) *G1Point {
	p.G1Affine.Add(p.G1Affine, p2.G1Affine)
	return p
}

// Sub another G1 point from this one
func (p *G1Point) Sub(p2 *G1Point) *G1Point {
	p.G1Affine.Sub(p.G1Affine, p2.G1Affine)
	return p
}

// VerifyEquivalence verifies G1Point is equivalent the G2Point
func (p *G1Point) VerifyEquivalence(p2 *G2Point) (bool, error) {
	return bn254utils.CheckG1AndG2DiscreteLogEquality(p.G1Affine, p2.G2Affine)
}

func (p *G1Point) Serialize() []byte {
	return bn254utils.SerializeG1(p.G1Affine)
}

func (p *G1Point) Deserialize(data []byte) *G1Point {
	return &G1Point{bn254utils.DeserializeG1(data)}
}

type G2Point struct {
	*bn254.G2Affine
}

func NewG2Point(X, Y [2]*big.Int) *G2Point {
	return &G2Point{
		&bn254.G2Affine{
			X: struct{ A0, A1 fp.Element }{
				// TODO: why are 1 and 0 swapped here?
				A0: newFpElement(X[1]),
				A1: newFpElement(X[0]),
			},
			Y: struct{ A0, A1 fp.Element }{
				A0: newFpElement(Y[1]),
				A1: newFpElement(Y[0]),
			},
		},
	}
}

func NewZeroG2Point() *G2Point {
	return NewG2Point([2]*big.Int{big.NewInt(0), big.NewInt(0)}, [2]*big.Int{big.NewInt(0), big.NewInt(0)})
}

// Add another G2 point to this one
func (p *G2Point) Add(p2 *G2Point) *G2Point {
	p.G2Affine.Add(p.G2Affine, p2.G2Affine)
	return p
}

// Sub another G2 point from this one
func (p *G2Point) Sub(p2 *G2Point) *G2Point {
	p.G2Affine.Sub(p.G2Affine, p2.G2Affine)
	return p
}

func (p *G2Point) Serialize() []byte {
	return bn254utils.SerializeG2(p.G2Affine)
}

func (p *G2Point) Deserialize(data []byte) *G2Point {
	return &G2Point{bn254utils.DeserializeG2(data)}
}

type Signature struct {
	*G1Point `json:"g1_point"`
}

func NewZeroSignature() *Signature {
	return &Signature{NewZeroG1Point()}
}

func (s *Signature) Add(otherS *Signature) *Signature {
	s.G1Point.Add(otherS.G1Point)
	return s
}

// Verify a message against a public key
func (s *Signature) Verify(pubkey *G2Point, message [32]byte) (bool, error) {
	ok, err := bn254utils.VerifySig(s.G1Affine, pubkey.G2Affine, message)
	if err != nil {
		return false, err
	}
	return ok, nil
}

type PrivateKey = fr.Element

func NewPrivateKey(sk string) (*PrivateKey, error) {
	ele, err := new(fr.Element).SetString(sk)
	if err != nil {
		return nil, err
	}
	return ele, nil
}

type KeyPair struct {
	PrivKey *PrivateKey
	PubKey  *G1Point
}

func NewKeyPair(sk *PrivateKey) *KeyPair {
	pk := bn254utils.MulByGeneratorG1(sk)
	return &KeyPair{sk, &G1Point{pk}}
}

func NewKeyPairFromString(sk string) (*KeyPair, error) {
	ele, err := new(fr.Element).SetString(sk)
	if err != nil {
		return nil, err
	}
	return NewKeyPair(ele), nil
}

func GenRandomBlsKeys() (*KeyPair, error) {

	//Max random value is order of the curve
	max := new(big.Int)
	max.SetString(fr.Modulus().String(), 10)

	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}

	sk := new(PrivateKey).SetBigInt(n)
	return NewKeyPair(sk), nil
}

// SaveToFile saves the private key in an encrypted keystore file
func (k *KeyPair) SaveToFile(path string, password string) error {
	data, err := k.EncryptedString(path, password)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0750); err != nil {
		fmt.Println("Error creating directories:", err)
		return err
	}
	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (k *KeyPair) EncryptedString(path string, password string) ([]byte, error) {
	sk32Bytes := k.PrivKey.Bytes()
	skBytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		skBytes[i] = sk32Bytes[i]
	}

	cryptoStruct, err := keystore.EncryptDataV3(
		skBytes,
		[]byte(password),
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	if err != nil {
		return nil, err
	}

	encryptedBLSStruct := encryptedBLSKeyJSONV3{
		k.PubKey.String(),
		cryptoStruct,
	}
	data, err := json.Marshal(encryptedBLSStruct)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadPrivateKeyFromFile(path string, password string) (*KeyPair, error) {
	keyStoreContents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	encryptedBLSStruct := &encryptedBLSKeyJSONV3{}
	err = json.Unmarshal(keyStoreContents, encryptedBLSStruct)
	if err != nil {
		return nil, err
	}

	// Check if pubkey is present, if not return error
	// There is an issue where if you specify ecdsa key file
	// it still works and returns a keypair since the format of storage is same.
	// This is to prevent and make sure pubkey is present.
	// ecdsa keys doesn't have that field
	if encryptedBLSStruct.PubKey == "" {
		return nil, fmt.Errorf("invalid bls key file. pubkey field not found")
	}

	skBytes, err := keystore.DecryptDataV3(encryptedBLSStruct.Crypto, password)
	if err != nil {
		return nil, err
	}

	privKey := new(fr.Element).SetBytes(skBytes)
	keyPair := NewKeyPair(privKey)
	return keyPair, nil
}

// This signs a message on G1, and so will require a G2Pubkey to verify
func (k *KeyPair) SignMessage(message [32]byte) *Signature {
	H := bn254utils.MapToCurve(message)
	sig := new(bn254.G1Affine).ScalarMultiplication(H, k.PrivKey.BigInt(new(big.Int)))
	return &Signature{&G1Point{sig}}
}

// This signs a message on G1, and so will require a G2Pubkey to verify
func (k *KeyPair) SignHashedToCurveMessage(g1HashedMsg *bn254.G1Affine) *Signature {
	sig := new(bn254.G1Affine).ScalarMultiplication(g1HashedMsg, k.PrivKey.BigInt(new(big.Int)))
	return &Signature{&G1Point{sig}}
}

func (k *KeyPair) GetPubKeyG2() *G2Point {
	return &G2Point{bn254utils.MulByGeneratorG2(k.PrivKey)}
}

func (k *KeyPair) GetPubKeyG1() *G1Point {
	return k.PubKey
}
