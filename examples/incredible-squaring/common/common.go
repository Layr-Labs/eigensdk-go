package examplecommon

import (
	"math/big"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// This function computes the square of a number
func Square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
	numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)

	return numberSquared, nil
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
