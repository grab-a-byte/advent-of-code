package day_04

import (
	"fmt"
	"os"
	"strings"
)

func Solution() {
	content, err := os.ReadFile("./day_04/input.txt")
	if err != nil {
		panic("unable to open day 4 input")
	}

	strContent := string(content)
	answer1 := partOneSolution(strContent)
	fmt.Printf("Day 4 - Part 1: %d \n", answer1)
	answer2 := partTwoSolution(strContent)
	fmt.Printf("Day 4 - Part 2: %d \n", answer2)
}

func partOneSolution(input string) int64 {
	var result int = 0
	sanatizedInput := strings.ReplaceAll(input, "\r\n", "\n")
	rawRows := strings.Split(sanatizedInput, "\n")
	rows := getRows(rawRows, 4)
	//Check Rows
	for _, row := range rows {
		if strings.Contains(row, "XMAS") || strings.Contains(row, "SAMX") {
			result++
		}
	}
	//Check colums
	columns := getColumns(rawRows, 4)
	for _, column := range columns {
		if strings.Contains(column, "XMAS") || strings.Contains(column, "SAMX") {
			result++
		}
	}
	//Check Diagonals
	diagonals := getDiagonals(rawRows, 4)
	for _, diag := range diagonals {
		if strings.Contains(diag, "XMAS") {
			result++
		}
	}

	return int64(result)
}

func getDiagonals(input []string, size int) []string {
	//Sometime -1 to size as index of 3 allowed to go (3,2,1,0), size 4 but index 3
	diagonals := []string{}
	for ri, column := range input {
		for ci := range column {
			//Down right
			if len(column) >= ci+size && len(input) >= ri+size {
				chars := []byte{}
				for i := 0; i < size; i++ {
					chars = append(chars, input[ci+i][ri+i])
				}
				diagonals = append(diagonals, string(chars))
			}
			//Down left
			if len(column) >= ci+size && 0 <= ri-(size-1) {
				chars := []byte{}
				for i := 0; i < size; i++ {
					chars = append(chars, input[ci+i][ri-i])
				}
				diagonals = append(diagonals, string(chars))
			}
			//Up right
			if 0 <= ci-(size-1) && len(input) >= ri+size {
				chars := []byte{}
				for i := 0; i < size; i++ {
					chars = append(chars, input[ci-i][ri+i])
				}
				diagonals = append(diagonals, string(chars))
			}
			//Up left
			if 0 <= ci-(size-1) && 0 <= ri-(size-1) {
				chars := []byte{}
				for i := 0; i < size; i++ {
					chars = append(chars, input[ci-i][ri-i])
				}
				diagonals = append(diagonals, string(chars))
			}
		}
	}
	return diagonals
}

func getColumns(input []string, length int) []string {
	columns := []string{}
	for ci := range input[0] {
		for ri := range input {
			if ri+length > len(input) {
				continue
			}
			column := []byte{}
			for l := 0; l < length; l++ {
				column = append(column, input[ri+l][ci])
			}
			columns = append(columns, string(column))
		}
	}
	return columns
}

func getRows(input []string, length int) []string {
	rows := []string{}
	for ri := range input {
		for ci := range input[0] {
			if ci+length > len(input[0]) {
				continue
			}
			row := []byte{}
			for l := 0; l < length; l++ {
				row = append(row, input[ri][ci+l])
			}
			rows = append(rows, string(row))
		}
	}
	return rows
}

func partTwoSolution(input string) int64 {
	var result int = 0
	sanatizedInput := strings.ReplaceAll(input, "\r\n", "\n")
	rawRows := strings.Split(sanatizedInput, "\n")
	for x := 1; x < len(rawRows[0])-1; x++ {
		for y := 1; y < len(rawRows)-1; y++ {
			if rawRows[y][x] != 'A' {
				continue
			}
			topLeftMas := rawRows[y-1][x-1] == 'M' && rawRows[y+1][x+1] == 'S'
			topLeftSam := rawRows[y-1][x-1] == 'S' && rawRows[y+1][x+1] == 'M'
			topRightMas := rawRows[y-1][x+1] == 'M' && rawRows[y+1][x-1] == 'S'
			topRightSam := rawRows[y-1][x+1] == 'S' && rawRows[y+1][x-1] == 'M'
			if (topLeftMas || topLeftSam) && (topRightMas || topRightSam) {
				result++
			}
		}
	}

	return int64(result)
}
