package day_17

import "testing"

const testInput string = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

func Test_DaySeventeen_PartOne(t *testing.T) {
	output := partOneSolution(testInput)
	var expected string = "4,6,3,5,6,3,5,2,1,0"

	if output != expected {
		t.Errorf("expected %s got %s", expected, output)
	}
}

func Test_Machine_Run(t *testing.T) {

}

func Test_DayFourteen_PartTwo(t *testing.T) {
	output := partTwoSolution(testInput)
	var expected string = "UNKNOWN"

	if output != expected {
		t.Errorf("expected %s got %s", expected, output)
	}
}
