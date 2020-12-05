package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkhan01/advent/2020/go/utils"
)

func main() {
	lines, err := utils.ReadRelativeFile2DString("../data/day02.txt", "\n", " ", true)
	if err != nil {
		panic(err)
	}

	partaAns := 0
	partbAns := 0
	for _, line := range lines {
		nums := strings.Split(string(line[0]), "-")
		lower, _ := strconv.Atoi(nums[0])
		upper, _ := strconv.Atoi(nums[1])
		characterRune := line[1][0]
		character := string(characterRune)
		password := line[2]
		partaAns += partaValidation(lower, upper, character, password)
		partbAns += partbValidation(lower, upper, characterRune, password)
	}
	fmt.Println(partaAns, partbAns)
}

func partaValidation(lower, upper int, character, password string) int {
	count := strings.Count(password, character)
	if count < lower || count > upper {
		return 0
	}
	return 1
}

func partbValidation(lower, upper int, characterRune byte, password string) int {
	lower, upper = lower-1, upper-1
	checkUpper := len(password) > upper
	checkLower := lower > -1

	if (checkLower && checkUpper && utils.Xor(password[lower] == characterRune, password[upper] == characterRune)) || (!checkUpper && checkLower && password[lower] == characterRune) || (!checkLower && checkUpper && password[upper] == characterRune) {
		if !(checkLower && checkUpper) {
			fmt.Println(lower, upper, password)
		}
		return 1
	}
	return 0
}
