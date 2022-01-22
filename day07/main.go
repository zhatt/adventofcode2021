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

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type fuel_algorithm int

const (
	linear fuel_algorithm = iota
	growing
)

func calculate_fuel(distance int, algorithm fuel_algorithm) int {
	if algorithm == linear {
		return distance
	}
	sum := 0
	for i := 0; i <= distance; i++ {
		sum += i
	}
	return sum
}

func simulate(crab_locations []int, algorithm fuel_algorithm) int {
	min_align_location := crab_locations[0]
	max_align_location := crab_locations[0]

	for _, crab_location := range crab_locations {
		if crab_location < min_align_location {
			min_align_location = crab_location
		} else if crab_location > max_align_location {
			max_align_location = crab_location
		}
	}

	min_fuel := aoc.MaxInt

	for align_location := min_align_location; align_location <= max_align_location; align_location++ {
		fuel := 0
		for _, crab_location := range crab_locations {
			fuel += calculate_fuel(absInt(crab_location-align_location), algorithm)
		}
		if fuel < min_fuel {
			min_fuel = fuel
		}
	}

	return min_fuel
}

func part1(inputList []string) string {
	crab_locations := parseInput(inputList)
	fuel := simulate(crab_locations, linear)
	return strconv.Itoa(fuel)
}

func part2(inputList []string) string {
	crab_locations := parseInput(inputList)
	fuel := simulate(crab_locations, growing)
	return strconv.Itoa(fuel)
}

func main() {
	aoc.MainFunc(part1, part2)
}
