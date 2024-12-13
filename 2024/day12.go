package main

import (
	"fmt"
	"os"
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
	// fmt.Printf("bs[%d][%d]=%s\n", r0, c0, string((*bs)[r0][c0]))
	plant := (*bs)[r0][c0]
	var q []int
	// fmt.Println(1)
	q = append(q, rCToInt(r0, c0, kDay12Mult))
	visited := make(map[int]bool)
	var ret []int
	// iters := 300
	for {
		// iters--
		// if iters == 0 {
		// 	fmt.Println("inf")
		// 	break
		// }
		if len(q) == 0 {
			break
		}
		// fmt.Println(q)
		q0 := q[0]
		q = q[1:]
		if visited[q0] {
			continue
		}
		r, c := rCFromInt(q0, kDay12Mult)
		// fmt.Printf("%d: (%d,%d)\n", q0, r, c)

		visited[q0] = true

		if (*bs)[r][c] == plant {
			// fmt.Printf("%d,%d is %s. Adding %d to result.\n", r, c, string(plant), q0)
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

func day12partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	bs := strListAs2dBytes(lines)
	h := len(bs)
	w := len(bs[0])

	// plot := floodPlot(&bs, 0, 0)
	// for _, v := range plot {
	// 	r, c := rCFromInt(v, kDay12Mult)
	// 	fmt.Printf("(%d,%d):%s\n", r, c, string(bs[r][c]))
	// }
	// fmt.Println(plot)

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
	var ret = 0
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
