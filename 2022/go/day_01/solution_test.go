package day_01

import (
	"fmt"
	"os"
	"testing"

	"github.com/parkeradam/aoc_2022/utils"
)

func Test_DayOne_PartOne(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartOneSolution(lines)
	fmt.Println("{}", result)

	if result != 24000 {
		t.Error("Wrong")
	}
}

func Test_DayOne_PartTwo(t *testing.T) {
	file, _ := os.Open("./test.txt")
	lines := utils.ReadFileAsLines(file)
	result := PartTwoSolution(lines)
	fmt.Println("{}", result)

	if result != 45000 {
		t.Error("Wrong")
	}
}
