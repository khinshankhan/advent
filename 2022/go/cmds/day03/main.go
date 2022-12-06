package main

import (
	"fmt"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	s := io.ReadRelativeFile("../data/day03.txt")
	input := parse(s)
	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parse(s string) []string {
	return util.TransformString1D(strings.TrimSpace(s), "\n", false)
}

func runePoint(r rune) int {
	if r >= 'a' {
		return int(r-'a') + 1
	}
	return int(r-'A') + 27
}

func splitStringEvenly(s string) (string, string) {
	midpoint := len(s) / 2
	return s[:midpoint], s[midpoint:]
}

func parta(input []string) int {
	points := 0
	for _, rucksack := range input {
		a, b := splitStringEvenly(rucksack)
		intersection := util.Intersection([]rune(a), []rune(b))
		for _, r := range util.Unique(intersection) {
			points += runePoint(r)
		}
	}
	return points
}

func partb(input []string) int {
	points := 0
	for i := 0; i < len(input); i += 3 {
		intersection := util.Intersection(
			[]rune(input[i]),
			[]rune(input[i+1]),
			[]rune(input[i+2]),
		)
		for _, r := range util.Unique(intersection) {
			points += runePoint(r)
		}
	}
	return points
}
