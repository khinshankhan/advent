package ds

type Set[T comparable] struct {
	m      map[T]struct{}
	keys   []T
	synced bool
}

func NewSet[T comparable](initial ...T) *Set[T] {
	set := Set[T]{}

	set.m = make(map[T]struct{}, len(initial))
	for _, e := range initial {
		set.m[e] = struct{}{}
	}

	set.synced = false

	return &set
}

func (set *Set[T]) Add(key T) {
	set.m[key] = struct{}{}
	set.synced = false
}

func (set *Set[T]) Remove(key T) {
	delete(set.m, key)
	set.synced = false
}

func (set *Set[T]) Has(key T) bool {
	_, ok := set.m[key]
	return ok
}

func (set *Set[T]) Len() int {
	return len(set.m)
}

func (set *Set[T]) GetKeys() []T {
	if set.synced {
		return set.keys
	}

	keys := make([]T, len(set.m))
	i := 0
	for key := range set.m {
		keys[i] = key
		i++
	}

	set.keys = keys
	set.synced = true

	return keys
}

func (set *Set[T]) GetMap() map[T]struct{} {
	return set.m
}
