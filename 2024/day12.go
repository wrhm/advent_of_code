package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// const day12example = `AAAA
// BBCD
// BBCC
// EEEC`

// const day12example = `OOOOO
// OXOXO
// OOOOO
// OXOXO
// OOOOO`

const day12example = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const kDay12Mult = 1000

// Compress a 2D int value into a single larger int, for easier
// processing of point locations.
func rCToInt(r int, c int, mult int) int {
	return r*mult + c
}

func rCFromInt(hash int, mult int) (int, int) {
	return hash / mult, hash % mult
}

func floodPlot(bs *([][]byte), r0 int, c0 int) []int {
	h := len(*bs)
	w := len((*bs)[0])
	plant := (*bs)[r0][c0]
	var q []int
	q = append(q, rCToInt(r0, c0, kDay12Mult))
	visited := make(map[int]bool)
	var ret []int
	for {
		if len(q) == 0 {
			break
		}
		q0 := q[0]
		q = q[1:]
		if visited[q0] {
			continue
		}
		r, c := rCFromInt(q0, kDay12Mult)
		visited[q0] = true
		if (*bs)[r][c] == plant {
			ret = append(ret, q0)
			dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
			for i := 0; i < 4; i++ {
				nextr := r + dirs[i][0]
				nextc := c + dirs[i][1]
				if inBounds(nextr, 0, h-1) && inBounds(nextc, 0, w-1) {
					q = append(q, rCToInt(nextr, nextc, kDay12Mult))
				}
			}
		}
	}
	return ret
}

func perimeter(rch []int) int {
	s := make(map[int]bool)
	for _, v := range rch {
		s[v] = true
	}
	ret := 0
	for _, v := range rch {
		ret += 4
		dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
		r, c := rCFromInt(v, kDay12Mult)
		for i := 0; i < 4; i++ {
			nextr := r + dirs[i][0]
			nextc := c + dirs[i][1]
			nrch := rCToInt(nextr, nextc, kDay12Mult)
			if s[nrch] {
				ret--
			}
		}
	}
	return ret
}

func makeCompassPoint(r int, c int, cmp string) string {
	return fmt.Sprintf("%d,%d,%s", r, c, cmp)
}
func unpackCompassPoint(s string) (int, int, string) {
	fs := strings.Split(s, ",")
	ri, _ := strconv.Atoi(fs[0])
	ci, _ := strconv.Atoi(fs[1])
	return ri, ci, fs[2]
}

func countTrues(m map[int]map[int]bool) int {
	ret := 0
	for k := range m {
		for x := range m[k] {
			if m[k][x] {
				ret++
			}
		}
	}
	return ret
}

func numSides(rch []int) int {
	// Find vertical edges in rows; find horizontal edges in columns; dedupe.
	//
	// First, identify all edges, regardless of length. Then, iteratively remove
	// the furthest segment along an edge from a consistent direction, until it
	// has a single segment.
	//
	// Then, count the edge segments.
	cells := make(map[int]bool)
	rmin := 1000
	rmax := 0
	cmin := 1000
	cmax := 0
	for _, v := range rch {
		cells[v] = true
		r, c := rCFromInt(v, kDay12Mult)
		if r < rmin {
			rmin = r
		}
		if r > rmax {
			rmax = r
		}
		if c < cmin {
			cmin = c
		}
		if c > cmax {
			cmax = c
		}
	}
	// Grow the bounding box by 1 in all directions, to allow detection of
	// outer edges.
	rmin--
	rmax++
	cmin--
	cmax++

	// wallRowsLeftOfCol. key:col, val: map of row to bool
	// goal: remove adjacent keys within the inner map, starting with highest.
	wallRowsLeftOfCol := make(map[int]map[int]bool)
	wallRowsRightOfCol := make(map[int]map[int]bool)

	wallColsAboveRow := make(map[int]map[int]bool)
	wallColsBelowRow := make(map[int]map[int]bool)

	// Number of times to attempt shrinking an edge set. Found by
	// experimentation.
	const kRemovals = 20
	for r := 0; r < rmax; r++ {
		for c := 0; c < cmax; c++ {
			rci := rCToInt(r, c, kDay12Mult)
			rcileft := rCToInt(r, c-1, kDay12Mult)
			rciright := rCToInt(r, c+1, kDay12Mult)
			if !cells[rcileft] && cells[rci] {
				// vertical wall at left of r,c
				if wallRowsLeftOfCol[c] == nil {
					wallRowsLeftOfCol[c] = make(map[int]bool)
				}
				wallRowsLeftOfCol[c][r] = true
			}
			if cells[rci] && !cells[rciright] {
				// vertical wall at right of r,c
				if wallRowsRightOfCol[c] == nil {
					wallRowsRightOfCol[c] = make(map[int]bool)
				}
				wallRowsRightOfCol[c][r] = true
			}
		}
	}
	for i := 0; i < kRemovals; i++ {
		for kc := range wallRowsLeftOfCol {
			for kr := range wallRowsLeftOfCol[kc] {
				if !wallRowsLeftOfCol[kc][kr] {
					continue
				}
				// The && ![...] clause ensures only the rearmost element is
				// removed, to prevent creating gaps.
				if wallRowsLeftOfCol[kc][kr-1] && !wallRowsLeftOfCol[kc][kr+1] {
					// remove redundant edge segment
					wallRowsLeftOfCol[kc][kr] = false
				}
			}
		}
	}
	for i := 0; i < kRemovals; i++ {
		for kc := range wallRowsRightOfCol {
			for kr := range wallRowsRightOfCol[kc] {
				if !wallRowsRightOfCol[kc][kr] {
					continue
				}
				if wallRowsRightOfCol[kc][kr-1] && !wallRowsRightOfCol[kc][kr+1] {
					// remove redundant edge segment
					wallRowsRightOfCol[kc][kr] = false
				}
			}
		}
	}

	for r := 0; r < rmax; r++ {
		for c := 0; c < cmax; c++ {
			rci := rCToInt(r, c, kDay12Mult)
			rciup := rCToInt(r-1, c, kDay12Mult)
			rcidown := rCToInt(r+1, c, kDay12Mult)
			if !cells[rciup] && cells[rci] {
				// horiztal wall above r,c
				if wallColsAboveRow[r] == nil {
					wallColsAboveRow[r] = make(map[int]bool)
				}
				wallColsAboveRow[r][c] = true
			}
			if cells[rci] && !cells[rcidown] {
				// horizontal wall below r,c
				if wallColsBelowRow[r] == nil {
					wallColsBelowRow[r] = make(map[int]bool)
				}
				wallColsBelowRow[r][c] = true
			}
		}
	}
	for i := 0; i < kRemovals; i++ {
		for kr := range wallColsAboveRow {
			for kc := range wallColsAboveRow[kr] {
				if !wallColsAboveRow[kr][kc] {
					continue
				}
				if wallColsAboveRow[kr][kc-1] && !wallColsAboveRow[kr][kc+1] {
					// remove redundant edge segment
					wallColsAboveRow[kr][kc] = false
				}
			}
		}
	}
	for i := 0; i < kRemovals; i++ {
		for kr := range wallColsBelowRow {
			for kc := range wallColsBelowRow[kr] {
				if !wallColsBelowRow[kr][kc] {
					continue
				}
				if wallColsBelowRow[kr][kc-1] && !wallColsBelowRow[kr][kc+1] {
					// remove redundant edge segment
					wallColsBelowRow[kr][kc] = false
				}
			}
		}
	}
	rl := countTrues(wallRowsLeftOfCol)
	rr := countTrues(wallRowsRightOfCol)
	ca := countTrues(wallColsAboveRow)
	cb := countTrues(wallColsBelowRow)

	return rl + rr + ca + cb
}

func day12partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	bs := strListAs2dBytes(lines)
	h := len(bs)
	w := len(bs[0])
	visited := make(map[int]bool)
	var ret = 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			rci := rCToInt(r, c, kDay12Mult)
			if visited[rci] {
				continue
			}
			plot := floodPlot(&bs, r, c)
			for _, v := range plot {
				visited[v] = true
			}
			price := perimeter(plot) * len(plot)
			ret += price
		}
	}
	LogPartOneResult(ret, start)
}

func day12partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	bs := strListAs2dBytes(lines)
	h := len(bs)
	w := len(bs[0])
	visited := make(map[int]bool)
	var ret = 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			rci := rCToInt(r, c, kDay12Mult)
			if visited[rci] {
				continue
			}
			plot := floodPlot(&bs, r, c)
			for _, v := range plot {
				visited[v] = true
			}
			price := numSides(plot) * len(plot)
			ret += price
		}
	}
	LogPartTwoResult(ret, start)
}

func day12main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day12partOne(day12example)
	day12partTwo(day12example)
	data, _ := os.ReadFile("inputs/day12.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day12partOne(content)
	day12partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
