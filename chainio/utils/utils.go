package utils

import (
	"math/big"

	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/BLSPublicKeyCompendium"
	"github.com/Layr-Labs/eigensdk-go-master/crypto/bls"
)

// BINDING UTILS - conversion from contract structs to golang structs

// BN254.sol is a library, so bindings for G1 Points and G2 Points are only generated
// in every contract that imports that library. Thus the output here will need to be
// type casted if G1Point is needed to interface with another contract (eg: BLSPublicKeyCompendium.sol)
func ConvertToBN254G1Point(input *bls.G1Point) blspubkeycompendium.BN254G1Point {
	output := blspubkeycompendium.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) blspubkeycompendium.BN254G2Point {
	output := blspubkeycompendium.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
