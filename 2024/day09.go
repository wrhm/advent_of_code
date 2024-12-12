package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const day09example = `2333133121414131402`

// const day09example = `12345`

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

func indFirstFreeBlockOfSize(vs *([]int), sz int) int {
	ret := -1
	for i := 0; i < len(*vs); i++ {
		valid := true
		for j := 0; j < sz; j++ {
			if (i+j) >= len(*vs) || (*vs)[i+j] != -1 {
				valid = false
				break
			}
		}
		if valid {
			return i
		}
	}
	return ret
}

func moveBlockOfSize(vs *([]int), sz int, src int, dest int) {
	for i := 0; i < sz; i++ {
		(*vs)[dest+i] = (*vs)[src+i]
		(*vs)[src+i] = -1
	}
}

func findBlockPosAndSize(vs *([]int), v int) (int, int) {
	pos := -1
	for i := 0; i < len(*vs); i++ {
		if (*vs)[i] == v {
			pos = i
			break
		}
	}
	if pos == -1 {
		return -1, 0
	}
	sz := 0
	for {
		sz++
		if pos+sz >= len(*vs) || (*vs)[pos+sz] != v {
			break
		}
	}
	return pos, sz
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
	dmtb := diskMapToIntBlocks(lines[0])
	for i := 0; i < 1e6; i++ {
		if !moveLastIntBlock(&dmtb) {
			break
		}
	}
	var ret = filesystemIntsChecksum(&dmtb)
	LogPartOneResult(ret, start)
}

func day09partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	dmtb := diskMapToIntBlocks(lines[0])
	mxv := 0
	for _, v := range dmtb {
		if v > mxv {
			mxv = v
		}
	}
	for bi := mxv; bi >= 0; bi-- {
		pos, sz := findBlockPosAndSize(&dmtb, bi)
		iffb := indFirstFreeBlockOfSize(&dmtb, sz)
		if iffb > -1 && iffb < pos {
			moveBlockOfSize(&dmtb, sz, pos, iffb)
		}
	}
	var ret = filesystemIntsChecksum(&dmtb)
	LogPartTwoResult(ret, start)
}

func day09main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day09partOne(day09example)
	day09partTwo(day09example)
	data, _ := os.ReadFile("inputs/day09.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day09partOne(content)
	day09partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
