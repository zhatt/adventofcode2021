package main

import (
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

type line struct {
	from coord.Coord
	to   coord.Coord
}

type ventMap struct {
	data map[coord.Coord]int
}

func newVentmap() ventMap {
	ventMap := ventMap{}
	ventMap.data = make(map[coord.Coord]int)

	return ventMap
}

func getIncrement(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

func (v *ventMap) addVent(line line) {

	xIncrement := getIncrement(line.from.X, line.to.X)
	yIncrement := getIncrement(line.from.Y, line.to.Y)

	x := line.from.X
	y := line.from.Y

	v.data[coord.Coord{X: x, Y: y}]++
	for {
		if x == line.to.X && y == line.to.Y {
			break
		}
		x += xIncrement
		y += yIncrement
		v.data[coord.Coord{X: x, Y: y}]++
	}
}

func (v *ventMap) numDangerous() int {
	numDangerous := 0

	for _, val := range v.data {
		if val > 1 {
			numDangerous++
		}
	}

	return numDangerous
}

func parseInput(inputList []string) []line {

	lines := make([]line, 0, len(inputList))

	for _, strVal := range inputList {

		tokens := strings.Fields(strVal)

		vals := strings.Split(tokens[0], ",")
		x, err := strconv.Atoi(vals[0])
		aoc.PanicOnError(err)
		y, err := strconv.Atoi(vals[1])
		aoc.PanicOnError(err)
		from := coord.Coord{X: x, Y: y}

		vals = strings.Split(tokens[2], ",")
		x, err = strconv.Atoi(vals[0])
		aoc.PanicOnError(err)
		y, err = strconv.Atoi(vals[1])
		aoc.PanicOnError(err)
		to := coord.Coord{X: x, Y: y}

		line := line{from: from, to: to}
		lines = append(lines, line)
	}
	return lines
}

func isManhattan(line line) bool {
	return line.from.X == line.to.X || line.from.Y == line.to.Y
}

func calculateNumDangerous(onlyManhattan bool, inputList []string) int {
	lines := parseInput(inputList)
	ventMap := newVentmap()

	for _, line := range lines {
		if !onlyManhattan || isManhattan(line) {
			ventMap.addVent(line)
		}
	}

	return ventMap.numDangerous()
}

func part1(inputList []string) string {
	numDangerous := calculateNumDangerous(true, inputList)
	return strconv.Itoa(numDangerous)
}

func part2(inputList []string) string {
	numDangerous := calculateNumDangerous(false, inputList)
	return strconv.Itoa(numDangerous)
}

func main() {
	aoc.MainFunc(part1, part2)
}
