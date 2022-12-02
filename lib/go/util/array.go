package util

func Sum1DInt(input []int) int {
	sum := 0
	for _, e := range input {
		sum += e
	}
	return sum
}
