package circularslice

type CircularSlice[T comparable] struct {
	slice []T
}

func New[T comparable](initial ...T) *CircularSlice[T] {
	cs := CircularSlice[T]{}

	cs.slice = make([]T, len(initial))
	for i, e := range initial {
		cs.slice[i] = e
	}

	return &cs
}

func (cs *CircularSlice[T]) Append(e T) {
	cs.slice = append(cs.slice, e)
}

func (cs *CircularSlice[T]) FindIndex(fn func(T) bool) int {
	for i, e := range cs.slice {
		if fn(e) {
			return i
		}
	}

	return -1
}

func (cs *CircularSlice[T]) NextIndex(i int) int {
	nextI := i + 1
	if nextI >= len(cs.slice) {
		return 0
	}
	return nextI
}

func (cs *CircularSlice[T]) PrevIndex(i int) int {
	prevI := i - 1
	if prevI < 0 {
		return len(cs.slice) - 1
	}
	return prevI
}

func (cs *CircularSlice[T]) Next(i int) T {
	nextI := cs.NextIndex(i)
	return cs.slice[nextI]
}

func (cs *CircularSlice[T]) Prev(i int) T {
	prevI := cs.PrevIndex(i)
	return cs.slice[prevI]
}
