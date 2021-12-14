package main

import (
	"fmt"
	"os"
	"strings"
)

type couple struct {
	value  string
	isLast bool
}

func dup(couples map[couple]int, transforms map[string]string) map[couple]int {
	newCouples := make(map[couple]int)
	for k, v := range couples {
		transformed := transforms[k.value]
		newCouples[couple{value: string(k.value[0]) + transformed}] += v
		if k.isLast {
			newCouples[couple{value: transformed + string(k.value[1]), isLast: true}] = v
		} else {
			newCouples[couple{value: transformed + string(k.value[1])}] += v
		}
	}
	return newCouples
}

func printCount(counts map[string]int) {

	leastCommonCount := 0
	leastCommon := ""
	mostCommonCount := 0
	mostCommon := ""
	for k, v := range counts {
		if mostCommonCount == 0 {
			mostCommon = k
			leastCommon = k
			mostCommonCount = v
			leastCommonCount = v
		} else {
			if mostCommonCount < v {
				mostCommonCount = v
				mostCommon = k
			}
			if leastCommonCount > v {
				leastCommon = k
				leastCommonCount = v
			}
		}
	}
	fmt.Println(mostCommon, mostCommonCount, leastCommon, leastCommonCount, mostCommonCount-leastCommonCount)
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(contents), "\n")
	s := lines[0]
	transforms := make(map[string]string)
	for _, line := range lines[2:] {
		split := strings.Split(line, " -> ")
		transforms[split[0]] = split[1]
	}

	couples := make(map[couple]int)
	for i := 0; i < len(s)-2; i++ {
		couples[couple{value: s[i : i+2]}] += 1
	}
	couples[couple{value: s[len(s)-2:], isLast: true}] = 1

	step := 40
	for i := 0; i < step; i++ {
		couples = dup(couples, transforms)
	}

	counts := make(map[string]int)
	for k, v := range couples {
		counts[string(k.value[0])] += v
		if k.isLast {
			counts[string(k.value[1])] += v
		}
	}
	printCount(counts)
}
