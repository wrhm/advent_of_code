package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const day06example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func findCaret(lines []string) (int, int) {
	h := len(lines)
	w := len(lines[0])
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if lines[r][c] == '^' {
				return r, c
			}
		}
	}
	return -1, -1
}

func day06partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")
	h := len(lines)
	w := len(lines[0])
	r, c := findCaret(lines)
	// dirs:=make([][]int,4)
	var img [][]byte
	for i := 0; i < h; i++ {
		var row []byte
		for j := 0; j < w; j++ {
			row = append(row, lines[i][j])
		}
		img = append(img, row)
	}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	pointers := "^>v<"
	di := 0
	m := make(map[string]bool)
	for {

		fmt.Println(r, c)
		k := strconv.Itoa(r) + ":" + strconv.Itoa(c)
		m[k] = true
		nextr := r + dirs[di][0]
		nextc := c + dirs[di][1]
		if !inBounds(nextr, 0, h-1) || !inBounds(nextc, 0, w-1) {
			break
		}
		// if lines[nextr][nextc] == '.' {
		// 	fmt.Println("going straight in dir", dirs[di])
		// 	img[r][c] = 'X'
		// 	r = nextr
		// 	c = nextc
		// } else {
		// 	fmt.Println("turning right because", nextr, nextc, "is", lines[nextr][nextc])
		// 	di = (di + 1) % 4
		// 	img[r][c] = pointers[di]
		// }
		if lines[nextr][nextc] == '#' {
			fmt.Println("turning right because", nextr, nextc, "is", lines[nextr][nextc])
			di = (di + 1) % 4
			img[r][c] = pointers[di]

		} else {
			fmt.Println("going straight in dir", dirs[di])
			img[r][c] = 'X'
			r = nextr
			c = nextc
		}
	}
	fmt.Println("exited at", r, c)
	nkeys := 0
	// for _ := range m {
	for i := 0; i < len(m); i++ {
		nkeys++
	}
	fmt.Println("nkeys", nkeys)
	for _, row := range img {
		// fmt.Println(row)
		fmt.Println(string(row))
	}

	var ret = nkeys
	LogPartOneResult(ret, start)
}

func day06partTwo(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day06main() {
	start := time.Now()
	fmt.Println("Example:")
	day06partOne(day06example)
	day06partTwo(day06example)
	data, _ := os.ReadFile("inputs/day06.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day06partOne(content)
	day06partTwo(content)
	LogTimingForDay(start)
}
