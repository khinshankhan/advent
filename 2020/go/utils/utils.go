package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadIntFile will return an array of ints from a given file separated by newlines
// https://stackoverflow.com/a/9863218
func ReadIntFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		// use scanf later days
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
