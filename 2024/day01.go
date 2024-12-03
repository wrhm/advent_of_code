package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const day01example = `3   4
4   3
2   5
1   3
3   9
3   3`

func numLists(contents string) ([]int, []int) {
	var list0 []int
	var list1 []int
	lines_ex := strings.Split(contents, "\n")
	for _, line := range lines_ex {
		vals := strings.Fields((line))
		vi0, _ := strconv.Atoi(vals[0])
		vi1, _ := strconv.Atoi(vals[1])
		list0 = append(list0, vi0)
		list1 = append(list1, vi1)
	}
	return list0, list1
}

func day01partOne(contents string) {
	start := time.Now()
	list0, list1 := numLists(contents)
	sort.Ints(list0)
	sort.Ints(list1)
	var total_diffs = 0
	for i, v := range list0 {
		total_diffs += int(math.Abs(float64(v - list1[i])))
	}
	LogPartOneResult(total_diffs, start)
}

func counts(nums []int) map[int]int {
	ret := make(map[int]int)
	for _, v := range nums {
		val, ok := ret[v]
		if ok {
			ret[v] = val + 1
		} else {
			ret[v] = 1
		}
	}
	return ret
}

func day01partTwo(contents string) {
	start := time.Now()
	list0, list1 := numLists(contents)
	right := counts(list1)
	ret := 0
	for _, v := range list0 {
		val, ok := right[v]
		if ok {
			ret += v * val
		}
	}
	LogPartTwoResult(ret, start)
}

func day01main() {
	start := time.Now()
	fmt.Println("Example:")
	day01partOne(day01example)
	day01partTwo(day01example)
	data, _ := os.ReadFile("inputs/day01.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day01partOne(content)
	day01partTwo(content)
	elapsed := time.Since(start)
	fmt.Println("total time: ", elapsed)
}
