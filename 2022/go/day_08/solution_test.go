package day_08

import (
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDayEight_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines)

	if result != 21 {
		t.Error("Incorrect Answer", result, "Should be", 21)
	}
}

func TestDayEight_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)

	if result != 8 {
		t.Error("Incorrect Answer", result)
	}
}
