package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/khinshankhan/advent/lib/go/ds"
	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
)

func main() {
	s := io.ReadRelativeFile("../data/day12.txt")
	// s = sample
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parta(input Input) int {
	return calc(input, false)
}

func partb(input Input) int {
	return calc(input, true)
}

func calc(input Input, part2 bool) int {
	// small trick, it's probably easier to go end to start because of constraints
	// even better for part 2 since we dont need to complete getting to the end
	start, end, dist := input.E, input.S, make(map[image.Point]int, input.Grid.Len())
	queue, seen := []image.Point{start}, ds.NewSet(start)

	shortestStart := start
	startVal := runePoint('a')
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if part2 && startVal == input.Grid.Get(c) {
			shortestStart = c
			break
		}
		neighbors := input.Grid.GetNeighbors(math.ManhattanMoves, c)
		for _, neighbor := range neighbors {
			if !seen.Has(neighbor) && input.IsTravserable(c, neighbor) {
				queue, dist[neighbor] = append(queue, neighbor), dist[c]+1
				seen.Add(neighbor)
			}

			// short circuit, minor optimization
			if neighbor.Eq(end) {
				break
			}

		}
	}

	if part2 {
		return dist[shortestStart]
	}
	return dist[end]
}

type Input struct {
	Grid *ds.PointGrid[int]
	S, E image.Point
}

func (input *Input) IsTravserable(p1, p2 image.Point) bool {
	return input.Grid.Get(p1) < input.Grid.Get(p2)+2
}

func runePoint(r rune) int {
	switch r {
	case 'S':
		return 0
	case 'E':
		return 27
	}

	return int(r-'a') + 1
}

func parse(s string) (ret Input) {
	alphaHeights := strings.Fields(s)
	ret.Grid = ds.NewPointGrid[int](len(alphaHeights), len(alphaHeights[0]))

	for y, line := range alphaHeights {
		for x, r := range line {
			point := image.Point{x, y}
			switch r {
			case 'S':
				ret.S = point
			case 'E':
				ret.E = point
			}
			ret.Grid.Set(point, runePoint(r))
		}
	}

	return
}

var sample = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
