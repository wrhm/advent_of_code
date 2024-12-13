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

// func numSides(rch []int) int {
// 	// Assume a plot of 'A'
// 	// start with a set of all corners of each cell
// 	// in the plot. remove:
// 	// * those with opposing vertical neighbors, or opposing
// 	// horizontal neighbors.
// 	// * those touching only other cells of the same plot.
// 	// Count the remaining corners.
// 	ret := 0
// 	cells := make(map[int]bool)
// 	corners := make(map[string]bool)
// 	compass := []string{"nw", "ne", "se", "sw"}
// 	for _, v := range rch {
// 		cells[v] = true
// 		r, c := rCFromInt(v, kDay12Mult)
// 		for i := 0; i < 4; i++ {
// 			cmp := compass[i]
// 			// fmt.Printf("adding %s corner of %d,%d\n", cmp, r, c)
// 			// k := fmt.Sprintf("%d,%d,%s", r, c, cmp)
// 			k := makeCompassPoint(r, c, cmp)
// 			corners[k] = true
// 		}
// 	}
// 	fmt.Println(corners)
// 	fmt.Println("removing horizontal")
// 	// new map because removing points from set breaks later
// 	// marking of removals. so rather than remove collinear,
// 	// only add non-collinear to new set.
// 	corners2 := make(map[string]bool)
// 	for v := range corners {
// 		// r, c := rCFromInt(v, kDay12Mult)
// 		r, c, co := unpackCompassPoint(v)
// 		// TODO: canonicalize corners as all top-left, for easier deduplication.
// 		// khl:=
// 		corners2[v] = true
// 	}

// 	// fmt.Println("removing vert")
// 	return ret
// }

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
	// find vertical edges in rows, find horizontal edges in columnds, dedupe.
	cells := make(map[int]bool)
	// corners := make(map[string]bool)
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
	// fmt.Println("cells", cells)
	rmin--
	rmax++
	cmin--
	cmax++
	// fmt.Println(rmin, rmax, cmin, cmax)

	// wallRowsLeftOfCol. key:col, val: map of row to bool
	// goal: remove adjacent keys within the inner map, starting with highest.
	wallRowsLeftOfCol := make(map[int]map[int]bool)
	// wallRowsLeftOfCol[0][0] = true
	// wallRowsLeftOfCol[2][3] = true
	wallRowsRightOfCol := make(map[int]map[int]bool)

	wallColsAboveRow := make(map[int]map[int]bool)
	wallColsBelowRow := make(map[int]map[int]bool)

	const kRemovals = 50
	for r := 0; r < rmax; r++ {
		for c := 0; c < cmax; c++ {
			rci := rCToInt(r, c, kDay12Mult)
			rcileft := rCToInt(r, c-1, kDay12Mult)
			rciright := rCToInt(r, c+1, kDay12Mult)
			if !cells[rcileft] && cells[rci] {
				// fmt.Println("vert wall at left of", r, c)
				if wallRowsLeftOfCol[c] == nil {
					wallRowsLeftOfCol[c] = make(map[int]bool)

				}
				wallRowsLeftOfCol[c][r] = true
			}
			if cells[rci] && !cells[rciright] {
				// fmt.Println("vert wall at right of", r, c)
				if wallRowsRightOfCol[c] == nil {
					wallRowsRightOfCol[c] = make(map[int]bool)

				}
				wallRowsRightOfCol[c][r] = true
			}
		}
	}
	// fmt.Println("wRLOC", wallRowsLeftOfCol)
	for i := 0; i < kRemovals; i++ {
		for kc := range wallRowsLeftOfCol {
			for kr := range wallRowsLeftOfCol[kc] {
				if !wallRowsLeftOfCol[kc][kr] {
					continue
				}
				if wallRowsLeftOfCol[kc][kr-1] && !wallRowsLeftOfCol[kc][kr+1] {
					// fmt.Println("removing redundant wRLOC", kr, kc)
					wallRowsLeftOfCol[kc][kr] = false
				}
			}
		}
	}
	// fmt.Println("wRLOC after", wallRowsLeftOfCol)

	// fmt.Println("wRROC", wallRowsRightOfCol)
	for i := 0; i < kRemovals; i++ {
		for kc := range wallRowsRightOfCol {
			for kr := range wallRowsRightOfCol[kc] {
				if !wallRowsRightOfCol[kc][kr] {
					continue
				}
				if wallRowsRightOfCol[kc][kr-1] && !wallRowsRightOfCol[kc][kr+1] {
					// fmt.Println("removing redundant wRROC", kr, kc)
					wallRowsRightOfCol[kc][kr] = false
				}
			}
		}
	}
	// fmt.Println("wRROC after", wallRowsRightOfCol)

	// TODO: debug wallCols for case "B"
	for r := 0; r < rmax; r++ {
		for c := 0; c < cmax; c++ {
			rci := rCToInt(r, c, kDay12Mult)
			rciup := rCToInt(r-1, c, kDay12Mult)
			rcidown := rCToInt(r+1, c, kDay12Mult)
			if !cells[rciup] && cells[rci] {
				// fmt.Println("horiz wall above", r, c)
				if wallColsAboveRow[r] == nil {
					wallColsAboveRow[r] = make(map[int]bool)

				}
				wallColsAboveRow[r][c] = true
			}
			if cells[rci] && !cells[rcidown] {
				// fmt.Println("horiz wall below", r, c)
				if wallColsBelowRow[r] == nil {
					wallColsBelowRow[r] = make(map[int]bool)

				}
				wallColsBelowRow[r][c] = true
			}
		}
	}

	// fmt.Println("wCAR", wallColsAboveRow)
	for i := 0; i < kRemovals; i++ {
		for kr := range wallColsAboveRow {
			for kc := range wallColsAboveRow[kr] {
				if !wallColsAboveRow[kr][kc] {
					continue
				}
				if wallColsAboveRow[kr][kc-1] && !wallColsAboveRow[kr][kc+1] {
					// fmt.Println("removing redundant wCAR", kr, kc)
					wallColsAboveRow[kr][kc] = false
				}
			}
		}
	}
	// fmt.Println("wCAR after", wallColsAboveRow)

	// fmt.Println("wCBR", wallColsBelowRow)
	for i := 0; i < kRemovals; i++ {
		for kr := range wallColsBelowRow {
			for kc := range wallColsBelowRow[kr] {
				if !wallColsBelowRow[kr][kc] {
					continue
				}
				if wallColsBelowRow[kr][kc-1] && !wallColsBelowRow[kr][kc+1] {
					// fmt.Println("removing redundant wCBR", kr, kc)
					wallColsBelowRow[kr][kc] = false
				}
			}
		}
	}
	// fmt.Println("wCBR after", wallColsBelowRow)

	rl := countTrues(wallRowsLeftOfCol)
	rr := countTrues(wallRowsRightOfCol)
	ca := countTrues(wallColsAboveRow)
	cb := countTrues(wallColsBelowRow)

	return rl + rr + ca + cb
}

func day12partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
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
			fmt.Printf("(%d,%d) %s: ", r, c, string(bs[r][c]))
			for _, v := range plot {
				visited[v] = true
				rf, cf := rCFromInt(v, kDay12Mult)
				fmt.Printf("%d,%d ", rf, cf)
			}
			fmt.Println()
			fmt.Println("perimeter", perimeter(plot))
			price := perimeter(plot) * len(plot)
			fmt.Println("price", price)
			ret += price
		}
	}
	LogPartOneResult(ret, start)
}

func day12partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
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
			fmt.Printf("(%d,%d) %s: ", r, c, string(bs[r][c]))
			for _, v := range plot {
				visited[v] = true
				rf, cf := rCFromInt(v, kDay12Mult)
				fmt.Printf("%d,%d ", rf, cf)
			}
			fmt.Println()
			// fmt.Println("perimeter", perimeter(plot))
			// price := perimeter(plot) * len(plot)
			// fmt.Println("price", price)
			// ret += price
			fmt.Println("sides", numSides(plot))
			price := numSides(plot) * len(plot)
			fmt.Println("price", price)
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
