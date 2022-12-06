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
}

func parta(s string) string {
	fm := ds.NewFrequencyMap[rune]()
	runes := []rune{}
	for i, r := range s {
		fm.Add(r)
		if len(fm.GetMap()) == 4 {
			return strconv.Itoa(i + 1)
		}

		runes = append(runes, r)
		if i > 2 {
			fm.SubtractOrDelete(runes[i-3])
		}
	}

	return "-1"
}
