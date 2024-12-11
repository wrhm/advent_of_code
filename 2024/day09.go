package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const day09example = `2333133121414131402`

// const day09example = `12345`

func digitRuneAsInt(r rune) int {
	v, _ := strconv.Atoi(string(r))
	return v
}

func digitByteAsInt(b byte) int {
	v, _ := strconv.Atoi(string(b))
	return v
}

func diskMapToIntBlocks(dm string) []int {
	var ret []int
	filenum := 0
	for i, r := range dm {
		v := digitRuneAsInt(r)
		if i%2 == 0 {
			for j := 0; j < v; j++ {
				ret = append(ret, filenum)
			}
			filenum++
		} else {
			for j := 0; j < v; j++ {
				ret = append(ret, -1)
			}
		}
	}
	return ret
}

func indFirstInt(vs []int, v int) int {
	for i := 0; i < len(vs); i++ {
		if vs[i] == v {
			return i
		}
	}
	return -1
}

func indLastNonInt(vs []int, v int) int {
	ret := -1
	for i := 0; i < len(vs); i++ {
		if vs[i] != v {
			ret = i
		}
	}
	return ret
}

func moveLastIntBlock(vs *([]int)) bool {
	indln := indLastNonInt(*vs, -1)
	indf := indFirstInt(*vs, -1)
	if indln < indf {
		return false
	}
	(*vs)[indf] = (*vs)[indln]
	(*vs)[indln] = -1
	return true
}

func filesystemIntsChecksum(vs *([]int)) int {
	ret := 0
	for i, v := range *vs {
		if v != -1 {
			ret += i * v
		}
	}
	return ret
}

func day09partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	fmt.Println(lines[0])
	dmtb := diskMapToIntBlocks(lines[0])
	fmt.Println(dmtb)
	for i := 0; i < 1e6; i++ {
		if !moveLastIntBlock(&dmtb) {
			break
		}
		fmt.Println("moving block", i, "among", len(dmtb), "bytes")
		// fmt.Println(dmtb)
	}
	fmt.Println("final", dmtb)
	var ret = filesystemIntsChecksum(&dmtb)
	LogPartOneResult(ret, start)
}

func day09partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day09main() {
	start := time.Now()
	fmt.Println("Example:")
	day09partOne(day09example)
	day09partTwo(day09example)
	data, _ := os.ReadFile("inputs/day09.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day09partOne(content)
	day09partTwo(content)
	LogTimingForDay(start)
}
