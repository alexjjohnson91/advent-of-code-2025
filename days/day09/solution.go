package day09

import (
	"aoc/internal/util"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func Part1(input string) int {
	points := parsePoints(util.ReadArray(input))
	log.Println(points)

	var maxArea int
	for i := range points {
		for j := i+1; j < len(points); j++ {
			curr := (math.Abs(float64(points[i].x - points[j].x)) + 1) * (math.Abs(float64(points[i].y - points[j].y)) + 1)
			if int(curr) > maxArea {
				maxArea = int(curr)
			}
		}
	}

	return maxArea
}

func Part2(input string) int {
	points := parsePoints(util.ReadArray(input))
	log.Println(points)

	var maxArea int

	pointsMap := make(map[int][2]int)
	crossings := make(map[int][]int)
	pointsLen := len(points)

	for i := range points {
		a := points[i]
		b := points[(i + 1) % pointsLen]

		if horizontalPoints(a, b) {
			y := a.y
			crossings[y] = append(crossings[y], a.x, b.x)
			continue
		}

		if verticalPoints(a, b) {
			x := a.x

			for y := min(a.y, b.y); y < max(a.y, b.y); y++ {
				crossings[y] = append(crossings[y], x)
			}

			continue
		}
	}

	for y, cross := range crossings {
		slices.Sort(cross)

		minX := cross[0]
		maxX := cross[len(cross) - 1]

		pointsMap[y] = [2]int{minX, maxX}
	}

	log.Println(pointsMap)

	for i := range points {
		for j := i+1; j < len(points); j++ {
			curr := (math.Abs(float64(points[i].x - points[j].x)) + 1) * (math.Abs(float64(points[i].y - points[j].y)) + 1)
			if isContained(pointsMap, points[i], points[j]) && int(curr) > maxArea {
				log.Printf("points are contained: %v, %v\n", points[i], points[j])
				maxArea = int(curr)
			}
		}
	}

	return maxArea
}

func isContained(pointsMap map[int][2]int, a Point, b Point) bool {
	left  := min(a.x, b.x)
	right := max(a.x, b.x)
	top   := min(a.y, b.y)
	bot   := max(a.y, b.y)

	for y := top; y <= bot; y++ {
		if interval, ok := pointsMap[y]; !ok {
			return false
		} else if interval[0] > left || interval[1] < right {
			return false
		}
	}

	return true
}

func parsePoints(input []string) []Point {
	var points []Point

	for _, s := range input {
		nums := strings.Split(s, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		points = append(points, Point{x, y})
	}

	return points
}

func horizontalPoints(a, b Point) bool {
	return a.y == b.y
}

func verticalPoints(a, b Point) bool {
	return a.x == b.x
}
