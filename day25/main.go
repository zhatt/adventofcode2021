package main

import (
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"
)

type cucumberMap struct {
	sizeX int
	sizeY int
	data  [][]string
}

func newCucumberMap(sizeX int, sizeY int) *cucumberMap {

	cucMap := &cucumberMap{
		sizeX: sizeX,
		sizeY: sizeY,
		data:  make([][]string, 0, sizeX),
	}

	for x := 0; x < sizeX; x++ {
		cucMap.data = append(cucMap.data, make([]string, sizeY))
	}

	return cucMap
}

func parseInput(inputLines []string) *cucumberMap {
	sizeX := len(inputLines[0])
	sizeY := len(inputLines)

	m := newCucumberMap(sizeX, sizeY)

	for y, line := range inputLines {
		for x := range line {
			m.data[x][y] = line[x : x+1]
		}
	}

	return m
}

func (m *cucumberMap) String() string {
	s := make([]string, 0, m.sizeX*m.sizeY)
	for y := 0; y < m.sizeY; y++ {
		if y != 0 {
			s = append(s, "\n")
		}
		for x := 0; x < m.sizeX; x++ {
			s = append(s, m.data[x][y])
		}
	}

	return strings.Join(s, "")
}

func (m *cucumberMap) moveEast() {
	newMap := newCucumberMap(m.sizeX, m.sizeY)

	for y := 0; y < m.sizeY; y++ {
		for x := 0; x < m.sizeX; x++ {
			var toX = x + 1
			if x == m.sizeX-1 {
				toX = 0
			}
			if m.data[x][y] == ">" && m.data[toX][y] == "." {
				newMap.data[x][y] = "."
				newMap.data[toX][y] = ">"
				x++ // Skip ahead since we've already filled.
			} else {
				newMap.data[x][y] = m.data[x][y]
			}
		}
	}

	m.data = newMap.data
}

func (m *cucumberMap) moveDown() {
	newMap := newCucumberMap(m.sizeX, m.sizeY)

	for x := 0; x < m.sizeX; x++ {
		for y := 0; y < m.sizeY; y++ {
			var toY = y + 1
			if y == m.sizeY-1 {
				toY = 0
			}
			if m.data[x][y] == "v" && m.data[x][toY] == "." {
				newMap.data[x][y] = "."
				newMap.data[x][toY] = "v"
				y++ // Skip ahead since we've already filled.
			} else {
				newMap.data[x][y] = m.data[x][y]
			}
		}
	}

	m.data = newMap.data
}

func (m *cucumberMap) move() {
	m.moveEast()
	m.moveDown()
}

func part1(inputList []string) string {
	cucumberMap := parseInput(inputList)

	steps := 0

	current := cucumberMap.String()

	for {
		steps++

		cucumberMap.move()

		next := cucumberMap.String()

		if current == next {
			break
		}

		current = next
	}

	return strconv.Itoa(steps)
}

func part2(inputList []string) string {
	// No part 2 on day 25

	return strconv.Itoa(0)
}

func main() {
	aoc.MainFunc(part1, part2)
}
