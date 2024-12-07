package main

import (
	"fmt"
	"os"
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

func opStrings(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"+", "*"}
	}
	rec := opStrings(n - 1)
	var ret []string
	for _, r := range rec {
		ret = append(ret, r+"+")
		ret = append(ret, r+"*")
	}
	return ret
}

func evalNumsAndOpsInOrder(nums []int, ops string) int {
	ret := nums[0]
	for i, o := range ops {
		x := nums[i+1]
		if o == '+' {
			ret += x
		} else {
			ret *= x
		}
	}
	return ret
}

// func testEvalNumsAndOpsInOrder() {
// 	v1 := evalNumsAndOpsInOrder([]int{1, 2}, "+")
// 	fmt.Println("EXPECTING", v1, "=", 3)
// 	v2 := evalNumsAndOpsInOrder([]int{1, 2}, "*")
// 	fmt.Println("EXPECTING", v2, "=", 2)
// 	v3 := evalNumsAndOpsInOrder([]int{3, 0, 5}, "*+")
// 	fmt.Println("EXPECTING", v3, "=", 5)
// 	v4 := evalNumsAndOpsInOrder([]int{3, 7, 5}, "+*")
// 	fmt.Println("EXPECTING", v4, "=", 50)
// }

func day07partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")

	var ret = 0
	for _, line := range lines {
		nums := parseAllNums(line)
		n_ops := len(nums) - 2
		// fmt.Println(line, nums, "has", len(nums)-1, "inputs and needs", n_ops, "ops")
		ops := opStrings(n_ops)
		// fmt.Println(n_ops, ops)
		has_sol := false
		for _, op := range ops {
			// fmt.Println(nums[0], nums[1:], evalNumsAndOpsInOrder(nums[1:],
			// op))
			if evalNumsAndOpsInOrder(nums[1:], op) == nums[0] {
				has_sol = true
				break
			}
		}
		if has_sol {
			ret += nums[0]
		}
	}
	// testEvalNumsAndOpsInOrder()
	// fmt.Println("all tests passed")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i, opStrings(i))
	// }
	LogPartOneResult(ret, start)
}

func day07partTwo(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day07main() {
	start := time.Now()
	fmt.Println("Example:")
	day07partOne(day07example)
	day07partTwo(day07example)
	data, _ := os.ReadFile("inputs/day07.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day07partOne(content)
	day07partTwo(content)
	LogTimingForDay(start)
}
