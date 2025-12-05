package main

import (
	"fmt"

	"aoc/days/day02"
	"aoc/internal/util"
)

func main() {
	input := util.ReadFile("days/day02/input.txt")

	fmt.Println("Part 1:", day02.Part1(input))
	fmt.Println("Part 2:", day02.Part2(input))
}
