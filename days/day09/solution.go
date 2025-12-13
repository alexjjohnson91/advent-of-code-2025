package day09

import (
	"aoc/internal/util"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

func Part1(input string) int {
	points := parsePoints(util.ReadArray(input))
	log.Println(points)

	var maxArea float64
	for i := range points {
		for j := i+1; j < len(points); j++ {
			curr := (math.Abs(points[i].x - points[j].x) + 1) * (math.Abs(points[i].y - points[j].y) + 1)
			if curr > maxArea {
				maxArea = curr
			}
		}
	}

	return int(maxArea)
}

func Part2(input string) int {
	// TODO
	return 0
}

func parsePoints(input []string) []Point {
	var points []Point

	for _, s := range input {
		nums := strings.Split(s, ",")
		x, _ := strconv.ParseFloat(nums[0], 64)
		y, _ := strconv.ParseFloat(nums[1], 64)
		points = append(points, Point{x, y})
	}

	return points
}
