package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const input_file string = "day02.txt"

var example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func numLists2d(contents string) [][]int {
	var ret [][]int
	lines_ex := strings.Split(contents, "\n")
	for _, line := range lines_ex {
		var row []int
		vals := strings.Fields((line))
		for _, vs := range vals {
			vi, _ := strconv.Atoi(vs)
			row = append(row, vi)
		}
		ret = append(ret, row)
	}
	return ret
}

func isIncreasing(vals []int) bool {
	if len(vals) < 2 {
		return true
	}
	for i, v := range vals {
		if i == 0 {
			continue
		}
		if v <= vals[i-1] {
			return false
		}
	}
	return true
}

func isDecreasing(vals []int) bool {
	if len(vals) < 2 {
		return true
	}
	for i, v := range vals {
		if i == 0 {
			continue
		}
		if v >= vals[i-1] {
			return false
		}
	}
	return true
}

func hasBoundedDiffs(vals []int) bool {
	if len(vals) < 2 {
		return true
	}
	for i, v := range vals {
		if i == 0 {
			continue
		}
		var diff = v - vals[i-1]
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isValid(vals []int) bool {
	return (isIncreasing(vals) || isDecreasing(vals)) && hasBoundedDiffs(vals)
}

func partOne(contents string) {
	start := time.Now()
	var lists = numLists2d(contents)
	var ret = 0
	for _, t := range lists {
		if isValid(t) {
			ret++
		}
	}
	fmt.Printf("part 1 result: %d\n", ret)
	elapsed := time.Since(start)
	fmt.Println("part 1 time: ", elapsed)
}

// returns all cases of removing a single element from vals.
func removals(vals []int) [][]int {
	var ret [][]int
	for i := range vals {
		var row []int
		left := vals[:i]
		right := vals[i+1:]
		for _, v := range left {
			row = append(row, v)
		}
		for _, v := range right {
			row = append(row, v)
		}
		ret = append(ret, row)
	}
	return ret

}

func partTwo(contents string) {
	start := time.Now()
	var lists = numLists2d(contents)
	var ret = 0
	for _, t := range lists {
		if isValid(t) {
			ret++
		} else {
			rems := removals(t)
			for _, v := range rems {
				if isValid(v) {
					ret++
					break
				}
			}
		}
	}
	fmt.Printf("part 2 result: %d\n", ret)
	elapsed := time.Since(start)
	fmt.Println("part 2 time: ", elapsed)
}

func main() {
	fmt.Println("Example:")
	partOne(example)
	partTwo(example)
	data, _ := os.ReadFile(input_file)
	content := string(data)
	fmt.Println("\nFrom file:")
	partOne(string(content))
	partTwo(string(content))
}
