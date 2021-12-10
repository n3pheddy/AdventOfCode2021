package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Find rating by filtering majority values from each position. Pass collectMajorityValues=true
// for oxygen generator rating and collectMajorityValues=false for CO2 scrubbing rating.
// Time complexity : O(n*m) where n=len(lines) and m=len(lines[0])
// Space complexity: O(m) where m=len(lines[0])
func findRating(lines []string, collectMajorityValues bool) (rating int) {
	bits := make([]int, len(lines[0]))

	// Short-circuits if there's only 1 element.
	for idx := 0; idx < len(lines[0]) && len(lines) > 1; idx++ {
		for _, nums := range lines {
			if nums[idx] == '1' {
				bits[idx] += 1
			}
		}

		onesAreMajority := false
		if float64(bits[idx]) >= float64(len(lines))/2 {
			onesAreMajority = true
		}

		// Flatten the count to the desired value.
		if onesAreMajority == collectMajorityValues {
			bits[idx] = 1
		} else {
			bits[idx] = 0
		}

		truncate := 0

		for jdx := 0; jdx+truncate < len(lines); {
			if bits[idx] != int(lines[jdx][idx]-'0') {
				// Moves invalid values to the end.
				lines[jdx], lines[len(lines)-truncate-1] = lines[len(lines)-truncate-1], lines[jdx]
				truncate++
			} else {
				jdx++
			}
		}

		// Invalid values are at the end.
		lines = lines[:len(lines)-truncate]
	}

	// Find majority after filtering down to just one line.
	for idx, value := range lines[0] {
		if value == '1' {
			power := len(lines[0]) - idx - 1
			rating += int(math.Pow(2, float64(power)))
		}
	}

	return rating
}

// Finds both the gamma and epsilon by calculating the majority value at each position.
// Time complexity:  O(n+m) where n=len(lines) and m=len(lines[0])
// Space complexity: O(m) where m=len(lines[0])
func findPowerConsumption(lines []string) (gamma, epsilon float64) {
	bits := make([]int, len(lines[0]))

	for _, nums := range lines {
		for idx, char := range nums {
			if char == '1' {
				bits[idx] += 1
			}
		}
	}

	// Find majority after parsing all lines and sum them.
	for power := range bits {
		idx := len(bits) - power - 1

		if bits[idx] > len(lines)/2 {
			gamma += math.Pow(2, float64(power))
		} else {
			epsilon += math.Pow(2, float64(power))
		}
	}

	return gamma, epsilon
}

func main() {
	input, _ := os.Open("var/aoc3_input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	gamma, epsilon := findPowerConsumption(lines)
	fmt.Printf("Submarine Power Consumption=%d\n", int(gamma*epsilon))

	oxygen, co2 := findRating(lines, true), findRating(lines, false)
	fmt.Printf("Submarine Life Support Rating=%d\n", int(oxygen*co2))
}
