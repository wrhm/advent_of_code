package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const day11example = `125 17`

func makeIntCounter() map[int]int {
	return make(map[int]int)
}

func incrementOrInsertInIntCounter(cts *(map[int]int), x int) {
	prev, exist := (*cts)[x]
	if exist {
		(*cts)[x] = prev + 1
	} else {
		(*cts)[x] = 1
	}
}

func decrementInIntCounter(cts *(map[int]int), x int) {
	prev, exist := (*cts)[x]
	if exist {
		(*cts)[x] = prev - 1
	}
	// delete key if count is 0?
}

func splitInHalfByDigits(n int) (int, int) {
	s := strconv.Itoa(n)
	nc := len(s) / 2
	// fmt.Println(s, nc, s[:nc], s[nc:])
	a, _ := strconv.Atoi(s[:nc])
	b, _ := strconv.Atoi(s[nc:])
	return a, b
}

func evolveStone(n int) []int {
	if n == 0 {
		return []int{1}
	}
	s := strconv.Itoa(n)
	if len(s)%2 == 0 {
		a, b := splitInHalfByDigits(n)
		return []int{a, b}
	}
	return []int{n * 2024}
}

func day11partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	// cts := makeIntCounter()
	// incrementOrInsertInIntCounter(&cts, 1)
	// incrementOrInsertInIntCounter(&cts, 0)
	// incrementOrInsertInIntCounter(&cts, 1)
	// incrementOrInsertInIntCounter(&cts, 1)
	// fmt.Println(cts)
	// decrementInIntCounter(&cts, 1)
	// fmt.Println(cts)
	// decrementInIntCounter(&cts, 0)
	// fmt.Println(cts)
	// a, b := splitInHalfByDigits(1234)
	// fmt.Println(a, b)
	// fmt.Println(0, evolveStone(0))
	// fmt.Println(1, evolveStone(1))
	// fmt.Println(10, evolveStone(10))
	// fmt.Println(99, evolveStone(99))
	// fmt.Println(999, evolveStone(999))
	cts := makeIntCounter()
	initial := parseAllNums(lines[0])
	for _, v := range initial {
		incrementOrInsertInIntCounter(&cts, v)
	}
	fmt.Println("initially")
	fmt.Println(cts)
	for i := 0; i < 25; i++ { //6
		next_cts := makeIntCounter()
		for k := range cts {
			ev := evolveStone(k)
			for _, e := range ev {
				for m := 0; m < cts[k]; m++ {
					// fmt.Println("inserting", e)
					incrementOrInsertInIntCounter(&next_cts, e)
				}
			}
		}
		cts = next_cts
		fmt.Println("after", i+1, "blinks")
		fmt.Println(cts)
	}
	var ret = 0
	for _, v := range cts {
		ret += v
	}
	LogPartOneResult(ret, start)
}

func day11partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day11main() {
	start := time.Now()
	fmt.Println("Example:")
	day11partOne(day11example)
	day11partTwo(day11example)
	data, _ := os.ReadFile("inputs/day11.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day11partOne(content)
	day11partTwo(content)
	LogTimingForDay(start)
}
