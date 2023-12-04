package main

import (
	"testing"
)

func TestExemplePart1(t *testing.T) {
	result := part1("input2.txt")
	expected := 4361

	if result != expected {
		t.Errorf("wrong answer %d, %d", result, expected)
	}
}

func TestExemplePart2(t *testing.T) {
	result := part2("input2.txt")
	expected := 467835

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}
