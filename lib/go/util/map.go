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
	freqs := make(map[T]int)
	for _, e := range iterable {
		if count, ok := freqs[e]; ok {
			freqs[e] = count + 1
		} else {
			freqs[e] = 1
		}
	}

	return freqs
}

func CloneMap[K comparable, V any](m map[K]V, clone func(V) V) map[K]V {
	c := make(map[K]V, len(m))
	for k, v := range m {
		c[k] = clone(v)
	}
	return c
}
