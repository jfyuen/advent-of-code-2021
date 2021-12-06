package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(contents), "\n")[0]
	values := make([]int, 0)
	for _, s := range strings.Split(lines, ",") {
		v, _ := strconv.Atoi(s)
		values = append(values, v)
	}

	counts := make(map[int]int)
	for _, v := range values {
		counts[v] += 1
	}

	days := 256
	for i := 0; i < days; i++ {
		newCounts := make(map[int]int)
		for k, v := range counts {
			if k == 0 {
				newCounts[6] += v
				newCounts[8] += v
			} else {
				newCounts[k-1] += v
			}
		}
		counts = newCounts
	}

	count := 0
	for _, c := range counts {
		count += c
	}
	fmt.Println(count)
}
