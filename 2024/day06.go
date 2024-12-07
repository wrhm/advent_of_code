package main

import (
	"fmt"
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

func findCaret(lines [][]byte) (int, int) {
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

func strListAs2dBytes(lines []string) [][]byte {
	var ret [][]byte
	h := len(lines)
	for i := 0; i < h; i++ {
		ret = append(ret, []byte(lines[i]))
	}
	return ret
}

func printStrList(lines []string) {
	for _, row := range lines {
		fmt.Println(string(row))
	}
}

func print2dBytesList(lines [][]byte) {
	for _, row := range lines {
		fmt.Println(string(row))
	}
}

// returns: (nkeys,inf_loop)
func simulateGuard(grid [][]byte) (int, bool) {

	// lines := strings.Split(contents, "\n")
	h := len(grid)
	w := len(grid[0])
	r, c := findCaret(grid)
	// img := strListAs2dBytes(grid)
	img := grid[:]
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	pointers := "^>v<"
	di := 0
	m := make(map[string]bool)
	inf_loop := true
	for {
		// fmt.Println(r, c)
		k := strconv.Itoa(r) + ":" + strconv.Itoa(c)
		m[k] = true
		nextr := r + dirs[di][0]
		nextc := c + dirs[di][1]
		if !inBounds(nextr, 0, h-1) || !inBounds(nextc, 0, w-1) {
			inf_loop = false
			break
		}
		if grid[nextr][nextc] == '#' {
			// fmt.Println("turning right because", nextr, nextc, "is", grid[nextr][nextc])
			di = (di + 1) % 4
			img[r][c] = pointers[di]

		} else {
			// fmt.Println("going straight in dir", dirs[di])
			img[r][c] = 'X'
			r = nextr
			c = nextc
		}
	}
	fmt.Println("exited at", r, c)
	nkeys := 0
	for i := 0; i < len(m); i++ {
		nkeys++
	}
	fmt.Println("nkeys", nkeys)
	print2dBytesList(img)
	// return r, c, nkeys
	return nkeys, inf_loop
}

func day06partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")

	img := strListAs2dBytes(lines)
	nkeys, _ := simulateGuard(img)
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
	// data, _ := os.ReadFile("inputs/day06.txt")
	// content := string(data)
	// fmt.Println("\nFrom file:")
	// day06partOne(content)
	// day06partTwo(content)
	LogTimingForDay(start)
}
