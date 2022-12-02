package main

import (
	"fmt"

	"github.com/khinshankhan/advent/lib/go/io"
)

func main() {
	input := io.Read2DString("../data/day02.txt", "\n", " ", true)
	parta(input)
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
