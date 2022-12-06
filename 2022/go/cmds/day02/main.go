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
	rock    = 0
	paper   = 1
	scissor = 2
)

// == draw, prev win, next lose
var handSigns = ds.NewCircularSlice("rock", "paper", "scissors")

var handPoints = map[int]int{
	rock:    1,
	paper:   2,
	scissor: 3,
}

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
		sum += handPoints[mine]
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

		// convert mine logic
		mine := 0
		if scenario == "X" {
			mine = handSigns.PrevIndex(opp)
		} else if scenario == "Y" {
			mine = opp
		} else {
			mine = handSigns.NextIndex(opp)
		}

		sum += handPoints[mine]
		if opp == mine {
			sum += 3
		} else if handSigns.PrevIndex(mine) == opp {
			sum += 6
		}
	}

	return sum
}
