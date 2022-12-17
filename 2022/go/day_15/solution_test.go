package day_15

import (
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDayFifteen_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines, 10)

	if result != 26 {
		t.Error("Incorrect Answer", result, "Should be", 26)
	}
}

func TestDayFifteen_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines, 0, 20)

	if result != 56000011 {
		t.Error("Incorrect Answer", result)
	}
}

func TestGetPointInDistance(t *testing.T) {
	p := getPointsInDistance(beacon{xPos: 3, yPos: 7}, sensor{xPos: 5, yPos: 10}, 100)
	if len(p) != 61 {
		t.Error()
	}
}
