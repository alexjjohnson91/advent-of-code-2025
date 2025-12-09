package day07

import (
	"aoc/internal/util"
)

func Part1(input string) int {
	manifold := util.ReadArray(input)

	manifold = startBeam(manifold)
	splits := moveBeam(manifold)

	return splits
}

func Part2(input string) int {
	// TODO
	return 0
}

func moveBeam(manifold []string) int {
	splits := 0
	beamLocation := 1
	for i := range manifold {
		s := manifold[i]
		for j := range s {
			if beamLocation == len(manifold) {
				return splits
			}
			if s[j] == '|' {
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
			}
		}
		beamLocation++
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
