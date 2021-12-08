package main

import (
	"fmt"
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"
)

type command int

const (
	forward command = iota
	down
	up
)

type instruction struct {
	command command
	amount  int
}

func parseInput(inputList []string) []instruction {
	retVal := make([]instruction, 0, len(inputList))

	for _, strVal := range inputList {
		var instruction instruction
		tokens := strings.Fields(strVal)
		if len(tokens) != 2 {
			aoc.PanicOnError(fmt.Errorf("Can't parse input:  %s", strVal))
		}
		switch tokens[0] {
		case "forward":
			instruction.command = forward
		case "down":
			instruction.command = down
		case "up":
			instruction.command = up
		default:
			aoc.PanicOnError(fmt.Errorf("Unknown command:  %s", tokens[0]))
		}
		intVal, err := strconv.Atoi(tokens[1])
		aoc.PanicOnError(err)
		instruction.amount = intVal

		retVal = append(retVal, instruction)
	}

	return retVal
}

func calculateLocation(instructions []instruction, useAim bool) int {
	distance := 0
	aimOrDepth := 0
	aimDepth := 0

	for _, instruction := range instructions {
		switch instruction.command {
		case forward:
			distance += instruction.amount
			aimDepth += instruction.amount * aimOrDepth
		case down:
			aimOrDepth += instruction.amount
		case up:
			aimOrDepth -= instruction.amount
		default:
			aoc.PanicOnError(fmt.Errorf("Unknown command:  %d", instruction.command))
		}
	}

	if useAim {
		return distance * aimDepth
	} else {
		return distance * aimOrDepth
	}
}

func part1(inputList []string) string {
	instructions := parseInput(inputList)
	retVal := calculateLocation(instructions /* useAim= */, false)
	return strconv.Itoa(retVal)
}

func part2(inputList []string) string {
	instructions := parseInput(inputList)
	retVal := calculateLocation(instructions /* useAim= */, true)
	return strconv.Itoa(retVal)
}

func main() {
	aoc.MainFunc(part1, part2)
}
