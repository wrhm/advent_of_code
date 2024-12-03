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
	// fmt.Printf("contents has size %d\n", len(contents))
	// clr, _ := regexp.Compile(`[^mul\(\),\d]`)
	fns, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	// fmt.Println(r.FindAllString(contents, -1))
	cleaned := contents
	// cleaned = clr.ReplaceAllString(cleaned, "")
	// fmt.Printf("cleaned: %s\n", cleaned)
	matches := fns.FindAllString(cleaned, -1)
	// fmt.Printf("matches: %s\n", matches)
	dr, _ := regexp.Compile(`\d+`)
	var total = 0
	for _, v := range matches {
		// fmt.Printf("%s: %s\n", v, dr.FindAllString(v, -1))
		// vi0, _ := strconv.Atoi(vals[0])
		num_strs := dr.FindAllString(v, -1)
		// for _, ns := range num_strs {
		// 	fmt.Println(ns)
		a, _ := strconv.Atoi(num_strs[0])
		b, _ := strconv.Atoi(num_strs[1])
		// fmt.Printf("adding %d * %d = %d\n", a, b, a*b)
		total += a * b
		// }
	}

	var ret = total
	fmt.Printf("part 1 result: %d\n", ret)
	elapsed := time.Since(start)
	fmt.Println("part 1 time: ", elapsed)
}

// func partTwo(contents string) {
// 	start := time.Now()
// 	fmt.Printf("contents has size %d\n", len(contents))
// 	var ret = 0
// 	fmt.Printf("part 2 result: %d\n", ret)
// 	elapsed := time.Since(start)
// 	fmt.Println("part 2 time: ", elapsed)
// }

func main() {
	start := time.Now()
	fmt.Println("Example:")
	partOne(example)
	// partTwo(example)
	data, _ := os.ReadFile(input_file)
	content := string(data)
	fmt.Println("\nFrom file:")
	partOne(string(content))
	// partTwo(string(content))
	elapsed := time.Since(start)
	fmt.Println("total time: ", elapsed)
}
