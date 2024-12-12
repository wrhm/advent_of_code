package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Logging

func LogPartOneResult(v int, start time.Time) {
	fmt.Printf("part 1 result: %d\n", v)
	elapsed := time.Since(start)
	fmt.Println("part 1 time: ", elapsed)
}

func LogPartTwoResult(v int, start time.Time) {
	fmt.Printf("part 2 result: %d\n", v)
	elapsed := time.Since(start)
	fmt.Println("part 2 time: ", elapsed)
}

func LogTimingForDay(elapsed time.Duration) {
	fmt.Println("total time for day: ", elapsed)
}

func printStrList(lines []string) {
	for _, row := range lines {
		fmt.Println(string(row))
	}
}

func print2dBytesList(lines [][]byte) {
	for _, row := range lines {
		fmt.Println(string(row))
	}
}

// Algorithms

func parseAllNums(s string) []int {
	dr, _ := regexp.Compile(`\d+`)
	num_strs := dr.FindAllString(s, -1)
	var nums []int
	for _, v := range num_strs {
		i, _ := strconv.Atoi(v)
		nums = append(nums, i)
	}
	return nums
}

func inBounds(x int, lo int, hi int) bool {
	return lo <= x && x <= hi
}

func gridHasByteAtPos(lines *([]string), r int, c int, b byte) int {
	w := len((*lines)[0])
	h := len(*lines)
	if inBounds(r, 0, h-1) && inBounds(c, 0, w-1) && (*lines)[r][c] == b {
		return 1
	}
	return 0
}

func strListAs2dBytes(lines []string) [][]byte {
	var ret [][]byte
	h := len(lines)
	for i := 0; i < h; i++ {
		ret = append(ret, []byte(lines[i]))
	}
	return ret
}

type GridPoint struct {
	r int
	c int
}

func GP(r int, c int) *GridPoint {
	return &GridPoint{r, c}
}

func make2dPointSet() map[string]bool {
	return make(map[string]bool)
}

func insertInto2dPointSet(ps *(map[string]bool), r int, c int) {
	pk := strconv.Itoa(r) + ":" + strconv.Itoa(c)
	(*ps)[pk] = true
}

func digitRuneAsInt(r rune) int {
	v, _ := strconv.Atoi(string(r))
	return v
}

func digitByteAsInt(b byte) int {
	v, _ := strconv.Atoi(string(b))
	return v
}
