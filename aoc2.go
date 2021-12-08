package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Part 1: Find horizontal position and depth by treating "up" and "down" instructions
// as traversing the depth.
// Time complexity : O(n) where n = len(scanner)
// Space complexity: O(1)
func findSimplePosition(scanner *bufio.Scanner) (horizontalPosition int, depth int) {
	for scanner.Scan() {
		instruction := scanner.Text()
		tokens := strings.Fields(instruction)
		steps, _ := strconv.Atoi(tokens[1])

		switch tokens[0] {
		case "up":
			depth -= steps
		case "down":
			depth += steps
		default:
			horizontalPosition += steps
		}
	}

	return horizontalPosition, depth
}

// Part 2: Find horizontal position and depth by treating "up" and "down" as aims, and
// "forward" moves both horizontal position and depth.
// Time complexity : O(n) where n = len(scanner)
// Space complexity: O(1)
func findAimPosition(scanner *bufio.Scanner) (horizontalPosition int, depth int) {
	aim := 0

	for scanner.Scan() {
		instruction := scanner.Text()
		tokens := strings.Fields(instruction)
		steps, _ := strconv.Atoi(tokens[1])

		switch tokens[0] {
		case "up":
			aim -= steps
		case "down":
			aim += steps
		default:
			horizontalPosition += steps
			depth += steps * aim
		}
	}

	return horizontalPosition, depth
}

func main() {
	input, _ := os.Open("var/aoc2_input.txt")
	defer input.Close()

	horizontalPosition, depth := findSimplePosition(bufio.NewScanner(input))
	fmt.Printf("(Part 1) Horizontal position * depth = %d\n", horizontalPosition*depth)

	input.Seek(0, io.SeekStart)
	horizontalPosition, depth = findAimPosition(bufio.NewScanner(input))
	fmt.Printf("(Part 2) Horizontal position * depth = %d\n", horizontalPosition*depth)
}
