package day01

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input string) int {
	moves := readArray(input);
	dial := 50

	for _, move := range moves {
		move_runes := []rune(move)
		direction := string(move_runes[0])
		number_string := string(move_runes[1:])
		move_number, err := strconv.Atoi(number_string)
		if err != nil {
			log.Fatalf("the move number is not a number")
		}
		log.Printf("direction %v, move number %v\n", direction, move_number)

		if direction == "L" {
			//TODO
		} else {
			//TODO
		}

		log.Printf("dial: %v\n", dial)
	}

	return 0
}

func Part2(input string) int {
	// TODO
	return 0
}

func readArray(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
