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

func parta(input [][]string) (interesting int) {
	interestingCyles := ds.NewSet(20, 60, 100, 140, 180, 220)

	cycle := 0
	strength := 1
	tick := func() {
		cycle += 1
		if interestingCyles.Has(cycle) {
			interesting += cycle * strength
		}
	}

	for _, l := range input {
		switch l[0] {
		case "addx":
			tick()
			tick()
			strength += util.FromStringToInt(l[1])
		case "noop":
			tick()
		}
	}

	return interesting
}

func partb(input [][]string) string {
	msg := ""

	cycle := 0
	sprite := 2
	tick := func() {
		cycle += 1
		crtpos := cycle % 40
		if crtpos == 1 && cycle != 1 {
			msg += "\n"
		}
		if math.Abs(sprite-crtpos) <= 1 {
			msg += "#"
		} else {
			msg += "."
		}
	}

	for _, l := range input {
		switch l[0] {
		case "addx":
			tick()
			tick()
			sprite += util.FromStringToInt(l[1])
		case "noop":
			tick()
		}
	}

	return msg
}
