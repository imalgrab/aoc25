package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Create a temporary file for testing
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	
	content := "line1\nline2\nline3"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	lines, err := ReadLines(tmpFile)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}
	if len(lines) != len(expected) {
		t.Fatalf("Expected %d lines, got %d", len(expected), len(lines))
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

func TestReadIntegers(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	
	content := "1\n2\n3\n\n4"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	numbers, err := ReadIntegers(tmpFile)
	if err != nil {
		t.Fatalf("ReadIntegers failed: %v", err)
	}

	expected := []int{1, 2, 3, 4}
	if len(numbers) != len(expected) {
		t.Fatalf("Expected %d numbers, got %d", len(expected), len(numbers))
	}

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("Number %d: expected %d, got %d", i, expected[i], num)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-100, 100},
	}

	for _, tt := range tests {
		result := Abs(tt.input)
		if result != tt.expected {
			t.Errorf("Abs(%d) = %d; expected %d", tt.input, result, tt.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{5, 5, 5},
		{-1, -2, -2},
	}

	for _, tt := range tests {
		result := Min(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Min(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{5, 5, 5},
		{-1, -2, -1},
	}

	for _, tt := range tests {
		result := Max(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Max(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
