package day_06

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DaySix_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 288 {
		t.Errorf("Expected 288, found, %d", answer)
	}
}

func Test_DaySix_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 71503 {
		t.Errorf("Expected 71503, found, %d", answer)
	}
}
