package main

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/set"
)

type cave struct {
	name                 string
	connectedCaveNameSet set.StringSet
}

func newCave(name string) cave {
	cave := cave{
		name:                 name,
		connectedCaveNameSet: set.NewStringSet(),
	}
	return cave
}

func (cave *cave) addConnection(otherCaveName string) {
	cave.connectedCaveNameSet.Add(string(otherCaveName))
}

func (cave *cave) getConnectedCaveNames() []string {
	return cave.connectedCaveNameSet.Values()
}

type caveData struct {
	caves map[string]cave
}

func newCaveData() caveData {
	data := caveData{}
	data.caves = make(map[string]cave)
	return data
}

func (data *caveData) addConnection(cave1, cave2 string) {
	cave, okay := data.caves[cave1]

	if !okay {
		cave = newCave(cave1)
		data.caves[cave1] = cave
	}
	cave.addConnection(cave2)

	cave, okay = data.caves[cave2]

	if !okay {
		cave = newCave(cave2)
		data.caves[cave2] = cave
	}
	cave.addConnection(cave1)
}

func parseInput(inputList []string) caveData {
	data := newCaveData()

	for _, line := range inputList {
		tokens := strings.Split(line, "-")
		data.addConnection(string(tokens[0]), string(tokens[1]))
	}

	return data
}

func isLegalPath(path string, canVisitOneSmallTwice bool) bool {
	smallCaveNames := set.NewStringMultiSet()
	visitedASmallCaveTwice := false

	for index, caveName := range strings.Split(path, ",") {
		if index != 0 && caveName == "start" {
			return false
		}

		if unicode.IsLower(rune(caveName[0])) {
			if !canVisitOneSmallTwice {
				if smallCaveNames.Contains(caveName) {
					return false
				}
				smallCaveNames.Add(caveName)
			} else {
				switch smallCaveNames.Count(caveName) {
				case 0:
					smallCaveNames.Add(caveName)
				case 1:
					if visitedASmallCaveTwice {
						return false
					}
					visitedASmallCaveTwice = true
					smallCaveNames.Add(caveName)
				case 2:
					// Visiting for the 3rd time
					return false
				}
			}
		}
	}

	return true
}

func simulate(caveData caveData, canVisitOneSmallTwice bool) int {

	type state struct {
		inCave string
		path   string
	}

	// Do breadth first search
	workQueue := make([]state, 0)
	workQueue = append(workQueue, state{inCave: "start", path: "start"})

	pathsVisited := set.NewStringSet()
	fullPaths := set.NewStringSet()

	for len(workQueue) != 0 {
		// Pop head
		currentState := workQueue[0]
		workQueue = workQueue[1:]

		pathsVisited.Add(currentState.path)

		if currentState.inCave == "end" {
			fullPaths.Add(currentState.path)
			continue
		}

		cave := caveData.caves[currentState.inCave]

		connectedCaveNames := cave.getConnectedCaveNames()
		sort.Strings(connectedCaveNames)
		for _, connectedCave := range connectedCaveNames {
			newPath := currentState.path + "," + connectedCave
			if !pathsVisited.Contains(newPath) && isLegalPath(newPath, canVisitOneSmallTwice) {
				newWork := state{inCave: connectedCave, path: newPath}
				workQueue = append(workQueue, newWork)
			}
		}
	}

	return fullPaths.Size()
}

func part1(inputList []string) string {
	caveData := parseInput(inputList)
	numPaths := simulate(caveData, false)
	return strconv.Itoa(numPaths)
}

func part2(inputList []string) string {
	caveData := parseInput(inputList)
	numPaths := simulate(caveData, true)
	return strconv.Itoa(numPaths)
}

func main() {
	aoc.MainFunc(part1, part2)
}
