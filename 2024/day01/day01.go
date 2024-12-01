package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

var example = `3   4
4   3
2   5
1   3
3   9
3   3`

func sumDiffs(contents string) {
	fmt.Println(contents)
	var list0 []int
	var list1 []int
	lines_ex := strings.Split(contents, "\n")
	for _, line := range lines_ex {
		// vals := strings.Split(line, " ")
		vals := strings.Fields((line))
		// fmt.Println(line, vals)
		// for _, v := range vals {
		// 	fmt.Println(v)
		// }
		vi0, _ := strconv.Atoi(vals[0])
		vi1, _ := strconv.Atoi(vals[1])
		list0 = append(list0, vi0)
		list1 = append(list1, vi1)
	}
	fmt.Println(list0)
	fmt.Println(list1)
	sort.Ints(list0)
	sort.Ints(list1)
	fmt.Println(list0)
	fmt.Println(list1)
	var total_diffs = 0
	for i, v := range list0 {
		total_diffs += int(math.Abs(float64(v - list1[i])))
	}
	fmt.Printf("total: %d\n", total_diffs)
}

func main() {
	// fmt.Printf("%d\n", 123)
	// fmt.Printf("%s\n", example)
	// dat, _ := os.ReadFile("day01.txt")
	// fmt.Println(string(dat))

	sumDiffs(example)

	// file, err := os.Open("day01.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }

	data, _ := ioutil.ReadFile("day01.txt")
	content := string(data)
	sumDiffs(string(content))
}
