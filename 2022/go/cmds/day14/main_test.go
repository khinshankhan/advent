package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay14(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day14.txt")

	tests := test.TestCases[Input, int, int]{
		{
			Name:  "Sample1",
			Input: sample,
			Ans1:  24,
			Ans2:  93,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  1199,
			Ans2:  23925,
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
