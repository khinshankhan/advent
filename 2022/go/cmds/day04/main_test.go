package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestDay4(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day04.txt")

	tests := test.TestCases[[][]string, int, int]{
		{
			Name:  "Sample",
			Input: sample,
			Ans1:  2,
			Ans2:  4,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  490,
			Ans2:  921,
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
