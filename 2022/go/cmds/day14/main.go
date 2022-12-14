package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sandMoves = []image.Point{
	// d, dl, dr for inverted grid
	{0, 1}, {-1, 1}, {1, 1},
}

func main() {
	s := io.ReadRelativeFile("../data/day14.txt")
	// s = sample
	input := parse(s)
	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parta(input Input) int {
	m := util.CloneMap(input.m, math.Identity[rune])
	input.m = m
	source := image.Point{500, 0}
	m[source] = '+'

	below := math.GetNeighbors(sandMoves, source)
	c, count := source, 0
	for c.Y <= input.maxY {
		neighbors := math.GetNeighbors(sandMoves, c)
		filled := true
		for _, neighbor := range neighbors {
			_, has := m[neighbor]
			if !has {
				filled = false
				c = math.ClonePoint(neighbor)
				break
			}
		}
		if filled {
			m[c] = 'o'
			c, count = source, count+1
		}

		filled = true
		for _, p := range below {
			_, has := m[p]
			if !has {
				filled = false
				break
			}
		}
		if filled {
			break
		}
	}

	debugGrid(input)
	return count
}

func partb(input Input) int {
	m, minX, minY, maxX, maxY := input.m, input.minX, input.minY, input.maxX, input.maxY
	source := image.Point{500, 0}
	m[source] = '+'

	below := math.GetNeighbors(sandMoves, source)
	c, count := source, 0
	for {
		neighbors := math.GetNeighbors(sandMoves, c)
		filled := true
		for _, neighbor := range neighbors {
			if neighbor.Y == input.maxY+2 {
				break
			}
			_, has := m[neighbor]
			if !has {
				filled = false
				c = math.ClonePoint(neighbor)
				if c.X < minX {
					minX = c.X
				}
				if c.X > maxX {
					maxX = c.X
				}
				if c.Y < minY {
					minY = c.Y
				}
				if c.Y > maxY {
					maxY = c.Y
				}
				break
			}
		}
		if filled {
			m[c] = 'o'
			c, count = source, count+1
		}

		filled = true
		for _, p := range below {
			_, has := m[p]
			if !has {
				filled = false
				break
			}
		}
		if filled {
			break
		}
	}

	input.minX, input.minY, input.maxX, input.maxY = minX-2, minY, maxX+2, maxY+1
	debugGrid(input)
	return count + 1
}

func debugGrid(input Input) {
	for y := input.minY; y < input.maxY+1; y++ {
		fmt.Printf("%02d: ", y)
		for x := input.minX; x < input.maxX; x++ {
			p := image.Point{x, y}
			v, has := input.m[p]
			if has {
				fmt.Printf(string(v))
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println(input.minX, input.minY, input.maxX, input.maxY)
}

type Input struct {
	m                      map[image.Point]rune
	minX, minY, maxX, maxY int
}

func parse(s string) Input {
	lines := strings.Split(s, "\n")

	m := make(map[image.Point]rune)
	var minX, minY, maxX, maxY int
	minX = 500
	for _, line := range lines {
		fmt.Println("line", line)
		coords := strings.Split(line, " -> ")
		for i := 1; i < len(coords); i++ {
			c1, c2 := strings.Split(coords[i-1], ","), strings.Split(coords[i], ",")
			x1, y1 := util.FromStringToInt(c1[0]), util.FromStringToInt(c1[1])
			x2, y2 := util.FromStringToInt(c2[0]), util.FromStringToInt(c2[1])
			p1, p2 := image.Point{x1, y1}, image.Point{x2, y2}

			diff := p2.Sub(p1)
			n := math.Max(math.Abs(diff.X), math.Abs(diff.Y))
			move := math.NormalizePoint(diff)
			c := p1
			for i := 0; i <= n; i++ {
				if i != 0 {
					c = c.Add(move)
				}
				m[c] = '#'
				if c.X < minX {
					minX = c.X
				}
				if c.X > maxX {
					maxX = c.X
				}
				if c.Y < minY {
					minY = c.Y
				}
				if c.Y > maxY {
					maxY = c.Y
				}
			}
		}
	}

	return Input{m, minX, minY, maxX, maxY}
}

var sample = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
