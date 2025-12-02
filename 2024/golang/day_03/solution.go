package day_03

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	content, err := os.ReadFile("./day_03/input.txt")
	if err != nil {
		panic("unable to open day 3 input")
	}

	strContent := string(content)
	answer1 := partOneSolution(strContent)
	fmt.Printf("Day 3 - Part 1: %d \n", answer1)
	answer2 := partTwoSolution(strContent)
	fmt.Printf("Day 3 - Part 2: %d \n", answer2)
}

func partOneSolution(input string) int64 {
	pairings := []Pairing{}
	var position int = 0
	for position < len(input) {
		if position+4 > len(input) {
			break
		}
		start := position + strings.Index(input[position:], "mul(")
		position = start + 4
		if isNumber(input[position]) {
			//Parse first number
			start = position
			for isNumber(input[position]) {
				position++
			}
			firstNum, _ := strconv.Atoi(string(input[start:position]))
			if input[position] != ',' {
				continue
			}
			position++

			//Parse second number
			start = position
			for isNumber(input[position]) {
				position++
			}
			secondNum, _ := strconv.Atoi(string(input[start:position]))
			if input[position] != ')' {
				continue
			}
			position++
			pairings = append(pairings, Pairing{First: firstNum, Second: secondNum})
		}
	}

	var result int64 = 0
	for _, pair := range pairings {
		result += int64(pair.First) * int64(pair.Second)
	}
	return result
}

func isNumber(c byte) bool {
	return c <= '9' && c >= '0'
}

type Pairing struct {
	First, Second int
}

func partTwoSolution(input string) int64 {
	pairings := []Pairing{}
	var position int = 0
	for position < len(input) {
		if position+4 > len(input) {
			break
		}
		dontInstruction := math.MaxInt
		if strings.Index(input[position:], "don't()") >= 0 {
			dontInstruction = position + strings.Index(input[position:], "don't()")
		}
		start := position + strings.Index(input[position:], "mul(")
		if dontInstruction < start {
			//If don't instruction before start, skip to next do instruction
			position = dontInstruction + 7
			doInstruction := position + strings.Index(input[position:], "do()")
			//Skip over teh do instruction
			position = doInstruction + 4
			continue
		}
		position = start + 4
		if isNumber(input[position]) {
			//Parse first number
			start = position
			for isNumber(input[position]) {
				position++
			}
			firstNum, _ := strconv.Atoi(string(input[start:position]))
			if input[position] != ',' {
				continue
			}
			position++

			//Parse second number
			start = position
			for isNumber(input[position]) {
				position++
			}
			secondNum, _ := strconv.Atoi(string(input[start:position]))
			if input[position] != ')' {
				continue
			}
			position++
			pairings = append(pairings, Pairing{First: firstNum, Second: secondNum})
		}
	}

	var result int64 = 0
	for _, pair := range pairings {
		result += int64(pair.First) * int64(pair.Second)
	}
	return result
}
