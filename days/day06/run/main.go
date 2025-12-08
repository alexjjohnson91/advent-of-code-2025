package main

import (
	"fmt"

	"aoc/days/day06"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	input := util.ReadFile("days/day06/input.txt")

	fmt.Println("Part 1:", day06.Part1(input))
	fmt.Println("Part 2:", day06.Part2(input))
}
