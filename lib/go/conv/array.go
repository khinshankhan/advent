package conv

import (
	"strconv"
)

func FromStringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return num
}

func From1DStringTo1DInt(input []string) []int {
	ret := make([]int, len(input))
	for j, e := range input {
		ret[j] = FromStringToInt(e)
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

func From1DStringTo2DRune(input []string) [][]rune {
	ret := make([][]rune, len(input))
	for j, e := range input {
		ret[j] = []rune(e)
	}

	return ret
}

func From1DStringTo1DRune(input []string) []rune {
	ret := []rune{}
	for _, e := range input {
		runes := []rune(e)
		ret = append(ret, runes...)
	}

	return ret
}
