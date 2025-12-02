package day_04

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_04/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 4")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

type Elf struct {
	startRange, endRange int
}

func (elf Elf) RangeSize() int {
	return (elf.endRange - elf.startRange) + 1
}

type ElfPair struct {
	elfOne, elfTwo Elf
}

func (pair ElfPair) OneCanContainOther() bool {
	var longerElf, shorterElf Elf
	if pair.elfOne.RangeSize() > pair.elfTwo.RangeSize() {
		longerElf = pair.elfOne
		shorterElf = pair.elfTwo
	} else {
		longerElf = pair.elfTwo
		shorterElf = pair.elfOne
	}

	return (longerElf.startRange <= shorterElf.startRange) && (longerElf.endRange >= shorterElf.endRange)
}

func (pair ElfPair) Overlaps() bool {
	var higherElf, lowerElf Elf
	if pair.elfOne.startRange >= pair.elfTwo.startRange {
		higherElf = pair.elfOne
		lowerElf = pair.elfTwo
	} else {
		higherElf = pair.elfTwo
		lowerElf = pair.elfOne
	}

	return lowerElf.endRange >= higherElf.startRange
}

func parseElf(line string) Elf {
	parts := strings.Split(line, "-")
	startInt, _ := strconv.Atoi(parts[0])
	endInt, _ := strconv.Atoi(parts[1])
	elf := Elf{startRange: startInt, endRange: endInt}
	return elf
}

func parseElfPairs(lines []string) []ElfPair {
	var elves []ElfPair
	for _, line := range lines {
		parts := strings.Split(line, ",")
		elfOne := parseElf(parts[0])
		elfTwo := parseElf(parts[1])
		pair := ElfPair{elfOne: elfOne, elfTwo: elfTwo}
		elves = append(elves, pair)
	}

	return elves
}

func PartOneSolution(lines []string) int {
	elves := parseElfPairs(lines)

	total := 0

	for _, pair := range elves {
		if pair.OneCanContainOther() {
			total++
		}
	}

	return total
}

func PartTwoSolution(lines []string) int {
	elves := parseElfPairs(lines)

	total := 0

	for _, pair := range elves {
		if pair.Overlaps() {
			total++
		}
	}

	return total
}
