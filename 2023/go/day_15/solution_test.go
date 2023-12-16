package day_15

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayFifteen_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 1320 {
		t.Errorf("Expected 1320, found, %d", answer)
	}
}

func Test_DayFifteen_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 281 {
		t.Errorf("Expected 281, found, %d", answer)
	}
}
