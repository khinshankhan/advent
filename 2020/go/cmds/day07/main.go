package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkhan01/advent/2020/go/utils"
)

func main() {
	lines, err := utils.ReadRelativeFile1DString("../data/day07.txt", "\n", true)
	if err != nil {
		panic(err)
	}

	golds := make(map[string]bool)
	lookup := map[string]map[string]int{}
	for _, line := range lines {
		input := strings.Split(line, "bags contain")
		bagType := input[0] + "bag"
		golds[bagType] = strings.Contains(line[5:], "shiny gold")
		bags := strings.Split(strings.ReplaceAll(input[1][1:len(input[1])-1], "bags", "bag"), ", ")

		stored := make(map[string]int)
		lookup[bagType] = stored

		for _, bag := range bags {
			rawCount := bag[0] // trust bag num < 9
			count, _ := strconv.Atoi(string(rawCount))
			if !strings.HasPrefix(bag, "no") {
				stored[bag[2:]] = count
			}
		}
	}

	fmt.Println(parta(lookup, golds))
	fmt.Println(partb(lookup))
}

func parta(lookup map[string]map[string]int, golds map[string]bool) (ans int) {
nested:
	for color, goldp := range golds {
		if !goldp {
			for goldColor := range lookup[color] {
				if golds[goldColor] {
					golds[color] = true
					goto nested
				}
			}
		}
	}

	for _, v := range golds {
		if v {
			ans += 1
		}
	}

	return ans
}

func partb(lookup map[string]map[string]int) (ans int) {
	for color, count := range lookup["shiny gold bag"] {
		ans += count + count*getPurses(color, lookup) // trust in the recursion
	}
	return ans
}

func getPurses(inColor string, lookup map[string]map[string]int) (ans int) {
	for color, count := range lookup[inColor] {
		ans += count + count*getPurses(color, lookup)
	}
	return ans
}
