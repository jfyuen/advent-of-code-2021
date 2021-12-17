package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type probe struct {
	x, y                 int
	xVelocity, yVelocity int
	maxY                 int
}

func (p *probe) step() {
	p.x += p.xVelocity
	p.y += p.yVelocity
	if p.y > p.maxY {
		p.maxY = p.y
	}
	p.yVelocity -= 1
	if p.xVelocity > 0 {
		p.xVelocity -= 1
	} else if p.xVelocity < 0 {
		p.xVelocity += 1
	}
}

type area struct {
	x0, x1, y0, y1 int
}

func (a area) isIn(b probe) bool {
	return a.x0 <= b.x && b.x <= a.x1 && a.y0 >= b.y && b.y >= a.y1
}

func (a area) isOut(b probe) bool {
	return b.x > a.x1 || b.y < a.y1
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	vals := strings.Split(string(contents), ",")
	xs := strings.Split(strings.Split(vals[0], "=")[1], "..")
	ys := strings.Split(strings.Split(vals[1], "=")[1], "..")
	x0, _ := strconv.Atoi(xs[0])
	x1, _ := strconv.Atoi(xs[1])
	y1, _ := strconv.Atoi(ys[0])
	y0, _ := strconv.Atoi(ys[1])
	a := area{x0: x0, x1: x1, y0: y0, y1: y1}

	probeCounts := 0
	for x := 0; x <= a.x1; x++ {
		for y := a.y1; y <= -a.y1*2; y++ {
			p := probe{xVelocity: x, yVelocity: y}
			for {
				p.step()
				if a.isIn(p) {
					probeCounts++
					break
				} else if a.isOut(p) {
					break
				}
			}
		}
	}

	fmt.Printf("%v\n", probeCounts)
}
