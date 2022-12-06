package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khinshankhan/advent/lib/go/conv"
	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/types"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	rawInput := io.Read1DString("../data/day05.txt", "\n\n", false)
	stacks, cmds := parseInput(rawInput)
	stacksClone := util.CloneMap(stacks, util.Clone1D[rune])

	fmt.Println(parta(stacksClone, cmds))
	fmt.Println(partb(stacks, cmds))
}

func parta(stacks map[rune][]rune, cmds [][][]rune) string {
	for _, cmd := range cmds {
		count := conv.FromStringToInt(string(cmd[0]))

		fromKey := cmd[1][0]
		toKey := cmd[2][0]
		from := stacks[fromKey]
		to := stacks[toKey]

		moveRange := math.Max(0, len(from)-count)
		newFrom := from[:moveRange]
		addTo := from[moveRange:]
		newTo := to
		for addI := len(addTo) - 1; addI > -1; addI-- {
			newTo = append(newTo, addTo[addI])
		}

		stacks[fromKey] = newFrom
		stacks[toKey] = newTo
	}

	_, orderedStacks := orderedStacksInfo(stacks)
	s := ""
	for _, stack := range orderedStacks {
		if len(stack)-1 < 0 {
			s += " "
		} else {
			s += string(stack[len(stack)-1])
		}
	}

	// delimiter !, in case some weird case happens
	return s + "!"
}

func partb(stacks map[rune][]rune, cmds [][][]rune) string {
	for _, cmd := range cmds {
		count := conv.FromStringToInt(string(cmd[0]))

		fromKey := cmd[1][0]
		toKey := cmd[2][0]
		from := stacks[fromKey]
		to := stacks[toKey]

		moveRange := math.Max(0, len(from)-count)
		newFrom := from[:moveRange]
		addTo := from[moveRange:]
		newTo := append(to, addTo...)

		stacks[fromKey] = newFrom
		stacks[toKey] = newTo
	}

	_, orderedStacks := orderedStacksInfo(stacks)
	s := ""
	for _, stack := range orderedStacks {
		if len(stack)-1 < 0 {
			s += " "
		} else {
			s += string(stack[len(stack)-1])
		}
	}

	// delimiter !, in case some weird case happens
	return s + "!"
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

func parseInput(rawInput []string) (map[rune][]rune, [][][]rune) {
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
		runeCmds := conv.From1DStringTo2DRune([]string{parts[1], parts[3], parts[5]})
		cmds[i] = runeCmds
	}

	return stacks, cmds
}
