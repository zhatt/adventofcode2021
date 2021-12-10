package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

var exampleInput = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}

func TestParseInput(t *testing.T) {
	lines := parseInput(exampleInput)

	assert.Equal(t, 10, len(lines))
	assert.Equal(t, coord.Coord{Xval: 8, Yval: 0}, lines[1].from)
	assert.Equal(t, coord.Coord{Xval: 1, Yval: 4}, lines[7].to)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "5", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "5585", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "12", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "17193", result)
}
