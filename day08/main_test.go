package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput1 = []string{
	"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
}

var exampleInput2 = []string{
	"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
	"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
	"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
	"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
	"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
	"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
	"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
	"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
	"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
	"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
}

func TestParseInput(t *testing.T) {
	data := parseInput(exampleInput2)

	assert.Len(t, data, 10)
	assert.Equal(t, "fgaebd", data[2].signal_values[0])
	assert.Equal(t, "cdgabef", data[2].signal_values[9])
	assert.Equal(t, "cg", data[2].output_values[0])
	assert.Equal(t, "cbg", data[2].output_values[3])
}

func TestPart1Example1(t *testing.T) {
	result := part1(exampleInput1)
	assert.Equal(t, "0", result)
}

func TestPart1Example2(t *testing.T) {
	result := part1(exampleInput2)
	assert.Equal(t, "26", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "412", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleInput1)
	assert.Equal(t, "5353", result)
}

func TestPart2Example2(t *testing.T) {
	result := part2(exampleInput2)
	assert.Equal(t, "61229", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "978171", result)
}
