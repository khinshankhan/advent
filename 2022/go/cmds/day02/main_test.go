package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample = `A Y
B X
C Z`

func TestDay2(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day02.txt")

	tests := test.TestCases[string, [][]string, int]{
		{
			Name:  "Sample",
			Input: sample,
			Ans1:  15,
			Ans2:  12,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  10994,
			Ans2:  12526,
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
	)
}
