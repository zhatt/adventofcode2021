package aoc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicOnError(t *testing.T) {
	assert.NotPanics(t, func() { PanicOnError(nil) })
	assert.Panics(t, func() { PanicOnError(fmt.Errorf("bad")) })
}

func part1(input []string) string {
	return "part1"
}

func part2(input []string) string {
	return "part2"
}

func part3(input []string) string {
	return input[0]
}

func TestMain(t *testing.T) {
	result := mainFunc([]string{}, part1, part2)
	assert.Equal(t, "part1", result)

	result = mainFunc([]string{"--part", "1"}, part1, part2)
	assert.Equal(t, "part1", result)

	result = mainFunc([]string{"--part", "2"}, part1, part2)
	assert.Equal(t, "part2", result)

	assert.Panics(t, func() { mainFunc([]string{"--part", "3"}, part1, part2) })

	result = mainFunc([]string{"--part", "3", "--input", "test_input.txt"}, part1, part2, part3)
	assert.Equal(t, "This file is for testing", result)
}
