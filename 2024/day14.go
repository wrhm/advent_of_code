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

func makeCountFromPoi(poi map[int]int) map[int]int {
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
	return count
}

func dispRobots(poi map[int]int, w int, h int, quad bool) {
	// rc_pos:count
	count := makeCountFromPoi(poi)
	fmt.Println()
	for r := 0; r < h; r++ {
		var row string
		for c := 0; c < w; c++ {
			x, prs := count[rCToInt(r, c, kDay14Mult)]
			if quad && (r == h/2 || c == w/2) {
				row += " "
			} else if prs {
				row += strconv.Itoa(x)
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
}

func quadrantsProduct(poi map[int]int, w int, h int, quad bool) int {
	// rc_pos:count
	count := makeCountFromPoi(poi)
	top_left := 0
	top_right := 0
	bottom_left := 0
	bottom_right := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			x, prs := count[rCToInt(r, c, kDay14Mult)]
			if quad && (r == h/2 || c == w/2) {
			} else if prs {
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
			}
		}
	}
	return top_left * top_right * bottom_left * bottom_right
}

func safetyFactor(pvs [][]int, w int, h int, sec int, sleep_millis int) int {
	// robot_id:rc_pos
	pos_of_id := make(map[int]int)
	for i, v := range pvs {
		c := v[0] // px
		r := v[1] // py
		pos_of_id[i] = rCToInt(r, c, kDay14Mult)
	}
	for n := 1; n <= sec; n++ {
		var rs []int
		var cs []int
		for i, v := range pvs {
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
			rs = append(rs, new_r)
			cs = append(cs, new_c)
			pos_of_id[i] = rCToInt(new_r, new_c, kDay14Mult)
		}
	}
	return quadrantsProduct(pos_of_id, w, h, true)
}

func sum(x []int) int {
	ret := 0
	for _, v := range x {
		ret += v
	}
	return ret
}

func mean(x []int) float64 {
	return float64(sum(x)) / float64(len(x))
}

func variance(x []int) float64 {
	mu := mean(x)
	num := 0.0
	for _, v := range x {
		diff := float64(v) - mu
		num += diff * diff
	}
	return num / float64(len(x))
}

func naiveModInverse(x int, n int, max_tries int) int {
	for i := 0; i < max_tries; i++ {
		if (i*x)%n == 1 {
			return i
		}
	}
	return -1
}

func findPictureDay(pvs [][]int, w int, h int, print_tree bool) int {
	// robot_id:rc_pos
	pos_of_id := make(map[int]int)
	for i, v := range pvs {
		c := v[0] // px
		r := v[1] // py
		pos_of_id[i] = rCToInt(r, c, kDay14Mult)
	}
	best_rv := 1e6
	rvn := 0
	best_cv := 1e6
	cvn := 0
	for n := 1; n <= 110; n++ {
		var rs []int
		var cs []int
		for i, v := range pvs {
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
			rs = append(rs, new_r)
			cs = append(cs, new_c)
			pos_of_id[i] = rCToInt(new_r, new_c, kDay14Mult)
		}
		rv := variance(rs)
		cv := variance(cs)
		if rv < best_rv {
			best_rv = rv
			rvn = n
		}
		if cv < best_cv {
			best_cv = cv
			cvn = n
		}
	}

	/*
		t = bc mod w
		t = br mod h

		bc+k*w=br (mod h)
		k*w=br-bc (mod h)
		k = inverse(w)*(br-bc) (mod h)

		t = bc + inverse(w)*(br-bc)*w

		bc is cvn
		br is rvn
	*/
	w_inverse := naiveModInverse(w, h, 100)
	ret := cvn + ((w_inverse*(rvn-cvn))%h)*w

	if len(pvs) < 100 || !print_tree {
		return ret
	}

	// rerun the simulation longer to print the tree
	for i, v := range pvs {
		c := v[0] // px
		r := v[1] // py
		pos_of_id[i] = rCToInt(r, c, kDay14Mult)
	}
	for n := 1; n <= ret; n++ {
		var rs []int
		var cs []int
		for i, v := range pvs {
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
			rs = append(rs, new_r)
			cs = append(cs, new_c)
			pos_of_id[i] = rCToInt(new_r, new_c, kDay14Mult)
		}

		if n == ret {
			dispRobots(pos_of_id, w, h, false)
			break
		}
	}
	return ret
}

func day14partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	var pvs [][]int
	for _, line := range lines {
		// [c,r,dc,dr]
		nums := parseAllNumsPosNeg(line)
		pvs = append(pvs, nums)
	}
	w := 11
	h := 7
	if len(lines) > 20 {
		w = 101
		h = 103
	}
	var ret = safetyFactor(pvs, w, h, 100, 0)
	LogPartOneResult(ret, start)
}

func day14partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	var pvs [][]int
	for _, line := range lines {
		// [c,r,dc,dr]
		nums := parseAllNumsPosNeg(line)
		pvs = append(pvs, nums)
	}
	w := 11
	h := 7
	if len(lines) > 20 {
		w = 101
		h = 103
	}
	var ret = findPictureDay(pvs, w, h, false)
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
