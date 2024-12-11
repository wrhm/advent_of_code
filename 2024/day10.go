package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// const day10example = `...0...
// ...1...
// ...2...
// 6543456
// 7.....7
// 8.....8
// 9.....9`

const day10example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func findTrailheads(grid *[][]byte) []*GridPoint {
	h := len(*grid)
	w := len((*grid)[0])
	var ret []*GridPoint
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if (*grid)[r][c] != '0' {
				continue
			}
			gp := GP(r, c)
			ret = append(ret, gp)
		}
	}
	return ret
}

type BFSClimbQElem struct {
	elev int
	gp   *GridPoint
}

func BFSClimb(grid *[][]byte, start_gp *GridPoint) []*GridPoint {
	h := len(*grid)
	w := len((*grid)[0])
	var q []*BFSClimbQElem
	q = append(q, &BFSClimbQElem{0, start_gp})
	visited := make2dPointSet()
	var ret []*GridPoint
	for {
		if len(q) == 0 {
			break
		}
		qe := q[0]
		q = q[1:]
		if qe.elev == 9 {
			ret = append(ret, qe.gp)
		}
		r := qe.gp.r
		c := qe.gp.c
		pk := strconv.Itoa(r) + ":" + strconv.Itoa(c)
		if visited[pk] {
			// fmt.Println("already saw", r, c)
			continue
		}
		insertInto2dPointSet(&visited, qe.gp.r, qe.gp.c)
		// fmt.Println("examining", qe.elev, *(qe.gp))
		dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
		for i := 0; i < 4; i++ {
			nextr := r + dirs[i][0]
			nextc := c + dirs[i][1]
			if inBounds(nextr, 0, h-1) && inBounds(nextc, 0, w-1) {
				nexth := digitByteAsInt((*grid)[nextr][nextc])
				if qe.elev+1 == nexth {
					// fmt.Println("adding to q:", nextr, nextc, "with h:", nexth)
					q = append(q, &BFSClimbQElem{nexth, &GridPoint{nextr, nextc}})
				}
			}
		}
	}
	return ret
}

func day10partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	bs := strListAs2dBytes(lines)
	// h := len(bs)
	// w := len(bs[0])
	trailheads := findTrailheads(&bs)
	var ret = 0
	for _, v := range trailheads {
		fmt.Println("trailhead", v)
		summits := BFSClimb(&bs, v)
		uniq := make2dPointSet()
		// fmt.Println(summits)
		for _, v := range summits {
			// fmt.Println(v.r, v.c)
			insertInto2dPointSet(&uniq, v.r, v.c)
		}
		fmt.Println(v.r, v.c, "can reach", len(uniq), "summits")
		ret += len(uniq)
	}
	LogPartOneResult(ret, start)
}

func day10partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day10main() {
	start := time.Now()
	fmt.Println("Example:")
	day10partOne(day10example)
	// day10partTwo(day10example)
	data, _ := os.ReadFile("inputs/day10.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day10partOne(content)
	// day10partTwo(content)
	LogTimingForDay(start)
}
