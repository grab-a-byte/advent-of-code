package day_06

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("./day_06/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day six part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day six part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	times := getInts(input[0])
	distances := getInts(input[1])
	if len(times) != len(distances) {
		panic("unable to run")
	}

	waysToWin := []int{}
	for i := 0; i < len(times); i++ {
		time := times[i]
		record := distances[i]
		_ = record
		recordAttempts := 0
		for h := 0; h <= time; h++ {
			dist := h * (time - h)
			if dist > record {
				recordAttempts++
			}
		}
		waysToWin = append(waysToWin, recordAttempts)
	}

	output := waysToWin[0]
	for _, i := range waysToWin[1:] {
		output *= i
	}
	return output
}

func getInts(s string) []int {
	numStr := strings.Split(strings.Split(s, ":")[1], " ")
	nums := []int{}
	for _, n := range numStr {
		n, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		nums = append(nums, n)
	}
	return nums
}

func partTwoSolution(input []string) int {
	time, err := strconv.Atoi(strings.ReplaceAll(strings.Split(input[0], ":")[1], " ", ""))
	if err != nil {
		panic("unable to parse time")
	}
	record, err := strconv.Atoi(strings.ReplaceAll(strings.Split(input[1], ":")[1], " ", ""))
	if err != nil {
		panic("unable to parse distance")
	}

	recordAttempts := 0
	for h := 0; h <= time; h++ {
		dist := h * (time - h)
		if dist > record {
			recordAttempts++
		}
	}
	return recordAttempts
}
