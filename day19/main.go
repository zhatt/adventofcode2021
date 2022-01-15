package main

import (
	"strconv"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"

	"zhatt/aoc2021/day19/scanner"
)

func parseInput(inputLines []string) []scanner.Scanner {
	inputLines = append(inputLines, "")

	scanners := make([]scanner.Scanner, 0)

	startIndex := 0
	for index := range inputLines {
		if inputLines[index] == "" {
			scanner := scanner.New()
			scanner.Load(inputLines[startIndex:index])
			startIndex = index + 1

			scanners = append(scanners, scanner)
		}
	}

	return scanners
}

func findBeacons(scanners []scanner.Scanner) ([]coord.Coord3d, []coord.Coord3d) {

	placedScanners := make(map[int]scanner.Scanner)
	unplacedScanners := make(map[int]scanner.Scanner)

	// We will reference all scanners to scanner 0.
	placedScanners[scanners[0].Number()] = scanners[0]
	for _, scanner := range scanners[1:] {
		unplacedScanners[scanner.Number()] = scanner
	}

RESTART:
	for len(unplacedScanners) != 0 {
		for _, unplacedScanner := range unplacedScanners {
			for _, placedScanner := range placedScanners {
				correlated, coord, rotation := scanner.Correlate(placedScanner, unplacedScanner)

				if correlated {
					unplacedScanner.SetLocation(coord)
					unplacedScanner.SetRotation(rotation)

					placedScanners[unplacedScanner.Number()] = unplacedScanner
					delete(unplacedScanners, unplacedScanner.Number())

					continue RESTART
				}
			}
		}
	}

	beaconMap := make(map[coord.Coord3d]bool)
	scannerLocations := make([]coord.Coord3d, 0, len(scanners))

	for _, placedScanner := range placedScanners {
		scannerLocations = append(scannerLocations, placedScanner.Location())
		for _, beacon := range placedScanner.Beacons() {
			beaconMap[beacon] = true
		}
	}

	beaconLocations := make([]coord.Coord3d, 0, len(beaconMap))

	for beacon := range beaconMap {
		beaconLocations = append(beaconLocations, beacon)
	}

	return scannerLocations, beaconLocations
}

func part1(inputList []string) string {
	scanners := parseInput(inputList)
	_, beacons := findBeacons(scanners)
	return strconv.Itoa(len(beacons))
}

func part2(inputList []string) string {
	scanners := parseInput(inputList)
	scannerLocations, _ := findBeacons(scanners)

	maxDistance := 0
	for _, b1 := range scannerLocations {
		for _, b2 := range scannerLocations {
			distance := coord.DistanceManhattan3d(b1, b2)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return strconv.Itoa(maxDistance)
}

func main() {
	aoc.MainFunc(part1, part2)
}
