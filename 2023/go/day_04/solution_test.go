package day_04

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayFour_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 13 {
		t.Errorf("Expected 13, found, %d", answer)
	}
}

func Test_DayFour_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 30 {
		t.Errorf("Expected 30, found, %d", answer)
	}
}
