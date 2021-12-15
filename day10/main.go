package main

import (
	"sort"
	"strconv"
	"zhatt/aoc2021/aoc"
)

var bracketPairs = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var syntaxErrorScore = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var autocompleteScore = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

type status int

const (
	corrupted status = iota
	incomplete
	good
)

func parseChunk(chunk string) (status, int) {
	tmp := ""
	score := 0
	status := good

loop:
	for _, c := range chunk {
		c := string(c)
		if openingBracket, ok := bracketPairs[c]; ok {
			if len(tmp) == 0 {
				score = syntaxErrorScore[c]
				status = corrupted
				break loop
			}

			if tmp[len(tmp)-1:] == openingBracket {
				// Remove matching bracket from tmp chunk string
				tmp = tmp[:len(tmp)-1]
			} else {
				score = syntaxErrorScore[c]
				status = corrupted
				break loop
			}
		} else {
			tmp += c
		}
	}

	if status != corrupted && len(tmp) != 0 {
		status = incomplete

		for index := len(tmp) - 1; index >= 0; index-- {
			c := string(tmp[index])
			score = score*5 + autocompleteScore[c]
		}
	}

	return status, score
}

func part1(inputList []string) string {
	score := 0
	for _, chunk := range inputList {
		status, chunkScore := parseChunk(chunk)
		if status == corrupted {
			score += chunkScore
		}
	}
	return strconv.Itoa(score)
}

func part2(inputList []string) string {
	scores := make([]int, 0)

	for _, chunk := range inputList {
		status, chunkScore := parseChunk(chunk)
		if status == incomplete {
			scores = append(scores, chunkScore)
		}
	}

	sort.Ints(scores)
	index := len(scores) / 2
	return strconv.Itoa(scores[index])
}

func main() {
	aoc.MainFunc(part1, part2)
}
