package day_02

import (
	"fmt"
	"os"
	"strings"

	"github.com/parkeradam/aoc_2022/utils"
	"golang.org/x/exp/slices"
)

func Solution() {
	file, err := os.Open("./day_02/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := utils.ReadFileAsLines(file)
	fmt.Println("DAY 2")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

func PartOneSolution(lines []string) int {

	var total int = 0

	for _, val := range lines {
		values := strings.Split(val, " ")
		opponentCall := values[0]
		ourCall := values[1]

		total += scorePoint(opponentCall, ourCall)
	}

	return total
}

type Game struct {
	ourCall      string
	opponentCall string
	value        int
}

func PartTwoSolution(lines []string) int {
	var total int = 0

	for _, val := range lines {
		values := strings.Split(val, " ")
		opponentCall := values[0]
		ourCall := values[1]

		total += scorePointGuesses(opponentCall, ourCall)
	}

	return total
}

func scorePoint(opponentCall, ourCall string) int {
	possibleGames := []Game{
		{"X", "A", 4},
		{"Y", "B", 5},
		{"Z", "C", 6},
		{"X", "B", 1},
		{"X", "C", 7},
		{"Y", "A", 8},
		{"Y", "C", 2},
		{"Z", "A", 3},
		{"Z", "B", 9},
	}

	index := slices.IndexFunc(possibleGames, func(g Game) bool {
		return g.opponentCall == opponentCall && g.ourCall == ourCall
	})

	return possibleGames[index].value
}

func scorePointGuesses(opponentCall, ourCall string) int {
	possibleGames := []Game{
		{"X", "A", 3},
		{"Y", "B", 5},
		{"Z", "C", 7},
		{"X", "B", 1},
		{"X", "C", 2},
		{"Y", "A", 4},
		{"Y", "C", 6},
		{"Z", "A", 8},
		{"Z", "B", 9},
	}

	index := slices.IndexFunc(possibleGames, func(g Game) bool {
		return g.opponentCall == opponentCall && g.ourCall == ourCall
	})

	return possibleGames[index].value
}
