package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func partOne(contents string) {
	var lists = numLists2d(contents)
	// fmt.Println(lists)
	var ret = 0
	for _, t := range lists {
		if (isIncreasing(t) || isDecreasing(t)) && hasBoundedDiffs(t) {
			ret++
		}
	}
	fmt.Printf("part 1 total: %d\n", ret)
}

func main() {
	fmt.Println("Example:")
	partOne(example)
	// partTwo(example)
	data, _ := ioutil.ReadFile("day02.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	partOne(string(content))
	// partTwo(string(content))
}
