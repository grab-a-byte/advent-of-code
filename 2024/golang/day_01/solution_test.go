package day_01

import "testing"

const testInput string = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_DayOne_PartOne(t *testing.T) {
	output := partOneSolution(testInput)
	var expected int64 = 11

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}

func Test_DayOne_PartTwo(t *testing.T) {
	output := partTwoSolution(testInput)
	var expected int64 = 31

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}
