package main

import (
	"fmt"
	"strings"

	"github.com/khinshankhan/advent/lib/go/ds"
	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	s := io.ReadRelativeFile("../data/day10.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parse(s string) [][]string {
	return util.TransformString2D(strings.TrimSpace(s), "\n", " ", true, true)
}

var moves = map[string]int{
	"addx": 2,
	"noop": 1,
}

func parta(input [][]string) int {
	cycle := 0
	strength := 1
	interestingCyles := ds.NewSet(20, 60, 100, 140, 180, 220)
	interesting := []int{}
	for _, l := range input {
		// prob smarter way
		for j := 0; j < moves[l[0]]; j++ {
			cycle += 1
			if interestingCyles.Has(cycle) {
				interesting = append(interesting, cycle*strength)
			}
		}
		if len(l) > 1 {
			strength += util.FromStringToInt(l[1])
		}
	}
	return math.SumNums(interesting...)
}

func partb(input [][]string) string {

	msg := ""
	cycle := 0
	pos := 2
	for _, l := range input {
		// prob smarter way
		for j := 0; j < moves[l[0]]; j++ {
			cycle += 1
			crtpos := cycle % 40
			if crtpos == 1 && cycle != 1 {
				msg += "\n"
			}
			if crtpos == pos || crtpos == pos+1 || crtpos == pos-1 {
				msg += "#"
			} else {
				msg += "."
			}
		}
		if len(l) > 1 {
			pos += util.FromStringToInt(l[1])
		}
	}
	return msg
}
