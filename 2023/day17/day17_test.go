package main

import (
	"testing"
)

func TestExemplePart1(t *testing.T) {
	result := part1("input2.txt")
	expected := 102

	if result != expected {
		t.Errorf("wrong answer %d, %d", result, expected)
	}
}

func TestExemplePart2_1(t *testing.T) {
	result := part2("input2.txt")
	expected := 94

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestExemplePart2_2(t *testing.T) {
	result := part2("input3.txt")
	expected := 71

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1("input.txt")
	expected := 771

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2("input.txt")
	expected := 930

	if result != expected {
		t.Errorf("wrong answer, %d, %d", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1("input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2("input.txt")
	}
}
