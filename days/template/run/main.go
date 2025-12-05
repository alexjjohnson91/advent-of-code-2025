package main

import (
	"fmt"

	"aoc/days/template"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	input := util.ReadFile("days/template/input.txt")

	fmt.Println("Part 1:", template.Part1(input))
	fmt.Println("Part 2:", template.Part2(input))
}
