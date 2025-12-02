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
	file, err := os.Open("./day_15/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 15")
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

func getPointsInDistance(b beacon, s sensor, yCheck int64) []point {
	var points []point

	distance := manhattanDistance(point(b), point(s))

	if s.yPos+distance < yCheck && s.yPos-distance > yCheck {
		return make([]point, 0)
	}

	for i := int64(0); i <= distance*2; i++ {

		currentY := s.yPos - distance + i
		if currentY != yCheck {
			continue
		}

		xAdd := int64(0)
		if i > distance {
			xAdd = (distance * 2) - i
		} else {
			xAdd = i
		}

		for x := s.xPos - xAdd; x <= s.xPos+xAdd; x++ {
			points = append(points, point{x, currentY})
		}

	}

	return points
}

type void struct{}

var instance void

func PartOneSolution(lines []string, yCheck int) int {
	pairs := utils.Map(lines, parseLine)

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
			points := getPointsInDistance(pair.beacon, pair.sensor, int64(yCheck))
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
