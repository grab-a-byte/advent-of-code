package day_02

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayTwo_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 8 {
		t.Errorf("Expected 8, found, %d", answer)
	}
}

func Test_DayTwo_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 2286 {
		t.Errorf("Expected 2286, found, %d", answer)
	}
}
