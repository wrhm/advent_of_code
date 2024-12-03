package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const day03example = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func parseNums(s string) (int, int) {
	dr, _ := regexp.Compile(`\d+`)
	num_strs := dr.FindAllString(s, -1)
	a, _ := strconv.Atoi(num_strs[0])
	b, _ := strconv.Atoi(num_strs[1])
	return a, b
}

func day03partOne(contents string) {
	start := time.Now()
	fns, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := fns.FindAllString(contents, -1)
	var total = 0
	for _, v := range matches {
		a, b := parseNums(v)
		total += a * b
	}
	LogPartOneResult(total, start)
}

func day03partTwo(contents string) {
	start := time.Now()
	fns, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	matches := fns.FindAllString(contents, -1)
	var total = 0
	var in_do bool = true
	for _, v := range matches {
		if v == "do()" {
			in_do = true
			continue
		}
		if v == "don't()" {
			in_do = false
			continue
		}
		if in_do {
			a, b := parseNums(v)
			total += a * b
		}
	}
	var ret = total
	LogPartTwoResult(ret, start)
}

func day03main() {
	start := time.Now()
	fmt.Println("Example:")
	day03partOne(day03example)
	day03partTwo(day03example)
	data, _ := os.ReadFile("inputs/day03.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day03partOne(content)
	day03partTwo(content)
	LogTimingForDay(start)
}
