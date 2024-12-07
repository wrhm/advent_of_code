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
	// print2dBytesList(img)
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	pointers := "^>v<"
	di := 0
	p_m := make(map[string]bool)
	// store position and direction, to detect loops
	pd_m := make(map[string]bool)
	inf_loop := true
	for {
		// fmt.Println(r, c)
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
	// if inf_loop {
	// 	fmt.Println("infinite loop detected")
	// } else {
	// 	fmt.Println("exited at", r, c)
	// }

	// count distinct positions, ignoring direction
	nkeys := len(p_m)

	// fmt.Println("nkeys", nkeys)
	// print2dBytesList(img)
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

func deepCopy2dBytes(inp [][]byte) [][]byte {
	var out [][]byte
	for _, r := range inp {
		// out = append(out, r[:])
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
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")
	img := strListAs2dBytes(lines)
	h := len(lines)
	w := len(lines[0])
	obst := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			// sim := img[:]

			// var sim [][]byte
			// copy(sim,img)
			sim := deepCopy2dBytes(img)
			// fmt.Println(len(sim))
			// fmt.Println(len(sim[0]))
			if sim[r][c] != '.' {
				continue
			}
			sim[r][c] = 'O'
			// fmt.Println("=== case", r, c, " ===")
			// fmt.Println("img")
			// print2dBytesList(img)
			// fmt.Println("sim")
			// print2dBytesList(sim)
			// fmt.Println()
			_, inf_loop := simulateGuard(sim)
			if inf_loop {
				// fmt.Println("obstacle at", r, c, "causes loop")
				obst++
				fmt.Println("loop in case", r, c, "total:", obst, "rows:", h)
			} else {
				// fmt.Println("exited")
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
