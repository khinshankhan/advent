package main

import (
	"fmt"
	"strings"

	"github.com/khinshankhan/advent/lib/go/conv"
	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	s := io.ReadRelativeFile("../data/day04.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parse(s string) [][]string {
	return util.TransformString2D(strings.TrimSpace(s), "\n", ",", false, false)
}

func parta(input [][]string) int {
	contained := 0
	for _, pair := range input {
		firstStr, secondStr := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")
		first, second := conv.From1DStringTo1DInt(firstStr), conv.From1DStringTo1DInt(secondStr)
		if (first[0] <= second[0] && first[1] >= second[1]) || (second[0] <= first[0] && second[1] >= first[1]) {
			contained += 1
		}
	}
	return contained
}

func partb(input [][]string) int {
	overlapped := 0
	for _, pair := range input {
		firstStr, secondStr := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")
		first, second := conv.From1DStringTo1DInt(firstStr), conv.From1DStringTo1DInt(secondStr)
		if (first[0] <= second[0] && first[1] >= second[0]) || (second[0] <= first[0] && second[1] >= first[0]) {
			overlapped += 1
		}
	}
	return overlapped
}
