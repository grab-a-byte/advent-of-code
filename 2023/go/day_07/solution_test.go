package day_07

import (
	"aoc/util"
	"fmt"
	"os"
	"testing"
)

func Test_DaySeven_PartOne(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	if answer != 6440 {
		t.Errorf("Expected 6440, found, %d", answer)
	}
}

func Test_DaySeven_PartTwo(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("test file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partTwoSolution(lines)
	if answer != 5905 {
		t.Errorf("Expected 5905, found, %d", answer)
	}
}

func TestCalculateHand(t *testing.T) {
	tests := []struct {
		input    string
		expected handType
	}{
		{"32T3K", ONE_PAIR},
		{"T55J5", THREE},
		{"KK677", TWO_PAIR},
		{"KTJJT", TWO_PAIR},
		{"QQQJA", THREE},
		{"12345", HIGH},
		{"99922", FULL},
		{"33222", FULL},
		{"44444", FIVE},
		{"55552", FOUR},
	}

	for _, test := range tests {
		res := calculateHandType(test.input)
		if res != test.expected {
			t.Errorf("expected %v but got %v for input %q", test.expected, res, test.input)
		}
	}
}
