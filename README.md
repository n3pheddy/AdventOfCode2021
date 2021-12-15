# AdventOfCode2021
Solutions for Advent of Code 2021 https://adventofcode.com/2021. My goal this year is to familiarize
myself with Go, Ruby, and Rust, so these would be the languages I would use this time. I optimize
my code for readability and modularity with documentation if possible.

1. go run aoc1.go  
Number of increasing (window size=1): 1446  
Number of increasing (window size=3): 1486

2. go run aoc2.go  
(Part 1) Horizontal position * depth = 1762050  
(Part 2) Horizontal position * depth = 1855892637

3. go run aoc3.go  
Submarine Power Consumption=3320834  
Submarine Life Support Rating=4481199

4. go run aoc4.go  
Score at first board: 99 * 899 = 89001  
Score at last board: 38 * 192 = 7296

Took me some time to do this. At first I created a reverse index to look up the numbers efficiently.
When I get to the second part, I realized I had to solve for all boards and (almost) all numbers
anyway, so I implemented the `solve()` algorithm to find the minimum index needed to win for each board.