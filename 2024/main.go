package main

// To run:
// $ go run *.go

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("\n== DAY 01 ==")
	day01main()
	fmt.Println("\n== DAY 02 ==")
	day02main()
	fmt.Println("\n== DAY 03 ==")
	day03main()
	fmt.Println("\n== DAY 04 ==")
	day04main()
	fmt.Println("\n== DAY 05 ==")
	day05main()
	fmt.Println("\n== DAY 06 ==")
	day06main()
	fmt.Println("\n== DAY 07 ==")
	day07main()
	fmt.Println("\n== DAY 08 ==")
	day08main()
	fmt.Println("\n== DAY 09 ==")
	day09main()
	fmt.Println("\n== DAY 10 ==")
	day10main()
	elapsed := time.Since(start)
	fmt.Println("\ngrand total time all days: ", elapsed)
}
