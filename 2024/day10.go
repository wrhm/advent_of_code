package main

import (
	"fmt"
	"os"
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

// const day10example = `.....0.
// ..4321.
// ..5..2.
// ..6543.
// ..7..4.
// ..8765.
// ..9....`

// const day10example = `..90..9
// ...1.98
// ...2..7
// 6543456
// 765.987
// 876....
// 987....`

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
	hist []*GridPoint
}

func BFSClimb(grid *[][]byte, start_gp *GridPoint) []*BFSClimbQElem {
	h := len(*grid)
	w := len((*grid)[0])
	var q []*BFSClimbQElem
	zhist := []*GridPoint{}
	zhist = append(zhist, start_gp)
	q = append(q, &BFSClimbQElem{0, start_gp, zhist})
	visited := make2dPointSet()
	var ret []*BFSClimbQElem
	for {
		if len(q) == 0 {
			break
		}
		qe := q[0]
		q = q[1:]
		if qe.elev == 9 {
			ret = append(ret, qe)
		}
		r := qe.gp.r
		c := qe.gp.c
		insertInto2dPointSet(&visited, qe.gp.r, qe.gp.c)
		dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
		for i := 0; i < 4; i++ {
			nextr := r + dirs[i][0]
			nextc := c + dirs[i][1]
			if inBounds(nextr, 0, h-1) && inBounds(nextc, 0, w-1) {
				nexth := digitByteAsInt((*grid)[nextr][nextc])
				if qe.elev+1 == nexth {
					hist := []*GridPoint{}
					for _, x := range qe.hist {
						hist = append(hist, x)
					}
					hist = append(hist, &GridPoint{nextr, nextc})
					q = append(q, &BFSClimbQElem{nexth, &GridPoint{nextr, nextc}, hist})
				}
			}
		}
	}
	return ret
}

func day10partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	bs := strListAs2dBytes(lines)
	trailheads := findTrailheads(&bs)
	var ret = 0
	for _, t := range trailheads {
		summits := BFSClimb(&bs, t)
		uniq := make2dPointSet()
		for _, v := range summits {
			insertInto2dPointSet(&uniq, v.gp.r, v.gp.c)
		}
		ret += len(uniq)
	}
	LogPartOneResult(ret, start)
}

func day10partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	bs := strListAs2dBytes(lines)
	trailheads := findTrailheads(&bs)
	var ret = 0
	for _, t := range trailheads {
		summits := BFSClimb(&bs, t)
		uniq := make(map[string]bool)
		for _, v := range summits {
			k := ""
			for _, x := range v.hist {
				k += fmt.Sprintf("%d,%d ", x.r, x.c)
			}
			uniq[k] = true
		}
		rating := len(uniq)
		ret += rating
	}
	LogPartTwoResult(ret, start)
}

func day10main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day10partOne(day10example)
	day10partTwo(day10example)
	data, _ := os.ReadFile("inputs/day10.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day10partOne(content)
	day10partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
