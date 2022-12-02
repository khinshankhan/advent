package main

import (
	"fmt"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	input := io.Read2DString("../data/day02.txt", "\n", " ", true)
	parta(input)
	partb(input)
}

var winningHand = map[string]string{
	"rock":    "scissor",
	"paper":   "rock",
	"scissor": "paper",
}

var handPoints = map[string]int{
	"rock":    1,
	"paper":   2,
	"scissor": 3,
}

var convertOpp = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissor",
}

var convertMine = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissor",
}

func parta(input [][]string) {
	sum := 0
	for _, hands := range input {
		opp, mine := convertOpp[hands[0]], convertMine[hands[1]]
		sum += handPoints[mine]
		if opp == mine {
			sum += 3
		} else if winningHand[mine] == opp {
			sum += 6
		}
	}

	fmt.Println(sum)
}

func partb(input [][]string) {
	losingHand := util.ReverseMap(winningHand)

	sum := 0
	for _, hands := range input {
		opp, scenario := convertOpp[hands[0]], hands[1]
		mine := ""
		switch scenario {
		case "X":
			mine = winningHand[opp]
		case "Y":
			mine = opp
		case "Z":
			mine = losingHand[opp]
		}

		sum += handPoints[mine]
		if opp == mine {
			sum += 3
		} else if winningHand[mine] == opp {
			sum += 6
		}
	}

	fmt.Println(sum)
}
