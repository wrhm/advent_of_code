package main

// To run:
// $ go run *.go

import (
	"fmt"
	"sort"
	"time"
)

type DayFn func() time.Duration

type DayAndFn struct {
	day int
	f   DayFn
}

type DayTiming struct {
	day     int
	elapsed time.Duration
}

func main() {
	start := time.Now()
	fns := []DayAndFn{
		// {1, day01main},
		// {2, day02main},
		// {3, day03main},
		// {4, day04main},
		// {5, day05main},
		// {6, day06main},
		// {7, day07main},
		// {8, day08main},
		// {9, day09main},
		// {10, day10main},
		// {11, day11main},
		// {12, day12main},
		// {13, day13main},
		{14, day14main},
	}
	timings := []DayTiming{}
	for _, v := range fns {
		fmt.Printf("\n== DAY %02d ==\n", v.day)
		dt := DayTiming{v.day, v.f()}
		LogTimingForDay(dt.elapsed)
		timings = append(timings, dt)
	}
	sort.Slice(timings, func(i, j int) bool {
		return timings[i].elapsed > timings[j].elapsed
	})
	elapsed := time.Since(start)
	fmt.Println("\ngrand total time all days: ", elapsed)
	fmt.Println("\nSorted, slowest first:")
	for _, v := range timings {
		fmt.Printf("Day %02d: %v\n", v.day, v.elapsed)
	}
}
