package ds

type FrequencyMap[T comparable] struct {
	m map[T]int
}

func NewFrequencyMap[T comparable](initial ...T) *FrequencyMap[T] {
	fm := FrequencyMap[T]{}

	fm.m = make(map[T]int)
	for _, e := range initial {
		fm.Add(e)
	}

	return &fm
}

func (fm *FrequencyMap[T]) Add(key T) int {
	count, has := fm.Get(key)
	if has {
		fm.m[key] = count + 1
	} else {
		fm.m[key] = 1
	}

	return fm.m[key]
}

func (fm *FrequencyMap[T]) Subtract(key T) int {
	count, has := fm.Get(key)
	if has {
		fm.m[key] = count - 1
	} else {
		fm.m[key] = -1
	}

	return fm.m[key]
}

func (fm *FrequencyMap[T]) SubtractOrDelete(key T) int {
	count := fm.Subtract(key)
	if count < 1 {
		fm.Remove(key)
		return -1
	}

	return count
}

func (fm *FrequencyMap[T]) Remove(key T) {
	delete(fm.m, key)
}

func (fm *FrequencyMap[T]) Get(key T) (int, bool) {
	count, ok := fm.m[key]
	return count, ok
}

func (fm *FrequencyMap[T]) GetMap() map[T]int {
	return fm.m
}

func (fm *FrequencyMap[T]) Len() int {
	return len(fm.m)
}
