package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay12(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day12.txt")

	tests := test.TestCases[Input, int, int]{
		{
			Name:  "Sample1",
			Input: sample,
			Ans1:  31,
			Ans2:  29,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  380,
			Ans2:  375,
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
