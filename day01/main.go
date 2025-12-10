package main

import (
	"fmt"
	"log"
	"os"

	"github.com/imalgrab/aoc25/utils"
)

func solvePart1(lines []string) int {
	// TODO: Implement part 1 solution
	return 0
}

func solvePart2(lines []string) int {
	// TODO: Implement part 2 solution
	return 0
}

func main() {
	// Default to reading input.txt from the same directory
	inputPath := "day01/input.txt"
	
	// Allow command line argument to specify input file
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	lines, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
