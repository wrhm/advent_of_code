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

func gridHasCharAtPos(lines *([]string), r int, c int, b byte) int {
	w := len((*lines)[0])
	h := len(*lines)
	if inBounds(r, 0, h-1) && inBounds(c, 0, w-1) && (*lines)[r][c] == b {
		return 1
	}
	return 0
}

func day04partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	w := len(lines[0])
	h := len(lines)
	total := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					x_r, x_c := r+0*dr, c+0*dc
					has_x := gridHasCharAtPos(&lines, x_r, x_c, 'X')
					m_r, m_c := r+1*dr, c+1*dc
					has_m := gridHasCharAtPos(&lines, m_r, m_c, 'M')
					a_r, a_c := r+2*dr, c+2*dc
					has_a := gridHasCharAtPos(&lines, a_r, a_c, 'A')
					s_r, s_c := r+3*dr, c+3*dc
					has_s := gridHasCharAtPos(&lines, s_r, s_c, 'S')
					if has_x+has_m+has_a+has_s == 4 {
						total++
					}
				}
			}
		}
	}
	LogPartOneResult(total, start)
}

func day04partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	w := len(lines[0])
	h := len(lines)
	total := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			a_r, a_c := r, c
			has_a := gridHasCharAtPos(&lines, a_r, a_c, 'A')
			if has_a == 0 {
				continue
			}
			nw_r, nw_c := r-1, c-1
			ne_r, ne_c := r-1, c+1
			sw_r, sw_c := r+1, c-1
			se_r, se_c := r+1, c+1
			nwm := gridHasCharAtPos(&lines, nw_r, nw_c, 'M')
			nws := gridHasCharAtPos(&lines, nw_r, nw_c, 'S')
			nem := gridHasCharAtPos(&lines, ne_r, ne_c, 'M')
			nes := gridHasCharAtPos(&lines, ne_r, ne_c, 'S')
			swm := gridHasCharAtPos(&lines, sw_r, sw_c, 'M')
			sws := gridHasCharAtPos(&lines, sw_r, sw_c, 'S')
			sem := gridHasCharAtPos(&lines, se_r, se_c, 'M')
			ses := gridHasCharAtPos(&lines, se_r, se_c, 'S')
			if nwm+nem+swm+sem != 2 {
				continue
			}
			if nws+nes+sws+ses != 2 {
				continue
			}
			if (nwm+sem == 2) || (nem+swm == 2) {
				continue
			}
			if (nws+ses == 2) || (nes+sws == 2) {
				continue
			}
			total++
		}
	}
	LogPartTwoResult(total, start)
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
