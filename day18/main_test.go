package main

import (
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

func TestDepth(t *testing.T) {
	var tests = []struct {
		input                 string
		pairStartBracketIndex int
		output                int
	}{
		{"[[[[[9,8],1],2],3],4]", 4, 4},
		{"[7,[6,[5,[4,[3,2]]]]]", 12, 4},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", 28, 4},
	}

	for _, test := range tests {
		depth := depth(test.input, test.pairStartBracketIndex)
		assert.Equal(t, test.output, depth)
	}
}

func TestExplode(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for _, test := range tests {
		output, exploded := explode(test.input)
		assert.Equal(t, test.output, output)
		assert.Equal(t, true, exploded)
	}
}

func TestReduce(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for _, test := range tests {
		output := reduce(test.input)
		assert.Equal(t, test.output, output)
	}
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"[[[[[9,10],1],2],3],4]", "[[[[[9,[5,5]],1],2],3],4]"},
		{"[[[[[11,10],1],2],3],4]", "[[[[[[5,6],10],1],2],3],4]"},
	}

	for _, test := range tests {
		output, splited := split(test.input)
		assert.Equal(t, test.output, output)
		assert.Equal(t, true, splited)
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		output string
	}{
		{"[1,2]", "[[3,4],5]", "[[1,2],[[3,4],5]]"},
		{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for _, test := range tests {
		output := add(test.input1, test.input2)
		assert.Equal(t, test.output, output)
	}
}

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		input     string
		magnutude int
	}{
		{"[9,1]", 29},
		{"[1,9]", 21},
		{"[[9,1],[1,9]]", 129},
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791.},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, test := range tests {
		m := magnitude(test.input)
		assert.Equal(t, test.magnutude, m)
	}
}

func TestPart1Example1(t *testing.T) {
	var nums = []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
	}

	expected := "[[[[1,1],[2,2]],[3,3]],[4,4]]"

	output := addNums(nums)

	assert.Equal(t, expected, output)
}

func TestPart1Example2(t *testing.T) {
	var nums = []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
	}

	expected := "[[[[3,0],[5,3]],[4,4]],[5,5]]"

	output := addNums(nums)

	assert.Equal(t, expected, output)
}

func TestPart1Example3(t *testing.T) {
	var nums = []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
		"[6,6]",
	}

	expected := "[[[[5,0],[7,4]],[5,5]],[6,6]]"

	output := addNums(nums)

	assert.Equal(t, expected, output)
}

func TestPart1Example4(t *testing.T) {
	var nums = []string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]",
	}

	expected := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"

	output := addNums(nums)

	assert.Equal(t, expected, output)
}

var exampleHomework = []string{
	"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
	"[[[5,[2,8]],4],[5,[[9,9],0]]]",
	"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
	"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
	"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
	"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
	"[[[[5,4],[7,7]],8],[[8,3],8]]",
	"[[9,3],[[9,9],[6,[4,9]]]]",
	"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
	"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
}

func TestPart1Example5(t *testing.T) {
	expectedSum := "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"
	expectedMagnitude := "4140"

	sum := addNums(exampleHomework)
	magnitude := part1(exampleHomework)

	assert.Equal(t, expectedSum, sum)
	assert.Equal(t, expectedMagnitude, magnitude)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "4235", result)
}

func TestPart2Example1(t *testing.T) {
	result := part2(exampleHomework)
	assert.Equal(t, "3993", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "4659", result)
}
