package day_01

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DayOne_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 142 {
		t.Errorf("Expected 142, found, %d", answer)
	}
}

func Test_DayOne_PartTwo(t *testing.T) {
	file, err := os.Open("./test_two.txt")
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

func TestMorphInput(t *testing.T) {

	tests := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteez":    76,
	}

	for input, expected := range tests {
		morphed := getNumber(morphInput(input))
		if morphed != expected {
			fmt.Printf("expected %q, got %q", expected, morphed)
		}
	}
}
