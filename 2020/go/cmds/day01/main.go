package main

import (
	"fmt"

	"github.com/kkhan01/advent/2020/go/utils"
)

func main() {
	nums, err := utils.ReadRelativeFile1DInt("../data/day01.txt", "\n")
	if err != nil {
		panic(err)
	}

	e, e2 := part1(nums)
	fmt.Println(e*e2, e, e2)
	e, e2, e3 := part2(nums)
	fmt.Println(e*e2*e3, e, e2, e3)
}

func part1(nums []int) (int, int) {
	for i, e := range nums {
		for _, e2 := range nums[i : len(nums)-1] {
			if e+e2 == 2020 {
				return e, e2
			}
		}
	}
	return -1, -1
}

func part2(nums []int) (int, int, int) {
	for i, e := range nums {
		for j, e2 := range nums[i : len(nums)-1] {
			for _, e3 := range nums[j : len(nums)-1] {
				if e+e2+e3 == 2020 {
					return e, e2, e3
				}
			}
		}
	}
	return -1, -1, -1
}
