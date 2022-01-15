package main

import (
	"fmt"
	"strconv"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
	"zhatt/aoc2021/day22/cuboid"
)

type command int

const (
	off command = iota
	on
)

type rebootStep struct {
	command command
	cuboid  cuboid.Cuboid
}

func parseInput(inputLines []string) []rebootStep {

	steps := make([]rebootStep, 0, len(inputLines))

	for _, line := range inputLines {
		cmd := on

		var xMin, xMax, yMin, yMax, zMin, zMax int

		// Parses: on x=-20..26,y=-36..17,z=-47..7
		// or:     off x=-20..26,y=-36..17,z=-47..7
		_, err := fmt.Sscanf(line, "on x=%d..%d,y=%d..%d,z=%d..%d", &xMin, &xMax, &yMin, &yMax, &zMin, &zMax)
		if err != nil {
			_, err = fmt.Sscanf(line, "off x=%d..%d,y=%d..%d,z=%d..%d", &xMin, &xMax, &yMin, &yMax, &zMin, &zMax)
			aoc.PanicOnError(err)
			cmd = off
		}

		step := rebootStep{
			command: cmd,
			cuboid:  cuboid.New(coord.Coord3d{X: xMin, Y: yMin, Z: zMin}, coord.Coord3d{X: xMax, Y: yMax, Z: zMax}),
		}

		steps = append(steps, step)
	}

	return steps
}

func part1(inputList []string) string {
	steps := parseInput(inputList)

	initializationArea := cuboid.New(coord.Coord3d{X: -50, Y: -50, Z: -50}, coord.Coord3d{X: 50, Y: 50, Z: 50})

	reactorCore := make(map[coord.Coord3d]struct{})

	for _, step := range steps {
		if cuboid.Overlap(step.cuboid, initializationArea) {
			for x := step.cuboid.MinVertex().X; x <= step.cuboid.MaxVertex().X; x++ {
				for y := step.cuboid.MinVertex().Y; y <= step.cuboid.MaxVertex().Y; y++ {
					for z := step.cuboid.MinVertex().Z; z <= step.cuboid.MaxVertex().Z; z++ {
						c := coord.Coord3d{X: x, Y: y, Z: z}
						if !coord.OutOfBounds3d(c, []coord.Coord3d{initializationArea.MinVertex(), initializationArea.MaxVertex()}) {
							if step.command == on {
								reactorCore[c] = struct{}{}
							} else {
								delete(reactorCore, c)
							}
						}
					}
				}
			}

		}
	}

	return strconv.Itoa(len(reactorCore))
}

func part2(inputList []string) string {
	steps := parseInput(inputList)

	reactorCore := cuboid.NewSet()

	for _, step := range steps {
		if step.command == on {
			reactorCore.Add(step.cuboid)
		} else {
			reactorCore.Remove(step.cuboid)
		}
	}

	return strconv.Itoa(reactorCore.Volume())
}

func main() {
	aoc.MainFunc(part1, part2)
}
