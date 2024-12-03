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

func partOne(contents string) {
	start := time.Now()
	fns, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	// fmt.Println(r.FindAllString(contents, -1))
	cleaned := contents
	matches := fns.FindAllString(cleaned, -1)
	// fmt.Printf("matches: %s\n", matches)
	dr, _ := regexp.Compile(`\d+`)
	var total = 0
	for _, v := range matches {
		num_strs := dr.FindAllString(v, -1)
		a, _ := strconv.Atoi(num_strs[0])
		b, _ := strconv.Atoi(num_strs[1])
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
	fmt.Println(fns.FindAllString(contents, -1))

	// dodont, _ := regexp.Compile(`do\(\)|don\'t\(\)`)
	// dodont, _ := regexp.Compile(`do`)
	// fmt.Println(dodont.FindAllString(contents, -1))

	cleaned := contents
	matches := fns.FindAllString(cleaned, -1)
	// fmt.Printf("matches: %s\n", matches)
	dr, _ := regexp.Compile(`\d+`)
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
			num_strs := dr.FindAllString(v, -1)
			a, _ := strconv.Atoi(num_strs[0])
			b, _ := strconv.Atoi(num_strs[1])
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
