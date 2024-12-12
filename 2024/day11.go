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

func incrementOrInsertInIntCounter(cts *(map[int]int), x int, inc int) {
	prev, exist := (*cts)[x]
	if exist {
		(*cts)[x] = prev + inc
	} else {
		(*cts)[x] = inc
	}
}

func splitInHalfByDigits(n int) (int, int) {
	s := strconv.Itoa(n)
	nc := len(s) / 2
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
	cts := makeIntCounter()
	initial := parseAllNums(lines[0])
	for _, v := range initial {
		incrementOrInsertInIntCounter(&cts, v, 1)
	}
	for i := 0; i < 25; i++ { //6
		next_cts := makeIntCounter()
		for k := range cts {
			ev := evolveStone(k)
			for _, e := range ev {
				for m := 0; m < cts[k]; m++ {
					incrementOrInsertInIntCounter(&next_cts, e, 1)
				}
			}
		}
		cts = next_cts
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
	cts := makeIntCounter()
	initial := parseAllNums(lines[0])
	for _, v := range initial {
		incrementOrInsertInIntCounter(&cts, v, 1)
	}
	for i := 0; i < 75; i++ {
		next_cts := makeIntCounter()
		for k := range cts {
			ev := evolveStone(k)
			for _, e := range ev {
				incrementOrInsertInIntCounter(&next_cts, e, cts[k])

			}
		}
		cts = next_cts
	}
	var ret = 0
	for _, v := range cts {
		ret += v
	}
	LogPartTwoResult(ret, start)
}

func day11main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day11partOne(day11example)
	day11partTwo(day11example)
	data, _ := os.ReadFile("inputs/day11.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day11partOne(content)
	day11partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
