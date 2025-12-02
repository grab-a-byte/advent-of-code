package day_01

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solution() {
	file, err := os.Open("./day_01/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day one part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day one part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	total := 0

	for _, line := range input {
		num := getNumber(line)
		total += num
	}

	return total
}

func partTwoSolution(input []string) uint64 {
	var morphedInput []string
	for _, line := range input {
		morphedInput = append(morphedInput, morphInput(line))
	}

	total := uint64(0)

	for _, line := range morphedInput {
		num := getNumber(line)
		total += uint64(num)
	}

	return total
}

func getNumber(s string) int {
	var nums []string
	for _, c := range s {
		if unicode.IsDigit(c) {
			nums = append(nums, string(c))
		}
	}

	output := nums[0] + nums[len(nums)-1]
	n, err := strconv.Atoi(output)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func morphInput(s string) string {
	morphed := s
	for key, value := range nameToIntConv {
		morphed = strings.Replace(morphed, key, value, -1)
	}
	return morphed
}

var nameToIntConv = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}
