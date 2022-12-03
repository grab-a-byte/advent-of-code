package day_03

import (
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDayThree_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines)

	if result != 157 {
		t.Error("Incorrect Answer")
	}
}

func TestDayThree_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)

	if result != 70 {
		t.Error("Incorrect Answer", result)
	}
}
