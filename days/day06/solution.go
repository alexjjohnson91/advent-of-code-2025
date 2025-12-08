package day06

import (
	"strconv"
	"strings"

	"aoc/internal/util"
)

func Part1(input string) int {
	data := parseData(util.ReadArray(input))
	sum := 0
	answers := make([]int, 0, len(data))
	operands := data[len(data)-1]

	for i := len(data) - 1; i >= 0; i-- {
		s := data[i]
		for j := range s {
			if s[j] == "*" {
				answers = append(answers, 1)
			} else if s[j] == "+" {
				answers = append(answers, 0)
			} else {
				num, _ := strconv.Atoi(s[j])
				if operands[j] == "*" {
					answers[j] *= num
				} else if operands[j] == "+" {
					answers[j] += num
				}
			}
		}
	}

	for i := range answers {
		sum += answers[i]
	}

	return sum
}

func Part2(input string) int {
	// TODO
	return 0
}

func parseData(data []string) [][]string {
	numbers := [][]string{}

	for _, s := range data {
		numbers = append(numbers, strings.Fields(s))
	}

	return numbers
}
