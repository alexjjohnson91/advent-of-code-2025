package day03

import (
	"aoc/internal/util"
	"testing"
)

func TestPart1_Example(t *testing.T) {
	input := util.ReadFile("input_test.txt")
	got := Part1(input)
	want := 357
	if got != want {
		t.Fatalf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2_Example(t *testing.T) {
	input := util.ReadFile("input_test.txt")
	got := Part2(input)
	want := 3121910778619
	if got != want {
		t.Fatalf("Part2() = %v, want %v", got, want)
	}
}
