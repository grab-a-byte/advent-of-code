package day_11

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_11/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 11")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

type operation func(int) int

type monkey struct {
	items       []int
	op          operation
	test        int
	trueMonkey  int
	falseMonkey int
}

func parseItems(item string) []int {
	parts := strings.Split(item, ": ")
	itemSet := strings.Split(parts[1], ",")
	return utils.Map(itemSet, func(x *string) int {
		trimmed := strings.Trim(*x, " ")
		val, _ := strconv.Atoi(trimmed)
		return val
	})
}

func parseOp(item string) operation {
	parts := strings.Split(item, "= old ")
	opParts := strings.Split(parts[1], " ")
	if opParts[1] == "old" {
		return func(x int) int {
			return x * x
		}
	}
	num, _ := strconv.Atoi(opParts[1])
	if opParts[0] == "*" {
		return func(x int) int {
			return x * num
		}
	} else if opParts[0] == "+" {
		return func(x int) int {
			return x + num
		}
	}
	return func(int) int { panic("NOT A VALID OPERATION") }
}

func parseInt(s string) int {
	trimmed := strings.Trim(s, " ")
	val, _ := strconv.Atoi(trimmed)
	return val
}

func parseMonkeys(lines []string) []monkey {
	var monkeys []monkey
	for i := 0; i <= len(lines)/7; i++ {
		startingIndex := i * 7
		currentSet := lines[startingIndex : startingIndex+7]
		newMonkey := monkey{
			items:       parseItems(currentSet[1]),
			op:          parseOp(currentSet[2]),
			test:        parseInt(strings.Split(strings.Trim(currentSet[3], " "), "Test: divisible by ")[1]),
			trueMonkey:  parseInt(strings.Split(strings.Trim(currentSet[4], " "), "If true: throw to monkey")[1]),
			falseMonkey: parseInt(strings.Split(strings.Trim(currentSet[5], " "), "If false: throw to monkey")[1]),
		}
		monkeys = append(monkeys, newMonkey)
	}

	return monkeys
}

func PartOneSolution(lines []string) int {
	monkeys := parseMonkeys(lines)
	inspectedItems := make(map[int]int, len(monkeys))

	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspectedItems[i] = inspectedItems[i] + 1
				newWorry := monkey.op(item)
				newWorry = newWorry / 3
				if (newWorry % monkey.test) == 0 {
					monkeys[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, newWorry)
				} else {
					monkeys[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, newWorry)
				}
			}
			monkeys[i].items = make([]int, 0)
		}
	}

	var nums []int
	for _, v := range inspectedItems {
		nums = append(nums, v)
	}

	sort.Ints(nums)
	return nums[len(nums)-1] * nums[len(nums)-2]
}

// func getDivisors(monkeys []monkey) int {
// 	divisors := make(map[int]struct{}, 0)
// 	var def struct{}
// 	for _, item := range monkeys {
// 		divisors[item.test] = def
// 	}

// 	total := 1
// 	for k := range divisors {
// 		total *= k
// 	}

// 	return total
// }

func PartTwoSolution(lines []string) int {
	return -1
	// monkeys := parseMonkeys(lines)
	// divisor := getDivisors(monkeys)

	// inspectedItems := make(map[int]int, len(monkeys))

	// for round := 0; round < 10_000; round++ {
	// 	for i, monkey := range monkeys {
	// 		for _, item := range monkey.items {
	// 			inspectedItems[i] = inspectedItems[i] + 1
	// 			newWorry := item % divisor
	// 			newWorry %= monkey.test
	// 			if newWorry == 0 {
	// 				monkeys[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, item)
	// 			} else {
	// 				monkeys[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, item)
	// 			}
	// 		}
	// 		monkeys[i].items = make([]int, 0)
	// 	}

	// 	// if round == 0 || round == 19 || round == 999 || round == 1999 || round == 2999 || round == 3999 {
	// 	var nums []int
	// 	for _, v := range inspectedItems {
	// 		nums = append(nums, v)
	// 	}
	// 	fmt.Println(round, nums)
	// 	// }
	// }

	// var nums []int
	// for _, v := range inspectedItems {
	// 	nums = append(nums, v)
	// }

	// sort.Ints(nums)
	// return nums[len(nums)-1] * nums[len(nums)-2]
}
