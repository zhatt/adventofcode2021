package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput1 = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}
var exampleInput2 = []string{
	"dc-end",
	"HN-start",
	"start-kj",
	"dc-start",
	"dc-HN",
	"LN-dc",
	"HN-end",
	"kj-sa",
	"kj-HN",
	"kj-dc",
}

var exampleInput3 = []string{
	"fs-end",
	"he-DX",
	"fs-he",
	"start-DX",
	"pj-DX",
	"end-zg",
	"zg-sl",
	"zg-pj",
	"pj-he",
	"RW-he",
	"fs-DX",
	"pj-RW",
	"zg-RW",
	"start-pj",
	"he-WI",
	"zg-he",
	"pj-fs",
	"start-RW",
}

func TestParseInput(t *testing.T) {
	caveData := parseInput(exampleInput1)
	assert.Equal(t, 6, len(caveData.caves))

	cave := caveData.caves["start"]
	names := cave.getConnectedCaveNames()
	sort.Strings(names)
	assert.Equal(t, []string{"A", "b"}, names)

	cave = caveData.caves["A"]
	names = cave.getConnectedCaveNames()
	sort.Strings(names)
	assert.Equal(t, []string{"b", "c", "end", "start"}, names)
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "10", result)
}

func TestPart1Example2(t *testing.T) {
	result := part1(exampleInput2)
	assert.Equal(t, "19", result)
}

func TestPart1Example3(t *testing.T) {
	result := part1(exampleInput3)
	assert.Equal(t, "226", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "3887", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput1)
	assert.Equal(t, "36", result)
}

func TestPart2Example2(t *testing.T) {
	result := part2(exampleInput2)
	assert.Equal(t, "103", result)
}

func TestPart2Example3(t *testing.T) {
	result := part2(exampleInput3)
	assert.Equal(t, "3509", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "104834", result)
}
