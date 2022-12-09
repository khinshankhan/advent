package math

import (
	"github.com/khinshankhan/advent/lib/go/types"
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Signed | constraints.Float](num T) T {
	if num < 0 {
		return num * T(-1)
	}
	return num
}

// https://en.wikipedia.org/wiki/Sign_function
func Sgn[T types.Number](num T) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}

func SumNums[T types.Number](nums ...T) T {
	sum := T(SumIdentity)
	for _, num := range nums {
		sum += num
	}
	return sum
}

func ProductNums[T types.Number](nums ...T) T {
	product := T(ProductIdentity)
	for _, num := range nums {
		product = product * num
	}
	return product
}
