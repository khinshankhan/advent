package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample1 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var sample2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

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
