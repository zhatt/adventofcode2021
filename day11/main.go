package main

import (
	"fmt"
	"strconv"
	"zhatt/aoc2021/aoc"
)

type mode int

const (
	untilCount mode = iota
	untilAllFlash
)

type grid struct {
	size  int
	octos [][]octo
}

func newGrid(size int) *grid {
	grid := &grid{size: size}
	grid.octos = make([][]octo, size)
	for i := 0; i < size; i++ {
		grid.octos[i] = make([]octo, size)
	}

	return grid
}

type octo struct {
	energy     int
	flashed    bool
	numFlashes int
}

func (o *octo) powerUp() {
	if !o.flashed {
		o.energy += 1
	}
}

func (o *octo) checkFlashed() bool {
	if o.energy >= 10 {
		o.flashed = true
		o.energy = 0
		o.numFlashes++
		return true
	}
	return false
}

func (o *octo) reset() {
	o.flashed = false
}

func (o *octo) numFlashesFlashed() int {
	return o.numFlashes
}

func (o *octo) setEnergy(energy int) {
	o.energy = energy
}

func parseInput(inputList []string) *grid {
	size := len(inputList[0])
	grid := newGrid(size)
	for y, line := range inputList {
		for x, energy := range line {
			energy, err := strconv.Atoi(string(energy))
			aoc.PanicOnError(err)

			grid.octos[x][y].setEnergy(energy)
		}
	}
	return grid
}

func (grid *grid) String() string {
	s := ""
	for y := range grid.octos {
		for x := range grid.octos[y] {
			// This adds newline to all lines except last
			if x != 0 && y != 0 {
				s += "\n"
			}
			energy := grid.octos[x][y].energy
			s += fmt.Sprintf("%d", energy)
		}
	}
	return s
}

func simulate(grid *grid) int {
	for y := range grid.octos {
		for x := range grid.octos[y] {
			grid.octos[x][y].reset()
			grid.octos[x][y].powerUp()
		}
	}

	numFlashed := 0

restart:
	for {
		for y := range grid.octos {
			for x := range grid.octos[y] {
				flashed := grid.octos[x][y].checkFlashed()

				if flashed {
					numFlashed++
					// Power up neighbors
					for yinc := -1; yinc <= 1; yinc++ {
						for xinc := -1; xinc <= 1; xinc++ {
							if xinc == 0 && yinc == 0 {
								continue
							}
							if x+xinc < 0 || y+yinc < 0 {
								continue
							}
							if x+xinc == len(grid.octos[y]) || y+yinc == len(grid.octos) {
								continue
							}
							grid.octos[x+xinc][y+yinc].powerUp()
						}
					}
					continue restart
				}
			}
		}
		// No more flashes
		break
	}

	return numFlashed
}

func simulateSteps(mode mode, grid *grid, steps int) int {
	numOctos := grid.size * grid.size
	for i := 0; i < steps || mode == untilAllFlash; i++ {
		numFlashes := simulate(grid)
		if mode == untilAllFlash && numFlashes == numOctos {
			return i + 1
		}
	}

	totalNumberOfFlashes := 0
	for y := range grid.octos {
		for x := range grid.octos[y] {
			totalNumberOfFlashes += grid.octos[x][y].numFlashesFlashed()
		}
	}
	return totalNumberOfFlashes
}

func part1(inputList []string) string {
	grid := parseInput(inputList)
	numFlashes := simulateSteps(untilCount, grid, 100)
	return strconv.Itoa(numFlashes)
}

func part2(inputList []string) string {
	grid := parseInput(inputList)
	numSteps := simulateSteps(untilAllFlash, grid, 0)
	return strconv.Itoa(numSteps)
}

func main() {
	aoc.MainFunc(part1, part2)
}
