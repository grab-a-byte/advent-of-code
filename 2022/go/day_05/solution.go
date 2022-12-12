package day_05

import (
	"fmt"
	"os"

	"github.com/parkeradam/aoc_2022/types"
	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_05/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 5")
	fmt.Println("Part 1:", PartOneSolution(lines, 9, 8))
	fmt.Println("Part 2:", PartTwoSolution(lines, 9, 8))
}

func parseStacks(length, height int, stackLines []string) []*types.Stack[string] {
	var stacks []*types.Stack[string]

	var parsed [][]string
	for i := 0; i < length; i++ {
		stack := utils.Map(stackLines, func(line *string) string {
			index := (4 * i) + 1
			str := *line
			return string(str[index])
		})

		parsed = append(parsed, stack)
	}

	for i := 0; i < length; i++ {
		var stack types.Stack[string]
		item := parsed[i]
		for i := len(item) - 1; i >= 0; i-- {
			char := item[i]
			if char == " " {
				continue
			}
			stack.Push(item[i])
		}
		stacks = append(stacks, &stack)
	}
	return stacks
}

type Instruction struct {
	from, to, count int
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction
	for _, line := range lines {
		var from, to, count int
		fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)

		instructions = append(instructions, Instruction{from: from, to: to, count: count})
	}
	return instructions
}

func PartOneSolution(lines []string, length, height int) string {
	stacks := parseStacks(length, height, lines[:height])
	instructions := parseInstructions(lines[height+2:])

	for _, ins := range instructions {
		fromStack := stacks[ins.from-1]
		toStack := stacks[ins.to-1]

		for i := 0; i < ins.count; i++ {
			val, _ := fromStack.Pop()
			toStack.Push(val)
		}

	}

	var output string
	for _, stack := range stacks {
		val, _ := stack.Pop()
		output += val
	}

	return output
}

func PartTwoSolution(lines []string, length, height int) string {
	stacks := parseStacks(length, height, lines[:height])
	instructions := parseInstructions(lines[height+2:])

	for _, ins := range instructions {
		fromStack := stacks[ins.from-1]
		toStack := stacks[ins.to-1]

		var toAdd []string
		for i := 0; i < ins.count; i++ {
			val, _ := fromStack.Pop()
			toAdd = append(toAdd, val)
		}

		for i := len(toAdd) - 1; i >= 0; i-- {
			toStack.Push(toAdd[i])
		}

	}

	var output string
	for _, stack := range stacks {
		val, _ := stack.Pop()
		output += val
	}

	return output
}
