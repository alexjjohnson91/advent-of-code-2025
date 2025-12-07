package main

import (
	"fmt"

	"aoc/days/day03"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	input := util.ReadFile("days/day03/input.txt")

	fmt.Println("Part 1:", day03.Part1(input))
	fmt.Println("Part 2:", day03.Part2(input))
}
