package day_01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solution() {
	content, err := os.ReadFile("./day_01/input.txt")
	if err != nil {
		panic("unable to open day 1 input")
	}

	strContent := string(content)
	answer1 := partOneSolution(strContent)
	fmt.Printf("Day 1 - Part 1: %d \n", answer1)
	answer2 := partTwoSolution(strContent)
	fmt.Printf("Day 1 - Part 2: %d \n", answer2)
}

func partOneSolution(input string) int64 {
	arr1, arr2 := getArrays(input)
	sort.Ints(arr1)
	sort.Ints(arr2)

	var total int64 = 0
	for i := 0; i < len(arr1); i++ {
		difference := arr1[i] - arr2[i]
		if difference < 0 {
			difference = -difference
		}
		total += int64(difference)
		fmt.Println(difference)
	}
	return int64(total)
}

func partTwoSolution(input string) int64 {
	arr1, arr2 := getArrays(input)
	counts := map[int]int{}
	for _, num := range arr2 {
		counts[num]++
	}

	var total int64 = 0
	for _, value := range arr1 {
		total += int64(counts[value] * value)
	}

	return total
}

func getArrays(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	arr1 := make([]int, len(lines))
	arr2 := make([]int, len(lines))

	for i, l := range lines {
		nums := strings.Split(l, "   ")
		if len(nums) != 2 {
			fmt.Println("missing in index" + strconv.Itoa(i))
			continue
		}
		strA := strings.ReplaceAll(nums[0], "\r", "")
		strB := strings.ReplaceAll(nums[1], "\r", "")
		num1, _ := strconv.Atoi(strA)
		num2, _ := strconv.Atoi(strB)
		arr1[i] = num1
		arr2[i] = num2
	}
	return arr1, arr2
}
