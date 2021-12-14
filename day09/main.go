package main

import (
	"sort"
	"strconv"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

type heightmap struct {
	data []string
}

func newHeightMap() *heightmap {
	hm := &heightmap{}
	hm.data = make([]string, 0)
	return hm
}

func (hm *heightmap) addRow(s string) {
	hm.data = append(hm.data, s)
}

func (hm *heightmap) getValue(c coord.Coord) int {
	value, err := strconv.Atoi(hm.data[c.Yval][c.Xval : c.Xval+1])
	aoc.PanicOnError(err)
	return value
}

func (hm *heightmap) getSize() (columns int, rows int) {
	columns = len(hm.data[0])
	rows = len(hm.data)
	return
}

func (hm *heightmap) isLowPoint(c coord.Coord) bool {
	numColumns, numRows := hm.getSize()

	for _, coordToCheck := range []coord.Coord{
		{Xval: 0, Yval: 1},
		{Xval: 0, Yval: -1},
		{Xval: 1, Yval: 0},
		{Xval: -1, Yval: 0},
	} {
		toCheck := coord.Add(c, coordToCheck)
		if toCheck.Xval < 0 {
			continue
		}
		if toCheck.Xval >= numColumns {
			continue
		}
		if toCheck.Yval < 0 {
			continue
		}
		if toCheck.Yval >= numRows {
			continue
		}

		if hm.getValue(c) >= hm.getValue(toCheck) {
			return false
		}
	}

	return true
}

func (hm *heightmap) getBasinSizeR(c coord.Coord, visited map[coord.Coord]struct{}) int {
	if _, alreadyVisited := visited[c]; alreadyVisited {
		return 0
	}

	numColumns, numRows := hm.getSize()

	if c.Xval < 0 || c.Yval < 0 || c.Xval >= numColumns || c.Yval >= numRows {
		return 0
	}

	if hm.getValue(c) == 9 {
		return 0
	}

	visited[c] = struct{}{}
	basinSize := 1

	for _, coordToCheck := range []coord.Coord{
		{Xval: 0, Yval: 1},
		{Xval: 0, Yval: -1},
		{Xval: 1, Yval: 0},
		{Xval: -1, Yval: 0},
	} {
		toCheck := coord.Add(c, coordToCheck)
		basinSize += hm.getBasinSizeR(toCheck, visited)
	}
	return basinSize
}

func (hm *heightmap) getBasinSize(c coord.Coord) int {

	visited := make(map[coord.Coord]struct{})
	basinSize := hm.getBasinSizeR(c, visited)

	return basinSize
}

func parseInput(inputList []string) *heightmap {
	hm := newHeightMap()

	for _, line := range inputList {
		hm.addRow(line)
	}

	return hm
}

func findRiskLevel(hm *heightmap) (int, []coord.Coord) {
	lowPoints := make([]coord.Coord, 0)
	riskLevel := 0
	numColumns, numRows := hm.getSize()
	for row := 0; row < numRows; row++ {
		for column := 0; column < numColumns; column++ {
			location := coord.Coord{Xval: column, Yval: row}
			if hm.isLowPoint(location) {
				lowPoints = append(lowPoints, location)
				riskLevel += hm.getValue(location) + 1
			}
		}
	}
	return riskLevel, lowPoints
}

func findThreeLargestBasins(hm *heightmap) int {
	basinSizes := make([]int, 0)
	_, basinLocations := findRiskLevel(hm)

	for _, basinLocation := range basinLocations {
		basinSize := hm.getBasinSize(basinLocation)
		basinSizes = append(basinSizes, basinSize)
	}

	sort.Ints(basinSizes)

	basinSize := 1
	for i := len(basinSizes) - 3; i < len(basinSizes); i++ {
		basinSize *= basinSizes[i]
	}

	return basinSize
}

func part1(inputList []string) string {
	heightmap := parseInput(inputList)
	riskLevel, _ := findRiskLevel(heightmap)
	return strconv.Itoa(riskLevel)
}

func part2(inputList []string) string {
	heightmap := parseInput(inputList)
	basinSize := findThreeLargestBasins(heightmap)
	return strconv.Itoa(basinSize)
}

func main() {
	aoc.MainFunc(part1, part2)
}
