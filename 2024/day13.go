package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const day13example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func day13Helper(lines []string, offset bool) int {
	ax := 0
	ay := 0
	bx := 0
	by := 0
	px := 0
	py := 0
	var ret = 0
	const offsetAmt = 10000000000000
	for _, line := range lines {
		nums := parseAllNums(line)
		if len(nums) != 2 {
			continue
		}
		if strings.Contains(line, "A") {
			ax = nums[0]
			ay = nums[1]
		} else if strings.Contains(line, "B") {
			bx = nums[0]
			by = nums[1]
		} else {
			px = nums[0]
			py = nums[1]

			if offset {
				px += offsetAmt
				py += offsetAmt
			}

			// fmt.Printf("A: %d,%d B: %d,%d P:%d,%d\n", ax, ay, bx, by, px, py)
			const a_tokens = 3
			const b_tokens = 1
			best_cost := 100000000
			best_a := -1
			// best_b := -1
			for am := 0; am <= 100; am++ {
				for bm := 0; bm <= 100; bm++ {
					if am*ax+bm*bx != px {
						continue
					}
					if am*ay+bm*by != py {
						continue
					}
					cost := am*a_tokens + bm*b_tokens
					if cost < best_cost {
						best_cost = cost
						best_a = am
						// best_b = bm
					}
				}
			}
			if best_a == -1 {
				// fmt.Println("no solution")
			} else {
				// fmt.Printf("cheapest: am=%d, bm=%d, cost=%d\n", best_a, best_b, best_cost)
				ret += best_cost
			}
		}
	}
	return ret
}

func day13partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	ax := 0
	ay := 0
	bx := 0
	by := 0
	px := 0
	py := 0
	ret := 0
	for _, line := range lines {
		nums := parseAllNums(line)
		if len(nums) != 2 {
			continue
		}
		if strings.Contains(line, "A") {
			ax = nums[0]
			ay = nums[1]
		} else if strings.Contains(line, "B") {
			bx = nums[0]
			by = nums[1]
		} else {
			px = nums[0]
			py = nums[1]

			const a_tokens = 3
			const b_tokens = 1
			best_cost := -1
			best_a := -1
			// best_b := -1
			for am := 0; am <= 100; am++ {
				for bm := 0; bm <= 100; bm++ {
					if am*ax+bm*bx != px {
						continue
					}
					if am*ay+bm*by != py {
						continue
					}
					cost := am*a_tokens + bm*b_tokens
					if best_cost < 0 || cost < best_cost {
						best_cost = cost
						best_a = am
						// best_b = bm
					}
				}
			}
			if best_a == -1 {
				// fmt.Println("no solution")
			} else {
				// fmt.Printf("cheapest: am=%d, bm=%d, cost=%d\n", best_a, best_b, best_cost)
				ret += best_cost
			}
		}
	}
	LogPartOneResult(ret, start)
}

/*
Need to solve simultaneous equations:

	a_m * a_x + b_m * b_x = p_x

	a_m * a_y + b_m * b_y = p_y

for a_m and b_m, while minimizing 3 * a_m + b_m.

minimize 3*a+b on a*p + b*q = r and a*s + b*t = u

https://www.wolframalpha.com/input?i=minimize+3*a%2Bb+on+a*94%2Bb*22%3D8400+and+a*34%2Bb*67%3D5400

Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=10000000018641, Y=10000000010279

https://www.wolframalpha.com/input?i=minimize+3*a%2Bb+on+a*69%2Bb*27%3D10000000018641+and+a*23%2Bb*71%3D10000000010279
a=102,851,800,151
b=107,526,881,786

https://math.libretexts.org/Bookshelves/Algebra/Intermediate_Algebra_1e_(OpenStax)/04%3A_Systems_of_Linear_Equations/4.06%3A_Solve_Systems_of_Equations_Using_Matrices

a*94+b*22=8400
a*34+b*67=5400

94a+22b=8400
34a+67b=5400
*/
func day13partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")

	ax := 0
	ay := 0
	bx := 0
	by := 0
	px := 0
	py := 0
	var ret = 0
	const offsetAmt = 10000000000000
	for _, line := range lines {
		nums := parseAllNums(line)
		if len(nums) != 2 {
			continue
		}
		if strings.Contains(line, "A") {
			ax = nums[0]
			ay = nums[1]
		} else if strings.Contains(line, "B") {
			bx = nums[0]
			by = nums[1]
		} else {
			px = nums[0]
			py = nums[1]

			px += offsetAmt
			py += offsetAmt

			const a_tokens = 3
			const b_tokens = 1
			// best_cost := -1
			// best_a := -1
			// best_b := -1
			// for am := 0; am <= 100; am++ {
			// 	for bm := 0; bm <= 100; bm++ {
			// 		if am*ax+bm*bx != px {
			// 			continue
			// 		}
			// 		if am*ay+bm*by != py {
			// 			continue
			// 		}
			// 		cost := am*a_tokens + bm*b_tokens
			// 		if best_cost < 0 || cost < best_cost {
			// 			best_cost = cost
			// 			best_a = am
			// 			// best_b = bm
			// 		}
			// 	}
			// }

			// "m" is am. "n" is bm.
			// "a1" is ax. "b1" is bx. "p1" is px.
			// "a2" is ay. "b2" is by. "p2" is py.

			m_num := bx*py - by*px
			m_denom := ay*bx - by*ax

			if m_num%m_denom != 0 {
				fmt.Println("no soln to m")
				continue
			}
			m := m_num / m_denom
			n_num := px - m*ax
			n_denom := bx

			if n_num%n_denom != 0 {
				fmt.Println("no soln to n")
				continue
			}
			n := n_num / n_denom
			fmt.Printf("m=%d n=%d\n", m, n)

			cost := m*a_tokens + n*b_tokens
			ret += cost
		}
	}

	// ret := 0
	LogPartTwoResult(ret, start)
}

func day13main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day13partOne(day13example)
	day13partTwo(day13example)
	data, _ := os.ReadFile("inputs/day13.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day13partOne(content)
	day13partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
