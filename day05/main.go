package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type coordinate struct {
	start, end point
}

func newCoordinate(l string) coordinate {
	split := strings.Split(l, " -> ")
	start := strings.Split(split[0], ",")
	end := strings.Split(split[1], ",")
	startX, _ := strconv.Atoi(start[0])
	startY, _ := strconv.Atoi(start[1])
	endX, _ := strconv.Atoi(end[0])
	endY, _ := strconv.Atoi(end[1])
	return coordinate{start: point{x: startX, y: startY}, end: point{x: endX, y: endY}}
}

func (c coordinate) getMinMax() (int, int) {
	var min, max int
	if c.start.x == c.end.x {
		if c.start.y > c.end.y {
			min = c.end.y
			max = c.start.y
		} else {
			min = c.start.y
			max = c.end.y
		}
	} else if c.start.x > c.end.x {
		min = c.end.x
		max = c.start.x
	} else {
		min = c.start.x
		max = c.end.x
	}
	return min, max
}

func (c coordinate) addPoints(points map[point]int, addDiagnonals bool) {
	if !addDiagnonals && c.start.x != c.end.x && c.start.y != c.end.y {
		return
	}
	var min, max = c.getMinMax()
	for i := 0; i <= max-min; i++ {
		x := c.start.x
		if c.start.x < c.end.x {
			x += i
		} else if c.start.x > c.end.x {
			x -= i
		}
		y := c.start.y
		if c.start.y < c.end.y {
			y += i
		} else if c.start.y > c.end.y {
			y -= i
		}
		points[point{x: x, y: y}] += 1
	}
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(contents), "\n")
	coords := make([]coordinate, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		coords = append(coords, newCoordinate(line))
	}

	addDiagonals := true
	points := make(map[point]int)
	for _, coord := range coords {
		coord.addPoints(points, addDiagonals)
	}

	overlap := 0
	for _, intersection := range points {
		if intersection >= 2 {
			overlap++
		}
	}
	fmt.Println(overlap)
}
