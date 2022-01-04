package main

import (
	"strconv"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

type algorithm []bool

type image struct {
	minBound coord.Coord
	maxBound coord.Coord
	pixels   map[coord.Coord]bool
}

func (image image) String() string {
	output := ""
	for y := image.minBound.Y; y <= image.maxBound.Y; y++ {
		if y != image.minBound.Y {
			output += "\n"
		}
		for x := image.minBound.X; x <= image.maxBound.X; x++ {
			if image.pixels[coord.Coord{X: x, Y: y}] {
				output += "#"
			} else {
				output += "."
			}
		}
	}

	return output
}

func parseInput(inputLines []string) (algorithm, image) {

	algorithm := algorithm(make([]bool, len(inputLines[0])))

	for index, command := range inputLines[0] {
		if command == '#' {
			algorithm[index] = true
		}
	}

	image := image{
		pixels: make(map[coord.Coord]bool),
	}

	for y, line := range inputLines[2:] {
		for x, pixel := range line {
			coord := coord.Coord{X: x, Y: y}
			image.maxBound = coord
			if pixel == '#' {
				image.pixels[coord] = true
			} else {
				image.pixels[coord] = false
			}
		}
	}

	return algorithm, image
}

func enhance(image *image, algorithm algorithm, assume bool) {
	newImage := *image
	// Each iteration will grow the image by 1 on each side.
	extra := 1
	newImage.minBound.X -= extra
	newImage.minBound.Y -= extra
	newImage.maxBound.X += extra
	newImage.maxBound.Y += extra

	newImage.pixels = make(map[coord.Coord]bool)

	for y := image.minBound.Y - extra; y <= image.maxBound.Y+extra; y++ {
		for x := image.minBound.X - extra; x <= image.maxBound.X+extra; x++ {

			value := uint(0)
			for y1 := -1; y1 <= 1; y1++ { // NB. Y increases downward
				for x1 := -1; x1 <= 1; x1++ {
					value <<= 1
					c := coord.Coord{X: x + x1, Y: y + y1}

					on, okay := image.pixels[c]
					if !okay {
						on = assume
					}
					if on {
						value |= 1
					}
				}
			}

			command := algorithm[value]
			if command {
				newImage.pixels[coord.Coord{X: x, Y: y}] = true
			} else {
				newImage.pixels[coord.Coord{X: x, Y: y}] = false
			}
		}
	}

	*image = newImage
}

func simulate(algorithm algorithm, image image, iterations int) int {
	// Figure out if we will be turning the infinite space on and off.  If
	// we are turning it on add off, we will toggle what we needToToggleAssume the
	// contents are each iteration.
	//
	// The example does not turn on for command 0 but the real data does.
	needToToggleAssume := algorithm[0] && !algorithm[511]
	currentAssume := false

	for i := 0; i < iterations; i++ {
		enhance(&image, algorithm, currentAssume)
		currentAssume = currentAssume != needToToggleAssume // toggle
	}
	count := 0
	for _, value := range image.pixels {
		if value {
			count++
		}
	}
	return count
}

func part1(inputList []string) string {
	algorithm, image := parseInput(inputList)
	count := simulate(algorithm, image, 2)
	return strconv.Itoa(count)
}

func part2(inputList []string) string {
	algorithm, image := parseInput(inputList)
	count := simulate(algorithm, image, 50)
	return strconv.Itoa(count)
}

func main() {
	aoc.MainFunc(part1, part2)
}
