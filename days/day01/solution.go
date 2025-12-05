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
	}

	return numberOfZeroPoints
}

func Part2(input string) int {
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

		// cycling the dial
		for moveNumber > 99 {
			moveNumber -= 100
			numberOfZeroPoints++
			log.Printf("cycling dial, moveNumber = %v, incrementing counter: %v", moveNumber, numberOfZeroPoints)
		}

		if direction == "L" {
			dial -= moveNumber
		} else {
			dial += moveNumber
		}

		// check bounds
		if dial < 0 || dial > 99 {
			// can only be once at this point
			numberOfZeroPoints++
			log.Printf("cyling past 0, incrementing counter %v", numberOfZeroPoints)
		}

		// wrapped mod
		dial = ((dial % 100) + 100) % 100

		log.Printf("dial: %v\n", dial)
	}

	return numberOfZeroPoints
}

func readArray(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
