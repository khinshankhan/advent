package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

func TestDay6(t *testing.T) {
	tests := test.TestCases[string, string]{
		{
			Name:  "Sample1",
			Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			Ans:   "7",
			First: true,
		},
		{
			Name:  "Sample2",
			Input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			Ans:   "5",
			First: true,
		},
		{
			Name:  "Sample3",
			Input: "nppdvjthqldpwncqszvftbrmjlhg",
			Ans:   "6",
			First: true,
		},
		{
			Name:  "Sample4",
			Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			Ans:   "10",
			First: true,
		},
		{
			Name:  "Sample5",
			Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			Ans:   "11",
			First: true,
		},
		{
			Name:  "Sample1",
			Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			Ans:   "19",
			First: false,
		},
		{
			Name:  "Sample2",
			Input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			Ans:   "23",
			First: false,
		},
		{
			Name:  "Sample3",
			Input: "nppdvjthqldpwncqszvftbrmjlhg",
			Ans:   "23",
			First: false,
		},
		{
			Name:  "Sample4",
			Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			Ans:   "29",
			First: false,
		},
		{
			Name:  "Sample5",
			Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			Ans:   "26",
			First: false,
		},
	}

	tests.Run(
		t,
		parta,
		partb,
		util.Eq[string],
	)
}
