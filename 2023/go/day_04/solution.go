package day_04

import (
	"aoc/util"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Solution() {
	file, err := os.Open("./day_04/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day four part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day four part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	total := 0
	for _, line := range input {
		nums := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning, ourNumbers := toNums(nums[0]), toNums(nums[1])
		total += scoreCard(winning, ourNumbers)
	}
	return total
}

func partTwoSolution(input []string) int {
	copies := make(map[int]int, len(input))
	for i, line := range input {
		nums := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning, ourNumbers := toNums(nums[0]), toNums(nums[1])
		//this card
		copies[i]++
		matches := cardMatches(winning, ourNumbers)
		for c := i + 1; c <= i+matches; c++ {
			copies[c] += copies[i]
		}
	}
	total := 0
	for _, val := range copies {
		total += val
	}
	return total
}

func cardMatches(winning, ours []int) int {
	matches := 0
	for _, n := range ours {
		if slices.Contains(winning, n) {
			matches += 1
		}
	}

	return matches
}

func scoreCard(winning, ours []int) int {
	matches := cardMatches(winning, ours)
	if matches > 0 {
		return int(math.Pow(2, float64(matches-1)))
	}
	return 0
}

func toNums(input string) []int {
	items := filterEmpty(strings.Split(input, " "))
	output := make([]int, len(items))
	for i, item := range items {
		stripped := stripWhitespace(item)
		val, _ := strconv.Atoi(stripped)
		output[i] = val
	}
	return output
}

func filterEmpty(in []string) []string {
	var out []string
	for _, s := range in {
		if len(s) == 0 {
			continue
		}
		out = append(out, stripWhitespace(s))
	}
	return out
}

func stripWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
