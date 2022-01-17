package main

import (
	"math"
	"strconv"
	"zhatt/aoc2021/aoc"

	"zhatt/aoc2021/day23/workqueue"
)

/*
This is layout of our locations.

#############
#12.3.4.5.67#  hall
###A#B#C#D###  1
  #A#B#C#D#    2
  #A#B#C#D#    3
  #A#B#C#D#    4
  #########
*/

type location int

const (
	roomA1 location = iota
	roomA2
	roomA3
	roomA4
	roomB1
	roomB2
	roomB3
	roomB4
	roomC1
	roomC2
	roomC3
	roomC4
	roomD1
	roomD2
	roomD3
	roomD4
	hall1
	hall2
	hall3
	hall4
	hall5
	hall6
	hall7
	numberOfLocations
)

var allLocations = [...]location{
	roomA1,
	roomA2,
	roomA3,
	roomA4,
	roomB1,
	roomB2,
	roomB3,
	roomB4,
	roomC1,
	roomC2,
	roomC3,
	roomC4,
	roomD1,
	roomD2,
	roomD3,
	roomD4,
	hall1,
	hall2,
	hall3,
	hall4,
	hall5,
	hall6,
	hall7,
}

type move struct {
	to       location
	path     []location
	distance int
}

type amphipodType int

const (
	noAmphipod amphipodType = iota
	amberAmphipod
	bronzeAmphipod
	copperAmphipod
	desertAmphipod
)

var toAmphipod = map[string]amphipodType{
	"A": amberAmphipod,
	"B": bronzeAmphipod,
	"C": copperAmphipod,
	"D": desertAmphipod,
}

type locationInfo struct {
	occupant  amphipodType
	peerRooms []location
	moves     []move // Legal moves from this location
}

type burrowType [numberOfLocations]amphipodType

var moveCost = map[amphipodType]int{
	amberAmphipod:  1,
	bronzeAmphipod: 10,
	copperAmphipod: 100,
	desertAmphipod: 1000,
}

type partEnum int

const (
	part1e partEnum = iota
	part2e
)

func parseInput(inputLines []string, part partEnum) burrowType {
	b := burrowType{}

	b[hall1] = noAmphipod
	b[hall2] = noAmphipod
	b[hall3] = noAmphipod
	b[hall4] = noAmphipod
	b[hall5] = noAmphipod
	b[hall6] = noAmphipod
	b[hall7] = noAmphipod

	b[roomA1] = toAmphipod[inputLines[2][3:4]]
	b[roomB1] = toAmphipod[inputLines[2][5:6]]
	b[roomC1] = toAmphipod[inputLines[2][7:8]]
	b[roomD1] = toAmphipod[inputLines[2][9:10]]

	if part == part1e {
		b[roomA2] = toAmphipod[inputLines[3][3:4]]
		b[roomB2] = toAmphipod[inputLines[3][5:6]]
		b[roomC2] = toAmphipod[inputLines[3][7:8]]
		b[roomD2] = toAmphipod[inputLines[3][9:10]]
	} else {
		b[roomA2] = desertAmphipod
		b[roomB2] = copperAmphipod
		b[roomC2] = bronzeAmphipod
		b[roomD2] = amberAmphipod
	}

	if part == part1e {
		b[roomA3] = amberAmphipod
		b[roomB3] = bronzeAmphipod
		b[roomC3] = copperAmphipod
		b[roomD3] = desertAmphipod
	} else {
		b[roomA3] = desertAmphipod
		b[roomB3] = bronzeAmphipod
		b[roomC3] = amberAmphipod
		b[roomD3] = copperAmphipod
	}

	if part == part1e {
		b[roomA4] = amberAmphipod
		b[roomB4] = bronzeAmphipod
		b[roomC4] = copperAmphipod
		b[roomD4] = desertAmphipod
	} else {
		b[roomA4] = toAmphipod[inputLines[3][3:4]]
		b[roomB4] = toAmphipod[inputLines[3][5:6]]
		b[roomC4] = toAmphipod[inputLines[3][7:8]]
		b[roomD4] = toAmphipod[inputLines[3][9:10]]
	}

	return b
}

func canMoveHere(amphipod amphipodType, peerRooms []location, burrow *burrowType) bool {

	if peerRooms == nil {
		return true
	}

	for _, room := range peerRooms {
		if moveData[room].occupant != amphipod {
			return false
		}

		if burrow[room] != noAmphipod && burrow[room] != amphipod {
			return false
		}
	}

	return true
}

func organized(burrow *burrowType) bool {

	return burrow[roomA1] == amberAmphipod &&
		burrow[roomA2] == amberAmphipod &&
		burrow[roomA3] == amberAmphipod &&
		burrow[roomA4] == amberAmphipod &&
		burrow[roomB1] == bronzeAmphipod &&
		burrow[roomB2] == bronzeAmphipod &&
		burrow[roomB3] == bronzeAmphipod &&
		burrow[roomB4] == bronzeAmphipod &&
		burrow[roomC1] == copperAmphipod &&
		burrow[roomC2] == copperAmphipod &&
		burrow[roomC3] == copperAmphipod &&
		burrow[roomC4] == copperAmphipod &&
		burrow[roomD1] == desertAmphipod &&
		burrow[roomD2] == desertAmphipod &&
		burrow[roomD3] == desertAmphipod &&
		burrow[roomD4] == desertAmphipod
}

type workType struct {
	burrow burrowType
	score  int
}

func simulate(burrow burrowType) int {

	minScore := math.MaxInt

	workQueue := workqueue.NewWorkQueue()
	workQueue.Push(workqueue.NewItem(workType{burrow: burrow}, minScore))

	beenThere := make(map[burrowType]int)
	beenThere[burrow] = 0

	for workQueue.Len() > 0 {
		work := workQueue.Pop().Item.(workType)

		if organized(&(work.burrow)) {
			if work.score < minScore {
				minScore = work.score
			}
			continue
		}

		for _, moveFrom := range allLocations {
			if work.burrow[moveFrom] == noAmphipod {
				continue
			}

		NEXTMOVE:
			for _, moveTo := range moveData[moveFrom].moves {

				amphipod := work.burrow[moveFrom]

				if !canMoveHere(amphipod, moveData[moveTo.to].peerRooms, &(work.burrow)) {
					continue
				}

				for _, l := range moveTo.path {
					if work.burrow[l] != noAmphipod {
						continue NEXTMOVE
					}
				}

				newScore := work.score + moveTo.distance*moveCost[amphipod]

				if newScore > minScore {
					continue NEXTMOVE
				}

				newWork := work
				newWork.score = newScore
				newWork.burrow[moveFrom] = noAmphipod
				newWork.burrow[moveTo.to] = amphipod

				if cost, okay := beenThere[newWork.burrow]; okay {
					if cost <= newWork.score {
						continue NEXTMOVE
					}
				}
				beenThere[newWork.burrow] = newWork.score

				workQueue.Push(workqueue.NewItem(newWork, newWork.score))
			}
		}
	}

	return minScore
}

func part1(inputList []string) string {
	burrow := parseInput(inputList, part1e)

	score := simulate(burrow)

	return strconv.Itoa(score)
}

func part2(inputList []string) string {
	burrow := parseInput(inputList, part2e)

	score := simulate(burrow)

	return strconv.Itoa(score)
}

func main() {
	aoc.MainFunc(part1, part2)
}
