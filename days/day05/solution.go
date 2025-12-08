package day05

import (
	"bufio"
	"sort"
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
	ranges = mergeRanges(ranges)

	totalFresh := 0

	for _, r := range ranges {
		totalFresh += (r[1] - r[0]) + 1
	}

	return totalFresh
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

func mergeRanges(spans [][2]int) [][2]int {
	sort.Slice(spans, func(i, j int) bool {
		if spans[i][0] == spans[j][0] {
			return spans[i][1] < spans[j][1]
		}
		return spans[i][0] < spans[j][0]
	})

	merged := make([][2]int, 0, len(spans))
	currSpan := spans[0]

	for _, span := range spans[1:] {
		if span[0] <= currSpan[1] {
			if span[1] > currSpan[1] {
				currSpan[1] = span[1]
			}
		} else {
			merged = append(merged, currSpan)
			currSpan = span
		}
	}

	merged = append(merged, currSpan)

	return merged
}
