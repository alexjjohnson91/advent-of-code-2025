package main

import (
	"fmt"

	"aoc/days/day04"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	input := util.ReadFile("days/day04/input.txt")

	fmt.Println("Part 1:", day04.Part1(input))
	fmt.Println("Part 2:", day04.Part2(input))
}
