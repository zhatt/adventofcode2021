package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"zhatt/aoc2021/aoc"
)

var exampleInput = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestIsCorruptedChunk(t *testing.T) {
	cases := []struct {
		chunk  string
		status status
		score  int
	}{
		{"()", good, 0},
		{"[]", good, 0},
		{"([])", good, 0},
		{"{()()()}", good, 0},
		{"<([{}])>", good, 0},
		{"[<>({}){}[([])<>]]", good, 0},
		{"(((((((((())))))))))", good, 0},

		{"(]", corrupted, syntaxErrorScore["]"]},
		{"{()()()>", corrupted, syntaxErrorScore[">"]},
		{"(((()))}", corrupted, syntaxErrorScore["}"]},
		{"<([]){()}[{}])", corrupted, syntaxErrorScore[")"]},

		{"[({(<(())[]>[[{[]{<()<>>", incomplete, 288957},
		{"[(()[<>])]({[<{<<[]>>(", incomplete, 5566},
		{"(((({<>}<{<{<>}{[]{[]{}", incomplete, 1480781},
		{"{<[[]]>}<{[{[{[]{()[[[]", incomplete, 995444},
		{"<{([{{}}[<[[[<>{}]]]>[]]", incomplete, 294},
	}

	for i, testCase := range cases {
		status, score := parseChunk(testCase.chunk)
		if testCase.status != status {
			fmt.Printf("Case %d %s %d\n", i, testCase.chunk, testCase.score)
		}
		assert.Equal(t, testCase.status, status)
		assert.Equal(t, testCase.score, score)
	}
}

func TestPart1Example(t *testing.T) {
	result := part1(exampleInput)
	assert.Equal(t, "26397", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "413733", result)
}

func TestPart2Example(t *testing.T) {
	result := part2(exampleInput)
	assert.Equal(t, "288957", result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "3354640192", result)
}
