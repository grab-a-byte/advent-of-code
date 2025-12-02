package day_02

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solution() {
	content, err := os.ReadFile("./day_02/input.txt")
	if err != nil {
		panic("unable to open day 1 input")
	}

	strContent := string(content)
	answer1 := partOneSolution(strContent)
	fmt.Printf("Day 2 - Part 1: %d \n", answer1)
	answer2 := partTwoSolution(strContent)
	fmt.Printf("Day 2 - Part 2: %d \n", answer2)
}

func partOneSolution(input string) int {
	sanatizedInput := strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(sanatizedInput, "\n")

	var result int = 0
	for _, l := range lines {
		parts := strings.Fields(l)
		nums := make([]int, len(parts))
		for i := 0; i < len(parts); i++ {
			num, _ := strconv.Atoi(parts[i])
			nums[i] = num
		}
		res := lineProceessor(nums)
		if res == 0 {
			result++
		}
	}
	return result
}

// TODO: Can 100% be optimized and read better but oh well
func lineProceessor(line []int) int {
	errors := 0
	if line[0] < line[1] {
		//Increasing
		for i := 0; i < len(line)-1; i++ {
			if line[i]-line[i+1] >= 0 {
				errors += 1
			}

			if math.Abs(float64(line[i]-line[i+1])) > 3 {
				errors += 1
			}
		}
		return errors
	}

	if line[1] < line[0] {
		//Increasing
		for i := 0; i < len(line)-1; i++ {
			if line[i+1]-line[i] >= 0 {
				errors += 1
			}

			if math.Abs(float64(line[i+1]-line[i])) > 3 {
				errors += 1
			}
		}
		return errors
	}

	return 1
}

func partTwoSolution(input string) int {
	sanatizedInput := strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(sanatizedInput, "\n")

	var result int = 0
	for _, l := range lines {
		parts := strings.Fields(l)
		nums := make([]int, len(parts))
		for i := 0; i < len(parts); i++ {
			num, _ := strconv.Atoi(parts[i])
			nums[i] = num
		}
		res := lineProceessor(nums)
		if res == 0 {
			result++
		} else {
			//Create arrays for which 1 element is removed from each
			newChecks := [][]int{}
			for i := 0; i < len(nums); i++ {
				before := nums[:i]
				after := nums[i+1:]
				newChecks = append(newChecks, slices.Concat(before, after))
			}
			for _, newLine := range newChecks {
				res = lineProceessor(newLine)
				if res == 0 {
					result++
					break
				}
			}
		}
	}
	return result
}
