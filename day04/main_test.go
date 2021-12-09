package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput = []string{
	"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
	"",
	"22 13 17 11  0",
	" 8  2 23  4 24",
	"21  9 14 16  7",
	" 6 10  3 18  5",
	" 1 12 20 15 19",
	"",
	" 3 15  0  2 22",
	" 9 18 13 17  5",
	"19  8  7 25 23",
	"20 11 10 24  4",
	"14 21 16 12  6",
	"",
	"14 21 17 24  4",
	"10 16 15  9 19",
	"18  8 23 26 20",
	"22 11 13  6  5",
	" 2  0 12  3  7",
}

func TestParseInput(t *testing.T) {
	draws, boards := parseInput(exampleInput)

	assert.Equal(t, 27, len(draws))
	assert.Equal(t, 7, draws[0])
	assert.Equal(t, 9, draws[2])
	assert.Equal(t, 1, draws[len(draws)-1])

	assert.Equal(t, 3, len(boards))
	assert.Equal(t, 4, boards[1][4][3].value)
	assert.Equal(t, 21, boards[0][0][2].value)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "4512", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "35670", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "1924", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "22704", result)
}
