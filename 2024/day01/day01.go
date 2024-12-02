package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var example = `3   4
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

func sumDiffs(contents string) {
	list0, list1 := numLists(contents)
	sort.Ints(list0)
	sort.Ints(list1)
	var total_diffs = 0
	for i, v := range list0 {
		total_diffs += int(math.Abs(float64(v - list1[i])))
	}
	fmt.Printf("part 1 total: %d\n", total_diffs)
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

func partTwo(contents string) {
	list0, list1 := numLists(contents)
	right := counts(list1)
	ret := 0
	for _, v := range list0 {
		val, ok := right[v]
		if ok {
			ret += v * val
		}
	}
	fmt.Printf("part 2 total: %d\n", ret)
}

func main() {
	fmt.Println("Example:")
	sumDiffs(example)
	partTwo(example)
	data, _ := os.ReadFile("day01.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	sumDiffs(string(content))
	partTwo(string(content))
}
