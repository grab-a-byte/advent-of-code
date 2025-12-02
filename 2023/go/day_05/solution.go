package day_05

import (
	"aoc/util"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("./day_05/input.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		panic("input file was unable to be opened")
	}
	lines := util.ReadFileAsLines(file)
	answer := partOneSolution(lines)
	fmt.Printf("Day five part 1 answer is : %d \n", answer)
	answer2 := partTwoSolution(lines)
	fmt.Printf("Day five part 2 answer is : %d \n", answer2)
}

func partOneSolution(input []string) int {

	seedsStr := strings.Split(strings.Split(input[0], ": ")[1], " ")
	var seeds []int
	for _, s := range seedsStr {
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		seeds = append(seeds, val)
	}

	return getLowestLocation(input, seeds)
}

func partTwoSolution(input []string) int {
	seedsStr := strings.Split(strings.Split(input[0], ": ")[1], " ")
	var seeds []int
	for i := 0; i < len(seedsStr); i += 2 {
		// 	fmt.Printf("running %d of %d \n", i, len(seedsStr))
		start, _ := strconv.Atoi(seedsStr[i])
		count, _ := strconv.Atoi(seedsStr[i+1])
		innerSeeds := make([]int, count)
		for s := 0; s < count; s++ {
			innerSeeds[s] = start + s
		}
		seeds = append(seeds, innerSeeds...)
	}
	return getLowestLocation(input, seeds)
}

func getLowestLocation(input []string, seeds []int) int {
	seedLine := slices.IndexFunc(input, func(x string) bool { return len(x) == 0 })
	seedToSoil, stsIndex := getSection(input, seedLine)
	seedToSoilMap := makeMap(seedToSoil)
	soilToFertilizer, stfIndex := getSection(input, stsIndex)
	soilToFertilizerMap := makeMap(soilToFertilizer)
	fertilizerToWater, ftwIndex := getSection(input, stfIndex)
	fertToWaterMap := makeMap(fertilizerToWater)
	waterToLight, wtlIndex := getSection(input, ftwIndex)
	waterToLightMap := makeMap(waterToLight)
	lightToTemp, lttIndex := getSection(input, wtlIndex)
	lightToTempMap := makeMap(lightToTemp)
	tempToHumid, tthIndex := getSection(input, lttIndex)
	tempToHumidMap := makeMap(tempToHumid)
	humidToLoc := input[tthIndex+1:]
	humidToLocMap := makeMap(humidToLoc)

	lowest := -1000
	// numSeeds := len(seeds)
	for _, seed := range seeds {
		// 	fmt.Printf("Seed %d of %d ,%.2f percent\n", i, numSeeds, float32(i)/float32(numSeeds)*100)
		soil := seedToSoilMap.Get(seed)
		fert := soilToFertilizerMap.Get(soil)
		water := fertToWaterMap.Get(fert)
		light := waterToLightMap.Get(water)
		temp := lightToTempMap.Get(light)
		humid := tempToHumidMap.Get(temp)
		final := humidToLocMap.Get(humid)
		if lowest < 0 {
			lowest = final
		} else if lowest > final {
			lowest = final
		}
	}

	return lowest
}

type source struct {
	sourceStart, offset, count int
}
type selfMap struct {
	sources []source
}

func (sm selfMap) Get(idx int) int {
	for _, s := range sm.sources {
		if idx >= s.sourceStart && idx < s.sourceStart+s.count {
			return idx + s.offset
		}
	}
	return idx
}

func makeMap(input []string) selfMap {
	m := selfMap{
		sources: []source{},
	}
	//Skip header line
	for _, line := range input[1:] {
		s := source{}
		var dest, src, num int
		_, err := fmt.Sscanf(line, "%d %d %d", &dest, &src, &num)
		if err != nil {
			fmt.Println(err)
		}
		s.sourceStart = src
		s.count = num
		s.offset = dest - src
		m.sources = append(m.sources, s)
	}

	return m
}

func getSection(input []string, startIndex int) ([]string, int) {
	// fmt.Println(input[startIndex+1:])
	nextIndex := startIndex + 1 + slices.IndexFunc(input[startIndex+1:], func(x string) bool { return len(x) == 0 || x[0] == 0 })
	// fmt.Println(startIndex, nextIndex)
	if nextIndex < startIndex {
		return input[startIndex:], -1
	}
	return input[startIndex+1 : nextIndex], nextIndex
}
