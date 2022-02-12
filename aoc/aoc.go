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

func mainFunc(args []string, partFuncs ...partFunc) string {
	var partOpt = 1
	var inputFile = ""

	flags := flag.FlagSet{}

	flags.StringVar(&inputFile, "input", "", "input file")
	flags.IntVar(&partOpt, "part", 1, "part to run")
	err := flags.Parse(args)
	PanicOnError(err)

	if partOpt > len(partFuncs) {
		panic("Invalid --part flag")
	}

	inputLines := ReadInput(inputFile)

	result := partFuncs[partOpt-1](inputLines)
	return result
}

// MainFunc implements standard Advent of Code main function.  It parses
// arguments and calls part1 or part2 function.
func MainFunc(partFuncs ...partFunc) {
	fmt.Println(mainFunc(os.Args[1:], partFuncs...))
}

// PanicOnError - panic on error.
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
