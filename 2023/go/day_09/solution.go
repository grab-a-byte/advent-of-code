package day_09

import (
	"aoc/util"
	"fmt"
	"os"
)

func Solution() {
	file, err := os.Open("./day_09/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day nine part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day nine part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	total := 0
	for _, line := range input {
		total += processLine(line)
	}
	return total
}

func partTwoSolution(input []string) int {
	total := 0
	for _, line := range input {
		total += processLinePartTwo(line)
	}
	return total
}

func processLine(line string) int {
	steps := buildSteps(line)
	return nextStep(steps)
}

func processLinePartTwo(line string) int {
	steps := buildSteps(line)
	return nextStepPart2(steps)
}

func buildSteps(line string) [][]int {
	values := util.StrToIntsSpace(line)
	steps := [][]int{}
	steps = append(steps, values)

	for !allZeroes(steps[len(steps)-1]) {
		currentToCalc := steps[len(steps)-1]
		steps = append(steps, []int{})
		for i := 0; i < len(currentToCalc)-1; i++ {
			valToAdd := currentToCalc[i+1] - currentToCalc[i]
			steps[len(steps)-1] = append(steps[len(steps)-1], valToAdd)
		}
	}
	return steps
}

func nextStep(steps [][]int) int {
	valueToAddTo := steps[0][len(steps[0])-1]
	if len(steps) == 2 {
		return valueToAddTo
	}
	return valueToAddTo + nextStep(steps[1:])
}

func nextStepPart2(steps [][]int) int {
	valueToAddTo := steps[0][0]
	if len(steps) == 2 {
		return valueToAddTo
	}
	return valueToAddTo - nextStepPart2(steps[1:])
}

func allZeroes(i []int) bool {
	for _, v := range i {
		if v != 0 {
			return false
		}
	}
	return true
}
