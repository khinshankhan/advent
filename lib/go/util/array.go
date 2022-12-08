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

func Clone1D[T comparable](list []T) []T {
	ret := make([]T, len(list))
	for i, e := range list {
		ret[i] = e
	}
	return ret
}

func Clone2D[T comparable](list [][]T) [][]T {
	ret := make([][]T, len(list))
	for i, e := range list {
		ret[i] = Clone1D(e)
	}
	return ret
}

// waiting for polymorphism after generics...
func Map1D[T any, U any](fn func(T) U, list []T) []U {
	ret := make([]U, len(list))
	for i, e := range list {
		ret[i] = fn(e)
	}
	return ret
}

func Map1DWithIndex[T any, U any](fn func(T, int) U, list []T) []U {
	ret := make([]U, len(list))
	for i, e := range list {
		ret[i] = fn(e, i)
	}
	return ret
}

func Map1DWithIndexAndList[T any, U any](fn func(T, int, []T) U, list []T) []U {
	ret := make([]U, len(list))
	for i, e := range list {
		ret[i] = fn(e, i, list)
	}
	return ret
}

func Filter1D[T any](fn func(T) bool, list []T) []T {
	ret := []T{}
	for _, e := range list {
		if fn(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func Filter1DWithIndex[T any](fn func(T, int) bool, list []T) []T {
	ret := []T{}
	for i, e := range list {
		if fn(e, i) {
			ret = append(ret, e)
		}
	}
	return ret
}

func Filter1DWithIndexAndList[T any](fn func(T, int, []T) bool, list []T) []T {
	ret := []T{}
	for i, e := range list {
		if fn(e, i, list) {
			ret = append(ret, e)
		}
	}
	return ret
}

func Reduce1D[T any, U any](fn func(U, T) U, list []T, initial U) U {
	acc := initial
	for _, e := range list {
		acc = fn(acc, e)
	}
	return acc
}

func Reduce1DWithIndex[T any, U any](fn func(U, T, int) U, list []T, initial U) U {
	acc := initial
	for i, e := range list {
		acc = fn(acc, e, i)
	}
	return acc
}

func Reduce1DWithIndexAndList[T any, U any](fn func(U, T, int, []T) U, list []T, initial U) U {
	acc := initial
	for i, e := range list {
		acc = fn(acc, e, i, list)
	}
	return acc
}
