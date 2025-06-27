package common

import "math/big"

// BigIntEqual receives two *big.Int and returns whether they are equal
func BigIntEqual(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}
