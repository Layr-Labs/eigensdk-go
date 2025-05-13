package examplecommon

import "math/big"

// This function computes the square of a number
func Square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
	numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)

	return numberSquared, nil
}

// BigIntEqual receives two *big.Int and returns whether they are equal
func BigIntEqual(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}
