package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns its lines as a slice of strings
func ReadLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// ReadIntegers reads a file and returns its lines as a slice of integers
func ReadIntegers(filepath string) ([]int, error) {
	lines, err := ReadLines(filepath)
	if err != nil {
		return nil, err
	}

	var numbers []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

// ReadFile reads entire file content as a string
func ReadFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// SplitByEmptyLines splits the content by empty lines
func SplitByEmptyLines(content string) []string {
	return strings.Split(strings.TrimSpace(content), "\n\n")
}

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
