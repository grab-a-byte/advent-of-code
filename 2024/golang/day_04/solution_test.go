package day_04

import "testing"

const testInput string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_DayFour_PartOne(t *testing.T) {
	output := partOneSolution(testInput)
	var expected int64 = 18

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}

func Test_DayFour_PartTwo(t *testing.T) {
	output := partTwoSolution(testInput)
	var expected int64 = 9

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}
