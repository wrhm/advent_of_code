package common

import (
	"fmt"
	"time"
)

func PrintPartOneIntResultWithTiming(v int, start time.Time) {
	fmt.Printf("part 1 result: %d\n", v)
	elapsed := time.Since(start)
	fmt.Println("part 1 time: ", elapsed)
}

func PrintPartTwoIntResultWithTiming(v int, start time.Time) {
	fmt.Printf("part 2 result: %d\n", v)
	elapsed := time.Since(start)
	fmt.Println("part 2 time: ", elapsed)
}
