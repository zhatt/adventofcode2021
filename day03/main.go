package main

import (
	"strconv"
	"zhatt/aoc2021/aoc"
)

type criteria int

const (
	mostCommon criteria = iota
	leastCommon
)

func parseInput(inputList []string) (int, []uint64) {
	numBits := len(inputList[0])

	retVal := make([]uint64, 0, len(inputList))

	for _, strVal := range inputList {
		intVal, err := strconv.ParseUint(strVal, 2, 64)
		aoc.PanicOnError(err)

		retVal = append(retVal, intVal)
	}
	return numBits, retVal
}

func calculatePower(numBits int, values []uint64) uint64 {
	gamma := uint64(0)
	epsilon := uint64(0)

	// Loop with 1 bit mask shifting left to right
	for mask := uint64(1) << (numBits - 1); mask != 0; mask >>= 1 {
		onesCount := 0
		for _, value := range values {
			if value&mask != 0 {
				onesCount++
			}
		}

		zerosCount := len(values) - onesCount

		if onesCount > zerosCount {
			gamma |= mask
		} else {
			epsilon |= mask
		}
	}
	return gamma * epsilon
}

func calculateRating(criteria criteria, numBits int, valuesIn []uint64) uint64 {

	rating := uint64(0)

	values := make([]uint64, len(valuesIn))
	copy(values, valuesIn)

	// Loop with 1 bit mask shifting left to right
	for mask := uint64(1) << (numBits - 1); mask != 0; mask >>= 1 {
		onesCount := 0

		for _, value := range values {
			if value&mask != 0 {
				onesCount += 1
			}
		}

		zerosCount := len(values) - onesCount

		ratingValue := uint64(0)
		if criteria == mostCommon {
			if onesCount >= zerosCount {
				ratingValue = mask
			}
		} else {
			if onesCount < zerosCount {
				ratingValue = mask
			}
		}

		newValues := make([]uint64, 0)
		for _, value := range values {
			if value&mask == uint64(ratingValue) {
				newValues = append(newValues, value)
			}
		}

		values = newValues
		if len(values) == 1 {
			rating = values[0]
			break
		}
	}
	return rating
}

func calculateLifeSupport(numBits int, valuesIn []uint64) uint64 {
	oxygenRating := calculateRating(mostCommon, numBits, valuesIn)
	co2ScrubberRating := calculateRating(leastCommon, numBits, valuesIn)
	return oxygenRating * co2ScrubberRating
}

func part1(inputList []string) string {
	numBits, values := parseInput(inputList)
	power := calculatePower(numBits, values)
	return strconv.FormatUint(power, 10)
}

func part2(inputList []string) string {
	numBits, values := parseInput(inputList)
	power := calculateLifeSupport(numBits, values)
	return strconv.FormatUint(power, 10)
}

func main() {
	aoc.MainFunc(part1, part2)
}
