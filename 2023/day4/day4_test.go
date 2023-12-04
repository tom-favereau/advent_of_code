package main

import (
	"testing"
)

func TestExemplePart1(t *testing.T) {
	result := part1("input2.txt")
	expected := 13

	if result != expected {
		t.Errorf("wrong answer %d, %d", result, expected)
	}
}

func TestExemplePart2(t *testing.T) {
	result := part2("input2.txt")
	expected := 30

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1("input.txt")
	expected := 24848

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2("input.txt")
	expected := 7258152

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}
