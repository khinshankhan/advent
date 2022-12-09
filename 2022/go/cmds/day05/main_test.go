package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestDay5(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day05.txt")

	tests := test.TestCases[Input, string, string]{
		{
			Name:  "Sample",
			Input: sample,
			Ans1:  "CMZ!",
			Ans2:  "MCD!",
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  "BWNCQRMDB!",
			Ans2:  "NHWZCBNBF!",
			Skip1: s == "",
			Skip2: s == "",
		},
	}

	tests.Run(
		t,
		parse,
		parta,
		partb,
		util.Eq[string],
		util.Eq[string],
	)
}
