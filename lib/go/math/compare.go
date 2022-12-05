package math

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](nums ...T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}

	m := nums[0]
	for _, v := range nums {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T constraints.Ordered](nums ...T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}

	m := nums[0]
	for _, v := range nums {
		if m > v {
			m = v
		}
	}
	return m
}
