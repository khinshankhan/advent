package math

import (
	"github.com/khinshankhan/advent/lib/go/types"
)

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
