package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const day07example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func opStrings(n int, base_ops []string) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return base_ops
	}
	rec := opStrings(n-1, base_ops)
	var ret []string
	for _, r := range rec {
		for _, c := range base_ops {
			ret = append(ret, r+c)
		}

	}
	return ret
}

// memoized version of opStrings. Returns prior results
// early if they are known.
func opStringsMem(n int, base_ops []string, mem *(map[int][]string)) []string {
	if n == 0 {
		(*mem)[n] = []string{}
		return (*mem)[n]
	}
	if n == 1 {
		(*mem)[n] = base_ops
		return (*mem)[n]
	}
	x, prs := (*mem)[n]
	if prs {
		return x
	}
	var rec []string
	xp, prsp := (*mem)[n-1]
	if prsp {
		rec = xp
		(*mem)[n-1] = rec
	} else {
		rec = opStringsMem(n-1, base_ops, mem)
	}
	var ret []string
	for _, r := range rec {
		for _, c := range base_ops {
			ret = append(ret, r+c)
		}

	}
	(*mem)[n] = ret
	return ret
}

func evalNumsAndOpsInOrder(nums []int, ops string) int {
	ret := nums[0]
	for i, o := range ops {
		x := nums[i+1]
		if o == '+' {
			ret += x
		} else if o == '*' {
			ret *= x
		} else {
			s := strconv.Itoa(ret) + strconv.Itoa(x)
			ret, _ = strconv.Atoi(s)
		}
	}
	return ret
}

func sumValidBridgeEqs(lines []string, base_ops []string) int {
	var ret = 0
	mem := make(map[int][]string)
	for _, line := range lines {
		nums := parseAllNums(line)
		n_ops := len(nums) - 2
		ops := opStringsMem(n_ops, base_ops, &mem)
		has_sol := false
		for _, op := range ops {
			if evalNumsAndOpsInOrder(nums[1:], op) == nums[0] {
				has_sol = true
				break
			}
		}
		if has_sol {
			ret += nums[0]
		}
	}
	return ret
}

func day07partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	var ret = sumValidBridgeEqs(lines, []string{"+", "*"})
	LogPartOneResult(ret, start)
}

func day07partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	var ret = sumValidBridgeEqs(lines, []string{"+", "*", "C"})
	LogPartTwoResult(ret, start)
}

func day07main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day07partOne(day07example)
	day07partTwo(day07example)
	data, _ := os.ReadFile("inputs/day07.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day07partOne(content)
	day07partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
