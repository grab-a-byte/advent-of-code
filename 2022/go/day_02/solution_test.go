package day_02

import (
	"fmt"
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func TestDayTwo_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines)

	if result != 15 {
		t.Error("Incorrect Answer")
	}
}

func TestDayTwo_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)

	if result != 12 {
		t.Error("Incorrect Answer")
	}
}

// X = Rock // 1
// Y = Paper // 2
// Z = Scissors // 3
func TestScorePoint(t *testing.T) {
	testCases := []struct {
		ourCall      string
		opponentCall string
		value        int
	}{
		{"X", "X", 4},
		{"Y", "Y", 5},
		{"Z", "Z", 6},
		{"X", "Y", 1},
		{"X", "Z", 7},
		{"Y", "X", 8},
		{"Y", "Z", 2},
		{"Z", "X", 3},
		{"Z", "Y", 9},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s vs %s score of %d", tc.ourCall, tc.opponentCall, tc.value), func(t *testing.T) {
			result := scorePoint(tc.opponentCall, tc.ourCall)
			if result != tc.value {
				t.Error("Incorrect call")
			}
		})
	}
}
