package day07

import (
	"log"

	"aoc/internal/util"
)

var memo map[[2]int]int

func Part1(input string) int {
	manifold := util.ReadArray(input)
	manifold = startBeam(manifold)
	splits := moveBeam(manifold)

	return splits
}

func Part2(input string) int {
	manifold := util.ReadArray(input)

	memo = make(map[[2]int]int)
	paths := countPaths(manifold, 1, (len(manifold[0])-1) / 2)

	return paths
}

func printManifold(manifold []string) {
	for _, s := range manifold {
		log.Println(s)
	}
}

func countPaths(manifold []string, pathDepth int, pathIndex int) int {
	key := [2]int{pathDepth, pathIndex}
	if v, ok := memo[key]; ok {
		return v
	}

	// base case
	if (pathDepth == len(manifold) - 1) {
		return 1
	} 

	// recursive case
	if (manifold[pathDepth+1][pathIndex] == '^') {
		result := 0
		result = countPaths(manifold, pathDepth+1, pathIndex-1) + countPaths(manifold, pathDepth+1, pathIndex+1)
		memo[key] = result
		return result
	} else {
		return countPaths(manifold, pathDepth+1, pathIndex)
	}
}

func moveBeam(manifold []string) int {
	splits := 0
	beamLocation := 1
	for i := range manifold {
		s := manifold[i]
		for j := range s {
			if beamLocation == len(manifold) {
				log.Printf("beam Location ended\n")
				return splits
			}
			if s[j] == '|' {
				log.Println("beam needs to move")
				bytes := []byte(manifold[i+1])
				// check for split
				if bytes[j] == '^' {
					bytes[j-1] = '|'
					bytes[j+1] = '|'
					splits++
				} else {
					bytes[j] = '|'
				}
				manifold[i+1] = string(bytes)
				log.Printf("beam Location: %v\n", beamLocation)
			}
		}
		beamLocation++
		printManifold(manifold)
	}

	return splits
}

func startBeam(manifold []string) []string {
	s := manifold[0]
	for j := range s {
		if s[j] == 'S' {
			bytes := []byte(manifold[1])
			bytes[j] = '|'
			manifold[1] = string(bytes)
			return manifold
		}
	}

	return manifold
}
