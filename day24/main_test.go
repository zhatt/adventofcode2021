package main

import (
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var expectedDiv1 = []int{1, 1, 1, 26, 1, 1, 26, 26, 1, 26, 1, 26, 26, 26}
	var expectedAdd1 = []int{11, 14, 13, -4, 11, 10, -4, -12, 10, -11, 12, -1, 0, -11}
	var expectedAdd2 = []int{3, 7, 1, 6, 14, 7, 9, 9, 6, 4, 0, 7, 12, 1}

	coefficients := parseInput(aoc.ReadInput("input.txt"))

	assert.Equal(t, expectedDiv1, coefficients.div1)
	assert.Equal(t, expectedAdd1, coefficients.add1)
	assert.Equal(t, expectedAdd2, coefficients.add2)
}

func TestMonad(t *testing.T) {
	coefficients := parseInput(aoc.ReadInput("input.txt"))

	value, err := monadValue(99_999_999_999_999, coefficients)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3904362732, value)
	assert.False(t, monad(99_999_999_999_999, coefficients))

	value, err = monadValue(12_345_678_912_345, coefficients)
	assert.Equal(t, nil, err)
	assert.Equal(t, 51827339, value)
	assert.False(t, monad(12_345_678_912_345, coefficients))

	// My largest value
	value, err = monadValue(92_967_699_949_891, coefficients)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, value)
	assert.True(t, monad(92_967_699_949_891, coefficients))

	// My smallest value
	value, err = monadValue(91_411_143_612_181, coefficients)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, value)
	assert.True(t, monad(91_411_143_612_181, coefficients))
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "92967699949891", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "91411143612181", result)
}
