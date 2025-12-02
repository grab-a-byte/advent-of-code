package day_11

import (
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDayEleven_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines)

	if result != 10605 {
		t.Error("Incorrect Answer", result, "Should be", 10605)
	}
}

func TestDayEleven_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)

	if result != 2713310158 {
		t.Error("Incorrect Answer", result)
	}
}
