package main

import (
	"container/list"
	"fmt"
	"strconv"
	"zhatt/aoc2021/aoc"
)

type deterministicDie struct {
	nextValue     int
	numberOfSides int
	numberOfRolls int
}

func newDeterministicDie(numberOfSides int) deterministicDie {

	return deterministicDie{
		nextValue:     1,
		numberOfSides: numberOfSides,
	}
}

func (d *deterministicDie) roll() int {
	value := d.nextValue

	d.nextValue++
	if d.nextValue > d.numberOfSides {
		d.nextValue = 1
	}

	d.numberOfRolls++

	return value
}

type player struct {
	position int
	score    int
}

func (p *player) move(distance int) int {
	// NB. position goes from 1 to 10.  Subract and add 1 to be able to use
	// modulus to wrap.
	p.position = ((p.position - 1 + distance) % 10) + 1
	p.score += p.position

	return p.score
}

func newPlayer(position int) player {
	return player{position: position}
}

func parseInput(inputLines []string) (int, int) {

	var player1Position int
	var player2Position int
	var player int

	// Parses: Player 1 starting position: 4
	_, err := fmt.Sscanf(inputLines[0], "Player %d starting position: %d", &player, &player1Position)
	aoc.PanicOnError(err)
	_, err = fmt.Sscanf(inputLines[1], "Player %d starting position: %d", &player, &player2Position)
	aoc.PanicOnError(err)

	return player1Position, player2Position
}

func part1(inputList []string) string {
	p1p, p2p := parseInput(inputList)

	players := [2]player{newPlayer(p1p), newPlayer(p2p)}

	die := newDeterministicDie(100)

	currentPlayer := 0
	for {
		count := die.roll()
		count += die.roll()
		count += die.roll()

		score := players[currentPlayer].move(count)

		currentPlayer = (currentPlayer + 1) % 2

		if score >= 1000 {
			break
		}

	}

	return strconv.Itoa(die.numberOfRolls * players[currentPlayer].score)
}

/*
	Precalculate the different roll sets and the die totals for each universe.

	111  3
	112  4
	113  5

	121  4
	122  5
	123  6

	131  5
	132  6
	133  7

	211  4
	212  5
	213  6

	221  5
	222  6
	223  7

	231  6
	232  7
	233  8

	311  5
	312  6
	313  7

	321  6
	322  7
	323  8

	331  7
	332  8
	333  9

	There are 27 universes are created each roll

	These are only 7 possible die throw set totals based on the above list.
*/
var quantumDieRolls = [...]struct {
	dieTotal          int
	numberOfUniverses int
}{
	{3, 1},
	{4, 3},
	{5, 6},
	{6, 7},
	{7, 6},
	{8, 3},
	{9, 1},
}

type workType struct {
	players       [2]player
	numUniverses  int
	currentPlayer int // zero based
}

func part2(inputList []string) string {
	p1p, p2p := parseInput(inputList)

	// Prime the work queue.
	workQueue := list.New()
	work := workType{
		players:      [2]player{newPlayer(p1p), newPlayer(p2p)},
		numUniverses: 1,
	}
	workQueue.PushBack(work)

	// Keep track of each players total wins.
	var wins [2]int

	for workQueue.Len() > 0 {
		work := workQueue.Remove(workQueue.Front()).(workType)

		for _, roll := range quantumDieRolls {
			newWork := work

			score := newWork.players[newWork.currentPlayer].move(roll.dieTotal)
			newWork.currentPlayer = (newWork.currentPlayer + 1) % 2

			if score >= 21 {
				// Add wins
				wins[work.currentPlayer] += work.numUniverses * roll.numberOfUniverses

			} else {
				newWork.numUniverses *= roll.numberOfUniverses
				workQueue.PushBack(newWork)
			}
		}

	}

	result := wins[0]
	if wins[1] > wins[0] {
		result = wins[1]
	}

	return strconv.Itoa(result)
}

func main() {
	aoc.MainFunc(part1, part2)
}
