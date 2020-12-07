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

	fmt.Println(parta(lines))
}

func parta(lines []string) int {
	ans := 0

	checked := make(map[string]bool)
	dict := map[string]map[string]int{}
	for _, line := range lines {
		occured := strings.Count(line[5:], "shiny gold")
		splitted := strings.Split(line, "bags contain")
		bagType := splitted[0] + "bag"
		bags := strings.Split(strings.ReplaceAll(splitted[1][1:len(splitted[1])-1], "bags", "bag"), ", ")

		stored := make(map[string]int)
		dict[bagType] = stored

		if occured > 0 {
			checked[bagType] = true
		} else {
			checked[bagType] = false
		}

		for _, bag := range bags {
			rawCount := bag[0]
			count, _ := strconv.Atoi(string(rawCount))
			if !strings.HasPrefix(bag, "no") {
				stored[bag[2:]] = count
			}
		}
	}

	checkedCount := 0
	for checkedCount < len(checked) {
		checkedCount = 0
		broken := false
		for k, v := range checked {
			if !v {
				for k1 := range dict[k] {
					if checked[k1] {
						checked[k] = true
						broken = true
						break
					}
				}
			}
			if broken {
				break
			}
			checkedCount += 1
		}
	}

	for _, v := range checked {
		if v {
			ans += 1
		}
	}
	return ans
}
