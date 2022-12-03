package util

func Sum1DInt(input []int) int {
	sum := 0
	for _, e := range input {
		sum += e
	}
	return sum
}

func Intersection[T comparable](a, b []T) []T {
	intersection := []T{}

	aCounts := make(map[T]int)
	for _, e := range a {
		if count, ok := aCounts[e]; ok {
			aCounts[e] = count + 1
		} else {
			aCounts[e] = 1
		}
	}

	for _, e := range b {
		if count, ok := aCounts[e]; ok && count > 0 {
			intersection = append(intersection, e)
			aCounts[e] = count - 1
		}
	}

	return intersection
}

func Unique[T comparable](list []T) []T{
	m  := make(map[T]bool)
	for _, e := range list {
		m[e] = true
	}

	ret := make([]T, len(m))
	i:= 0
	for k := range m{
		ret[i] = k
		i++
	}

	return ret
}
