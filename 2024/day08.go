package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const day08example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

type GridPoint struct {
	r int
	c int
}

func GP(r int, c int) *GridPoint {
	return &GridPoint{r, c}
}

func make2dPointSet() map[string]bool {
	return make(map[string]bool)
}

func insertInto2dPointSet(ps *(map[string]bool), r int, c int) {
	pk := strconv.Itoa(r) + ":" + strconv.Itoa(c)
	(*ps)[pk] = true
}

func antinode(r1 int, c1 int, r2 int, c2 int) (int, int) {
	return r2 + (r2 - r1), c2 + (c2 - c1)
}

func day08partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")
	bs := strListAs2dBytes(lines)
	h := len(bs)
	w := len(bs[0])
	// from byte to slice of points
	points_per_ch := make(map[byte][]*GridPoint)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			ch := bs[r][c]
			gp := GP(r, c)
			points_per_ch[ch] = append(points_per_ch[ch], gp)
		}
	}
	antinodes := make2dPointSet()
	for k := range points_per_ch {
		if k == '.' {
			continue
		}
		// fmt.Printf("\n%c: ", k)
		// for _, gp := range points_per_ch[k] {
		// 	fmt.Printf("(%d,%d) ", gp.r, gp.c)
		// }
		for _, pi := range points_per_ch[k] {
			for _, pj := range points_per_ch[k] {
				if pi == pj {
					continue
				}
				anr, anc := antinode(pi.r, pi.c, pj.r, pj.c)
				// fmt.Printf("\nantinode of (%d,%d),(%d,%d) is (%d,%d)", pi.r, pi.c, pj.r, pj.c, anr, anc)
				if inBounds(anr, 0, h-1) && inBounds(anc, 0, w-1) {
					// fmt.Printf("\ninserting #%d: (%d,%d)", len(antinodes)+1, anr, anc)
					insertInto2dPointSet(&antinodes, anr, anc)

				}
			}
		}
	}
	fmt.Println()
	// fmt.Println(antinodes)
	// var aks []string
	// for k := range antinodes {
	// 	aks = append(aks, k)
	// }
	// sort.Strings(aks)
	// fmt.Println(aks)
	var ret = len(antinodes)
	LogPartOneResult(ret, start)
}

func collinear(r1 int, c1 int, r2 int, c2 int, r3 int, c3 int) bool {
	dr := r2 - r1
	dc := c2 - c1
	// return (r3-r1)%dr == 0 && (c3-c1)%dc == 0
	if (r3-r1)%dr != 0 || (c3-c1)%dc != 0 {
		return false
	}
	return (r3-r1)/dr == (c3-c1)/dc
}

func day08partTwo(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")
	bs := strListAs2dBytes(lines)
	h := len(bs)
	w := len(bs[0])
	// from byte to slice of points
	points_per_ch := make(map[byte][]*GridPoint)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			ch := bs[r][c]
			gp := GP(r, c)
			points_per_ch[ch] = append(points_per_ch[ch], gp)
		}
	}
	antinodes := make2dPointSet()
	for k := range points_per_ch {
		if k == '.' {
			continue
		}
		// fmt.Printf("\n%c: ", k)
		// for _, gp := range points_per_ch[k] {
		// 	fmt.Printf("(%d,%d) ", gp.r, gp.c)
		// }
		for _, pi := range points_per_ch[k] {
			for _, pj := range points_per_ch[k] {
				if pi == pj {
					continue
				}
				// anr, anc := antinode(pi.r, pi.c, pj.r, pj.c)
				// // fmt.Printf("\nantinode of (%d,%d),(%d,%d) is (%d,%d)", pi.r, pi.c, pj.r, pj.c, anr, anc)
				// if inBounds(anr, 0, h-1) && inBounds(anc, 0, w-1) {
				// 	// fmt.Printf("\ninserting #%d: (%d,%d)", len(antinodes)+1, anr, anc)
				// 	insertInto2dPointSet(&antinodes, anr, anc)
				// }
				for r := 0; r < h; r++ {
					for c := 0; c < w; c++ {
						// include if equally spaced on a line with pi and pj

						if collinear(pi.r, pi.c, pj.r, pj.c, r, c) {
							fmt.Printf("\nline (%d,%d),(%d,%d) includes (%d,%d)", pi.r, pi.c, pj.r, pj.c, r, c)
							if inBounds(r, 0, h-1) && inBounds(c, 0, w-1) {
								insertInto2dPointSet(&antinodes, r, c)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println()
	// fmt.Println(antinodes)
	// var aks []string
	// for k := range antinodes {
	// 	aks = append(aks, k)
	// }
	// sort.Strings(aks)
	// fmt.Println(aks)
	var ret = len(antinodes)
	LogPartTwoResult(ret, start)
}

func day08main() {
	start := time.Now()
	fmt.Println("Example:")
	day08partOne(day08example)
	day08partTwo(day08example)
	data, _ := os.ReadFile("inputs/day08.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day08partOne(content)
	day08partTwo(content)
	LogTimingForDay(start)
}
