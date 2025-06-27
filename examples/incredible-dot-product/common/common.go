package examplecommon

import (
	"math/big"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

// This function computes the dot product of a pair of points
func DotProduct(taskIndex uint32, points DotProductInput) (*big.Int, error) {
	totalSum := big.NewInt(0)
	for i := range points.X {
		currentSum := big.NewInt(0).Mul(points.X[i], points.Y[i])
		totalSum.Add(totalSum, currentSum)
	}

	return totalSum, nil
}

// This function reads the config from the .toml file at the path received as a parameter
// and returns a config with those values
func ReadTomlConfig(path string, config any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(data, config)
	return err
}
