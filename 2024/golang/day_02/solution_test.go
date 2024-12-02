package day_02

import "testing"

const testInput string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_DayOje_PartOne(t *testing.T) {
	output := partOneSolution(testInput)
	var expected int = 2

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}

func Test_P1LineProcessor(t *testing.T) {
	input := []int{38, 41, 44, 47, 50, 47}
	result := lineProceessor(input)
	if result != 1 {
		t.Error("Invalid input marked as valid")
	}
}

func Test_DayOne_PartTwo(t *testing.T) {
	output := partTwoSolution(testInput)
	var expected int = 4

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}
