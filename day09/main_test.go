package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

var exampleInput = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

func TestParseInput(t *testing.T) {
	hm := parseInput(exampleInput)

	assert.Equal(t, 3, hm.getValue(coord.Coord{X: 0, Y: 1}))
	assert.Equal(t, 7, hm.getValue(coord.Coord{X: 3, Y: 3}))
	assert.Equal(t, 8, hm.getValue(coord.Coord{X: 4, Y: 3}))
}

func TestPart1Example(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "15", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "570", result)
}

func TestPart2Example(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "1134", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "899392", result)
}
