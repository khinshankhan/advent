package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
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

type TestInfo struct {
	name  string
	first bool
	input string
	ans   string
}

func TestDay5(t *testing.T) {
	actualInput := io.ReadTestFile("..", "..", "data", "day05.txt")
	tcs := []TestInfo{
		{
			name:  "Sample1",
			first: true,
			input: sample,
			ans:   "CMZ!",
		},
		{
			name:  "Sample2",
			first: false,
			input: sample,
			ans:   "MCD!",
		},
		{
			name:  "Actual1",
			first: true,
			input: actualInput,
			ans:   "BWNCQRMDB!",
		},
		{
			name:  "Actual2",
			first: false,
			input: actualInput,
			ans:   "NHWZCBNBF!",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input == "" {
				t.Skip("Input for day doesn't exist")
			}
			rawInput := util.TransformString1D(tc.input, "\n\n", false)
			stacks, cmds := parseInput(rawInput)
			var res string
			if tc.first {
				res = parta(stacks, cmds)
			} else {
				res = partb(stacks, cmds)
			}

			if tc.ans != res {
				t.Fatalf("expected: %v, got: %v", tc.ans, res)
			}
		})
	}
}
