package main

import (
	"fmt"
	"os"
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
	lines := strings.Split(contents, "\n")
	rules, updates := parseRulesAndUpdates(lines)
	var total = 0
	for _, upd := range updates {
		if updateIsCorrectlyOrdered(&rules, &upd) {
			nu := len(upd)
			total += upd[nu/2]
		}
	}
	LogPartOneResult(total, start)
}

func indexOf(list []int, val int) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}

func sortUpdateByRules(rules *([][]int), update []int) []int {
	var upd []int
	for _, v := range update {
		upd = append(upd, v)
	}
	// Bubble sort
	for j := 0; j < len(upd); j++ {
		for _, rule := range *rules {
			a := rule[0]
			b := rule[1]
			ioa := indexOf(upd, a)
			iob := indexOf(upd, b)
			if ioa != -1 && iob != -1 && ioa > iob {
				upd[ioa] = b
				upd[iob] = a
			}
		}
	}
	return upd
}

func day05partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	rules, updates := parseRulesAndUpdates(lines)
	var total = 0
	for _, upd := range updates {
		if !updateIsCorrectlyOrdered(&rules, &upd) {
			std_upd := sortUpdateByRules(&rules, upd)
			total += std_upd[len(upd)/2]
		}
	}
	LogPartTwoResult(total, start)
}

func day05main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day05partOne(day05example)
	day05partTwo(day05example)
	data, _ := os.ReadFile("inputs/day05.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day05partOne(content)
	day05partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
