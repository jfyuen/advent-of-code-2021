package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type path struct {
	route        []string
	visited      map[string]int
	visitedTwice bool
}

func buildCaves(contents string) map[string][]string {
	lines := strings.Split(contents, "\n")
	caves := make(map[string][]string)
	for _, line := range lines {
		split := strings.Split(line, "-")
		caves[split[0]] = append(caves[split[0]], split[1])
		caves[split[1]] = append(caves[split[1]], split[0])
	}
	return caves
}

func isLower(s string) bool {
	for _, r := range s {
		return unicode.IsLower(r)
	}
	return false
}

func newPath() path {
	return path{visited: make(map[string]int), route: make([]string, 0)}
}

func (p *path) visit(node string) {
	p.route = append(p.route, node)
	if isLower(node) {
		p.visited[node] += 1
		if p.visited[node] == 2 {
			p.visitedTwice = true
		}
		if p.visited[node] > 2 {
			panic(p)
		}
	}
}

func (p *path) copy() *path {
	p2 := newPath()
	for k, v := range p.visited {
		p2.visited[k] = v
	}
	for _, v := range p.route {
		p2.route = append(p2.route, v)
	}
	p2.visitedTwice = p.visitedTwice
	return &p2
}

func (p *path) hasVisited(s string) bool {
	if s == "start" || s == "end" {
		if _, ok := p.visited[s]; ok {
			return true
		}
	}
	if _, ok := p.visited[s]; ok && p.visitedTwice {
		return true
	}
	return false
}

func explore(caves map[string][]string, p *path, start string, end string) []*path {
	paths := make([]*path, 0)
	p.visit(start)
	if start == end {
		paths = append(paths, p)
		return paths
	}
	for _, next := range caves[start] {
		if isLower(next) {
			if p.hasVisited(next) {
				continue
			}
		}
		for _, p := range explore(caves, p.copy(), next, end) {
			paths = append(paths, p)
		}
	}
	return paths
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	caves := buildCaves(string(contents))
	path := newPath()
	paths := explore(caves, &path, "start", "end")
	// for _, p := range paths {
	// 	fmt.Printf("%v\n", p.route)
	// }
	fmt.Println(len(paths))
}
