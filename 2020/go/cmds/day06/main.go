package main

import (
	"fmt"
	"strings"

	"github.com/kkhan01/advent/2020/go/utils"
)

func main() {
	lines, err := utils.ReadRelativeFile1DString("../data/day06.txt", "\n\n", false)
	if err != nil {
		panic(err)
	}
	fmt.Println(parta(lines))
	fmt.Println(partb(lines))
}

func parta(lines []string) int {
	totalAns := 0
	for _, answers := range lines {
		uniqueAnswers := make(map[rune]int, 26)
		answers = strings.ReplaceAll(answers, "\n", "")
		for _, answer := range answers {
			uniqueAnswers[answer] += 1
		}
		for _, uniqueAnswer := range uniqueAnswers {
			if uniqueAnswer != 0 {
				totalAns += 1
			}
		}
	}
	return totalAns
}

func partb(lines []string) int {
	totalAns := 0
	for _, answers := range lines {
		uniqueAnswers := make(map[rune]int, 26)
		peopleAnswers := strings.Split(answers, "\n")
		for _, personAnswers := range peopleAnswers {
			for _, answer := range personAnswers {
				uniqueAnswers[answer] += 1
			}
		}
		for _, uniqueAnswer := range uniqueAnswers {
			if uniqueAnswer == len(peopleAnswers) {
				totalAns += 1
			}
		}
	}
	return totalAns
}
