package main

import (
	"fmt"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample = `30373
25512
65332
33549
35390`

func main() {
	s := io.ReadRelativeFile("../data/day08.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func getVisibles(heights [][]int) [][]bool {
	visibles := make([][]bool, len(heights))
	for i, line := range heights {
		visibles[i] = make([]bool, len(line))
	}

	tallestCol := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		tallestLeft := 0
		for j := 0; j < len(heights[i]); j++ {
			visible := false

			if i == 0 || heights[i][j] > tallestCol[j] {
				tallestCol[j] = heights[i][j]
				visible = true
			}
			if j == 0 || heights[i][j] > tallestLeft {
				tallestLeft = heights[i][j]
				visible = true
			}
			if i == len(heights)-1 || j == len(heights[i])-1 {
				visible = true
			}

			visibles[i][j] = visible
		}
	}

	for i := range tallestCol {
		tallestCol[i] = 0
	}
	for i := len(heights[0]) - 1; i > -1; i-- {
		tallestRight := 0
		for j := len(heights[i]) - 1; j > -1; j-- {
			visible := visibles[i][j]

			if i == len(heights)-1 || heights[i][j] > tallestCol[j] {
				tallestCol[j] = heights[i][j]
				visible = true
			}
			if j == len(heights[i])-1 || heights[i][j] > tallestRight {
				tallestRight = heights[i][j]
				visible = true
			}

			visibles[i][j] = visible
		}
	}

	return visibles
}

func parta(heights [][]int) int {
	visibles := getVisibles(heights)

	total := 0
	for _, r := range visibles {
		for _, c := range r {
			if c {
				total += 1
			}
		}
	}
	return total
}

type DirectionalValues struct {
	Top    int
	Left   int
	Bottom int
	Right  int
}

func getScenic(heights [][]int) [][]int {
	scenic := make([][]int, len(heights))
	for i, line := range heights {
		scenic[i] = make([]int, len(line))
	}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			see := DirectionalValues{}
			if i == 0 || j == 0 || i == len(heights)-1 || j == len(heights[i])-1 {
				continue
			}

			for k := i - 1; k > -1; k-- {
				see.Top += 1
				if heights[i][j] <= heights[k][j] {
					break
				}
			}
			for k := j - 1; k > -1; k-- {
				see.Left += 1
				if heights[i][j] <= heights[i][k] {
					break
				}
			}
			for k := i + 1; k < len(heights); k++ {
				see.Bottom += 1
				if heights[i][j] <= heights[k][j] {
					break
				}
			}
			for k := j + 1; k < len(heights[i]); k++ {
				see.Right += 1
				if heights[i][j] <= heights[i][k] {
					break
				}
			}

			scenic[i][j] = math.ProductNums(
				see.Top,
				see.Left,
				see.Bottom,
				see.Right,
			)
		}
	}

	return scenic
}

func partb(heights [][]int) int {
	scenic := getScenic(heights)

	r := 0
	for _, line := range scenic {
		r = math.Max(r, math.Max(line...))
	}

	return r
}

func parse(s string) [][]int {
	lines := util.TransformString1D(strings.TrimSpace(s), "\n", true)
	heights := make([][]int, len(lines))
	runes := util.From1DStringTo2DRune(lines)
	for i, runeLine := range runes {
		heights[i] = util.Map1D(util.FromRuneToInt, runeLine)
	}

	return heights
}
