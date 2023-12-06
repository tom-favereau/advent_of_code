package main

import (
	"testing"
)

func TestExemplePart1(t *testing.T) {
	result := part1("input2.txt")
	expected := 288

	if result != expected {
		t.Errorf("wrong answer %d, %d", result, expected)
	}
}

func TestExemplePart2(t *testing.T) {
	result := part2("input2.txt")
	expected := 71503

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1("input.txt")
	expected := 2374848

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2("input.txt")
	expected := 39132886

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}
