package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/khinshankhan/advent/lib/go/ds"
	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample1 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var sample2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func main() {
	s := io.ReadRelativeFile("../data/day09.txt")
	// s := sample1
	input := parse(s)
	fmt.Println(parta(input))

	// s = sample2
	// input = parse(s)
	fmt.Println(partb(input))
}

func parse(s string) [][]string {
	return util.TransformString2D(strings.TrimSpace(s), "\n", " ", true, true)
}

var moves = map[string]image.Point{
	"U": {0, 1},
	"D": {0, -1},
	"R": {1, 0},
	"L": {-1, 0},
}

func parta(input [][]string) int {
	h, t := image.Point{0, 0}, image.Point{0, 0}
	set := ds.NewSet(math.ClonePoint(t))
	for _, line := range input {
		d, n := line[0], util.FromStringToInt(line[1])
		for m := 0; m < n; m++ {
			prevH := math.ClonePoint(h)
			h = h.Add(moves[d])
			if math.ChebyshevDistance(h, t) > 1 {
				t = prevH
				set.Add(math.ClonePoint(t))
			}
		}
	}

	return set.Len()
}

func partb(input [][]string) int {
	rope := make([]image.Point, 10)
	set := ds.NewSet(math.ClonePoint(image.Point{}))
	for _, line := range input {
		d, n := line[0], util.FromStringToInt(line[1])
		for m := 0; m < n; m++ {
			rope[0] = rope[0].Add(moves[d])
			for i := 1; i < len(rope); i++ {
				if math.ChebyshevDistance(rope[i-1], rope[i]) > 1 {
					rope[i] = math.ChebyshevMove(rope[i], rope[i-1])
				}
			}
			set.Add(math.ClonePoint(rope[len(rope)-1]))
		}
	}

	return set.Len()
}
