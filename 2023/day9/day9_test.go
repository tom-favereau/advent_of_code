package main

import (
	"testing"
)

func TestExemplePart2(t *testing.T) {
	res := part2("input2.txt")
	expected := 2

	if res != expected {
		t.Errorf("wrong answer, %d, %d", res, expected)
	}
}

func TestPart1(t *testing.T) {
	res := part1("input.txt")
	expected := 1953784198

	if res != expected {
		t.Errorf("wrong answer, %d, %d", res, expected)
	}
}

func TestPart2(t *testing.T) {
	res := part2("input.txt")
	expected := 957

	if res != expected {
		t.Errorf("wrong answer, %d, %d", res, expected)
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
