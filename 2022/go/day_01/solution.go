package day_01

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/parkeradam/aoc_2022/utils"
)

type Elf struct {
	items []int
}

func (elf *Elf) GetTotalCalories() int {
	var total int
	for _, item := range elf.items {
		total += item
	}

	return total
}

func Solution() {
	file, err := os.Open("./day_01/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 1")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

func PartOneSolution(lines []string) int {
	elves := getElves(lines)
	totalCals := utils.Map(elves, func(x *Elf) int { return x.GetTotalCalories() })
	return utils.Max(totalCals)
}

func PartTwoSolution(lines []string) int {
	elves := getElves(lines)
	totalCals := utils.Map(elves, func(x *Elf) int { return x.GetTotalCalories() })
	sort.Ints(totalCals)
	calsLength := len(totalCals)
	return totalCals[calsLength-1] + totalCals[calsLength-2] + totalCals[calsLength-3]
}

func getElves(lines []string) []Elf {
	var elves []Elf
	var currentElfItems []int
	for _, line := range lines {
		if len(line) == 0 {
			elves = append(elves, Elf{items: currentElfItems})
			currentElfItems = []int{}
			continue
		}
		val, _ := strconv.Atoi(line)
		currentElfItems = append(currentElfItems, val)
	}

	elves = append(elves, Elf{items: currentElfItems})
	return elves
}
