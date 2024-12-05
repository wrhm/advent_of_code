package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const day04example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func inBounds(x int, lo int, hi int) bool {
	return lo <= x && x <= hi
}

func day04partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")
	fmt.Printf("and %d lines of length %d\n", len(lines), len(lines[0]))
	w := len(lines[0])
	h := len(lines)
	total := 0
	// dirs = {{-1,-1},{-1,0},{}}
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					// fmt.Printf("r=%d,c=%d,dr=%d,dc=%d : ", r, c, dr, dc)
					// mr := r + dr
					// mc := c + dc
					x_r, x_c := r+0*dr, c+0*dc
					// has_x := x_r >= 0 && x_r < h && x_c >= 0 && x_c < w && lines[x_r][x_c] == 'X'
					has_x := inBounds(x_r, 0, h-1) && inBounds(x_c, 0, w-1) && lines[x_r][x_c] == 'X'
					// fmt.Println(has_x)
					m_r, m_c := r+1*dr, c+1*dc
					has_m := inBounds(m_r, 0, h-1) && inBounds(m_c, 0, w-1) && lines[m_r][m_c] == 'M'
					// fmt.Println(has_m)
					a_r, a_c := r+2*dr, c+2*dc
					has_a := inBounds(a_r, 0, h-1) && inBounds(a_c, 0, w-1) && lines[a_r][a_c] == 'A'
					// fmt.Println(has_a)
					s_r, s_c := r+3*dr, c+3*dc
					has_s := inBounds(s_r, 0, h-1) && inBounds(s_c, 0, w-1) && lines[s_r][s_c] == 'S'
					// fmt.Println(has_s)
					// fmt.Println(has_x, has_m, has_a, has_s)
					if has_x && has_m && has_a && has_s {
						total++
					}
				}
			}
		}
	}
	var ret = total
	LogPartOneResult(ret, start)
}

func day04partTwo(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")
	fmt.Printf("and %d lines of length %d\n", len(lines), len(lines[0]))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day04main() {
	start := time.Now()
	fmt.Println("Example:")
	day04partOne(day04example)
	day04partTwo(day04example)
	data, _ := os.ReadFile("inputs/day04.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day04partOne(content)
	day04partTwo(content)
	LogTimingForDay(start)
}
