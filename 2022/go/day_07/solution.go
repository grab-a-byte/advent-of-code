package day_07

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_07/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 7")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

func parseDirectories(lines []string) map[string]int {

	directorySizes := make(map[string]int)

	cwd := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			dir := strings.Trim(line[5:], " ")
			cwd = path.Join(cwd, dir)
			if directorySizes[cwd] == 0 {
				directorySizes[cwd] = 0
			}
		} else if strings.HasPrefix(line, "$ ls") || strings.HasPrefix(line, "dir") {
			continue
		} else {
			file := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(file[0])
			directorySizes[cwd] += fileSize
		}
	}

	return directorySizes
}

func PartOneSolution(lines []string) uint64 {
	directorySizes := parseDirectories(lines)

	totalMap := make(map[string]int)

	for key, value := range directorySizes {
		total := value
		for innerKey, innerValue := range directorySizes {
			if key != innerKey && strings.HasPrefix(innerKey, key) {
				total += innerValue
			}
		}
		totalMap[key] = total
	}

	var finalTotal uint64
	for _, v := range totalMap {
		if v < 100_000 {
			finalTotal += uint64(v)
		}
	}

	return finalTotal
}

func PartTwoSolution(lines []string) int {
	return -1
}
