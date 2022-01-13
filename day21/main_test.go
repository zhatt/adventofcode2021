package main

import (
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"Player 1 starting position: 4",
	"Player 2 starting position: 8",
}

func TestDeterministicDie(t *testing.T) {
	die := newDeterministicDie(3)

	assert.Equal(t, 1, die.roll())
	assert.Equal(t, 1, die.numberOfRolls)
	assert.Equal(t, 2, die.roll())
	assert.Equal(t, 2, die.numberOfRolls)
	assert.Equal(t, 3, die.roll())
	assert.Equal(t, 3, die.numberOfRolls)
	assert.Equal(t, 1, die.roll())
	assert.Equal(t, 4, die.numberOfRolls)
	assert.Equal(t, 2, die.roll())
	assert.Equal(t, 5, die.numberOfRolls)
}

func TestParseInput(t *testing.T) {
	player1Position, player2Position := parseInput(exampleInput)

	assert.Equal(t, 4, player1Position)
	assert.Equal(t, 8, player2Position)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "739785", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "925605", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "444356092776315", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "486638407378784", result)
}
