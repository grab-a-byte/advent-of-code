package day_03

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 3")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

func PartOneSolution(lines []string) int {
	var total int
	for _, line := range lines {
		lineLength := len(line)
		halfWidth := lineLength / 2
		compartment1 := line[:halfWidth]
		compartment2 := line[halfWidth:]
		intersections := intersection(compartment1, compartment2)

		for _, c := range intersections {
			total += toValue(c)
		}
	}

	return total
}

func toValue(char rune) int {
	if unicode.IsLower(char) {
		return int(char) - 96
	} else {
		return int(char) - (65 - 27)
	}
}

func intersection(a, b string) []rune {
	set := make(map[rune]bool)
	for _, c := range a {
		for _, c2 := range b {
			if c2 == c {
				set[c] = true
			}
		}
	}

	var keys []rune
	for r, _ := range set {
		keys = append(keys, r)
	}

	return keys
}

func PartTwoSolution(lines []string) int {
	var total int

	for i := 0; i < len(lines)/3; i++ {
		index := i * 3
		for _, c := range lines[index] {
			nextContains := strings.ContainsRune(lines[index+1], c)
			secondNextContains := strings.ContainsRune(lines[index+2], c)

			if nextContains && secondNextContains {
				total += toValue(c)
				break
			}
		}
	}

	return total
}
