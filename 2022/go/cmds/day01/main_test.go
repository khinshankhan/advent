package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestDay1(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day01.txt")

	tests := test.TestCases[[][]int, int, int]{
		{
			Name:  "Sample",
			Input: sample,
			Ans1:  24000,
			Ans2:  45000,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  69626,
			Ans2:  206780,
			Skip1: s == "",
			Skip2: s == "",
		},
	}

	tests.Run(
		t,
		parse,
		parta,
		partb,
		util.Eq[int],
		util.Eq[int],
	)
}
