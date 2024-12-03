package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const input_file string = "day03.txt"

const example = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func parseNums(s string) (int, int) {
	dr, _ := regexp.Compile(`\d+`)
	num_strs := dr.FindAllString(s, -1)
	a, _ := strconv.Atoi(num_strs[0])
	b, _ := strconv.Atoi(num_strs[1])
	return a, b
}

func partOne(contents string) {
	start := time.Now()
	fns, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := fns.FindAllString(contents, -1)
	var total = 0
	for _, v := range matches {
		a, b := parseNums(v)
		total += a * b
	}
	var ret = total
	fmt.Printf("part 1 result: %d\n", ret)
	elapsed := time.Since(start)
	fmt.Println("part 1 time: ", elapsed)
}

func partTwo(contents string) {
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
	fmt.Printf("part 2 result: %d\n", ret)
	elapsed := time.Since(start)
	fmt.Println("part 2 time: ", elapsed)
}

func main() {
	start := time.Now()
	fmt.Println("Example:")
	partOne(example)
	partTwo(example)
	data, _ := os.ReadFile(input_file)
	content := string(data)
	fmt.Println("\nFrom file:")
	partOne(string(content))
	partTwo(string(content))
	elapsed := time.Since(start)
	fmt.Println("total time: ", elapsed)
}
