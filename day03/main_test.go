package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "198", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "2743844", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "230", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "6677951", result)
}
