package day_05

import (
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDayFive_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines, 3, 3)

	if result != "CMZ" {
		t.Error("Incorrect Answer", result)
	}
}

func TestDayFive_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines, 3, 3)

	if result != "MCD" {
		t.Error("Incorrect Answer", result)
	}
}
