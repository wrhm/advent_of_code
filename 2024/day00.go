package main

import (
	"fmt"
	"os"
)

var example = `EXAMPLE CASE HERE
`

func partOne(contents string) {
	fmt.Printf("contents has size %d", len(contents))
	var ret = 0
	fmt.Printf("part 1 total: %d\n", ret)
}

func partTwo(contents string) {
	fmt.Printf("contents has size %d", len(contents))
	var ret = 0
	fmt.Printf("part 2 total: %d\n", ret)
}

func main() {
	fmt.Println("Example:")
	partOne(example)
	partTwo(example)
	data, _ := os.ReadFile("day00.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	partOne(string(content))
	partTwo(string(content))
}
