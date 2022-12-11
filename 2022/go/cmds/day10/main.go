package main

import (
	"bytes"
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

func parta(input [][]string) (interesting int) {
	interestingCyles := ds.NewSet(20, 60, 100, 140, 180, 220)

	cycle, strength := 0, 1
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
	var crt bytes.Buffer

	cycle, sprite := 0, 2
	tick := func() {
		cycle += 1
		crtpos := (cycle-1)%40 + 1
		if crtpos == 1 && cycle != 1 {
			crt.WriteString("\n")
		}
		if math.Abs(sprite-crtpos) <= 1 {
			crt.WriteString("#")
		} else {
			crt.WriteString(".")
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

	return crt.String()
}

func parse(s string) [][]string {
	return util.TransformString2D(strings.TrimSpace(s), "\n", " ", true, true)
}

var sample = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`
