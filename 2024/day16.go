package main

import (
	"fmt"
	"strings"
	"time"
)

const day16example = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

func day16partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	for _, line := range lines {
		fmt.Println(line)
	}
	var ret = 0
	LogPartOneResult(ret, start)
}

func day16partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day16main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day16partOne(day16example)
	day16partTwo(day16example)
	// data, _ := os.ReadFile("inputs/day16.txt")
	// content := string(data)
	// fmt.Println("\nFrom file:")
	// day16partOne(content)
	// day16partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
