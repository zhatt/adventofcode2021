package main

import (
	"fmt"
	"strconv"
	"zhatt/aoc2021/aoc"
)

type recursiveMode int

const (
	findLargest recursiveMode = iota
	findSmallest
)

type coefficients struct {
	div1 []int
	add1 []int
	add2 []int
}

func parseInput(inputLines []string) *coefficients {
	coefficients := &coefficients{}

	for lineNumber, line := range inputLines {
		var div1 int
		var add1 int
		var add2 int

		if lineNumber%18 == 4 {
			_, err := fmt.Sscanf(line, "div z %d", &div1)
			aoc.PanicOnError(err)
			coefficients.div1 = append(coefficients.div1, div1)
		}

		if lineNumber%18 == 5 {
			_, err := fmt.Sscanf(line, "add x %d", &add1)
			aoc.PanicOnError(err)
			coefficients.add1 = append(coefficients.add1, add1)
		}

		if lineNumber%18 == 15 {
			_, err := fmt.Sscanf(line, "add y %d", &add2)
			aoc.PanicOnError(err)
			coefficients.add2 = append(coefficients.add2, add2)
		}
	}

	return coefficients
}

// digitIteration is a reimplementation of the assembly language algorithm for
// each digit.
func digitIterationX(digit int, z int, div1 int, add1 int, add2 int) int {

	if (z%26)+add1 == digit {
		z /= div1
	} else {
		z /= div1
		z = z*26 + digit + add2
	}

	return z
}

type pruningSet struct {
	index int
	z     int
}

func recursiveFindModelNumber(mode recursiveMode, index int, z int, number int, coefficients *coefficients, pruned map[pruningSet]bool) (int, error) {
	if pruned == nil {
		pruned = make(map[pruningSet]bool)
	}

	// If we have already been called with a givin index and z then we can
	// stop checking this subtree.
	pruningSet := pruningSet{index: index, z: z}
	if pruned[pruningSet] {
		return -1, fmt.Errorf("not found")
	}

	if index == len(coefficients.add1) {
		if z == 0 {
			return number, nil
		} else {
			return -1, fmt.Errorf("not found 1")
		}
	}

	var start int
	var end int // One past end
	var increment int
	if mode == findLargest {
		start = 9
		end = 0
		increment = -1
	} else {
		start = 1
		end = 10
		increment = 1
	}

	for digit := start; digit != end; digit += increment {
		newZ := digitIterationX(digit, z, coefficients.div1[index], coefficients.add1[index], coefficients.add2[index])

		newNumber := number*10 + digit
		result, err := recursiveFindModelNumber(mode, index+1, newZ, newNumber, coefficients, pruned)

		if err == nil {
			return result, nil
		}
	}

	// This subtree is not valid to prune it in case another subtree has the
	// same z accumulated at this depth (index)
	pruned[pruningSet] = true

	return -1, fmt.Errorf("not found")
}

func monadValue(modelNumber int, coefficients *coefficients) (int, error) {
	var digits [14]int

	if modelNumber > 99_999_999_999_999 || modelNumber < 11_111_111_111_111 {
		return -1, fmt.Errorf("bad model number")
	}

	for i := 13; i >= 0; i-- {
		digit := modelNumber % 10

		if digit == 0 {
			// Returning -1 so that it doesn't accidentally get used as 0
			return -1, fmt.Errorf("model number contains zero")
		}

		digits[i] = digit
		modelNumber /= 10
	}

	z := 0

	for i := 0; i <= 13; i++ {
		z = digitIterationX(digits[i], z, coefficients.div1[i], coefficients.add1[i], coefficients.add2[i])
	}

	return z, nil
}

// This was used when testing with a brute force algorithm to learn about
// Monads.
func monad(modelNumber int, coefficients *coefficients) bool {
	value, err := monadValue(modelNumber, coefficients)
	return err == nil && 0 == value
}

func part1(inputList []string) string {
	coefficients := parseInput(inputList)

	serialNumber, err := recursiveFindModelNumber(findLargest, 0, 0, 0, coefficients, nil)
	aoc.PanicOnError(err)

	return strconv.Itoa(serialNumber)
}

func part2(inputList []string) string {
	coefficients := parseInput(inputList)

	serialNumber, err := recursiveFindModelNumber(findSmallest, 0, 0, 0, coefficients, nil)
	aoc.PanicOnError(err)

	return strconv.Itoa(serialNumber)
}

func main() {
	aoc.MainFunc(part1, part2)
}
