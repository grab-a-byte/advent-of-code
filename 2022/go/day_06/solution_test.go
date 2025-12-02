package day_06

import (
	"fmt"
	"testing"
)

func TestDaySix_PartOne(t *testing.T) {
	testCases := []struct {
		input  string
		answer int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s returns %d", tc.input, tc.answer), func(t *testing.T) {
			result := PartOneSolution(tc.input)
			if result != tc.answer {
				t.Error("Incorrect answer", result)
			}
		})
	}
}

func TestDaySix_PartTwo(t *testing.T) {
	testCases := []struct {
		input  string
		answer int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s returns %d", tc.input, tc.answer), func(t *testing.T) {
			result := PartTwoSolution(tc.input)
			if result != tc.answer {
				t.Error("Incorrect answer", result)
			}
		})
	}
}
