package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// TODO: use maps of locations of #/O/X/^
// to avoid string copies

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

// returns: (nkeys,inf_loop)
func simulateGuard(grid [][]byte) (int, bool) {
	h := len(grid)
	w := len(grid[0])
	r, c := findCaret(grid)
	img := grid[:]
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	pointers := "^>v<"
	di := 0
	// TODO: refactor p_m and pd_m to use make2dPointSet() and
	// insertInto2dPointSet() like day08.go.
	p_m := make(map[string]bool)
	// store position and direction, to detect loops
	pd_m := make(map[string]bool)
	inf_loop := true
	for {
		pk := strconv.Itoa(r) + ":" + strconv.Itoa(c)
		p_m[pk] = true
		pdk := strconv.Itoa(r) + ":" + strconv.Itoa(c) + ":" + strconv.Itoa(di)
		if pd_m[pdk] {
			inf_loop = true
			break
		}
		pd_m[pdk] = true
		nextr := r + dirs[di][0]
		nextc := c + dirs[di][1]
		if !inBounds(nextr, 0, h-1) || !inBounds(nextc, 0, w-1) {
			inf_loop = false
			break
		}
		next_char := grid[nextr][nextc]
		if next_char == '#' || next_char == 'O' {
			di = (di + 1) % 4
			img[r][c] = pointers[di]
		} else {
			img[r][c] = 'X'
			r = nextr
			c = nextc
		}
	}
	// count distinct positions, ignoring direction
	nkeys := len(p_m)
	return nkeys, inf_loop
}

func day06partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	img := strListAs2dBytes(lines)
	nkeys, _ := simulateGuard(img)
	var ret = nkeys
	LogPartOneResult(ret, start)
}

func deepCopy2dBytes(inp [][]byte) [][]byte {
	var out [][]byte
	for _, r := range inp {
		var row []byte
		for _, b := range r {
			row = append(row, b)
		}
		out = append(out, row)
	}
	return out
}

func day06partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	img := strListAs2dBytes(lines)
	h := len(lines)
	w := len(lines[0])
	obst := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			sim := deepCopy2dBytes(img)
			if sim[r][c] != '.' {
				continue
			}
			sim[r][c] = 'O'
			_, inf_loop := simulateGuard(sim)
			if inf_loop {
				obst++
				// fmt.Println("loop in case", r, c, "total:", obst, "rows:", h)
			}
			sim[r][c] = '.'
		}
	}
	var ret = obst
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
