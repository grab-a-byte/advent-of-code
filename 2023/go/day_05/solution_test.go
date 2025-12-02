package day_05

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayFive_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 35 {
		t.Errorf("Expected 35, found, %d", answer)
	}
}

func Test_DayFive_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 46 {
		t.Errorf("Expected 46, found, %d", answer)
	}
}
