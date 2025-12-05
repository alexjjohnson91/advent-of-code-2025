package day01

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input string) int {
	moves := readArray(input);
	dial := 50
	numberOfZeroPoints := 0

	for _, move := range moves {
		moveRunes := []rune(move)
		direction := string(moveRunes[0])
		numberString := string(moveRunes[1:])
		moveNumber, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatalf("the move number is not a number")
		}
		log.Printf("direction %v, move number %v\n", direction, moveNumber)

		if direction == "L" {
			dial -= moveNumber
		} else {
			dial += moveNumber
		}

		// wrapped mod
		dial = ((dial % 100) + 100) % 100

		if dial == 0 {
			numberOfZeroPoints++
		}

		log.Printf("dial: %v\n", dial)
	}

	return numberOfZeroPoints
}

func Part2(input string) int {
	// TODO
	return 0
}

func readArray(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
