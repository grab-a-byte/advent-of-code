package day_15

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_12/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 12")
	fmt.Println("Part 1:", PartOneSolution(lines, 2_000_000))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

type sensor struct{ xPos, yPos int64 }
type beacon struct{ xPos, yPos int64 }

func parseSensor(part string) sensor {
	sections := strings.Split(part, ", ")
	x, _ := strconv.ParseInt(strings.TrimPrefix(sections[0], "Sensor at x="), 10, 64)
	y, _ := strconv.ParseInt(strings.TrimPrefix(sections[1], "y="), 10, 64)

	return sensor{x, y}
}

func parseBeacon(part string) beacon {
	sections := strings.Split(part, ", ")
	x, _ := strconv.ParseInt(strings.TrimPrefix(sections[0], "closest beacon is at x="), 10, 64)
	y, _ := strconv.ParseInt(strings.TrimPrefix(sections[1], "y="), 10, 64)

	return beacon{x, y}
}

type beaconSensorPair struct {
	beacon
	sensor
}

func parseLine(line *string) beaconSensorPair {
	sections := strings.Split(*line, ": ")
	return beaconSensorPair{
		parseBeacon(sections[1]),
		parseSensor(sections[0]),
	}
}

type point struct{ xPos, yPos int64 }

func manhattanDistance(a, b point) int64 {
	x := math.Abs(float64(a.xPos) - float64(b.xPos))
	y := math.Abs(float64(b.yPos) - float64(a.yPos))
	return int64(x + y)
}

func getPointsInDistance(b beacon, s sensor) []point {
	var points []point

	distance := manhattanDistance(point(b), point(s))
	for x := s.xPos - distance; x < s.xPos+distance; x++ {
		for y := s.yPos - distance; y < s.yPos+distance; y++ {
			pointDistance := manhattanDistance(point(s), point{x, y})
			if pointDistance <= distance {
				points = append(points, point{x, y})
			}
		}
	}
	return points
}

type void struct{}

var instance void

func PartOneSolution(lines []string, yCheck int) int {
	pairs := utils.Map(lines, parseLine)
	fmt.Println(pairs)

	set := make(map[point]void)

	var wg sync.WaitGroup
	wg.Add(len(pairs))

	c := make(chan point)

	go func() {
		wg.Wait()
		close(c)
	}()

	for _, p := range pairs {
		go func(pair beaconSensorPair) {
			defer wg.Done()
			points := getPointsInDistance(pair.beacon, pair.sensor)
			for _, point := range points {
				c <- point
			}
		}(p)
	}

	for p := range c {
		set[p] = instance
	}

	total := 0
	for k := range set {
		if k.yPos == int64(yCheck) {
			total++
		}
	}

	return total - 1
}

func PartTwoSolution(lines []string) int {
	return -1
}
