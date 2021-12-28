package main

import (
	"fmt"
	"strconv"
	"zhatt/aoc2021/aoc"
)

type computer struct {
	totalVersion uint64
	stream       string
}

func (bits *computer) getTotalVersion() uint64 { return bits.totalVersion }

func parseInput(inputLines []string) string {
	data := ""

	for index := range inputLines[0] {
		dataByte, err := strconv.ParseUint(inputLines[0][index:index+1], 16, 8)
		aoc.PanicOnError(err)
		data += fmt.Sprintf("%04b", dataByte)
	}

	return data
}

type typeID uint

const (
	typeSum          typeID = 0
	typeProduct      typeID = 1
	typeMin          typeID = 2
	typeMax          typeID = 3
	typeLiteralValue typeID = 4
	typeGreaterThan  typeID = 5
	typeLessThan     typeID = 6
	typeEqual        typeID = 7
)

func (c *computer) parseLiteralValueBits() uint64 {
	value := uint64(0)
	for {
		bits := c.stream[0:5]
		c.stream = c.stream[5:]

		b, err := strconv.ParseUint(bits[1:5], 2, 4)
		aoc.PanicOnError(err)
		value = (value << 4) + b
		if bits[0:1] == "0" {
			break
		}
	}

	return value
}

func (c *computer) parseOperatorTotalLengthMode() []uint64 {
	values := make([]uint64, 0)

	length, err := strconv.ParseUint(c.stream[0:15], 2, 16)
	aoc.PanicOnError(err)
	c.stream = c.stream[15:]

	targetLength := len(c.stream) - int(length)

	for len(c.stream) > targetLength {
		values = append(values, c.parseBits()...)
	}

	return values
}

func (c *computer) parseOperationNumberOfPacketsMode() []uint64 {
	values := make([]uint64, 0)
	numSubPackets, err := strconv.ParseUint(c.stream[0:11], 2, 16)
	aoc.PanicOnError(err)
	c.stream = c.stream[11:]

	for i := 0; i < int(numSubPackets); i++ {
		values = append(values, c.parseBits()...)
	}

	return values
}

func (c *computer) parseBits() []uint64 {
	values := make([]uint64, 0)

	version, err := strconv.ParseUint(c.stream[0:3], 2, 8)
	aoc.PanicOnError(err)
	c.stream = c.stream[3:]

	c.totalVersion += version

	typeidInt, err := strconv.ParseUint(c.stream[0:3], 2, 8)
	aoc.PanicOnError(err)
	c.stream = c.stream[3:]

	typeid := typeID(typeidInt)

	switch typeid {
	case typeLiteralValue:
		value := c.parseLiteralValueBits()
		values = append(values, value)

	default: // Operator packets
		lengthTypeIDBits := c.stream[0:1]
		c.stream = c.stream[1:]

		var opValues []uint64
		if lengthTypeIDBits == "0" { // total length in bits of sub-packets mode
			opValues = c.parseOperatorTotalLengthMode()
		} else {
			opValues = c.parseOperationNumberOfPacketsMode()
		}

		switch typeid {
		case typeSum:
			sum := uint64(0)
			for _, opValue := range opValues {
				sum += opValue
			}
			values = append(values, sum)
		case typeProduct:
			sum := uint64(1)
			for _, opValue := range opValues {
				sum *= opValue
			}
			values = append(values, sum)
		case typeMin:
			min := opValues[0]
			for _, opValue := range opValues {
				if opValue < min {
					min = opValue
				}
			}
			values = append(values, min)
		case typeMax:
			max := opValues[0]
			for _, opValue := range opValues {
				if opValue > max {
					max = opValue
				}
			}
			values = append(values, max)
		case typeGreaterThan:
			if opValues[0] > opValues[1] {
				values = append(values, 1)
			} else {
				values = append(values, 0)
			}
		case typeLessThan:
			if opValues[0] < opValues[1] {
				values = append(values, 1)
			} else {
				values = append(values, 0)
			}
		case typeEqual:
			if opValues[0] == opValues[1] {
				values = append(values, 1)
			} else {
				values = append(values, 0)
			}
		default:
			panic(fmt.Errorf("unimplemented %d", typeid))
		}
	}

	return values
}

func (c *computer) run() uint64 {
	values := c.parseBits()
	return values[0]
}

func part1(inputList []string) string {
	stream := parseInput(inputList)
	c := computer{stream: stream}

	c.run()
	totalVersion := c.getTotalVersion()

	return strconv.FormatUint(totalVersion, 10)
}

func part2(inputList []string) string {
	stream := parseInput(inputList)
	c := computer{stream: stream}

	value := c.run()

	return strconv.FormatUint(value, 10)
}

func main() {
	aoc.MainFunc(part1, part2)
}
