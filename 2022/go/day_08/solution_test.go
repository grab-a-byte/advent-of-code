package day_08

import (
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDaySeven_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines)

	if result != 21 {
		t.Error("Incorrect Answer", result, "Should be", 21)
	}
}

func TestDaySeven_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)

	if result != -100 {
		t.Error("Incorrect Answer", result)
	}
}
