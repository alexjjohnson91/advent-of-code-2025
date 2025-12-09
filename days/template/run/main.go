package main

import (
	"fmt"
	"io"
	"log"

	"aoc/days/template"
	"aoc/internal/util"
)

func main() {
	// NOTE: when copied, imports get rewritten by script
	log.SetOutput(io.Discard)
	input := util.ReadFile("days/template/input.txt")

	fmt.Println("Part 1:", template.Part1(input))
	fmt.Println("Part 2:", template.Part2(input))
}
