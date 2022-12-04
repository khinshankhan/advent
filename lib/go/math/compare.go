package math

func MaxInt(nums ...int) int {
	max := nums[0]
	for _, e := range nums {
		if max < e {
			max = e
		}
	}

	return max
}

func MinInt(nums ...int) int {
	min := nums[0]
	for _, e := range nums {
		if min > e {
			min = e
		}
	}

	return min
}
