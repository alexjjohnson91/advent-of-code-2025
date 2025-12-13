package day08

import (
	"sort"
	"strconv"
	"strings"

	"aoc/internal/util"
)

type Point struct {
	x float64
	y float64
	z float64
}

type Edge struct {
	i        int
	j        int
	distance float64
}

var (
	parent []int
	rank   []int
)

func Part1(input string) int {
	points := parsePoints(input)
	edges := parseEdges(points)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	parent = make([]int, len(points))
	rank = make([]int, len(points))

	for i := range points {
		parent[i] = i
		rank[i] = 0
	}

	conn := 0
	for _, edge := range edges {
		if find(edge.i) != find(edge.j) {
			union(edge.i, edge.j)
		}

		conn++
		if conn == min(1000, len(edges)) {
			break
		}
	}

	sizes := make([]int, len(points))

	for i := range points {
		root := find(i)
		sizes[root]++
	}

	// var circuits []int
	circuits := make([]int, 0)
	for _, size := range sizes {
		if size > 0 {
			circuits = append(circuits, size)
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i] > circuits[j]
	})

	total := circuits[0] * circuits[1] * circuits[2]

	return total
}

func Part2(input string) int {
	// TODO
	return 0
}

func parsePoints(input string) []Point {
	var points []Point

	for _, s := range util.ReadArray(input) {
		coords := strings.Split(s, ",")
		x, _ := strconv.ParseFloat(coords[0], 64)
		y, _ := strconv.ParseFloat(coords[1], 64)
		z, _ := strconv.ParseFloat(coords[2], 64)

		newPoint := Point{x, y, z}

		points = append(points, newPoint)
	}

	return points
}

func parseEdges(points []Point) []Edge {
	var edges []Edge

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			distance := calculateDistance(points[i], points[j])
			edges = append(edges, Edge{i, j, distance})
		}
	}

	return edges
}

func calculateDistance(point1, point2 Point) float64 {
	dx := point1.x - point2.x
	dy := point1.y - point2.y
	dz := point1.z - point2.z
	return dx*dx + dy*dy + dz*dz
}

func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func union(x, y int) {
	rootX := find(x)
	rootY := find(y)

	if rootX == rootY {
		return
	}

	if rank[rootX] < rank[rootY] {
		parent[rootX] = rootY
	} else if rank[rootX] > rank[rootY] {
		parent[rootY] = rootX
	} else {
		parent[rootY] = rootX
		rank[rootX]++
	}
}
