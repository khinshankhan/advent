package util

import (
	"github.com/khinshankhan/advent/lib/go/ds/set"
)

func Sum1DInt(input []int) int {
	sum := 0
	for _, e := range input {
		sum += e
	}
	return sum
}

func Intersection[T comparable](a, b []T) []T {
	aFreqs := FrequencyMap(a)

	intersection := []T{}
	for _, e := range b {
		if count, ok := aFreqs[e]; ok && count > 0 {
			intersection = append(intersection, e)
			aFreqs[e] = count - 1
		}
	}

	return intersection
}

func Unique[T comparable](list []T) []T {
	s := set.New(list...)
	return s.GetKeys()

}
