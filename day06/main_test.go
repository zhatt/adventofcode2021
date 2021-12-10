package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput = []string{
	"3,4,3,1,2",
}

func TestParseInput(t *testing.T) {
	data := parseInput(exampleInput)

	assert.Equal(t, 5, len(data))
	assert.Equal(t, 4, data[1])
}

func TestPart1Example1(t *testing.T) {
	data := parseInput(exampleInput)
	num_fish := simulate(18, data)
	assert.Equal(t, 26, num_fish)
}

func TestPart1Example2(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "5934", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "352872", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "26984457539", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "1604361182149", result)
}
