package main

import (
	"strconv"
	"testing"
	"zhatt/aoc2021/aoc"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []string{
	"8A004A801A8002F478",
	"620080001611562C8802118E34",
	"C0015000016115A2E0802F182340",
	"A0016C880162017C3686B18A3D4780",
}

func TestParseInput(t *testing.T) {
	input := []string{exampleInput1[0][0:5]}
	data := parseInput(input)

	assert.Equal(t, "10001010000000000100", data)
}

func TestPart1Example1(t *testing.T) {
	input := []string{"D2FE28"}
	result := part1(input)
	assert.Equal(t, "6", result)
}

func TestPart1Example2(t *testing.T) {
	input := []string{"38006F45291200"}
	result := part1(input)
	assert.Equal(t, strconv.Itoa(1+6+2), result)
}

func TestPart1Example3(t *testing.T) {
	input := []string{"EE00D40C823060"}
	result := part1(input)
	assert.Equal(t, strconv.Itoa(7+2+4+1), result)
}

func TestPart1Example4(t *testing.T) {
	result := part1(exampleInput1[0:1])
	assert.Equal(t, "16", result)

	result = part1(exampleInput1[1:2])
	assert.Equal(t, "12", result)

	result = part1(exampleInput1[2:3])
	assert.Equal(t, "23", result)

	result = part1(exampleInput1[3:4])
	assert.Equal(t, "31", result)
}

func TestPart1Input(t *testing.T) {
	result := part1(aoc.ReadInput("input.txt"))
	assert.Equal(t, "901", result)
}

func TestPart2Example1(t *testing.T) {
	input := []string{"C200B40A82"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(3), result)
}

func TestPart2Example2(t *testing.T) {
	input := []string{"04005AC33890"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(54), result)
}

func TestPart2Example3(t *testing.T) {
	input := []string{"880086C3E88112"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(7), result)
}

func TestPart2Example4(t *testing.T) {
	input := []string{"CE00C43D881120"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(9), result)
}

func TestPart2Example5(t *testing.T) {
	input := []string{"D8005AC2A8F0"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(1), result)
}

func TestPart2Example6(t *testing.T) {
	input := []string{"F600BC2D8F"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(0), result)
}

func TestPart2Example7(t *testing.T) {
	input := []string{"9C0141080250320F1802104A08"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(1), result)
}

func TestPart2Example8(t *testing.T) {
	input := []string{"9C005AC2F8F0"}
	result := part2(input)
	assert.Equal(t, strconv.Itoa(0), result)
}

func TestPart2Input(t *testing.T) {
	result := part2(aoc.ReadInput("input.txt"))
	assert.Equal(t, "110434737925", result)
}
