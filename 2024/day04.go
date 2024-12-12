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

func day04partOne(contents string) {
	// For all starting positions in the grid, search for the string in all 8
	// possible directions.
	start := time.Now()
	lines := strings.Split(contents, "\n")
	w := len(lines[0])
	h := len(lines)
	total := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					x_r, x_c := r+0*dr, c+0*dc
					has_x := gridHasByteAtPos(&lines, x_r, x_c, 'X')
					m_r, m_c := r+1*dr, c+1*dc
					has_m := gridHasByteAtPos(&lines, m_r, m_c, 'M')
					a_r, a_c := r+2*dr, c+2*dc
					has_a := gridHasByteAtPos(&lines, a_r, a_c, 'A')
					s_r, s_c := r+3*dr, c+3*dc
					has_s := gridHasByteAtPos(&lines, s_r, s_c, 'S')
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
	// Consider each position as the center of the snowflake, which must be an
	// "A". Next, grab all diagonal neighbors. These must consist of 2 "M"s and
	// 2 "S"s. Additionally, the "M"s must not be diagonally opposite each
	// other.
	start := time.Now()
	lines := strings.Split(contents, "\n")
	w := len(lines[0])
	h := len(lines)
	total := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			a_r, a_c := r, c
			has_a := gridHasByteAtPos(&lines, a_r, a_c, 'A')
			if has_a == 0 {
				continue
			}
			nw_r, nw_c := r-1, c-1
			ne_r, ne_c := r-1, c+1
			sw_r, sw_c := r+1, c-1
			se_r, se_c := r+1, c+1
			nwm := gridHasByteAtPos(&lines, nw_r, nw_c, 'M')
			nws := gridHasByteAtPos(&lines, nw_r, nw_c, 'S')
			nem := gridHasByteAtPos(&lines, ne_r, ne_c, 'M')
			nes := gridHasByteAtPos(&lines, ne_r, ne_c, 'S')
			swm := gridHasByteAtPos(&lines, sw_r, sw_c, 'M')
			sws := gridHasByteAtPos(&lines, sw_r, sw_c, 'S')
			sem := gridHasByteAtPos(&lines, se_r, se_c, 'M')
			ses := gridHasByteAtPos(&lines, se_r, se_c, 'S')
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

func day04main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day04partOne(day04example)
	day04partTwo(day04example)
	data, _ := os.ReadFile("inputs/day04.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day04partOne(content)
	day04partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
