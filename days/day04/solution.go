package day04

import (
	"aoc/internal/util"
)

func Part1(input string) int {
	chart := toChart(util.ReadArray(input))
	rolls := 0

	for i := range chart {
		for j := range chart[i] {
			if string(chart[i][j]) == "@" && checkSpot(chart, i, j) {
				rolls++
			}
		}
	}

	return rolls
}

func Part2(input string) int {
	chart := toChart(util.ReadArray(input))
	rollsAccessed := 0
	lastCycleAccessedRolls := -1

	for lastCycleAccessedRolls != 0 {
		chart, lastCycleAccessedRolls = processChart(chart)
		rollsAccessed += lastCycleAccessedRolls
	}

	return rollsAccessed
}

func processChart(chart [][]rune) ([][]rune, int) {
	rolls := 0

	for i := range chart {
		for j := range chart[i] {
			if string(chart[i][j]) == "@" && checkSpot(chart, i, j) {
				chart[i][j] = 'x'
				rolls++
			}
		}
	}

	return chart, rolls
}

func checkSpot(chart [][]rune, i, j int) bool {

	checkSet := [][2]int{
		{i + 1, j + 1},
		{i + 1, j},
		{i + 1, j - 1},
		{i, j - 1},
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j + 1},
	}
	adjRolls := 0

	for i, set := range checkSet {
		if isOutOfBounds(set, len(chart), len(chart[i])) {
			continue
		}
		if string(chart[set[0]][set[1]]) == "@" {
			adjRolls++
		}
	}

	return adjRolls < 4
}

func toChart(lines []string) [][]rune {
	runes := make([][]rune, len(lines))
	for i, s := range lines {
		runes[i] = []rune(s)
	}
	return runes
}

func isOutOfBounds(set [2]int, x int, y int) bool {
	if set[0] < 0 || set[1] < 0 || set[0] >= x || set[1] >= y {
		return true
	}

	return false
}
