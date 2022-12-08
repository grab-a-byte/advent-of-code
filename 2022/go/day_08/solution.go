package day_08

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/parkeradam/aoc_2022/utils"
)

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

type Item struct {
	height, posX, posY int
}

func checkItem(
	item Item,
	trees [][]Tree,
	startX int,
	startY int,
	endX int,
	endY int,
) bool {
	visible := true
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			tree := trees[x][y]
			if tree.height >= item.height {
				visible = false
			}
		}
	}

	return visible
}

func checkHorizontal(item Item, trees [][]Tree) bool {
	visibleLeft := checkItem(item, trees, 0, item.posY, item.posX-1, item.posY)
	visibleRight := checkItem(item, trees, item.posX+1, item.posY, len(trees)-1, item.posY)

	return visibleLeft || visibleRight
}

func checkVertical(item Item, trees [][]Tree) bool {
	visibleUp := checkItem(item, trees, item.posX, 0, item.posX+1, item.posY-1)
	visibleDown := checkItem(item, trees, item.posX, item.posY+1, item.posX, len(trees[0])-1)

	return visibleDown || visibleUp
}

func check(item Item, trees [][]Tree) bool {
	return checkHorizontal(item, trees) || checkVertical(item, trees)
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

func PartOneSolution(lines []string) int {
	var trees [][]Tree

	for _, line := range lines {
		items := strings.Split(line, "")
		lineNums := utils.Map(items, func(s *string) Tree {
			x, _ := strconv.Atoi(*s)
			return Tree{height: x}
		})
		trees = append(trees, lineNums)
	}

	for x := 1; x < len(trees)-1; x++ {
		for y := 1; y < len(trees[0])-1; y++ {
			isVisible := check(Item{trees[x][y].height, x, y}, trees)
			trees[x][y].visible = isVisible
		}
	}

	setEdgesVisible(trees)

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

func PartTwoSolution(lines []string) int {
	return -1
}
