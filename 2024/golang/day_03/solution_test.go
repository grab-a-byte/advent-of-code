package day_03

import "testing"

func Test_DayThree_PartOne(t *testing.T) {
	testInput := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	output := partOneSolution(testInput)
	var expected int64 = 161

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}

func Test_DayThree_PartTwo(t *testing.T) {
	testInput := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	output := partTwoSolution(testInput)
	var expected int64 = 48

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}
