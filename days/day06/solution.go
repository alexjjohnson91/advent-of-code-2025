package day06

import (
	"log"
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
	data := util.ReadArray(input)

	problemSet := []int{}

	// this is the counter we will use to iterate through the strings
	dataIt := len(data[0]) - 1

	sum := 0
	// loop until our iterator makes it to the start of the arrays
	for dataIt >= 0 {
		// then we continuously loop through the data set until we have the stuff
		chars := make([]rune, 0, len(data))
		for _, s := range data {
			if dataIt >= len(s) {
				continue
			}
			chars = append(chars, rune(s[dataIt]))
		}
		if string(chars[len(chars) - 1]) == "+" || string(chars[len(chars) - 1]) == "*" {
			num, _ := strconv.Atoi(strings.TrimSpace(string(chars[:len(chars) - 1])))
			problemSet = append(problemSet, num)
			operand := string(chars[len(chars) - 1])
			log.Printf("problem %v\n", problemSet)
			log.Printf("operand %v\n", operand)
			sum += processProblem(problemSet, operand)
			problemSet = problemSet[:0]
			dataIt--
		} else {
			num, _ := strconv.Atoi(strings.TrimSpace(string(chars)))
			problemSet = append(problemSet, num)
			log.Printf("problem %v\n", problemSet)
		}
		dataIt--
	}

	return sum
}

func processProblem(problemSet []int, operand string) int {
	var answer int
	if operand == "*" {
		answer = 1
	} else {
		answer = 0
	}

	for _, num := range problemSet {
		if operand == "*" {
			answer *= num
		} else {
			answer += num 
		}
	}

	return answer
}

func parseData(data []string) [][]string {
	numbers := [][]string{}

	for _, s := range data {
		numbers = append(numbers, strings.Fields(s))
	}

	return numbers
}
