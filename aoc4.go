package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Add all the values in an []int.
func sum(numbers []int) (value int) {
	for _, number := range numbers {
		value += number
	}

	return value
}

// Finds the index of the needle in the haystack.
func indexOf(needle int, haystack []int) int {
	for idx, value := range haystack {
		if needle == value {
			return idx
		}
	}

	return -1
}

// Finds the index of the minimum value.
func min(numbers []int) (index int) {
	lowestValue := math.MaxInt
	for idx, value := range numbers {
		if value < lowestValue {
			index, lowestValue = idx, value
		}
	}

	return index
}

// Finds the index of the maximum value.
func max(numbers []int) (index int) {
	highestValue := math.MinInt
	for idx, value := range numbers {
		if value > highestValue {
			index, highestValue = idx, value
		}
	}

	return index
}

// Create the boards given the lines of numbers.
func makeBoards(scanner *bufio.Scanner) (boards [][][]int) {
	boards = make([][][]int, 0)
	board := make([][]int, 0)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			if len(board) > 0 {
				boards = append(boards, board)
				board = make([][]int, 0, len(board))
			}
			continue
		}

		numbers = make([]int, 0)
		for _, key := range strings.Split(line, " ") {
			if len(key) == 0 {
				continue
			}

			key, _ := strconv.Atoi(key)
			numbers = append(numbers, key)
		}

		board = append(board, numbers)
	}

	if len(board) > 0 {
		boards = append(boards, board)
	}

	return boards
}

// Finds the lowest index to win for each board.
func solve(numbers []int, boards [][][]int) (winningIndex []int) {
	winningIndex = make([]int, 0, len(boards))
	for _, board := range boards {
		lowestIdx := len(numbers)
		highestColIdx := make([]int, 0)

		// Get lowest index to win for row.
		for _, row := range board {
			highestIdx := -1

			for colIdx, value := range row {
				valIdx := indexOf(value, numbers)

				if valIdx > highestIdx {
					highestIdx = valIdx
				}

				if colIdx == len(highestColIdx) {
					highestColIdx = append(highestColIdx, valIdx)
				} else if valIdx > highestColIdx[colIdx] {
					highestColIdx[colIdx] = valIdx
				}
			}

			if highestIdx < lowestIdx {
				lowestIdx = highestIdx
			}
		}

		// Compare with lowest index to win for column.
		for _, colIdx := range highestColIdx {
			if colIdx < lowestIdx {
				lowestIdx = colIdx
			}
		}

		winningIndex = append(winningIndex, lowestIdx)
	}

	return winningIndex
}

// Finds unmarked numbers in a board.
func findUnmarked(board [][]int, marked []int) (unmarked []int) {
	unmarked = make([]int, 0)

	for _, rows := range board {
		for _, value := range rows {
			if indexOf(value, marked) < 0 {
				unmarked = append(unmarked, value)
			}
		}
	}

	return unmarked
}

func main() {
	input, _ := os.Open("var/aoc4_input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	numberStr := strings.Split(scanner.Text(), ",")
	numbers := make([]int, 0, len(numberStr))

	for _, numStr := range numberStr {
		numStr, _ := strconv.Atoi(numStr)
		numbers = append(numbers, numStr)
	}

	boards := makeBoards(scanner)
	results := solve(numbers, boards)

	firstIdx, lastIdx := min(results), max(results)
	firstWinningNum, lastWinningNum := numbers[results[firstIdx]], numbers[results[lastIdx]]
	firstBoardUnmarked := findUnmarked(boards[firstIdx], numbers[:results[firstIdx]+1])
	lastBoardUnmarked := findUnmarked(boards[lastIdx], numbers[:results[lastIdx]+1])
	firstBoardUnmarkedSum, lastBoardUnmarkedSum := sum(firstBoardUnmarked), sum(lastBoardUnmarked)

	fmt.Printf("Score at first board: %d * %d = %d\n", firstWinningNum, firstBoardUnmarkedSum, firstWinningNum*firstBoardUnmarkedSum)
	fmt.Printf("Score at last board: %d * %d = %d\n", lastWinningNum, lastBoardUnmarkedSum, lastWinningNum*lastBoardUnmarkedSum)
}
