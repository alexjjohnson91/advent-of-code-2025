package main

import (
	"fmt"

	"aoc/days/day01"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	input := util.ReadFile("days/day01/input.txt")

	fmt.Println("Part 1:", day01.Part1(input))
	fmt.Println("Part 2:", day01.Part2(input))
}
