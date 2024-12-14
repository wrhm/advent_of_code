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
					}
				}
			}
			if best_a == -1 {
				// no solution
			} else {
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

The minimization constraint is a distraction. For each input,
there is either 1 exact solution, or none.

# Rewrite the 2 equations as

ma1+nb1=p1  (Eq. 1)
ma2+nb2=p2  (Eq. 2)

Add the equations and isolate n.

m(a1+a2)+n(b1+b2)=p2+p2
n=(p1-ma1)/b1  (Eq. 3)

Then plug that expression for n back into Eq 1.

ma2+b2(p1-ma1)/b1=p2

Solve for m, rationalizing by multiplying through by (b1/b1).

m=(b1p2-b2p1)/(a2b1-b2a1)  (Eq. 4)

Then substitue Eq. 4 into Eq. 3 to compute n from m.
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
			// "m" is am. "n" is bm.
			// "a1" is ax. "b1" is bx. "p1" is px.
			// "a2" is ay. "b2" is by. "p2" is py.
			m_num := bx*py - by*px
			m_denom := ay*bx - by*ax
			if m_num%m_denom != 0 {
				// no soln to m
				continue
			}
			m := m_num / m_denom
			n_num := px - m*ax
			n_denom := bx
			if n_num%n_denom != 0 {
				// no soln to n
				continue
			}
			n := n_num / n_denom
			const a_tokens = 3
			const b_tokens = 1
			cost := m*a_tokens + n*b_tokens
			ret += cost
		}
	}
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
