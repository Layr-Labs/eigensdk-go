package examplecommon

import "math/big"

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
