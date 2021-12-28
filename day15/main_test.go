package main

import (
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
}

func TestParseInput(t *testing.T) {
	data := parseInput(exampleInput1)

	assert.Equal(t, 10, len(data))
	assert.Equal(t, 10, len(data[0]))
	assert.Equal(t, []int{1, 1, 6, 3, 7, 5, 1, 7, 4, 2}, data[0])
	assert.Equal(t, []int{2, 3, 1, 1, 9, 4, 4, 5, 8, 1}, data[9])
}

func TestCreateFullMap(t *testing.T) {
	data := parseInput(exampleInput1)
	data = makeFullMap(data)

	assert.Equal(t, []int{1, 1, 6, 3, 7, 5, 1, 7, 4, 2, 2, 2, 7, 4, 8, 6, 2, 8, 5, 3, 3, 3, 8, 5, 9, 7, 3, 9, 6, 4, 4, 4, 9, 6, 1, 8, 4, 1, 7, 5, 5, 5, 1, 7, 2, 9, 5, 2, 8, 6}, data[0])
	assert.Equal(t, []int{6, 7, 5, 5, 4, 8, 8, 9, 3, 5, 7, 8, 6, 6, 5, 9, 9, 1, 4, 6, 8, 9, 7, 7, 6, 1, 1, 2, 5, 7, 9, 1, 8, 8, 7, 2, 2, 3, 6, 8, 1, 2, 9, 9, 8, 3, 3, 4, 7, 9}, data[len(data)-1])
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "40", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "498", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput1)
	assert.Equal(t, "315", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "2901", result)
}
