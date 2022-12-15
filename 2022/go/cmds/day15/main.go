package main

import (
	"flag"
	"fmt"
	"image"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
)

var (
	Actual = false
	Part2  = false

	RelevantSampleY = 10
	RelevantActualY = 2000000

	MaxSample = 20
	MaxActual = 4000000

	Up   = image.Point{0, 1}
	Down = image.Point{0, -1}
)

func main() {
	flag.BoolVar(&Actual, "actual", false, "run actual or sample")
	flag.BoolVar(&Part2, "part2", false, "part 1 or 2")
	flag.Parse()

	var s string
	if Actual {
		s = io.ReadRelativeFile("../data/day15.txt")
	} else {
		s = sample
	}

	input := parse(s)

	if !Part2 {
		fmt.Println(parta(input))
	} else {
		fmt.Println(partb(input))
	}
}

func parta(input Input) int {
	relevantY := RelevantSampleY
	if Actual {
		relevantY = RelevantActualY
	}
	relevantRow := make(map[image.Point]struct{})
	for sensor, beacon := range input.sensors {
		dist := math.ManhattanDistance(sensor, beacon)
		up, down := Up.Mul(dist).Add(sensor), Down.Mul(dist).Add(sensor)
		if relevantY >= down.Y && relevantY <= up.Y {
			for x := input.minX - dist; x <= input.maxX+dist; x++ {
				potentialPoint := image.Point{x, relevantY}
				sensorToPotential := math.ManhattanDistance(sensor, potentialPoint)
				if sensorToPotential <= dist {
					relevantRow[potentialPoint] = struct{}{}
				}
			}
		}
	}

	arebeacon := 0
	for potential := range relevantRow {
		_, isbeacon := input.beacons[potential]
		if isbeacon {
			arebeacon += 1
		}
	}

	return len(relevantRow) - arebeacon
}

func partb(input Input) int {
	maxCoord := MaxSample
	if Actual {
		maxCoord = MaxActual
	}

	potential := make(map[image.Point]struct{})
	addPotential := func(point image.Point) {
		if point.X >= 0 && point.X <= maxCoord && point.Y >= 0 && point.Y <= maxCoord {
			potential[point] = struct{}{}
		}
	}
	for sensor, beacon := range input.sensors {
		dist := math.ManhattanDistance(sensor, beacon)

		// 0-3 are manhattan moves
		// dist + 1 to get diamond, each edge's points should be out of current sensor range
		up := math.ChebyshevMoves[0].Mul(dist + 1).Add(sensor)
		down := math.ChebyshevMoves[1].Mul(dist + 1).Add(sensor)
		right := math.ChebyshevMoves[2].Mul(dist + 1).Add(sensor)
		left := math.ChebyshevMoves[3].Mul(dist + 1).Add(sensor)

		// get points along each edge of the diamond
		// we can use the rest of chebyshev moves for this
		point := left
		for done := false; !done; {
			done = point.Eq(up)
			addPotential(point)
			point = point.Add(math.ChebyshevMoves[4])
		}

		point = left
		for !point.Eq(down) {
			addPotential(point)
			point = point.Add(math.ChebyshevMoves[5])
		}

		point = right
		for !point.Eq(up) {
			addPotential(point)
			point = point.Add(math.ChebyshevMoves[6])
		}

		point = right
		for !point.Eq(down) {
			addPotential(point)
			point = point.Add(math.ChebyshevMoves[7])
		}
	}

	// points mightve been in another sensor's range, remove them
	for point := range potential {
		for sensor, beacon := range input.sensors {
			dist := math.ManhattanDistance(sensor, beacon)
			sensorToPotential := math.ManhattanDistance(sensor, point)
			if sensorToPotential <= dist {
				delete(potential, point)
				break
			}
		}
	}

	distressBeacon := image.Point{}
	for point := range potential { // only one point, but no way to retrieve
		distressBeacon = point
	}

	return distressBeacon.X*4000000 + distressBeacon.Y
}

type Input struct {
	sensors, beacons       map[image.Point]image.Point
	minX, minY, maxX, maxY int
}

func (i *Input) UpdateBounds(p image.Point) {
	if p.X < i.minX {
		i.minX = p.X
	}
	if p.X > i.maxX {
		i.maxX = p.X
	}
	if p.Y < i.minY {
		i.minY = p.Y
	}
	if p.Y > i.maxY {
		i.maxY = p.Y
	}
}

func parse(s string) Input {
	input := Input{}
	input.sensors = make(map[image.Point]image.Point)
	input.beacons = make(map[image.Point]image.Point)
	input.minY = 9999999 // not used but might be interesting

	split := strings.Split(strings.TrimSpace(s), "\n")
	for _, line := range split {
		var sensorx, sensory, beaconx, beacony int
		fmt.Sscanf(
			line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensorx, &sensory, &beaconx, &beacony,
		)
		sensor, beacon := image.Point{sensorx, sensory}, image.Point{beaconx, beacony}
		input.sensors[sensor], input.beacons[beacon] = beacon, sensor
		input.UpdateBounds(sensor)
		input.UpdateBounds(beacon)
	}

	return input
}

var sample = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`
