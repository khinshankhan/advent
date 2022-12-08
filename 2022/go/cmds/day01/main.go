package main

import (
	"fmt"
	"sort"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	s := io.ReadRelativeFile("../data/day01.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parse(s string) [][]int {
	ri := util.TransformString2D(s, "\n\n", "\n", true, false)
	return util.From2DStringTo2DInt(ri)
}

func parta(input [][]int) int {
	m := -1
	for _, carrying := range input {
		sum := math.SumNums(carrying...)
		m = math.Max(m, sum)
	}

	return m
}

func partb(input [][]int) int {
	sums := []int{}
	for _, carrying := range input {
		sum := math.SumNums(carrying...)
		sums = append(sums, sum)
	}

	sort.Ints(sums)
	top3 := sums[len(sums)-3:]
	sum := math.SumNums(top3...)

	return sum
}
