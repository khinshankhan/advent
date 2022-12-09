package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay9(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day09.txt")

	tests := test.TestCases[[][]string, int, int]{
		{
			Name:  "Sample1",
			Input: sample1,
			Ans1:  13,
			Ans2:  1,
		},
		{
			Name:  "Sample2",
			Input: sample2,
			Ans1:  0,
			Ans2:  36,
			Skip1: true,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  6376,
			Ans2:  2607,
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
