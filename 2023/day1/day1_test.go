package main

import (
	"testing"
)

func TestExemplePart1(t *testing.T) {
	result := part1("input2.txt")
	expected := 142

	if result != expected {
		t.Errorf("wrong answer %d, %d", result, expected)
	}
}

func TestExemplePart2(t *testing.T) {
	result := part2("input2.txt")
	expected := 142

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
