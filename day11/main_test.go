package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput1 = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

var exampleInput2 = []string{
	"11111",
	"19991",
	"19191",
	"19991",
	"11111",
}

func TestParseInput(t *testing.T) {
	grid := parseInput(exampleInput1)
	assert.Equal(t, 10, len(grid.octos))
	assert.Equal(t, 10, len(grid.octos[0]))
	assert.Equal(t, 4, grid.octos[2][1].energy)
}

func TestPart1Example1(t *testing.T) {
	iteration1 := []string{
		"34543",
		"40004",
		"50005",
		"40004",
		"34543",
	}
	grid := parseInput(exampleInput2)
	gridIter1 := parseInput(iteration1)
	simulate(grid)
	assert.Equal(t, gridIter1.String(), grid.String())
	simulate(grid)
}

func TestPart1Example2(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "1656", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "1637", result)
}

func TestPart2Example(t *testing.T) {
	result := part2(exampleInput1)
	assert.Equal(t, "195", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "242", result)
}
