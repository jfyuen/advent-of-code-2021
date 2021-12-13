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

func foldHorizontal(points map[point]struct{}, n int) map[point]struct{} {
	newPoints := make(map[point]struct{})
	for p, _ := range points {
		if p.x < n {
			newPoints[p] = struct{}{}
		} else if p.x > n {
			x := p.x - n - 1
			newPoints[point{x: n - x - 1, y: p.y}] = struct{}{}
		}
	}
	return newPoints
}

func foldVertical(points map[point]struct{}, n int) map[point]struct{} {
	newPoints := make(map[point]struct{})
	for p, _ := range points {
		if p.y < n {
			newPoints[p] = struct{}{}
		} else if p.y > n {
			y := p.y - n - 1
			newPoints[point{x: p.x, y: n - y - 1}] = struct{}{}
		}
	}
	return newPoints
}

func printMap(points map[point]struct{}) {
	var maxX, maxY int
	for p, _ := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	fmt.Println(maxX, maxY)
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if _, ok := points[point{x: j, y: i}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(contents), "\n")
	folds := make([]string, 0)
	points := make(map[point]struct{})
	var maxX, maxY int
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			folds = append(folds, line)
			continue
		}

		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		points[point{x: x, y: y}] = struct{}{}

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	for _, fold := range folds {
		fold = strings.Trim(fold, "fold along ")
		split := strings.Split(fold, "=")
		if split[0] == "y" {
			val, _ := strconv.Atoi(split[1])
			points = foldVertical(points, val)
		} else if split[0] == "x" {
			val, _ := strconv.Atoi(split[1])
			points = foldHorizontal(points, val)
		}
	}
	printMap(points)
}
