package day02

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input string) int {
	inputRanges := parseRanges(input)
	summedInvalidIds := 0

	for start, end := range inputRanges {
		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			start, end := splitInHalf(id)

			if start == end {
				log.Printf("invalid id: %v\n", id)
				summedInvalidIds += i
			} 
		}
	}

	return summedInvalidIds
}

func Part2(input string) int {
	// TODO
	return 0
}

func parseRanges(input string) map[int]int {
	tokens := strings.Split(input, ",")

	ranges := make(map[int]int)

	for _, t := range tokens {
		t = strings.TrimSpace(t)
		parts := strings.Split(t, "-")

		start, startErr := strconv.Atoi(parts[0])
		end, endErr := strconv.Atoi(parts[1])

		if startErr != nil || endErr != nil {
			continue
		}

		ranges[start] = end
	}

	return ranges
}

func splitInHalf(s string) (string, string) {
	r := []rune(s)
	mid := len(r) / 2
	return string(r[:mid]), string(r[mid:])
}
