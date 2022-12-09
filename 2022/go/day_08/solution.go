package day_08

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/parkeradam/aoc_2022/utils"
)

// TODO :- While it works, this can 100% be refactored and optimised. DO IT!

func Solution() {
	file, err := os.Open("./day_08/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 8")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

type Tree struct {
	height  int
	visible bool
}

func setEdgesVisible(trees [][]Tree) {
	for x := 0; x < len(trees); x++ {
		trees[x][0].visible = true
		trees[x][len(trees[0])-1].visible = true
	}

	for y := 0; y < len(trees[0]); y++ {
		trees[0][y].visible = true
		trees[len(trees)-1][y].visible = true
	}
}

func checkUp(trees [][]Tree, posX, posY, height int) bool {
	for y := 0; y < posY; y++ {
		treeHeight := trees[posX][y].height
		if treeHeight >= height {
			return false
		}
	}
	return true
}

func checkDown(trees [][]Tree, posX, posY, height int) bool {
	for y := posY + 1; y < len(trees[0]); y++ {
		treeHeight := trees[posX][y].height
		if treeHeight >= height {
			return false
		}
	}
	return true
}

func checkLeft(trees [][]Tree, posX, posY, height int) bool {
	for x := 0; x < posX; x++ {
		treeHeight := trees[x][posY].height
		if treeHeight >= height {
			return false
		}
	}
	return true
}

func checkRight(trees [][]Tree, posX, posY, height int) bool {
	for x := posX + 1; x < len(trees); x++ {
		treeHeight := trees[x][posY].height
		if treeHeight >= height {
			return false
		}
	}
	return true
}

func parseTrees(lines []string) [][]Tree {
	var trees [][]Tree = make([][]Tree, len(lines))

	for _, line := range lines {
		items := strings.Split(line, "")

		for x, val := range items {
			treeHeight, _ := strconv.Atoi(val)
			trees[x] = append(trees[x], Tree{height: treeHeight})
		}
	}
	return trees
}

func PartOneSolution(lines []string) int {
	var trees [][]Tree = make([][]Tree, len(lines))

	for _, line := range lines {
		items := strings.Split(line, "")

		for x, val := range items {
			treeHeight, _ := strconv.Atoi(val)
			trees[x] = append(trees[x], Tree{height: treeHeight})
		}
	}

	setEdgesVisible(trees)

	for x := 1; x < len(trees)-1; x++ {
		for y := 1; y < len(trees[0])-1; y++ {
			treeHeight := trees[x][y].height
			leftVisibility := checkLeft(trees, x, y, treeHeight)
			rightVisibility := checkRight(trees, x, y, treeHeight)
			upVisibility := checkUp(trees, x, y, treeHeight)
			downVisibility := checkDown(trees, x, y, treeHeight)

			if leftVisibility || rightVisibility || upVisibility || downVisibility {
				trees[x][y].visible = true
			}
		}
	}

	total := 0

	for x := 0; x < len(trees); x++ {
		for y := 0; y < len(trees[0]); y++ {
			if trees[x][y].visible {
				total++
			}
		}
	}

	return total
}

///////
///////
// PART 2
///////
///////

func checkUpCount(trees [][]Tree, posX, posY, height int) int {
	count := 0
	for y := posY - 1; y >= 0; y-- {
		treeHeight := trees[posX][y].height
		count++
		if treeHeight >= height {
			break
		}
	}
	return count
}

func checkDownCount(trees [][]Tree, posX, posY, height int) int {
	count := 0
	for y := posY + 1; y < len(trees[0]); y++ {
		treeHeight := trees[posX][y].height
		count++
		if treeHeight >= height {
			break
		}
	}
	return count
}

func checkLeftCount(trees [][]Tree, posX, posY, height int) int {
	count := 0
	for x := posX - 1; x >= 0; x-- {
		treeHeight := trees[x][posY].height
		count++
		if treeHeight >= height {
			break
		}
	}
	return count
}

func checkRightCount(trees [][]Tree, posX, posY, height int) int {
	count := 0
	for x := posX + 1; x < len(trees); x++ {
		treeHeight := trees[x][posY].height
		count++
		if treeHeight >= height {
			break
		}
	}
	return count
}

func PartTwoSolution(lines []string) int {
	trees := parseTrees(lines)

	maxSoFar := 0

	for x := 1; x < len(trees)-1; x++ {
		for y := 1; y < len(trees[0])-1; y++ {
			treeHeight := trees[x][y].height
			leftVisibility := checkLeftCount(trees, x, y, treeHeight)
			rightVisibility := checkRightCount(trees, x, y, treeHeight)
			upVisibility := checkUpCount(trees, x, y, treeHeight)
			downVisibility := checkDownCount(trees, x, y, treeHeight)

			scenicScore := leftVisibility * rightVisibility * upVisibility * downVisibility
			if scenicScore > maxSoFar {
				maxSoFar = scenicScore
			}
		}
	}

	return maxSoFar
}
