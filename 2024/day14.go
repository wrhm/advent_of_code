package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const day14example = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

// p=x,y. x is col, y is row

const kDay14Mult = 1000

func dispRobots(poi map[int]int, w int, h int, quad bool) int {
	// rc_pos:count
	count := make(map[int]int)
	for _, v := range poi {
		x, prs := count[v]
		if prs {
			count[v] = x + 1
		} else {
			count[v] = 1
		}
	}
	fmt.Println()
	top_left := 0
	top_right := 0
	bottom_left := 0
	bottom_right := 0
	for r := 0; r < h; r++ {
		var row string
		for c := 0; c < w; c++ {
			x, prs := count[rCToInt(r, c, kDay14Mult)]
			if quad && (r == h/2 || c == w/2) {
				row += " "
			} else if prs {
				row += strconv.Itoa(x)
				if r < h/2 {
					if c < w/2 {
						top_left += x
					} else {
						top_right += x
					}
				} else {
					if c < w/2 {
						bottom_left += x
					} else {
						bottom_right += x
					}
				}
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
	if quad {
		fmt.Println(top_left, top_right, bottom_left, bottom_right)
	}
	return top_left * top_right * bottom_left * bottom_right
}

func safetyFactor(pvs [][]int, w int, h int, sec int) int {
	// robot_id:rc_pos
	pos_of_id := make(map[int]int)
	for i, v := range pvs {
		// if i != 10 {
		// 	continue //debugging the example
		// }
		c := v[0] // px
		r := v[1] // py
		// dc := v[2] // vx
		// dr := v[3] // vy
		pos_of_id[i] = rCToInt(r, c, kDay14Mult)
	}
	dispRobots(pos_of_id, w, h, false)
	for n := 1; n <= sec; n++ {
		for i, v := range pvs {
			// if i != 10 {
			// 	continue //debugging the example
			// }
			r, c := rCFromInt(pos_of_id[i], kDay14Mult)
			dc := v[2] // vx
			dr := v[3] // vy
			new_r := (r + dr) % h
			for {
				if new_r < 0 {
					new_r += h
				} else {
					break
				}
			}
			new_c := (c + dc) % w
			for {
				if new_c < 0 {
					new_c += w
				} else {
					break
				}
			}
			fmt.Printf("n=%d. r#%d (dr=%d,dc=%d) moved from %d,%d to %d,%d\n", n, i, dr, dc, r, c, new_r, new_c)
			pos_of_id[i] = rCToInt(new_r, new_c, kDay14Mult)
		}
		fmt.Printf("after %d seconds:\n", n)
		dispRobots(pos_of_id, w, h, false)
	}
	dispRobots(pos_of_id, w, h, false)

	return dispRobots(pos_of_id, w, h, true)
}

func day14partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	var pvs [][]int
	for _, line := range lines {
		// [c,r,dc,dr]
		nums := parseAllNumsPosNeg(line)
		// fmt.Println(line, nums)
		pvs = append(pvs, nums)
	}
	w := 11
	h := 7
	if len(lines) > 20 {
		w = 101
		h = 103
	}
	var ret = safetyFactor(pvs, w, h, 100)
	LogPartOneResult(ret, start)
}

func day14partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day14main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day14partOne(day14example)
	day14partTwo(day14example)
	data, _ := os.ReadFile("inputs/day14.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day14partOne(content)
	day14partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
