package main

import (
	"fmt"
	"image"
	"strings"

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
	start, end := input.E, input.S
	queue, dist := []image.Point{start}, make(map[image.Point]int, len(input.Grid))

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if part2 && 'a' == input.Grid[c] {
			end = c
			break
		}
		neighbors := math.GetNeighbors(math.ManhattanMoves, c)
		for _, neighbor := range neighbors {
			_, has := input.Grid[neighbor]
			_, seen := dist[neighbor]
			if has && !seen && input.Grid[c] < input.Grid[neighbor]+2 {
				queue, dist[neighbor] = append(queue, neighbor), dist[c]+1
			}

			// short circuit, minor optimization
			if neighbor.Eq(end) {
				break
			}

		}
	}

	return dist[end]
}

type Input struct {
	Grid map[image.Point]rune
	S, E image.Point
}

func parse(s string) (ret Input) {
	alphaHeights := strings.Fields(s)
	ret.Grid = make(map[image.Point]rune, len(alphaHeights)*len(alphaHeights[0]))

	for y, line := range alphaHeights {
		for x, r := range line {
			point := image.Point{x, y}
			ret.Grid[point] = r
			switch r {
			case 'S':
				ret.S = point
				ret.Grid[point] = 'a'
			case 'E':
				ret.E = point
				ret.Grid[point] = 'z' + 1
			}
		}
	}

	return
}

var sample = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
