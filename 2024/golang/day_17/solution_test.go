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
	check := machine{
		c:       9,
		program: []int{2, 6},
	}
	check.run()
	if check.b != 1 {
		t.Error("check  failed")
	}

	check2 := machine{
		a:       10,
		program: []int{5, 0, 5, 1, 5, 4},
	}
	check2.run()
	if !testEq(check2.output, []int{0, 1, 2}) {
		t.Error("check 2 failed")
	}

	check3 := machine{
		a:       2024,
		program: []int{0, 1, 5, 4, 3, 0},
	}
	check3.run()

	if !testEq(check3.output, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}) || check3.a != 0 {
		t.Error("check 3 failed")
	}

	check4 := machine{
		b:       29,
		program: []int{1, 7},
	}
	check4.run()

	if check4.b != 26 {
		t.Error("check 4 failed")
	}

	check5 := machine{
		b:       2024,
		c:       43690,
		program: []int{4, 0},
	}
	check5.run()

	if check5.b != 44354 {
		t.Error("check 5 failed")
	}

}

func Test_DayFourteen_PartTwo(t *testing.T) {
	output := partTwoSolution(testInput)
	var expected int = 117440

	if output != 0 {
		t.Errorf("expected %d got %d", expected, output)
	}
}
