package main

import (
	"fmt"
	"regexp"
	"strconv"
	"zhatt/aoc2021/aoc"
)

func parseInput(inputLines []string) []string {
	return inputLines
}

func depth(input string, start int) int {
	depth := 0

	for index := 0; index < start; index++ {
		if input[index] == '[' {
			depth++
		}
		if input[index] == ']' {
			depth--
		}
	}

	return depth
}

var pairRe = regexp.MustCompile(`\[(\d+),(\d+)\]`)
var numberRe = regexp.MustCompile(`\d+`)

func explode(input string) (string, bool) {
	output := input

	matchPairs := pairRe.FindAllStringSubmatchIndex(input, -1)

	for _, match := range matchPairs {
		bracketOpen := match[0]
		bracketClose := match[1]
		leftNumStart := match[2]
		leftNumEnd := match[3]
		rightNumStart := match[4]
		rightNumEnd := match[5]
		depth := depth(input, bracketOpen)
		if depth >= 4 {
			leftBegin := 0
			leftEnd := 0
			rightBegin := 0
			rightEnd := 0

			matchNumbers := numberRe.FindAllStringIndex(input, -1)
			for index, numMatch := range matchNumbers {
				if numMatch[0] == leftNumStart {
					if index > 0 {
						leftBegin = matchNumbers[index-1][0]
						leftEnd = matchNumbers[index-1][1]
						break
					}
				}
			}
			for index, numMatch := range matchNumbers {
				if numMatch[0] == rightNumStart {
					if len(matchNumbers) > index+1 {
						rightBegin = matchNumbers[index+1][0]
						rightEnd = matchNumbers[index+1][1]
						break
					}
				}
			}

			if leftBegin != leftEnd {
				leftNumber, err := strconv.Atoi(input[leftBegin:leftEnd])
				aoc.PanicOnError(err)
				pairFirstNumber, err := strconv.Atoi(input[leftNumStart:leftNumEnd])
				aoc.PanicOnError(err)

				output = input[0:leftBegin] + strconv.Itoa(leftNumber+pairFirstNumber) + input[leftEnd:bracketOpen]
			} else {
				output = input[0:bracketOpen]
			}

			output += "0"

			if rightBegin != rightEnd {
				rightNumber, err := strconv.Atoi(input[rightBegin:rightEnd])
				aoc.PanicOnError(err)
				pairSecondNumber, err := strconv.Atoi(input[rightNumStart:rightNumEnd])
				aoc.PanicOnError(err)

				output += input[bracketClose:rightBegin] + strconv.Itoa(rightNumber+pairSecondNumber) + input[rightEnd:]
			} else {
				output += input[bracketClose:]
			}
			break
		}
	}

	return output, output != input
}

func split(input string) (string, bool) {
	output := input

	matchNumbers := numberRe.FindAllStringIndex(input, -1)
	for _, numMatch := range matchNumbers {
		number, err := strconv.Atoi(input[numMatch[0]:numMatch[1]])
		aoc.PanicOnError(err)

		if number >= 10 {
			left := number / 2
			right := (number + 1) / 2

			output = input[0:numMatch[0]] + fmt.Sprintf("[%d,%d]", left, right) + input[numMatch[1]:]
			break
		}
	}
	return output, output != input
}

func reduce(input string) string {
	output := input
	for {
		var reduced bool

		output, reduced = explode(output)
		if reduced {
			continue
		}

		output, reduced = split(output)
		if reduced {
			continue
		}

		break
	}
	return output
}

func add(input1, input2 string) string {
	output := fmt.Sprintf("[%s,%s]", input1, input2)
	output = reduce(output)
	return output
}

func magnitude(input string) int {
	output := input
	magnitude := 0

	for {
		match := pairRe.FindStringSubmatchIndex(output)
		if match == nil {
			break
		}

		bracketOpen := match[0]
		bracketClose := match[1]
		leftNumStart := match[2]
		leftNumEnd := match[3]
		rightNumStart := match[4]
		rightNumEnd := match[5]

		leftNum, err := strconv.Atoi(output[leftNumStart:leftNumEnd])
		aoc.PanicOnError(err)

		rightNum, err := strconv.Atoi(output[rightNumStart:rightNumEnd])
		aoc.PanicOnError(err)

		magnitude = 3*leftNum + 2*rightNum

		output = output[0:bracketOpen] + strconv.Itoa(magnitude) + output[bracketClose:]
	}

	return magnitude
}

func addNums(numbers []string) string {
	output := numbers[0]
	for _, num := range numbers[1:] {
		output = add(output, num)
	}
	return output
}

func part1(inputList []string) string {
	numbers := parseInput(inputList)
	output := addNums(numbers)
	return strconv.Itoa(magnitude(output))
}

func part2(inputList []string) string {
	numbers := parseInput(inputList)

	largestMagnitude := 0

	for i1, number1 := range numbers {
		for i2, number2 := range numbers {
			if i1 == i2 {
				continue
			}

			sum := add(number1, number2)
			magnitude := magnitude(sum)

			if magnitude > largestMagnitude {
				largestMagnitude = magnitude
			}
		}
	}

	return strconv.Itoa(largestMagnitude)
}

func main() {
	aoc.MainFunc(part1, part2)
}
