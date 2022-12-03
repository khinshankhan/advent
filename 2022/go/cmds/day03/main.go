package main

import (
	"fmt"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	input := io.Read1DString("../data/day03.txt", "\n", true)
	parta(input)
	partb(input)
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

func parta(input []string) {
	points := 0
	for _, rucksack := range input {
		a, b := splitStringEvenly(rucksack)
		intersection := util.Intersection([]rune(a), []rune(b))
		for _, r := range util.Unique(intersection) {
			points += runePoint(r)
		}
	}
	fmt.Println(points)
}

func partb(input []string) {
	points := 0
	for i := 0; i < len(input); i += 3 {
		firstIntersection := util.Intersection([]rune(input[i]), []rune(input[i+1]))
		secondIntersection := util.Intersection(firstIntersection, []rune(input[i+2]))
		for _, r := range util.Unique(secondIntersection) {
			points += runePoint(r)
		}
	}
	fmt.Println(points)
}
