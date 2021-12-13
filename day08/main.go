package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"

	"github.com/gitchander/permutation"
)

type data struct {
	signal_values [10]string
	output_values [4]string
}

func parseInput(inputList []string) []data {
	all_data := make([]data, 0)

	for _, line := range inputList {
		var line_data data
		vals := strings.Fields(line)
		copy(line_data.signal_values[0:10], vals[0:0+10])
		copy(line_data.output_values[0:4], vals[11:11+4])
		all_data = append(all_data, line_data)
	}

	return all_data
}

func count1478(data []data) int {
	count := 0
	for _, one_displays_data := range data {
		for _, output_value := range one_displays_data.output_values {
			num_on := len(output_value)
			switch num_on {
			case
				/* display 1 */ 2,
				/* display 4 */ 4,
				/* display 7 */ 3,
				/* display 8 */ 7:
				count++
			}
		}
	}

	return count
}

func mapSignalsToSegments(signals string, order []string) string {
	segments := make([]string, 0, len(signals))

	for _, signal := range signals {
		switch signal {
		case 'a':
			segments = append(segments, order[0])
		case 'b':
			segments = append(segments, order[1])
		case 'c':
			segments = append(segments, order[2])
		case 'd':
			segments = append(segments, order[3])
		case 'e':
			segments = append(segments, order[4])
		case 'f':
			segments = append(segments, order[5])
		case 'g':
			segments = append(segments, order[6])
		default:
			aoc.PanicOnError(fmt.Errorf("Error:  Unknown signal"))
		}
	}

	sort.Strings(segments)

	return strings.Join(segments, "")
}

var segmentCombinations = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func isLegalSegmentCombination(segments string) bool {
	// segments must be sorted
	_, valid := segmentCombinations[segments]
	return valid
}

func getDigit(segments string) int {
	// must be legal combination
	value, valid := segmentCombinations[segments]
	if !valid {
		aoc.PanicOnError(fmt.Errorf("bad segments"))
	}
	return value
}

func simulate(data []data) int {
	sum := 0

	for _, data_set := range data {
		wires := []string{"a", "b", "c", "d", "e", "f", "g"}
		permutations := permutation.New(permutation.StringSlice(wires))

	PERMUTATION:
		for permutations.Next() {
			for _, input := range data_set.signal_values {

				segments := mapSignalsToSegments(input, wires)
				if !isLegalSegmentCombination(segments) {
					continue PERMUTATION
				}
			}
			break
		}

		displayVal := 0
		for _, signals := range data_set.output_values {
			segments := mapSignalsToSegments(signals, wires)
			digit := getDigit(segments)

			displayVal = displayVal*10 + digit
		}
		sum += displayVal
	}

	return sum
}

func part1(inputList []string) string {
	data := parseInput(inputList)
	count := count1478(data)
	return strconv.Itoa(count)
}

func part2(inputList []string) string {
	data := parseInput(inputList)
	total := simulate(data)
	return strconv.Itoa(total)
}

func main() {
	aoc.MainFunc(part1, part2)
}
