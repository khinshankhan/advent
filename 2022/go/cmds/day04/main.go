package main

import (
	"fmt"
	"strings"

	"github.com/khinshankhan/advent/lib/go/cast"
	"github.com/khinshankhan/advent/lib/go/io"
)

func main() {
	input := io.Read2DString("../data/day04.txt", "\n", ",", true)
	parta(input)
	partb(input)
}

func parta(input [][]string) {
	contained := 0
	for _, pair := range input {
		firstStr, secondStr := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")
		first, second := cast.From1DStringTo1DInt(firstStr), cast.From1DStringTo1DInt(secondStr)
		if (first[0] <= second[0] && first[1] >= second[1]) || (second[0] <= first[0] && second[1] >= first[1]) {
			contained += 1
		}
	}
	fmt.Println(contained)
}

func partb(input [][]string) {
	contained := 0
	for _, pair := range input {
		firstStr, secondStr := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")
		first, second := cast.From1DStringTo1DInt(firstStr), cast.From1DStringTo1DInt(secondStr)
		if (first[0] <= second[0] && first[1] >= second[0]) || (second[0] <= first[0] && second[1] >= first[0]) {
			contained += 1
		}
	}
	fmt.Println(contained)
}
