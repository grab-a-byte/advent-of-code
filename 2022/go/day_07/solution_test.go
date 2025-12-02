package day_07

import (
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDaySeven_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines)

	if result != 95437 {
		t.Error("Incorrect Answer", result)
	}
}

func TestDaySeven_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)

	if result != 24933642 {
		t.Error("Incorrect Answer", result)
	}
}
