package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Add all the values in an []int.
func sum(nums []int) (result int) {
	for _, num := range nums {
		result += num
	}

	return result
}

// Groups a list of numbers into sliding measurement windows of size [winsize], then
// returns the number of windows whose sum is greater than its previous.
// Time complexity:  O(n) where n = len(scanner)
// Space complexity: winsize
func increasingWithWindow(scanner *bufio.Scanner, winsize int) (increasing int) {
	lastNums := make([]int, 0)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lastNums = append(lastNums, num)

		if len(lastNums) > winsize {
			// Opportunity for memoizing here, but choosing not to for readability.
			prev, next := sum(lastNums[:winsize]), sum(lastNums[1:])

			if next > prev {
				increasing++
			}

			lastNums = lastNums[1:]
		}
	}

	return increasing
}

func main() {
	input, _ := os.Open("var/aoc1_input.txt")
	defer input.Close()

	oneIncreasing := increasingWithWindow(bufio.NewScanner(input), 1)
	input.Seek(0, io.SeekStart)
	threeIncreasing := increasingWithWindow(bufio.NewScanner(input), 3)

	fmt.Printf("Number of increasing (window size=1): %d\n", oneIncreasing)
	fmt.Printf("Number of increasing (window size=3): %d\n", threeIncreasing)
}
