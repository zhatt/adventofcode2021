package main

import (
	"fmt"
	"strconv"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

type bounds struct {
	min coord.Coord
	max coord.Coord
}

func parseInput(inputLines []string) bounds {
	bounds := bounds{}

	// Format: target area: x=20..30, y=-10..-5
	_, err := fmt.Sscanf(inputLines[0], "target area: x=%d..%d, y=%d..%d",
		&bounds.min.Xval, &bounds.max.Xval,
		&bounds.min.Yval, &bounds.max.Yval,
	)
	aoc.PanicOnError(err)

	return bounds
}

func calculateInitialXRange(bounds bounds) (int, int) {

	var minInitialXVelocity int
	var maxInitialVelocity int

	for initialXVelocity := 0; ; initialXVelocity++ {
		velocity := initialXVelocity
		x := 0
		for ; velocity != 0; x += velocity {
			velocity--
		}

		if x >= bounds.min.Xval {
			minInitialXVelocity = initialXVelocity - 1
			break
		}
	}

	maxInitialVelocity = bounds.max.Xval

	return minInitialXVelocity, maxInitialVelocity
}

func simulateY(bounds bounds, initialSteps int, totalSteps int, velocity int) (
	int, // max height
	bool, // in range
	bool, // too fast
	bool, // too slow
) {
	y := 0
	maxHeight := 0
	inRange := 0
	for s := 0; s < totalSteps; s++ {
		y += velocity
		velocity--

		if y > maxHeight {
			maxHeight = y
		}

		if y >= bounds.min.Yval && y <= bounds.max.Yval && s >= initialSteps-1 {
			inRange++
			break
		}

		if y < bounds.min.Yval {
			break
		}
	}

	if inRange != 0 {
		return maxHeight, true, false, false
	}

	if y > bounds.max.Yval {
		return 0, false, true, false
	}

	if y < bounds.min.Yval {
		return 0, false, false, true
	}

	return 0, false, false, false
}

func simulate(bounds bounds, xVel int, inRangeSet map[coord.Coord]bool) int {

	initialXvel := xVel
	steps := 0
	maxHeight := 0

	for x := 0; ; {
		steps++
		x += xVel
		xVel -= 1

		if x > bounds.max.Xval {
			break
		}

		if xVel < 0 {
			break
		}

		if x >= bounds.min.Xval {
			// In region
			// Need to simulate y
			for yVel := 0; ; {
				simSteps := steps
				if xVel == 0 {
					// NB.  This hack is probably not correct for 100% of inputs.
					simSteps = 1000
				}

				height, inRange, tooFast, _ := simulateY(bounds, steps, simSteps, yVel)

				if tooFast {
					break
				}
				if inRange {
					if height > maxHeight {
						maxHeight = height
					}
					inRangeSet[coord.Coord{Xval: initialXvel, Yval: yVel}] = true
				}
				yVel++

			}
			for yVel := -1; ; {
				simSteps := steps
				if xVel == 0 {
					// NB.  This hack is probably not correct for 100% of inputs.
					simSteps = 1000
				}

				height, inRange, _, tooSlow := simulateY(bounds, steps, simSteps, yVel)

				if tooSlow {
					break
				}
				if inRange {
					if height > maxHeight {
						maxHeight = height
					}
					inRangeSet[coord.Coord{Xval: initialXvel, Yval: yVel}] = true
				}

				yVel--
			}
		}
	}

	return maxHeight
}

func part1(inputList []string) string {
	bounds := parseInput(inputList)

	maxHeight := 0
	inRangeSet := make(map[coord.Coord]bool)

	minX, maxX := calculateInitialXRange(bounds)
	for xVel := minX; xVel <= maxX; xVel++ {
		height := simulate(bounds, xVel, inRangeSet)
		if height > maxHeight {
			maxHeight = height
		}
	}

	return strconv.Itoa(maxHeight)
}

func part2(inputList []string) string {
	bounds := parseInput(inputList)

	inRangeSet := make(map[coord.Coord]bool)

	minX, maxX := calculateInitialXRange(bounds)
	for xVel := minX; xVel <= maxX; xVel++ {
		_ = simulate(bounds, xVel, inRangeSet)
	}

	return strconv.Itoa(len(inRangeSet))
}

func main() {
	aoc.MainFunc(part1, part2)
}
