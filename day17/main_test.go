package main

import (
	"testing"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []string{
	"target area: x=20..30, y=-10..-5",
}

func TestParseInput(t *testing.T) {
	bounds := parseInput(exampleInput1)

	assert.Equal(t, coord.Coord{Xval: 20, Yval: -10}, bounds.min)
	assert.Equal(t, coord.Coord{Xval: 30, Yval: -5}, bounds.max)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "45", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "3570", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput1)
	assert.Equal(t, "112", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "1919", result)
}
