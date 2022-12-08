package util

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

func FromRuneToInt(r rune) int {
	return int(r - '0')
}

func FromStringToRunes(s string) []rune {
	return []rune(s)
}

func From2DTo1D[T any](matrix [][]T) []T {
	ret := []T{}
	for _, list := range matrix {
		ret = append(ret, list...)
	}
	return ret
}

func From1DStringTo1DInt(list []string) []int {
	return Map1D(FromStringToInt, list)
}

func From2DStringTo2DInt(list [][]string) [][]int {
	return Map1D(From1DStringTo1DInt, list)
}

func From1DStringTo2DRune(list []string) [][]rune {
	return Map1D(FromStringToRunes, list)
}
