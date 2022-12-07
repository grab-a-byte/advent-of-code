package day_06

import (
	"fmt"
	"os"

	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_06/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 6")
	fmt.Println("Part 1:", PartOneSolution(lines[0]))
	fmt.Println("Part 2:", PartTwoSolution(lines[0]))
}

func findDisctinctIndex(line string, uniqueChars int) int {
	for i := range line {
		check := line[i : i+uniqueChars]
		values := make(map[rune]int, 0)

		for _, c := range check {
			values[c] = (values[c]) + 1
		}
		anyOver := false
		for _, v := range values {
			if v > 1 {
				anyOver = true
			}
		}
		if !anyOver {
			return i + uniqueChars
		}
	}
	return -1
}

func PartOneSolution(line string) int {
	return findDisctinctIndex(line, 4)
}

func PartTwoSolution(line string) int {
	return findDisctinctIndex(line, 14)
}
