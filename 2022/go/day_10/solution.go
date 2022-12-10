package day_10

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_10/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 10")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:")
	fmt.Println(PartTwoSolution(lines))
}

func parseCycles(lines []string) []int {
	var cycles []int

	xRegister := 1
	cycles = append(cycles, xRegister)
	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			cycles = append(cycles, xRegister)
		} else if strings.HasPrefix(line, "addx") {
			cycles = append(cycles, xRegister)
			value, _ := strconv.Atoi(strings.Split(line, " ")[1])
			xRegister += value
			cycles = append(cycles, xRegister)
		}
	}

	return cycles
}

func PartOneSolution(lines []string) int {
	cycles := parseCycles(lines)
	total := 0
	for _, x := range []int{20, 60, 100, 140, 180, 220} {
		total += (cycles[x-1] * x)
	}
	return total
}

func PartTwoSolution(lines []string) string {
	cycles := parseCycles(lines)
	var renderedLines []string

	currentScreen := ""
	for c := 0; c < len(cycles); c++ {
		pixel := len(currentScreen)
		currentPosition := cycles[c]
		if (pixel == currentPosition) || (pixel == currentPosition-1) || (pixel == currentPosition+1) {
			currentScreen += "#"
		} else {
			currentScreen += "."
		}
		if len(currentScreen) == 40 {
			renderedLines = append(renderedLines, currentScreen)
			currentScreen = ""
		}
	}
	return strings.Join(renderedLines, "\r\n")
}
