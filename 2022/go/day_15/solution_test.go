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
		t.Error("Incorrect Answer", result, "Should be", 31)
	}
}

func TestDayFifteen_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)

	if result != 29 {
		t.Error("Incorrect Answer", result)
	}
}
