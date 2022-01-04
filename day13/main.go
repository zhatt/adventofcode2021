package main

import (
	"fmt"
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

type direction int

const (
	foldX direction = iota
	foldY
)

type instruction struct {
	direction direction
	location  int
}

func newInstruction(direction direction, location int) instruction {
	return instruction{direction: direction, location: location}
}

type paper struct {
	points map[coord.Coord]struct{}
}

func newPaper() paper {
	return paper{points: make(map[coord.Coord]struct{})}
}

func (paper *paper) addPoint(c coord.Coord) {
	paper.points[c] = struct{}{}
}

func (paper *paper) fold(instruction instruction) {
	if instruction.direction == foldX {
		for point := range paper.points {
			if point.X > instruction.location {
				delete(paper.points, point)
				newPoint := coord.Coord{
					X: instruction.location - (point.X - instruction.location),
					Y: point.Y,
				}

				paper.points[newPoint] = struct{}{}
			}
		}

	}
	if instruction.direction == foldY {
		for point := range paper.points {
			if point.Y > instruction.location {
				delete(paper.points, point)
				newPoint := coord.Coord{
					X: point.X,
					Y: instruction.location - (point.Y - instruction.location),
				}

				paper.points[newPoint] = struct{}{}
			}
		}

	}
}

func (paper *paper) numberOfPoints() int {
	return len(paper.points)
}

func (paper paper) String() string {
	coords := make([]coord.Coord, 0, len(paper.points))

	for coord := range paper.points {
		coords = append(coords, coord)
	}

	minBound := coord.MinBound(coords)
	maxBound := coord.MaxBound(coords)

	output := ""
	for y := minBound.Y; y <= maxBound.Y; y++ {
		if y != 0 {
			output += "\n"
		}
		for x := minBound.X; x <= maxBound.X; x++ {
			if _, exists := paper.points[coord.Coord{X: x, Y: y}]; exists {
				output += "#"
			} else {
				output += " "
			}
		}
	}
	return output
}

func parseInput(inputLines []string) (paper, []instruction) {
	paper := newPaper()
	instructions := make([]instruction, 0)

	for index, line := range inputLines {
		if line == "" {
			inputLines = inputLines[index+1:]
			break
		}
		tokens := strings.Split(line, ",")
		xval, err := strconv.Atoi(tokens[0])
		aoc.PanicOnError(err)
		yval, err := strconv.Atoi(tokens[1])
		aoc.PanicOnError(err)
		paper.addPoint(coord.Coord{X: xval, Y: yval})
	}

	for _, line := range inputLines {
		tokens := strings.Split(line, "=")
		direction := foldX
		if tokens[0][len(tokens[0])-1] == 'y' {
			direction = foldY
		}
		amount, err := strconv.Atoi(tokens[1])
		aoc.PanicOnError(err)

		instruction := newInstruction(direction, amount)
		instructions = append(instructions, instruction)
	}

	return paper, instructions
}

func part1(inputList []string) string {
	paper, instructions := parseInput(inputList)
	paper.fold(instructions[0])
	count := paper.numberOfPoints()
	return strconv.Itoa(count)
}

func part2(inputList []string) string {
	paper, instructions := parseInput(inputList)
	for _, instruction := range instructions {
		paper.fold(instruction)
	}

	return fmt.Sprintln(paper)
}

func main() {
	aoc.MainFunc(part1, part2)
}
