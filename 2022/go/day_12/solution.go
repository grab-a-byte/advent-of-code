package day_12

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"unicode"

	"github.com/parkeradam/aoc_2022/types"
	"github.com/parkeradam/aoc_2022/utils"
)

func Solution() {
	file, err := os.Open("./day_12/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 12")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

type point struct {
	height          int
	currentShortest int
	xPos, yPos      int
}

type points []*point

func (p *points) getStartPoint() *point {
	for _, point := range *p {
		if point.height == 'S' {
			return point
		}
	}
	return nil
}

func (p *points) getEndPoint() *point {
	for _, point := range *p {
		if point.height == 'E' {
			return point
		}
	}
	return nil
}

func (p *points) getPointChecked(posX, posY int, checkP *point) (*point, error) {
	for _, searchPoint := range *p {
		if searchPoint.xPos == posX && searchPoint.yPos == posY {
			calc := searchPoint.height - checkP.height
			if calc <= 1 || (searchPoint.height == 'E' && checkP.height == (int('z')-97)) {
				return searchPoint, nil
			} else {
				return nil, errors.New("point too high")
			}
		}
	}
	return nil, errors.New("no point")
}

func parseInput(lines []string) *points {
	var points points
	for y, line := range lines {
		for x, r := range line {

			var height int
			if unicode.IsUpper(r) {
				height = int(r)
			} else {
				height = int(r) - 97
			}

			points = append(points, &point{
				xPos:            x,
				yPos:            y,
				height:          height,
				currentShortest: maxInt,
			})
		}
	}

	return &points
}

func runSolutionFromPoint(p *points, xPos, yPos int) int {
	var point *point
	for _, x := range *p {
		if x.xPos == xPos && x.yPos == yPos {
			point = x
		}
	}

	return runSolution(p, point)
}

func runSolution(points *points, start *point) int {
	start.currentShortest = 0

	var pointStack types.Stack[*point]
	pointStack.Push(start)

	point, popped := pointStack.Pop()
	for popped {
		newEffort := point.currentShortest + 1
		left, err := points.getPointChecked(point.xPos-1, point.yPos, point)
		if err == nil && newEffort < left.currentShortest {
			left.currentShortest = newEffort
			pointStack.Push(left)
		}
		right, err := points.getPointChecked(point.xPos+1, point.yPos, point)
		if err == nil && newEffort < right.currentShortest {
			right.currentShortest = newEffort
			pointStack.Push(right)
		}
		up, err := points.getPointChecked(point.xPos, point.yPos-1, point)
		if err == nil && newEffort < up.currentShortest {
			up.currentShortest = newEffort
			pointStack.Push(up)
		}
		down, err := points.getPointChecked(point.xPos, point.yPos+1, point)
		if err == nil && newEffort < down.currentShortest {
			down.currentShortest = newEffort
			pointStack.Push(down)
		}

		point, popped = pointStack.Pop()
	}

	endPoint := points.getEndPoint()
	return endPoint.currentShortest
}

func PartOneSolution(lines []string) int {
	points := parseInput(lines)
	start := points.getStartPoint()
	return runSolution(points, start)
}

// NOT OPTIMISED!! Tales about 15 mins to run, could be parralellised
func PartTwoSolution(lines []string) int {
	points := parseInput(lines)
	possibleRouteLength := make([]int, 0)
	for _, point := range *points {
		if point.height == 0 {
			newPoints := parseInput(lines)
			result := runSolutionFromPoint(newPoints, point.xPos, point.yPos)
			possibleRouteLength = append(possibleRouteLength, result)
		}
	}

	sort.Ints(possibleRouteLength)

	return possibleRouteLength[0]
}
