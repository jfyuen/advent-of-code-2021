package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type board struct {
	values [5][5]int
	marked [5][5]bool
}

func (b board) String() string {
	s := make([]string, 0)
	for i := 0; i < len(b.values); i++ {
		for j := 0; j < len(b.values[i]); j++ {
			s = append(s, fmt.Sprintf("%d", b.values[i][j]))
		}
		s = append(s, "\n")
	}
	return strings.Join(s, " ")
}

func (b board) hasWon() bool {
	for i := 0; i < len(b.values); i++ {
		rowOk := true
		colOk := true
		for j := 0; j < len(b.values[i]); j++ {
			rowOk = rowOk && b.marked[i][j]
			colOk = colOk && b.marked[j][i]
		}
		if rowOk || colOk {
			return true
		}
	}
	return false
}

func (b *board) draw(number int) {
	for i := 0; i < len(b.values); i++ {
		for j := 0; j < len(b.values[i]); j++ {
			if b.values[i][j] == number {
				b.marked[i][j] = true
			}
		}
	}
}

func (b board) score(lastDrawn int) int {
	score := 0
	for i := 0; i < len(b.values); i++ {
		for j := 0; j < len(b.values[i]); j++ {
			if !b.marked[i][j] {
				score += b.values[i][j]
			}
		}
	}
	return score * lastDrawn
}

func newBoard(lines []string) board {
	b := board{}
	space := regexp.MustCompile(`\s+`)
	for i, row := range lines {
		row = space.ReplaceAllString(strings.TrimSpace(row), " ")
		for j, col := range strings.Split(row, " ") {
			val, _ := strconv.Atoi(col)
			b.values[i][j] = val
		}
	}
	return b
}

func allWOn(boards []bool) bool {
	for _, b := range boards {
		if !b {
			return false
		}
	}
	return true
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(contents), "\n")
	draws := strings.Split(lines[0], ",")
	lines = lines[2:]
	boards := make([]board, 0)
	for i := 0; i < len(lines); i += 6 {
		b := newBoard(lines[i : i+5])
		boards = append(boards, b)
	}

	wonBoardsIndexes := make([]bool, len(boards))
	lastWonIndex := -1
	lastDrawn := -1
	for _, drawStr := range draws {
		if allWOn(wonBoardsIndexes) {
			break
		}
		draw, _ := strconv.Atoi(strings.TrimSpace(drawStr))
		for i := range boards {
			b := &boards[i]
			b.draw(draw)
			if b.hasWon() {
				wonBoardsIndexes[i] = true
				lastDrawn = draw
				if allWOn(wonBoardsIndexes) {
					lastWonIndex = i
					break
				}
			}
		}
	}
	fmt.Println(boards[lastWonIndex].score(lastDrawn))

	// vals = filterEmpty(vals)
	// fmt.Println(toInt(selectStr(vals, true)) * toInt(selectStr(vals, false)))
}
