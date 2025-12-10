# Advent of Code 2025 - Go Solutions

This repository contains solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Project Structure

```
aoc25/
├── day01/           # Day 1 solution
│   ├── main.go      # Solution code
│   └── input.txt    # Puzzle input (gitignored)
├── day02/           # Day 2 solution
│   └── ...
├── utils/           # Common utility functions
│   └── input.go     # File reading helpers
└── README.md
```

## Setup

### Prerequisites

- Go 1.24 or later

### Installation

Clone the repository:

```bash
git clone https://github.com/imalgrab/aoc25.git
cd aoc25
```

## Running Solutions

Each day has its own directory with a `main.go` file. To run a solution:

```bash
# Run day 1
go run day01/main.go

# Or specify a custom input file
go run day01/main.go path/to/input.txt

# Build a specific day's solution
go build -o bin/day01 day01/main.go
./bin/day01
```

## Adding a New Day

1. Create a new directory for the day (e.g., `day02`)
2. Copy the template from `day01/main.go`
3. Add your puzzle input to `dayXX/input.txt`
4. Implement the solution in the `solvePart1` and `solvePart2` functions

Example template:

```go
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
	inputPath := "dayXX/input.txt"
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
```

## Utility Functions

The `utils` package provides common helper functions:

- `ReadLines(filepath)` - Read file as slice of strings
- `ReadIntegers(filepath)` - Read file as slice of integers
- `ReadFile(filepath)` - Read entire file as string
- `SplitByEmptyLines(content)` - Split content by empty lines
- `Abs(x)`, `Min(a, b)`, `Max(a, b)` - Math utilities

## Testing

Run tests for all packages:

```bash
go test ./...
```

Run tests for a specific package:

```bash
go test ./utils
```

## Notes

- Input files (`input.txt`) are gitignored as per Advent of Code guidelines
- Each day's solution is independent and can be run separately
- Solutions are designed to read from `dayXX/input.txt` by default