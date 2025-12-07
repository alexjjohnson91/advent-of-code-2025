package day03

import (
	"aoc/internal/util"
	"log"
	"strconv"
)

func Part1(input string) int {
	banks := util.ReadArray(input)
	combinedJoltage := 0

	for _, bank := range banks {
		maxJoltage, err := strconv.Atoi(findMaxDigits([]rune{}, toIntArray(bank), 2)) 
		if err != nil {
			log.Fatalln("error finding max joltage")
		}

		combinedJoltage += maxJoltage 
	}

	return combinedJoltage
}

func Part2(input string) int {
	banks := util.ReadArray(input)
	combinedJoltage := 0

	for _, bank := range banks {
		maxJoltage, err := strconv.Atoi(findMaxDigits([]rune{}, toIntArray(bank), 12)) 
		if err != nil {
			log.Fatalln("error finding max joltage")
		}

		combinedJoltage += maxJoltage 
	}

	return combinedJoltage
}

func findMaxDigits(joltage []rune, digits []int, digitsRemaining int) string {
	if digitsRemaining == 0 {
		return string(joltage)
	}

	maxDigit, maxDigitIndex := -1, -1
	for i := 0; i < len(digits) - digitsRemaining + 1; i++ {
		if digits[i] > maxDigit {
			maxDigit = digits[i]
			maxDigitIndex = i
		}
	}

	return findMaxDigits(append(joltage, rune(maxDigit)), digits[maxDigitIndex+1:], digitsRemaining - 1)
}

func toIntArray(input string) []int {
	ints := []int{}

	for _, s := range input {
		ints = append(ints, int(s))
	}

	return ints
}
