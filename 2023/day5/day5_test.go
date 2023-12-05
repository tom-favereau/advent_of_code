package main

import (
	"testing"
)

func TestExemplePart1(t *testing.T) {
	result := part1("input2.txt")
	expected := 35

	if result != expected {
		t.Errorf("wrong answer %d, %d", result, expected)
	}
}

func TestExemplePart2(t *testing.T) {
	result := part2("input2.txt")
	expected := 46

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1("input.txt")
	expected := 26273516

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2("input.txt")
	expected := 34062605

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}
