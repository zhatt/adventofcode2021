package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput = []string{
	"16,1,2,0,4,2,7,1,2,14",
}

func TestParseInput(t *testing.T) {
	data := parseInput(exampleInput)

	assert.Equal(t, 10, len(data))
	assert.Equal(t, 1, data[1])
}

func TestAbsInt(t *testing.T) {
	assert.Equal(t, 10, absInt(10))
	assert.Equal(t, 10, absInt(-10))
	assert.Equal(t, math.MaxInt, absInt(math.MaxInt))
	assert.Equal(t, math.MaxInt, absInt(math.MinInt+1))
	// Maximum negative int has no corresponding positive int
	assert.Equal(t, math.MinInt, absInt(math.MinInt))
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "37", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "344535", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "168", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "95581659", result)
}
