package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/khinshankhan/advent/lib/go/ds"
	"github.com/khinshankhan/advent/lib/go/io"
)

func main() {
	input := io.ReadRelativeFile("../data/day06.txt")
	s := strings.TrimSpace(input)
	fmt.Println(parta(s))
	fmt.Println(partb(s))
}

func parta(s string) string {
	pos := findNConsecutiveUnique(4, s)
	if pos == -1 {
		return "-1"
	}
	return strconv.Itoa(pos + 1)
}

func partb(s string) string {
	pos := findNConsecutiveUnique(14, s)
	if pos == -1 {
		return "-1"
	}
	return strconv.Itoa(pos + 1)
}

func findNConsecutiveUnique(n int, s string) int {
	fm := ds.NewFrequencyMap[rune]()
	runes := []rune{}
	for i, r := range s {
		fm.Add(r)
		if len(fm.GetMap()) == n {
			return i
		}

		runes = append(runes, r)
		if i > n-2 {
			fm.SubtractOrDelete(runes[i-n+1])
		}
	}

	return -1
}
