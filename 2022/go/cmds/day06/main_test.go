package main

import (
	"testing"
)

type TestInfo struct {
	name  string
	first bool
	input string
	ans   string
}

func TestDay6(t *testing.T) {
	tcs := []TestInfo{
		{
			name:  "Sample1",
			first: true,
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			ans:   "7",
		},
		{
			name:  "Sample2",
			first: true,
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			ans:   "5",
		},
		{
			name:  "Sample3",
			first: true,
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			ans:   "6",
		},
		{
			name:  "Sample4",
			first: true,
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			ans:   "10",
		},
		{
			name:  "Sample5",
			first: true,
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			ans:   "11",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input == "" {
				t.Skip("Input for day doesn't exist")
			}
			var res string
			if tc.first {
				res = parta(tc.input)
			}
			// else {
			// res = partb(stacks, cmds)
			// }

			if tc.ans != res {
				t.Fatalf("expected: %v, got: %v", tc.ans, res)
			}
		})
	}
}
