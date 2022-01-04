package main

import (
	"math"
	"strconv"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"

	"zhatt/aoc2021/day15/workqueue"
)

func parseInput(inputLines []string) [][]int {
	data := make([][]int, 0, len(inputLines))

	for _, line := range inputLines {
		row := make([]int, 0, len(line))
		for _, digit := range line {
			riskLevel, err := strconv.Atoi(string(digit))
			aoc.PanicOnError(err)
			row = append(row, riskLevel)
		}
		data = append(data, row)
	}

	return data
}

// Calculate shortest distance using Dijkstra's algorithm.
func simulate(data [][]int) int {
	minBound := coord.Coord{X: 0, Y: 0}
	maxBound := coord.Coord{X: len(data[0]) - 1, Y: len(data) - 1}
	bounds := []coord.Coord{minBound, maxBound}

	workQueue := workqueue.NewWorkQueue()
	unvisited := make(map[coord.Coord]bool)
	locations := make(map[coord.Coord]*workqueue.Item)

	// Initialize datastructures
	for y := minBound.Y; y <= maxBound.Y; y++ {
		for x := minBound.X; x <= maxBound.X; x++ {
			distance := math.MaxInt
			if x == 0 && y == 0 {
				distance = 0
			}
			coord := coord.Coord{X: x, Y: y}
			work := workqueue.NewItem(coord, distance)
			unvisited[coord] = true
			locations[coord] = work
			workQueue.Push(work)
		}
	}

	// NB.  It may be possible to stop right after we find end distance but
	// we are will calculate all distances instead.
	for workQueue.Len() != 0 {
		work := workQueue.Pop()
		currentLocation := work.Location

		for _, delta := range []coord.Coord{
			{X: 1, Y: 0},
			{X: -1, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: -1},
		} {
			neighborLocation := coord.Add(currentLocation, delta)
			if coord.OutOfBounds(neighborLocation, bounds) {
				continue
			}

			if !unvisited[neighborLocation] {
				continue
			}

			newNeighborDistance := work.Distance +
				data[neighborLocation.Y][neighborLocation.X]

			neighbor := locations[neighborLocation]
			if newNeighborDistance < neighbor.Distance {
				neighbor.Distance = newNeighborDistance
				workQueue.Update(neighbor)
			}
		}
		unvisited[currentLocation] = false
	}

	return locations[maxBound].Distance
}

func makeFullMap(scanData [][]int) [][]int {
	mapData := make([][]int, 0, len(scanData)*5)
	scanSizeX := len(scanData[0])
	scanSizeY := len(scanData)

	for y := 0; y < 5*scanSizeY; y++ {
		row := make([]int, 0, 5*scanSizeX)
		for x := 0; x < 5*scanSizeX; x++ {
			adjustedValue := scanData[y%scanSizeY][x%scanSizeX] + y/scanSizeY + x/scanSizeX
			for adjustedValue > 9 {
				adjustedValue -= 9
			}

			row = append(row, adjustedValue)
		}
		mapData = append(mapData, row)
	}
	return mapData
}

func part1(inputList []string) string {
	data := parseInput(inputList)
	risk := simulate(data)
	return strconv.Itoa(risk)
}

func part2(inputList []string) string {
	data := parseInput(inputList)
	data = makeFullMap(data)
	risk := simulate(data)
	return strconv.Itoa(risk)
}

func main() {
	aoc.MainFunc(part1, part2)
}
