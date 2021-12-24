package main

import (
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"
)

type pairInsertion struct {
	pair      string
	insertion string
}

func parseInput(inputLines []string) (string, []pairInsertion) {
	polymerTemplate := inputLines[0]
	pairInsertions := make([]pairInsertion, 0, len(inputLines)-2)

	for _, line := range inputLines[2:] {
		tokens := strings.Split(line, " -> ")
		pairInsertion := pairInsertion{pair: tokens[0], insertion: tokens[1]}
		pairInsertions = append(pairInsertions, pairInsertion)
	}

	return polymerTemplate, pairInsertions
}

func polymerize(polymer string, pairInsertions []pairInsertion, iterations int) int {

	// Create a map of the insertion pairs for fast lookup.
	insertionPairs := make(map[string]pairInsertion)
	for _, pair := range pairInsertions {
		insertionPairs[pair.pair] = pair
	}

	// Add _ to start and end of polymer so that we don't have to special
	// case the element at the start and end.
	polymer = "_" + polymer + "_"

	// Create a map or counts of all element pairs in polymer.  We will then
	// be able to simulate the polymerization by applying the insertion
	// instructions to the pairs.
	polymerPairCounts := make(map[string]int)
	notSplitPolymerPairs := make(map[string]int)
	for index := 0; index < len(polymer)-1; index++ {
		pair := polymer[index : index+2]
		polymerPairCounts[pair]++
		notSplitPolymerPairs[pair]++
	}

	// Simlate the polymerization by splitting pairs.
	for iteration := 0; iteration < iterations; iteration++ {
		newPolymerPairs := make(map[string]int)
		for polymerPair, polymerPairCount := range polymerPairCounts {
			if insertionPair, exists := insertionPairs[polymerPair]; exists {
				pair1 := polymerPair[0:1] + insertionPair.insertion
				pair2 := insertionPair.insertion + polymerPair[1:2]
				newPolymerPairs[pair1] += polymerPairCount
				newPolymerPairs[pair2] += polymerPairCount
				delete(notSplitPolymerPairs, polymerPair)
			} else {
				newPolymerPairs[polymerPair] = polymerPairCount
			}
		}
		polymerPairCounts = newPolymerPairs
	}

	// Count elements but ignore _ added at start and end of polymer
	// template.
	counts := make(map[rune]int)
	for pair, count := range polymerPairCounts {
		for _, element := range pair {
			if element != '_' {
				counts[element] += count
			}
		}
	}

	leastCommonCount := 0
	mostCommonCount := 0

	for _, count := range counts {
		leastCommonCount = count
		mostCommonCount = count
		break
	}

	for _, count := range counts {
		if count < leastCommonCount {
			leastCommonCount = count
		}
		if count > mostCommonCount {
			mostCommonCount = count
		}
	}

	if len(notSplitPolymerPairs) != 2 {
		// If there was a pair that never got split then we would need
		// to make an adjustment to the /2 below.
		panic("need to handle not split")
	} else {
		// The counts double counts the elements.  Adjust.
		mostCommonCount /= 2
		leastCommonCount /= 2
	}

	return mostCommonCount - leastCommonCount
}

func part1(inputList []string) string {
	polymer, pairInsertions := parseInput(inputList)
	count := polymerize(polymer, pairInsertions, 10)
	return strconv.Itoa(count)
}

func part2(inputList []string) string {
	polymer, pairInsertions := parseInput(inputList)
	count := polymerize(polymer, pairInsertions, 40)
	return strconv.Itoa(count)
}

func main() {
	aoc.MainFunc(part1, part2)
}
