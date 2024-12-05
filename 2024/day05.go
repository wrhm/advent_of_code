package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const day05example = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func parseAllNums(s string) []int {
	dr, _ := regexp.Compile(`\d+`)
	num_strs := dr.FindAllString(s, -1)
	var nums []int
	for _, v := range num_strs {
		i, _ := strconv.Atoi(v)
		nums = append(nums, i)
	}
	return nums
}

func parseRulesAndUpdates(lines []string) ([][]int, [][]int) {
	var rules [][]int
	var updates [][]int
	var add_to_rules = true
	for _, v := range lines {
		var nums = parseAllNums(v)
		if len(nums) == 0 {
			add_to_rules = false
		} else if add_to_rules {
			rules = append(rules, nums)
		} else {
			updates = append(updates, nums)
		}
	}
	return rules, updates
}

func updateIsCorrectlyOrdered(rules *([][]int), update *([]int)) bool {
	nr := len(*rules)
	nu := len(*update)
	for i := 0; i < nu; i++ {
		for j := i + 1; j < nu; j++ {
			ui := (*update)[i]
			uj := (*update)[j]
			for k := 0; k < nr; k++ {
				a := (*rules)[k][0]
				b := (*rules)[k][1]
				if a == uj && b == ui {
					return false
				}
			}
		}
	}
	return true
}

func day05partOne(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	lines := strings.Split(contents, "\n")
	// fmt.Println(parseRulesAndUpdates(lines))
	rules, updates := parseRulesAndUpdates(lines)
	var total = 0
	for _, upd := range updates {
		// fmt.Printf("%v,%v\n", upd, updateIsCorrectlyOrdered(&rules, &upd))
		if updateIsCorrectlyOrdered(&rules, &upd) {
			nu := len(upd)
			total += upd[nu/2]
		}
	}
	LogPartOneResult(total, start)
}

func day05partTwo(contents string) {
	start := time.Now()
	fmt.Printf("contents has size %d\n", len(contents))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day05main() {
	start := time.Now()
	fmt.Println("Example:")
	day05partOne(day05example)
	day05partTwo(day05example)
	data, _ := os.ReadFile("inputs/day05.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day05partOne(content)
	day05partTwo(content)
	LogTimingForDay(start)
}
