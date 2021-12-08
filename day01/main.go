package main

import (
	"strconv"
	"zhatt/aoc2021/aoc"
)

func inputToIntList(inputList []string) []int {
	retVal := make([]int, 0, len(inputList))

	for _, strVal := range inputList {
		intVal, err := strconv.Atoi(strVal)
		aoc.PanicOnError(err)

		retVal = append(retVal, intVal)
	}
	return retVal
}

func getNumIncreasing(depths []int, size int) int {
	greaterCount := 0

	for index := range depths {
		if index < size {
			continue
		}

		currentWindow := 0
		lastWindow := 0

		for i := 0; i < size; i++ {
			currentWindow += depths[index-i]
			lastWindow += depths[index-i-1]
		}
		if currentWindow > lastWindow {
			greaterCount++
		}
	}
	return greaterCount
}

func part1(inputList []string) string {
	depthList := inputToIntList(inputList)
	count := getNumIncreasing(depthList, 1)
	return strconv.Itoa(count)
}

func part2(inputList []string) string {
	depthList := inputToIntList(inputList)
	count := getNumIncreasing(depthList, 3)
	return strconv.Itoa(count)
}

func main() {
	aoc.MainFunc(part1, part2)
}
