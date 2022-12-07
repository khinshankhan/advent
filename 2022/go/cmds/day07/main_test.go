package main

import (
	"testing"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/test"
	"github.com/khinshankhan/advent/lib/go/util"
)

var sample = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestDay7(t *testing.T) {
	s := io.ReadTestFile("..", "..", "data", "day07.txt")

	tests := test.TestCases[string, map[string]FsNode, int]{
		{
			Name:  "Sample",
			Input: sample,
			Ans1:  95437,
			Ans2:  24933642,
		},
		{
			Name:  "Actual",
			Input: s,
			Ans1:  1423358,
			Ans2:  545729,
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
