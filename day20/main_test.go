package main

import (
	"testing"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##" +
		"#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###" +
		".######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#." +
		".#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#....." +
		".#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.." +
		"...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#....." +
		"..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
	"",
	"#..#.",
	"#....",
	"##..#",
	"..#..",
	"..###",
}

func TestParseInput(t *testing.T) {
	algorithm, image := parseInput(exampleInput)

	assert.Equal(t, 512, len(algorithm))
	assert.Equal(t, coord.Coord{X: 0, Y: 0}, image.minBound)
	assert.Equal(t, coord.Coord{X: 4, Y: 4}, image.maxBound)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "35", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "5682", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "3351", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "17628", result)
}
