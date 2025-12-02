package day_13

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayThirteen_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 405 {
		t.Errorf("Expected 405, found, %d", answer)
	}
}

func Test_DayThirteen_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 6 {
		t.Errorf("Expected 6, found, %d", answer)
	}
}
