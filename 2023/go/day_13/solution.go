package day_13

import (
	"aoc/util"
	"fmt"
	"os"
)

func Solution() {
	file, err := os.Open("./day_13/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day thirteen part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day thirteen part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	grids := getGrids(input)
	total := 0
	for _, grid := range grids {
		result := processGrid(grid)
		total += result
	}
	return total
}

func partTwoSolution(input []string) int {
	return 0
}

func processGrid(g grid) int {
	//Check for vertical mirror
	for i := 0; i < len(g.lines[0])-1; i++ {
		res := checkLineVertical(g, i)
		// fmt.Println(res)
		if res {
			return i + 1 // + 1 as 0 index based but puzzle 1 index based
		}
	}

	//Check for horizontal mirror
	for i := 0; i < len(g.lines)-1; i++ {
		res := checkLineHorizontal(g, i)
		if res {
			return (i + 1) * 100
		}
		// fmt.Printf("Index %d found %v\n", i, res)
	}

	return 0
}

func checkLineVertical(g grid, i int) bool {
	for c := 0; c <= len(g.lines[0]); c++ {
		//Some wierd hack to stop it going over ends, dont ask
		if i+c >= len(g.lines[0])-1 || i-c < 0 {
			return true
		}
		leftCol := g.getColumn(i - c)
		rightCol := g.getColumn(i + 1 + c)
		if rightCol != leftCol {
			return false
		}
	}
	return true
}

func checkLineHorizontal(g grid, i int) bool {
	for c := 0; c <= len(g.lines); c++ {
		//Some wierd hack to stop it going over ends, dont ask
		if i+c >= len(g.lines)-1 || i-c < 0 {
			return true
		}
		leftRow := g.getRow(i - c)
		rightRow := g.getRow(i + 1 + c)
		if rightRow != leftRow {
			return false
		}
	}
	return true
}

func getGrids(lines []string) []grid {
	var inputGrids [][]string
	currentGrid := []string{}
	for _, line := range lines {
		if len(line) != 0 {
			currentGrid = append(currentGrid, line)
		} else {
			inputGrids = append(inputGrids, currentGrid)
			currentGrid = []string{}
		}
	}
	inputGrids = append(inputGrids, currentGrid)
	grids := []grid{}
	for _, g := range inputGrids {
		grids = append(grids, grid{lines: g})
	}
	return grids
}

type grid struct {
	lines []string
}

func (g *grid) getRow(index int) string {
	return g.lines[index]
}

func (g *grid) getColumn(index int) string {
	column := []rune{}
	for _, col := range g.lines {
		runes := []rune(col)
		column = append(column, runes[index])
	}
	return string(column)
}
