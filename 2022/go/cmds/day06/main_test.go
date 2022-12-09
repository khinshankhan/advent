package main

import (
	"strings"
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay6(t *testing.T) {
	input := io.ReadTestFile("..", "..", "data", "day06.txt")
	s := strings.TrimSpace(input)

	tests := test.TestCases[string, int, int]{
		{
			Name:  "Sample1",
			Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			Ans1:  7,
			Ans2:  19,
		},
		{
			Name:  "Sample2",
			Input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			Ans1:  5,
			Ans2:  23,
		},
		{
			Name:  "Sample3",
			Input: "nppdvjthqldpwncqszvftbrmjlhg",
			Ans1:  6,
			Ans2:  23,
		},
		{
			Name:  "Sample4",
			Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			Ans1:  10,
			Ans2:  29,
		},
		{
			Name:  "Sample5",
			Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			Ans1:  11,
			Ans2:  26,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  1802,
			Ans2:  3551,
			Skip1: s == "",
			Skip2: s == "",
		},
	}

	tests.Run(
		t,
		math.Identity[string],
		parta,
		partb,
		util.Eq[int],
		util.Eq[int],
	)
}
