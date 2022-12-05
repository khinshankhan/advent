package math

import (
	"golang.org/x/exp/constraints"
)

func SumNums[T constraints.Ordered](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
