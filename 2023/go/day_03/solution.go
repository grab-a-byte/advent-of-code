package day_03

import (
	"aoc/util"
	"fmt"
	"os"
)

func Solution() {
	panic("NOT IMPLEMNTD")
	file, err := os.Open("./day_03/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day three part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day three part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	return 0
}

func partTwoSolution(input []string) int {
	return 0
}
