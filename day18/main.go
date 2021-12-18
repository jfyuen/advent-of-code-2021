package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type number struct {
	value       int
	left, right *number
	parent      *number
}

func (n *number) getLevel() int {
	if n.parent == nil {
		return 0
	}
	return 1 + n.parent.getLevel()
}

func (n *number) getNearestLeft(maxLevel int) *number {
	if n.parent == nil || maxLevel < 0 {
		return nil
	}
	if n.parent.left != nil && n.parent.left != n {
		return n.parent.left
	}
	return n.parent.getNearestLeft(maxLevel - 1)
}

func (n *number) getNearestRight(maxLevel int) *number {
	if n.parent == nil || maxLevel < 0 {
		return nil
	}
	if n.parent.right != nil && n.parent.right != n {
		return n.parent.right
	}
	return n.parent.getNearestRight(maxLevel - 1)
}

func (n *number) getLeft() *number {
	if n.left != nil {
		return n.left.getLeft()
	}
	return n
}

func (n *number) getRight() *number {
	if n.right != nil {
		return n.right.getRight()
	}
	return n
}

func (n *number) explode() bool {
	if n.isSimplePair() && n.getLevel() >= 4 {
		right := n.getNearestRight(4)
		if right != nil {
			nearestRight := right.getLeft()
			nearestRight.value += n.right.value
		}
		left := n.getNearestLeft(4)
		if left != nil {
			nearestLeft := left.getRight()
			nearestLeft.value += n.left.value
		}
		n.right = nil
		n.left = nil
		n.value = 0
		return true
	}
	if n.left != nil && n.left.explode() {
		return true
	}
	if n.right != nil && n.right.explode() {
		return true
	}
	return false
}

func (n *number) isLeaf() bool {
	return n != nil && n.left == nil && n.right == nil
}

func (n *number) isSimplePair() bool {
	return n.left.isLeaf() && n.right.isLeaf()
}

func (n *number) split() bool {
	if n.left != nil && n.left.split() {
		return true
	}
	if n.right != nil && n.right.split() {
		return true
	}
	if n.isLeaf() && n.value >= 10 {
		n.left = &number{value: int(math.Floor(float64(n.value) / 2)), parent: n}
		n.right = &number{value: int(math.Ceil(float64(n.value) / 2)), parent: n}
		n.value = 0
		return true
	}
	return false
}

func (n *number) reduce() bool {
	if n.explode() {
		return true
	}
	if n.split() {
		return true
	}
	return false
}
func (n *number) reduceAll() {
	for n.reduce() {
	}
}

func (n number) String() string {
	if n.left == nil && n.right == nil {
		return strconv.Itoa(n.value)
	}
	return fmt.Sprintf("[%v,%v]", n.left, n.right)
}

func splitNumber(s string) (string, string) {
	if s[0] != '[' {
		panic(s + " has a value")
	}
	subPairsCount := 0
	midIndex := 0
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '[' {
			subPairsCount++
		} else if s[i] == ']' {
			subPairsCount--
		} else if s[i] == ',' && subPairsCount == 0 {
			midIndex = i
			break
		}
	}
	return s[1:midIndex], s[midIndex+1 : len(s)-1]
}

func (n *number) magnitude() int {
	if n.isLeaf() {
		return n.value
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func newNumber(s string) *number {
	r := number{}
	if s[0] == '[' {
		left, right := splitNumber(s)
		r.left = newNumber(left)
		r.left.parent = &r
		r.right = newNumber(right)
		r.right.parent = &r
	} else {
		val, _ := strconv.Atoi(s)
		r.value = val
	}
	return &r
}

func getMagnitude(s1, s2 string) int {
	r := fmt.Sprintf("[%v,%v]", s1, s2)
	n := newNumber(r)
	n.reduceAll()
	return n.magnitude()
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	vals := strings.Split(string(contents), "\n")

	maxMagnitude := 0
	for i := 0; i < len(vals)-1; i++ {
		for j := i; j < len(vals); j++ {
			max1 := getMagnitude(vals[i], vals[j])
			if max1 > maxMagnitude {
				maxMagnitude = max1
			}
			max2 := getMagnitude(vals[j], vals[i])
			if max2 > maxMagnitude {
				maxMagnitude = max2
			}
		}
	}

	fmt.Println(maxMagnitude)
}
