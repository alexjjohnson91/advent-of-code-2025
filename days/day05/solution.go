package day05

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func Part1(input string) int {
	ranges, numbers := parseInput(input)

	fresh := 0

	// check each number
	for _, num := range numbers {
		// check all the ranges
		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				fresh++
				break
			}
		}
	}

	return fresh
}

func Part2(input string) int {
	ranges, _ := parseInput(input)
	totalFresh := make(map[int]int)

	ranges := mergeRanges(ranges)

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if totalFresh[i] != 1 {
				totalFresh[i] = 1
			}
		}
	}

	return len(totalFresh)
}

func parseInput(input string) ([][2]int, []int) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var ranges [][2]int
	var numbers []int
	parsingRanges := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, [2]int{start, end})
		} else {
			n, _ := strconv.Atoi(line)
			numbers = append(numbers, n)
		}
	}

	return ranges, numbers
}

func mergeRanges(ranges [][2]int) [][2]int {

	return ranges
}
