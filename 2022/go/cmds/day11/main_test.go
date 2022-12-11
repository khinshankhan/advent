package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay11(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day11.txt")

	tests := test.TestCases[Monkeys, int64, int64]{
		{
			Name:  "Sample1",
			Input: sample,
			Ans1:  10605,
			Ans2:  2713310158,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  56350,
			Ans2:  13954061248,
			Skip1: s == "",
			Skip2: s == "",
		},
	}

	tests.Run(
		t,
		parse,
		parta,
		partb,
		util.Eq[int64],
		util.Eq[int64],
	)
}
