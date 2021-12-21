package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type image struct {
	points                 map[point]int
	minX, maxX, minY, maxY int
	background             int
}

func (img *image) add(x, y, val int) {
	img.points[point{x: x, y: y}] = val
	if len(img.points) == 1 {
		img.maxX = x
		img.minX = x
		img.minY = y
		img.maxY = y
	} else {
		if x > img.maxX {
			img.maxX = x
		}
		if x < img.minX {
			img.minX = x
		}
		if y > img.maxY {
			img.maxY = y
		}
		if y < img.minY {
			img.minY = y
		}
	}
}

func (img image) String() string {
	r := make([]string, 0)
	for i := img.minY; i <= img.maxY; i++ {
		for j := img.minX; j <= img.maxX; j++ {
			val := img.valueAt(j, i)
			if val == 1 {
				r = append(r, "#")
			} else {
				r = append(r, ".")
			}
		}
		r = append(r, "\n")
	}
	return strings.Join(r, "")
}

func (img image) enhance(algorithm []int) image {
	newImg := image{points: make(map[point]int)}
	if img.background == 0 {
		newImg.background = algorithm[0]
	} else {
		newImg.background = algorithm[511]
	}

	for i := img.minY - 1; i <= img.maxY+1; i++ {
		for j := img.minX - 1; j <= img.maxX+1; j++ {
			index := img.getTransformAt(j, i)
			if algorithm[index] == 1 {
				newImg.add(j, i, 1)
			}
		}
	}

	return newImg
}

func (img image) valueAt(x, y int) int {
	if x >= img.minX && x <= img.maxX && y >= img.minY && y <= img.maxY {
		return img.points[point{x: x, y: y}]
	}
	return img.background
}

func (img image) getTransformAt(x, y int) int {
	res := 0
	n := 8
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			val := img.valueAt(x+i, y+j)
			res |= (val << n)
			n--
		}
	}
	return res
}

func (img image) countLit() int {
	return len(img.points)
}

func newImage(imageVals []string) image {
	img := image{points: make(map[point]int)}

	for i := 0; i < len(imageVals); i++ {
		for j := 0; j < len(imageVals[i]); j++ {
			if imageVals[i][j] == '#' {
				img.add(j, i, 1)
			}
		}
	}
	return img
}

func parseAlgorithm(in string) []int {
	algorithm := make([]int, len(in))
	for i, c := range in {
		if c == '#' {
			algorithm[i] = 1
		} else {
			algorithm[i] = 0
		}
	}
	return algorithm
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	vals := strings.Split(string(contents), "\n")
	algorithm := parseAlgorithm(vals[0])
	img := newImage(vals[2:])
	for i := 0; i < 50; i++ {
		img = img.enhance(algorithm)
	}
	fmt.Println(img.countLit())
}
