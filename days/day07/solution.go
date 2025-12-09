package day07

import (
	"aoc/internal/util"
	"log"
)

func Part1(input string) int {
	manifold := util.ReadArray(input)

	printManifold(manifold)

	return 0
}

func Part2(input string) int {
	// TODO
	return 0
}

func printManifold(manifold []string) {
	for _, s := range manifold {
		log.Println(s)
	}
}

func moveBeam(manifold []string) {
	printManifold(manifold)

	for _, s := range manifold {
	}
}
