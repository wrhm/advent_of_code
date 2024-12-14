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
	var ret = 0
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
			fmt.Printf("A: %d,%d B: %d,%d P:%d,%d\n", ax, ay, bx, by, px, py)
			const a_tokens = 3
			const b_tokens = 1
			best_cost := 100000000
			best_a := -1
			best_b := -1
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
						best_b = bm
					}
				}
			}
			if best_a == -1 {
				fmt.Println("no solution")
			} else {
				fmt.Printf("cheapest: am=%d, bm=%d, cost=%d\n", best_a, best_b, best_cost)
				ret += best_cost
			}
		}
	}
	LogPartOneResult(ret, start)
}

func day13partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var ret = 0
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
