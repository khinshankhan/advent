package main

import (
	"fmt"
	"strings"

	"github.com/khinshankhan/advent/lib/go/ds"
	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	s := io.ReadRelativeFile("../data/day02.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parse(s string) [][]string {
	return util.TransformString2D(strings.TrimSpace(s), "\n", " ", false, false)
}

const (
	rock = iota
	paper
	scissor
)

// == draw, prev win, next lose
// technically slice contents don't actually matter
var handSigns = ds.NewCircularSlice("rock", "paper", "scissors")

var convertOpp = map[string]int{
	"A": rock,
	"B": paper,
	"C": scissor,
}

func parta(input [][]string) int {
	var convertMine = map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissor,
	}

	sum := 0
	for _, hands := range input {
		opp, mine := convertOpp[hands[0]], convertMine[hands[1]]
		sum += mine + 1

		if opp == mine {
			sum += 3
		} else if handSigns.PrevIndex(mine) == opp {
			sum += 6
		}
	}

	return sum
}

func partb(input [][]string) int {
	sum := 0
	for _, hands := range input {
		opp, scenario := convertOpp[hands[0]], hands[1]

		// convert my hand logic
		mine := 0
		switch scenario {
		case "X":
			mine = handSigns.PrevIndex(opp)
		case "Y":
			mine = opp
			sum += 3
		case "Z":
			mine = handSigns.NextIndex(opp)
			sum += 6
		}

		sum += mine + 1
	}

	return sum
}
