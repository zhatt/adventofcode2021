package main

import (
	"strings"
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []string{
	"...>>>>>...",
}

var exampleInput2 = []string{
	"..........",
	".>v....v..",
	".......>..",
	"..........",
}

var exampleInput3 = []string{
	"...>...",
	".......",
	"......>",
	"v.....>",
	"......>",
	".......",
	"..vvv..",
}

var exampleInput4 = []string{
	"v...>>.vv>",
	".vv>>.vv..",
	">>.>v>...v",
	">>v>>.>.v.",
	"v>v.vv.v..",
	">.>>..v...",
	".vv..>.>v.",
	"v.v..>>v.v",
	"....v..v.>",
}

func TestParseInput(t *testing.T) {
	cucumberMap1 := parseInput(exampleInput1)
	assert.Equal(t, strings.Join(exampleInput1, "\n"), cucumberMap1.String())

	cucumberMap2 := parseInput(exampleInput2)
	assert.Equal(t, strings.Join(exampleInput2, "\n"), cucumberMap2.String())
}

func TestPart1Example1(t *testing.T) {
	cucumberMap := parseInput(exampleInput1)

	cucumberMap.move()
	assert.Equal(t, "...>>>>.>..", cucumberMap.String())

	cucumberMap.move()
	assert.Equal(t, "...>>>.>.>.", cucumberMap.String())
}

func TestPart2Example2(t *testing.T) {
	cucumberMap := parseInput(exampleInput2)

	expectedStep1 := "..........\n.>........\n..v....v>.\n.........."
	cucumberMap.move()
	assert.Equal(t, expectedStep1, cucumberMap.String())
}

func TestPart2Example3(t *testing.T) {
	cucumberMap := parseInput(exampleInput3)

	expectedStep4 := ">......\n..v....\n..>.v..\n.>.v...\n...>...\n.......\nv......"

	cucumberMap.move()
	cucumberMap.move()
	cucumberMap.move()
	cucumberMap.move()

	assert.Equal(t, expectedStep4, cucumberMap.String())
}

func TestPart2Example4(t *testing.T) {
	result := part1(exampleInput4)

	assert.Equal(t, "58", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "482", result)
}
