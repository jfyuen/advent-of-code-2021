package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	hp "container/heap"
)

//heap
type path struct {
	cost int
	node point
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

type point struct {
	x, y int
}

func dijkstra(matrix [][]int) int {
	h := newHeap()
	h.push(path{cost: 0, node: point{x: 0, y: 0}})
	dist := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dist[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 0

	for len(*h.values) > 0 {
		p := h.pop()
		node := p.node
		for _, next := range [4]point{{x: 0, y: 1}, {x: 1, y: 0}, {x: 0, y: -1}, {x: -1, y: 0}} {
			x := node.x + next.x
			y := node.y + next.y
			if x < 0 || y < 0 || x >= len(matrix[0]) || y >= len(matrix) {
				continue
			}
			if dist[y][x] == -1 || dist[y][x] > dist[node.y][node.x]+matrix[y][x] {
				dist[y][x] = dist[node.y][node.x] + matrix[y][x]
				newPath := path{cost: dist[y][x], node: point{x: x, y: y}}
				h.push(newPath)
			}
		}
	}
	return dist[len(dist)-1][len(dist[0])-1]
}

func makeBigMatrix(matrix [][]int) [][]int {
	r := make([][]int, len(matrix)*5)
	for i := 0; i < len(r); i++ {
		r[i] = make([]int, len(matrix[0])*5)
		for j := 0; j < len(matrix[0])*5; j++ {
			val := matrix[i%len(matrix)][j%len(matrix[0])]
			val += i/len(matrix) + j/len(matrix[0])
			if val > 9 {
				val %= 9
			}
			r[i][j] = val
		}
	}
	return r
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	matrix := make([][]int, 0)

	for _, line := range strings.Split(string(contents), "\n") {
		row := make([]int, 0)
		for _, c := range line {
			val, _ := strconv.Atoi(string(c))
			row = append(row, val)
		}
		matrix = append(matrix, row)
	}

	big := makeBigMatrix(matrix)
	fmt.Println(dijkstra(big))
}
