package day02

import (
	"log"
	"strconv"
	"strings"
)

type InputRange struct {
	Start 	int
	End 		int
}

func Part1(input string) int {

	inputRanges := parseRanges(input)

	for _, r := range inputRanges {
		log.Printf("Range %v\n", r)
		for i := r.Start; i <= r.End; i++ {
			log.Println(i)
		}
	}


	return 0
}

func Part2(input string) int {
	// TODO
	return 0
}

func parseRanges(input string) []InputRange {
	tokens := strings.Split(input, ",")

	ranges := make([]InputRange, 0, len(tokens))

	for _, t := range tokens {
		parts := strings.Split(t, "-")

		start, startErr := strconv.Atoi(parts[0])
		end, endErr := strconv.Atoi(parts[1])

		if startErr != nil || endErr != nil {
			continue
		}

		ranges = append(ranges, InputRange{Start: start, End: end})
	}

	return ranges
}
