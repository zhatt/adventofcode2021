package main

import (
	"strings"
	"testing"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []string{
	"6,10",
	"0,14",
	"9,10",
	"0,3",
	"10,4",
	"4,11",
	"6,0",
	"6,12",
	"4,1",
	"0,13",
	"10,12",
	"3,4",
	"3,0",
	"8,4",
	"1,10",
	"2,14",
	"8,10",
	"9,0",
	"",
	"fold along y=7",
	"fold along x=5",
}

func TestParseInput(t *testing.T) {
	paper, instructions := parseInput(exampleInput1)
	assert.Equal(t, 18, len(paper.points))
	assert.Contains(t, paper.points, coord.Coord{X: 6, Y: 10})
	assert.Contains(t, paper.points, coord.Coord{X: 9, Y: 0})
	assert.Contains(t, paper.points, coord.Coord{X: 0, Y: 13})

	assert.Equal(t, []instruction{
		newInstruction(foldY, 7),
		newInstruction(foldX, 5)}, instructions)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "17", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "743", result)
}

func TestPart2Example1(t *testing.T) {
	var expected = []string{
		"#####\n",
		"#   #\n",
		"#   #\n",
		"#   #\n",
		"#####\n",
	}

	result := part2(exampleInput1)
	assert.Equal(t, strings.Join(expected, ""), result)
}

func TestPart2Input(t *testing.T) {
	var expected = []string{
		"###   ##  ###  #     ##  #  # #  # #   \n",
		"#  # #  # #  # #    #  # # #  #  # #   \n",
		"#  # #    #  # #    #  # ##   #### #   \n",
		"###  #    ###  #    #### # #  #  # #   \n",
		"# #  #  # #    #    #  # # #  #  # #   \n",
		"#  #  ##  #    #### #  # #  # #  # ####\n",
	}
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, strings.Join(expected, ""), result)
}
