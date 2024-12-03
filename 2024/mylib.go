package main

import (
	"fmt"
	"time"
)

func LogPartOneResult(v int, start time.Time) {
	fmt.Printf("part 1 result: %d\n", v)
	elapsed := time.Since(start)
	fmt.Println("part 1 time: ", elapsed)
}

func LogPartTwoResult(v int, start time.Time) {
	fmt.Printf("part 2 result: %d\n", v)
	elapsed := time.Since(start)
	fmt.Println("part 2 time: ", elapsed)
}

func LogTimingForDay(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("total time for day: ", elapsed)
}
