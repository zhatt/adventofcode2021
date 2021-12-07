// Package aoc implements part 1 and part 2 runner for AOC.
//
package aoc

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type partFunc func(inputLines []string) string

// ReadInput reads an Advent of Code input file and converts it to
// a slice of strings, one line per entry.  Newlines are removed.
func ReadInput(inputFile string) []string {
	var lines []string
	file := os.Stdin

	if inputFile != "" {
		var err error
		file, err = os.Open(filepath.Clean(inputFile))
		if err != nil {
			panic(err)
		}
		defer func() {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
	return lines
}

// MainFunc is implements standard Advent of Code main function.
// It parses arguments and calls part1 or part2 function.
func MainFunc(part1, part2 partFunc) {
	var partOpt = 1
	var inputFile = ""

	flag.StringVar(&inputFile, "input", "", "input file")
	flag.IntVar(&partOpt, "part", 1, "part 1 or 2")
	flag.Parse()

	if partOpt != 1 && partOpt != 2 {
		panic("Invalid --part flag")
	}

	inputLines := ReadInput(inputFile)
	result := ""
	if partOpt == 1 {
		result = part1(inputLines)
	} else {
		result = part2(inputLines)
	}
	fmt.Println(result)
}

// PanicOnError - panic on error.
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
