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

func main() {
	s := io.ReadRelativeFile("../data/day09.txt")
	input := parse(s)

	fmt.Println(parta(input))
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
	return calcUniqueTailPositions(2, input)
}

func partb(input [][]string) int {
	return calcUniqueTailPositions(10, input)
}

func calcUniqueTailPositions(ropeLength int, input [][]string) int {
	rope := make([]image.Point, ropeLength)
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
