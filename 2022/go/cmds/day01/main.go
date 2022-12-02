package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	fname := "input.txt"

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		panic(err)
	}
	inputString := string(b)

	rawLines := strings.Split(inputString, "\n\n")
	lines := make([][]int, len(rawLines))
	for i, rawLine := range rawLines {
		rawNums := strings.Split(rawLine, "\n")
		nums := []int{}
		for _, rawNum := range rawNums {
			if rawNum == "" {
				continue
			}
			num, err := strconv.Atoi(rawNum)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		lines[i] = nums
	}

	return lines
}

func main() {
	input := parseInput()
	parta(input)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func parta(input [][]int) {
	m := -1
	for _, carrying := range input {
		sum := 0
		for _, calories := range carrying {
			sum += calories
		}
		m = max(m, sum)
	}
	fmt.Println(m)
}
