package main

import (
	"fmt"

	"aoc/days/day07"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	input := util.ReadFile("days/day07/input.txt")

	fmt.Println("Part 1:", day07.Part1(input))
	fmt.Println("Part 2:", day07.Part2(input))
}
