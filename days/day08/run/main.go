package main

import (
	"fmt"
	"io"
	"log"

	"aoc/days/day08"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	log.SetOutput(io.Discard)
	input := util.ReadFile("days/day08/input.txt")

	fmt.Println("Part 1:", day08.Part1(input))
	fmt.Println("Part 2:", day08.Part2(input))
}
