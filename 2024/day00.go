package main

import (
	"fmt"
	"os"
	"time"
)

const input_file string = "day00.txt"

const example = `EXAMPLE CASE HERE
`

func partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	var ret = 0
	fmt.Printf("part 1 result: %d\n", ret)
	elapsed := time.Since(start)
	fmt.Println("part 1 time: ", elapsed)
}

func partTwo(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	var ret = 0
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
