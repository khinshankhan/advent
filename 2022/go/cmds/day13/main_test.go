package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay13(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day13.txt")

	tests := test.TestCases[[]Packet, int, int]{
		{
			Name:  "Sample1",
			Input: sample,
			Ans1:  13,
			Ans2:  140,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  5580,
			Ans2:  26200,
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
