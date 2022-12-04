package util

import (
	"github.com/khinshankhan/advent/lib/go/ds"
)

func intersection2Lists[T comparable](a, b []T) []T {
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

func Intersection[T comparable](lists ...[]T) []T {
	intersection := lists[0]
	for _, list := range lists[1:] {
		intersection = intersection2Lists(intersection, list)
	}
	return intersection
}

func Unique[T comparable](list []T) []T {
	return ds.NewSet(list...).GetKeys()
}
