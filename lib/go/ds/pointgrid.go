package ds

import (
	"fmt"
	"image"
)

type PointGrid[T comparable] struct {
	Grid                   map[image.Point]T
	h, w                   int
	minX, minY, maxX, maxY int
}

func NewPointGrid[T comparable](h, w int) *PointGrid[T] {
	ptg := PointGrid[T]{}
	if w > -1 && h > -1 {
		ptg.Grid = make(map[image.Point]T, w*h)
	} else {
		ptg.Grid = make(map[image.Point]T)
	}
	ptg.h, ptg.w = h, w

	return &ptg
}

func (ptg *PointGrid[T]) Set(pt image.Point, v T) {
	if pt.X < ptg.minX {
		ptg.minX = pt.X
	}
	if pt.X > ptg.maxX {
		ptg.maxX = pt.X
	}
	if pt.Y < ptg.minY {
		ptg.minY = pt.Y
	}
	if pt.Y > ptg.maxY {
		ptg.maxY = pt.Y
	}

	ptg.Grid[pt] = v
}

func (ptg *PointGrid[T]) Has(pt image.Point) bool {
	_, ok := ptg.Grid[pt]
	return ok
}

func (ptg *PointGrid[T]) Get(pt image.Point) T {
	v, _ := ptg.Grid[pt]
	return v
}

func (ptg *PointGrid[T]) Len() int {
	return len(ptg.Grid)
}

func (ptg *PointGrid[T]) GetNeighbors(moves []image.Point, p image.Point) []image.Point {
	neighbors := []image.Point{}
	for _, move := range moves {
		neighbor := p.Add(move)
		if ptg.Has(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func (ptg *PointGrid[T]) Print(
	stringify func(T) string,
	defaultStr func(p image.Point) string,
) {
	for y := ptg.minY; y < ptg.maxY; y++ {
		for x := ptg.minX; x < ptg.maxX; x++ {
			point := image.Point{x, y}
			if ptg.Has(point) {
				fmt.Printf(stringify(ptg.Get(point)))
			} else {
				fmt.Printf(defaultStr(point))
			}
		}
		fmt.Printf("\n")
	}
}

// sample debug stringify fns
// stringifyNum := func(n int) string { return fmt.Sprintf("%02d ", n) }
// stringifyLetter := func(n int) string { return string(rune(n + int('a') - 1)) }
