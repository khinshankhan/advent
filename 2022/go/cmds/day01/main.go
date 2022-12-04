package main

import (
	"fmt"
	"sort"

	"github.com/khinshankhan/advent/lib/go/conv"
	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
)

func main() {
	rawInput := io.Read2DString("../data/day01.txt", "\n\n", "\n", false)
	input := conv.From2DStringTo2DInt(rawInput)
	parta(input)
	partb(input)
}

func parta(input [][]int) {
	m := -1
	for _, carrying := range input {
		sum := math.SumInts(carrying...)
		m = math.MaxInt(m, sum)
	}

	fmt.Println(m)
}

func partb(input [][]int) {
	sums := []int{}
	for _, carrying := range input {
		sum := math.SumInts(carrying...)
		sums = append(sums, sum)
	}

	sort.Ints(sums)
	top3 := sums[len(sums)-3:]
	sum := math.SumInts(top3...)

	fmt.Println(sum)
}
