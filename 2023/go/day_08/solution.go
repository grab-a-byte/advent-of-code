package day_08

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
)

func Solution() {
	file, err := os.Open("./day_08/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day eight part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day eight part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	navigation := strings.Repeat(input[0], 100000)
	network := getNetwork(input[2:])
	steps := calculateSteps("AAA", navigation, network, func(s string) bool { return s == "ZZZ" })
	return steps
}

func partTwoSolution(input []string) int {
	navigation := strings.Repeat(input[0], 100000)
	network := getNetwork(input[2:])

	nodes := []string{}
	for key := range network {
		if strings.HasSuffix(key, "A") {
			nodes = append(nodes, key)
		}
	}

	stepRange := make([]int, len(nodes))
	for i, node := range nodes {
		stepRange[i] = calculateSteps(node, navigation, network, func(s string) bool { return strings.HasSuffix(s, "Z") })
	}

	if len(stepRange) < 3 {
		return stepRange[0] * stepRange[1]
	}

	result := LCM(stepRange[0], stepRange[1], stepRange[2:]...)
	return result
}

func calculateSteps(start, navigation string, network map[string]pair, endFunc func(string) bool) int {
	currentNode := start
	steps := 0
	for i, n := range navigation {
		if n == 'L' {
			currentNode = network[currentNode].left
		} else if n == 'R' {
			currentNode = network[currentNode].right
		}

		if endFunc(currentNode) {
			steps = i + 1
			break
		}
	}
	return steps
}

func getNetwork(networkLines []string) map[string]pair {
	network := map[string]pair{}
	for _, line := range networkLines {
		var start, left, right string
		_, err := fmt.Sscanf(line, "%s = (%3s, %3s)", &start, &left, &right)
		if err != nil {
			panic("failed to parse a network line")
		}
		network[start] = pair{
			left:  left,
			right: right,
		}
	}
	return network
}

func allEndWithZ(input []string) bool {
	for _, i := range input {
		if !strings.HasSuffix(i, "Z") {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type pair struct {
	left, right string
}
