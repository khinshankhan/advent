package util

import (
	"github.com/khinshankhan/advent/lib/go/ds"
)

func ReverseMap[K comparable, V comparable](m map[K]V) map[V]K {
	n := make(map[V]K, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func ExistenceMap[T comparable](iterable []T) map[T]struct{} {
	return ds.NewSet(iterable...).GetMap()
}

func FrequencyMap[T comparable](iterable []T) map[T]int {
	return ds.NewFrequencyMap(iterable...).GetMap()
}

func CloneMap[K comparable, V any](m map[K]V, clone func(V) V) map[K]V {
	c := make(map[K]V, len(m))
	for k, v := range m {
		c[k] = clone(v)
	}
	return c
}
