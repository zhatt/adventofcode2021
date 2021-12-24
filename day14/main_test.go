package main

import (
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []string{
	"NNCB",
	"",
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C",
}

func TestParseInput(t *testing.T) {
	polymerTemplate, pairInsertions := parseInput(exampleInput1)
	assert.Equal(t, "NNCB", polymerTemplate)
	assert.Equal(t, 16, len(pairInsertions))
	assert.Equal(t, pairInsertion{pair: "CH", insertion: "B"}, pairInsertions[0])
	assert.Equal(t, pairInsertion{pair: "CN", insertion: "C"}, pairInsertions[15])
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "1588", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "2112", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput1)
	assert.Equal(t, "2188189693529", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "3243771149914", result)
}
