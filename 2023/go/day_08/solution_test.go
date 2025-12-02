package day_08

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayEight_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 2 {
		t.Errorf("Expected 2, found, %d", answer)
	}
}

func Test_DayEight_PartTwo(t *testing.T) {
	file, err := os.Open("./test2.txt")
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
