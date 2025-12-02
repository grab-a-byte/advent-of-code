package day_14

import "testing"

const testInput string = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func Test_DayFourteen_PartOne(t *testing.T) {
	output := partOneSolution(testInput, 7, 11)
	var expected int64 = 12

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}

func Test_Step(t *testing.T) {
	g := guard{
		currentX: 2,
		currentY: 4,
		moveX:    2,
		moveY:    -3,
	}

	expected := []struct{ x, y int }{
		{4, 1},
		{6, 5},
		{8, 2},
		{10, 6},
		{1, 3},
	}

	for i, e := range expected {
		g.Step(11, 7)
		if g.currentX != e.x {
			t.Errorf("expexted x of %d, got %d at index %d", e.x, g.currentX, i)
		}
		if g.currentY != e.y {
			t.Errorf("expexted y of %d, got %d at index %d", e.y, g.currentY, i)
		}
	}
}

func Test_DayFourteen_PartTwo(t *testing.T) {
	output := partTwoSolution(testInput)
	var expected int64 = 9

	if output != expected {
		t.Errorf("expected %d got %d", expected, output)
	}
}
