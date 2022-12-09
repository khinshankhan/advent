package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay8(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day08.txt")

	tests := test.TestCases[[][]int, int, int]{
		{
			Name:  "Sample",
			Input: sample,
			Ans1:  21,
			Ans2:  8,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  1798,
			Ans2:  259308,
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
