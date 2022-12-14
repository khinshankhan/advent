package main

import (
	"encoding/json"
	"fmt"
	"image"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/util"
)

const (
	Rock   = '#'
	Sand   = 'o'
	Source = '+'
	Air    = '.'
)

var (
	SourcePoint = image.Point{500, 0}
	sandMoves   = []image.Point{
		// d, dl, dr for inverted grid
		{0, 1}, {-1, 1}, {1, 1},
	}
)

func main() {
	s := io.ReadRelativeFile("../data/day14.txt")
	input := parse(s)
	// since part b reuses input
	m := util.CloneMap(input.m, math.Identity[rune])
	input2 := input
	input.m = m

	fmt.Println(parta(input))
	fmt.Println(partb(input2))
}

func parta(input Input) int {
	curr, count, floor := SourcePoint, 0, input.maxY+1
	for curr.Y < floor {
		neighbors, filled := math.GetNeighbors(sandMoves, curr), true
		for _, neighbor := range neighbors {
			_, has := input.m[neighbor]
			if !has {
				filled = false
				curr = math.ClonePoint(neighbor)
				break
			}
		}

		if filled {
			input.m[curr], count = Sand, count+1 // input.UpdateBounds(curr) // debugging
			if curr.Eq(SourcePoint) {
				break
			}
			curr = SourcePoint
		}
	} // debugGrid(input.m, input.minX-1, input.minY, input.maxX+1, floor)

	return count
}

func partb(input Input) int {
	source := image.Point{500, 0}
	input.m[SourcePoint], input.maxY = Source, input.maxY+2

	curr, count, floor := source, 0, input.maxY
	for {
		neighbors, filled := math.GetNeighbors(sandMoves, curr), true
		for _, neighbor := range neighbors {
			if neighbor.Y == floor {
				break
			}
			_, has := input.m[neighbor]
			if !has {
				filled = false
				curr = math.ClonePoint(neighbor)
				break
			}
		}

		if filled {
			input.m[curr], count = Sand, count+1 // input.UpdateBounds(curr) // debugging
			if curr.Eq(source) {
				break
			}
			curr = source
		}
	} // debugGrid(input.m, input.minX-2, input.minY, input.maxX+2, floor)

	return count
}

type Input struct {
	m                      map[image.Point]rune
	minX, minY, maxX, maxY int
}

func (i *Input) UpdateBounds(p image.Point) {
	if p.X < i.minX {
		i.minX = p.X
	}
	if p.X > i.maxX {
		i.maxX = p.X
	}
	if p.Y < i.minY {
		i.minY = p.Y
	}
	if p.Y > i.maxY {
		i.maxY = p.Y
	}
}

func parse(s string) Input {
	cleanedStr := "[[{\"X\":" +
		strings.Join(
			strings.Fields(
				strings.NewReplacer(
					" -> ", "},{\"X\":",
					",", ",\"Y\":",
				).Replace(strings.TrimSpace(s))),
			"}],[{\"X\":") +
		"}]]"
	var lines [][]image.Point
	json.Unmarshal([]byte(cleanedStr), &lines)

	m := make(map[image.Point]rune)
	input := Input{m, 500, 0, 0, 0}
	for _, line := range lines {
		for i := 1; i < len(line); i++ {
			p1, p2 := line[i-1], line[i]
			diff := p2.Sub(p1)
			n := math.Max(math.Abs(diff.X), math.Abs(diff.Y))
			move := math.NormalizePoint(diff)
			curr := p1
			for i := 0; i <= n; i++ {
				input.UpdateBounds(curr) // helpful to see 'floor'
				input.m[curr] = Rock
				curr = curr.Add(move)
			}
		}
	}
	input.m[SourcePoint] = Source

	return input
}

func debugGrid(m map[image.Point]rune, minX, minY, maxX, maxY int) {
	for y := minY; y < maxY+1; y++ {
		fmt.Printf("%02d: ", y)
		for x := minX; x < maxX; x++ {
			p := image.Point{x, y}
			v, has := m[p]
			if has {
				fmt.Printf(string(v))
			} else {
				fmt.Printf(string(Air))
			}
		}
		fmt.Printf("\n")
	}
}

var sample = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
