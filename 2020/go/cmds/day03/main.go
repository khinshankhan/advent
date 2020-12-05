package main

import (
	"fmt"

	"github.com/kkhan01/advent/2020/go/utils"
)

func main() {
	lines, err := utils.ReadRelativeFile1DString("../data/day03.txt", "\n", true)
	if err != nil {
		panic(err)
	}

	rows, cols := len(lines), len(lines[0])
	fmt.Println(parta(lines, rows, cols))
	fmt.Println(partb(lines, rows, cols))
}

func parta(lines []string, rows, cols int) int {
	return transverse(lines, rows, cols, 3, 1)
}

func partb(lines []string, rows, cols int) int {
	slopes := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}
	ans := 1
	for _, slope := range slopes {
		ans *= transverse(lines, rows, cols, slope[0], slope[1])
	}
	return ans
}

func transverse(topography []string, rows, cols, right, down int) int {
	treesEncountered := 0
	for r, c := 0, 0; r < rows; r, c = r+down, c+right {
		if c >= cols {
			c = c % cols
		}
		if topography[r][c] == '#' {
			treesEncountered += 1
		}
	}
	return treesEncountered
}
