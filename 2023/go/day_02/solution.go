package day_02

import (
	"aoc/util"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("./day_02/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day two part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day two part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	games := parseGames(input)
	// fmt.Println(games)
	return gamesWithMax(games, 12, 14, 13)
}

func partTwoSolution(input []string) int {
	games := parseGames(input)
	fullTotal := 0
	for _, g := range games {
		red := g.Max("red")
		blue := g.Max("blue")
		green := g.Max("green")
		total := red * blue * green
		fullTotal += total
	}
	return fullTotal
}

func gamesWithMax(games []Game, maxRed, maxBlue, maxGreen int) int {
	idTotal := 0
	for _, g := range games {
		if g.Max("blue") > maxBlue {
			continue
		}
		if g.Max("red") > maxRed {
			continue
		}
		if g.Max("green") > maxGreen {
			continue
		}
		idTotal += g.Id
	}

	return idTotal
}

func parseGames(input []string) []Game {
	var games []Game
	for _, line := range input {
		game := Game{}
		var parts = strings.Split(line, ": ")
		game.Id, _ = strconv.Atoi(strings.Split(parts[0], " ")[1])
		rounds := strings.Split(parts[1], "; ")
		//3 blue, 4 red
		for _, round := range rounds {
			roundMap := make(map[string]int)
			for _, part := range strings.Split(round, ", ") {
				//3 blue
				innerparts := strings.Split(part, " ")
				amount, name := innerparts[0], innerparts[1]
				amountInt, _ := strconv.Atoi(amount)
				roundMap[name] = amountInt
			}
			game.Rounds = append(game.Rounds, roundMap)
		}
		games = append(games, game)
	}
	return games
}

type Game struct {
	Id     int
	Rounds []map[string]int
}

func (g *Game) Max(colour string) int {
	var values []int
	for _, round := range g.Rounds {
		if _, ok := round[colour]; ok {
			values = append(values, round[colour])
		}
	}
	return slices.Max(values)
}
