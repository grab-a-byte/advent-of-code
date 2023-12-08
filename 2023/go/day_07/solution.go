package day_07

import (
	"aoc/util"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("./day_07/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day seven part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day seven part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {
	hands := []hand{}
	for _, l := range input {
		hands = append(hands, parseHand(l))
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if a.handType != b.handType {
			return int(b.handType) - int(a.handType)
		}
		for i, c := range a.handStr {
			if c == rune(b.handStr[i]) {
				continue
			}
			return slices.Index(cardTypes, string(b.handStr[i])) - slices.Index(cardTypes, string(c))
		}
		return 0
	})

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bet
	}

	return total
}

// TODO: INCOMPLETE PART 2
func partTwoSolution(input []string) int {
	return 0
}

type handType int

const (
	FIVE     handType = 0
	FOUR     handType = 1
	FULL     handType = 2
	THREE    handType = 3
	TWO_PAIR handType = 4
	ONE_PAIR handType = 5
	HIGH     handType = 6
)

var cardTypes []string = []string{
	"A",
	"K",
	"Q",
	"J",
	"T",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
}

type hand struct {
	handStr  string
	bet      int
	handType handType
}

func calculateHandType(s string) handType {
	bestType := HIGH
	for _, r := range cardTypes {
		// fmt.Printf("Chekcing char %q current is %v \n", r, bestType)
		switch strings.Count(s, string(r)) {
		case 5:
			return FIVE
		case 4:
			return FOUR
		case 3:
			if bestType == ONE_PAIR {
				return FULL
			}
			bestType = THREE
		case 2:
			if bestType == THREE {
				return FULL
			} else if bestType == ONE_PAIR {
				return TWO_PAIR
			}
			bestType = ONE_PAIR
		}
	}
	return bestType
}

func parseHand(line string) hand {
	parts := strings.Split(line, " ")
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("unable to parse bet")
	}
	return hand{
		handStr:  parts[0],
		bet:      val,
		handType: calculateHandType(parts[0]),
	}
}
