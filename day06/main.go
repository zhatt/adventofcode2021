package main

import (
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"
)

func parseInput(inputList []string) []int {
	data := make([]int, 0)

	vals := strings.Split(inputList[0], ",")
	for _, valStr := range vals {
		val, err := strconv.Atoi(valStr)
		aoc.PanicOnError(err)
		data = append(data, val)
	}

	return data
}

func simulate(days int, fishes []int) int {
	const (
		recycleTimer = 6
		newFishTimer = recycleTimer + 2
		maxFishTimer = recycleTimer + 2
	)

	var data [maxFishTimer + 1]int

	for _, fish := range fishes {
		data[fish]++
	}

	for i := 0; i < days; i++ {
		var newData [maxFishTimer + 1]int

		for index, numFish := range data {
			if index == 0 {
				// Reset the fish
				newData[recycleTimer] += numFish
				// Add some new fish
				newData[newFishTimer] += numFish
			} else {
				newData[index-1] += numFish
			}
		}

		data = newData
	}

	sum := 0
	for _, numFish := range data {
		sum += numFish
	}
	return sum
}

func part1(inputList []string) string {
	data := parseInput(inputList)
	num_fish := simulate(80, data)
	return strconv.Itoa(num_fish)
}

func part2(inputList []string) string {
	data := parseInput(inputList)
	num_fish := simulate(256, data)
	return strconv.Itoa(num_fish)
}

func main() {
	aoc.MainFunc(part1, part2)
}
