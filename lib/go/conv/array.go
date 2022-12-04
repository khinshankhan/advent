package conv

import "strconv"

func From1DStringTo1DInt(input []string) []int {
	ret := make([]int, len(input))
	for j, e := range input {
		num, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		ret[j] = num
	}

	return ret
}

func From2DStringTo2DInt(input [][]string) [][]int {
	ret := make([][]int, len(input))
	for i, row := range input {
		ret[i] = From1DStringTo1DInt(row)
	}

	return ret
}
