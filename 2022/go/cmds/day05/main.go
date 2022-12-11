package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/types"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	s := io.ReadRelativeFile("../data/day05.txt")
	input := parse(s)
	input2 := Input{util.CloneMap(input.Stacks, util.Clone1D[rune]), input.Cmds}

	fmt.Println(parta(input))
	fmt.Println(partb(input2))
}

type Input struct {
	Stacks map[rune][]rune
	Cmds   [][][]rune
}

func parta(input Input) string {
	stacks, cmds := input.Stacks, input.Cmds

	for _, cmd := range cmds {
		count := util.FromStringToInt(string(cmd[0]))

		fromKey, toKey := cmd[1][0], cmd[2][0]
		from, to := stacks[fromKey], stacks[toKey]

		moveRange := math.Max(0, len(from)-count)
		newFrom, addTo := from[:moveRange], from[moveRange:]
		newTo := to
		for addI := len(addTo) - 1; addI > -1; addI-- {
			newTo = append(newTo, addTo[addI])
		}

		stacks[fromKey], stacks[toKey] = newFrom, newTo
	}

	// delimiter !, in case some weird case happens
	return stacksMsg(stacks) + "!"
}

func partb(input Input) string {
	stacks, cmds := input.Stacks, input.Cmds

	for _, cmd := range cmds {
		count := util.FromStringToInt(string(cmd[0]))

		fromKey, toKey := cmd[1][0], cmd[2][0]
		from, to := stacks[fromKey], stacks[toKey]

		moveRange := math.Max(0, len(from)-count)
		newFrom, addTo := from[:moveRange], from[moveRange:]
		newTo := append(to, addTo...)

		stacks[fromKey], stacks[toKey] = newFrom, newTo
	}

	// delimiter !, in case some weird case happens
	return stacksMsg(stacks) + "!"
}

func stacksMsg(stacks map[rune][]rune) string {
	_, orderedStacks := orderedStacksInfo(stacks)
	var s bytes.Buffer
	for _, stack := range orderedStacks {
		if len(stack)-1 < 0 {
			s.WriteRune(' ')
		} else {
			s.WriteRune(stack[len(stack)-1])
		}
	}
	return s.String()
}

func orderedStacksInfo(stacks map[rune][]rune) ([]rune, [][]rune) {
	keys := []rune{}
	for k := range stacks {
		keys = append(keys, k)
	}

	var r types.ByRune = keys
	sort.Sort(r)

	orderedStacks := make([][]rune, len(stacks))
	for i, key := range keys {
		orderedStacks[i] = stacks[key]
	}

	return keys, orderedStacks
}

func parse(s string) Input {
	rawInput := util.TransformString1D(s, "\n\n", false)
	rawStackInfo := strings.Split(rawInput[0], "\n")

	// index -> key
	keys := make(map[int]rune)
	for i, e := range rawStackInfo[len(rawStackInfo)-1] {
		if e != ' ' {
			keys[i] = e
		}
	}

	stacks := make(map[rune][]rune, len(keys))
	for _, v := range keys {
		stacks[v] = []rune{}
	}
	for i := len(rawStackInfo) - 2; i > -1; i-- {
		for coli, r := range rawStackInfo[i] {
			if stackKey, ok := keys[coli]; ok && r != ' ' {
				stacks[stackKey] = append(stacks[stackKey], r)
			}
		}
	}

	cmdsInfo := strings.Split(strings.TrimSpace(rawInput[1]), "\n")
	cmds := make([][][]rune, len(cmdsInfo))
	for i, e := range cmdsInfo {
		parts := strings.Split(e, " ")
		runeCmds := util.From1DStringTo2DRune([]string{parts[1], parts[3], parts[5]})
		cmds[i] = runeCmds
	}

	return Input{stacks, cmds}
}
