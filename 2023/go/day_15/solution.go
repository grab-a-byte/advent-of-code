package day_15

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
)

func Solution() {
	file, err := os.Open("./day_15/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day fifteen part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day fifteen part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	in := strings.Split(strings.Join(input, ""), ",")
	total := 0
	for _, part := range in {
		total += solvePart(part)
	}
	return total
}

func solvePart(s string) int {
	current := 0
	for _, r := range s {
		current += int(r)
		current *= 17
		current %= 256
	}
	return current
}

func partTwoSolution(input []string) uint64 {
	return 0
}
