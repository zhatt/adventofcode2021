package main

import (
	"fmt"
	"strconv"
	"strings"
	"zhatt/aoc2021/aoc"
)

const board_size = 5

type square struct {
	value  int
	called bool
}

// board[column][row] Upper left square is 0,0
type board [board_size][board_size]square

func parseInput(inputList []string) ([]int, []*board) {

	drawNumbers := make([]int, 0)
	boards := make([]*board, 0)

	lineNum := 0
	var thisBoard *board = &board{}
	for line, strVal := range inputList {
		if line == 0 {
			for _, numStr := range strings.Split(strVal, ",") {
				number, err := strconv.Atoi(numStr)
				aoc.PanicOnError(err)
				drawNumbers = append(drawNumbers, number)
			}
			continue
		}

		if strVal == "" {
			continue
		}

		for index, valStr := range strings.Fields(strVal) {
			value, err := strconv.Atoi(valStr)
			aoc.PanicOnError(err)
			thisBoard[index][lineNum] = square{value: value}
		}

		lineNum++

		if lineNum == board_size {
			boards = append(boards, thisBoard)
			lineNum = 0
			thisBoard = &board{}
		}

	}
	return drawNumbers, boards
}

func callNumber(numberCalled int, board *board) {
	for row := 0; row < board_size; row++ {
		for column := 0; column < board_size; column++ {
			if board[row][column].value == numberCalled {
				board[row][column].called = true
			}
		}
	}
}

func calculateWinningScore(winningNumber int, board *board) int {
	sum := 0
	for column := 0; column < board_size; column++ {
		for row := 0; row < board_size; row++ {
			if !board[column][row].called {
				sum += board[column][row].value
			}
		}
	}
	return sum * winningNumber
}

func isWinner(board *board) bool {
	for column := 0; column < board_size; column++ {
		numCalled := 0
		for row := 0; row < board_size; row++ {
			if board[column][row].called {
				numCalled++
			}
		}
		if numCalled == board_size {
			return true
		}
	}

	for row := 0; row < board_size; row++ {
		numCalled := 0
		for column := 0; column < board_size; column++ {
			if board[column][row].called {
				numCalled++
			}
		}
		if numCalled == board_size {
			return true
		}
	}

	return false
}

func findWinningBoard(whichWinner int, draw []int, boards []*board) (int, *board) {
	winning_boards := make(map[int]bool)

	for _, numberToCall := range draw {
		for boardNum, board := range boards {
			if !winning_boards[boardNum] {
				callNumber(numberToCall, board)
				isWinner := isWinner(board)
				if isWinner {
					winning_boards[boardNum] = true
				}
			}

			if len(winning_boards) == whichWinner {
				return numberToCall, board
			}
		}
	}
	aoc.PanicOnError(fmt.Errorf("no winner"))
	return 0, nil
}

func part1(inputList []string) string {
	draw, boards := parseInput(inputList)
	winningNumber, winningBoard := findWinningBoard(1, draw, boards)
	score := calculateWinningScore(winningNumber, winningBoard)
	return strconv.Itoa(score)
}

func part2(inputList []string) string {
	draw, boards := parseInput(inputList)
	winningNumber, winningBoard := findWinningBoard(len(boards), draw, boards)
	score := calculateWinningScore(winningNumber, winningBoard)
	return strconv.Itoa(score)
}

func main() {
	aoc.MainFunc(part1, part2)
}
