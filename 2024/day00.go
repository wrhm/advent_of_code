package main

import (
	"fmt"
	"os"
	"time"
)

const day00example = `EXAMPLE CASE HERE
`

func day00partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	var ret = 0
	LogPartOneResult(ret, start)
}

func day00partTwo(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day00main() {
	start := time.Now()
	fmt.Println("Example:")
	day00partOne(day00example)
	day00partTwo(day00example)
	data, _ := os.ReadFile("day00.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day00partOne(string(content))
	day00partTwo(string(content))
	elapsed := time.Since(start)
	fmt.Println("total time: ", elapsed)
}
