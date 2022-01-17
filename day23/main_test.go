package main

import (
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"#############",
	"#...........#",
	"###B#C#B#D###",
	"  #A#D#C#A#",
	"  #########",
}

func TestParseInput(t *testing.T) {
	burrow := parseInput(exampleInput, part1e)

	assert.Equal(t, bronzeAmphipod, burrow[roomA1])
	assert.Equal(t, amberAmphipod, burrow[roomA2])

	assert.Equal(t, copperAmphipod, burrow[roomB1])
	assert.Equal(t, desertAmphipod, burrow[roomB2])

	assert.Equal(t, bronzeAmphipod, burrow[roomC1])
	assert.Equal(t, copperAmphipod, burrow[roomC2])

	assert.Equal(t, desertAmphipod, burrow[roomD1])
	assert.Equal(t, amberAmphipod, burrow[roomD2])
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "12521", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "16157", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "44169", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "43481", result)
}
