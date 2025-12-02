package day_09

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayNine_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 114 {
		t.Errorf("Expected 114, found, %d", answer)
	}
}

func TestProcessLine(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"0 3 6 9 12 15", 18},
		{"1 3 6 10 15 21", 28},
		{"10 13 16 21 30 45", 68},
	}

	for _, test := range tests {
		result := processLine(test.input)
		if result != test.expected {
			t.Errorf("For %s expected %d but got %d", test.input, test.expected, result)
		}
	}
}

func Test_DayNine_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 2 {
		t.Errorf("Expected 2, found, %d", answer)
	}
}
