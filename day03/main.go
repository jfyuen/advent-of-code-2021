package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func countZeroOnes(vals []string) ([]int, []int) {
	size := len(vals[0])
	ones := make([]int, size)
	zeros := make([]int, size)
	for _, val := range vals {
		for i, c := range val {
			if c == '1' {
				ones[i] += 1
			} else {
				zeros[i] += 1
			}
		}
	}
	return ones, zeros
}

func filterValues(vals []string, c string, pos int) []string {
	res := make([]string, 0)
	for _, v := range vals {
		if string(v[pos]) == c {
			res = append(res, v)
		}
	}
	return res
}

func filterEmpty(s []string) []string {
	var r []string
	for _, s := range s {
		if strings.TrimSpace(s) != "" {
			r = append(r, s)
		}
	}
	return r
}

func toInt(s string) int {
	r := 0
	for i, c := range s {
		if string(c) == "1" {
			r += int(math.Pow(2, float64(len(s)-i-1)))
		}
	}
	return r
}

func selectStr(vals []string, mostCommon bool) string {
	filtered := vals
	for i := 0; i < len(vals[0]); i++ {
		ones, zeros := countZeroOnes(filtered)
		check := ones[i] >= zeros[i]
		if !mostCommon {
			check = !check
		}
		c := "1"
		if !check {
			c = "0"
		}
		filtered = filterValues(filtered, c, i)
		if len(filtered) == 1 {
			break
		}
	}
	return filtered[0]
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	vals := strings.Split(string(contents), "\n")
	vals = filterEmpty(vals)
	fmt.Println(toInt(selectStr(vals, true)) * toInt(selectStr(vals, false)))
}
