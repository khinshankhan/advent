package main

import (
	"fmt"

	"github.com/khinshankhan/advent/lib/go/ds/circularslice"
	"github.com/khinshankhan/advent/lib/go/io"
)

func main() {
	input := io.Read2DString("../data/day02.txt", "\n", " ", true)
	parta(input)
	partb(input)
}

const (
	rock    = 0
	paper   = 1
	scissor = 2
)

// == draw, prev win, next lose
var handSigns = circularslice.New("rock", "paper", "scissors")

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

func parta(input [][]string) {
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

	fmt.Println(sum)
}

func partb(input [][]string) {
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

	fmt.Println(sum)
}
