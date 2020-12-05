package main

import (
	"fmt"
	"sync"

	"github.com/kkhan01/advent/2020/go/utils"
)

func main() {
	lines, err := utils.ReadRelativeFileString("../data/day05.txt", "\n")
	if err != nil {
		panic(err)
	}

	fmt.Println(parta(lines))
	fmt.Println(partb(lines))
}

func parta(lines []string) int {
	ans := 0
	for _, line := range lines {
		ans = utils.Max(ans, findSeatNumber(line))
	}
	return ans
}

func partb(lines []string) int {
	minSeat, maxSeat := 10000, 0
	takenSeats := make(map[int]bool, len(lines))

	// not sure if this actually parallelized properly
	wg := sync.WaitGroup{}
	for _, line := range lines {
		wg.Add(1)
		go func(line string) {
			seat := findSeatNumber(line)
			minSeat = utils.Min(minSeat, seat)
			maxSeat = utils.Max(maxSeat, seat)
			takenSeats[seat] = true
			wg.Done()
		}(line)
	}
	wg.Wait()

	for i := minSeat; i < maxSeat; i++ {
		if takenSeats[i-1] && takenSeats[i+1] && !takenSeats[i] {
			return i
		}
	}
	return -1
}

func findSeatNumber(line string) int {
	rows, cols := []int{0, 127}, []int{0, 7}
	for _, rune := range line {
		if rune == 'F' {
			rows[1] = rows[0] + (rows[1]-rows[0])/2
		} else if rune == 'B' {
			rows[0] += (rows[1] - rows[0] + 1) / 2
		} else if rune == 'L' {
			cols[1] = cols[0] + (cols[1]-cols[0])/2
		} else if rune == 'R' {
			cols[0] += (cols[1] - cols[0] + 1) / 2
		} else {
			fmt.Println("debug: unexpected input:", rune)
			return -1
		}
	}

	if rows[0] != rows[1] || cols[0] != cols[1] {
		fmt.Println("debug: invalid ticket number:", rows, cols)
		return -1
	}
	return rows[0]*8 + cols[0]
}
