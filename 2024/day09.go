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

func nRepsOf(n int, s string) string {
	var ret string
	for i := 0; i < n; i++ {
		ret = ret + s
	}
	return ret
}

func digitRuneAsInt(r rune) int {
	v, _ := strconv.Atoi(string(r))
	return v
}

func digitByteAsInt(b byte) int {
	v, _ := strconv.Atoi(string(b))
	return v
}

func bytesToInts(bs []byte) []int {
	var ret []int
	for _, v := range bs {
		ret = append(ret, digitByteAsInt(v))
	}
	return ret
}

func diskMapToBlocks(dm string) string {
	ret := ""
	filenum := 0
	for i, r := range dm {
		// v, _ := strconv.Atoi(string(c))
		v := digitRuneAsInt(r)
		if i%2 == 0 {
			part := nRepsOf(v, strconv.Itoa(filenum))
			filenum++
			// filenum = (filenum + 1) % 10
			// fmt.Println("file", v, part)
			ret += part
		} else {
			part := nRepsOf(v, ".")
			// fmt.Println("space", v, part)
			ret += part
		}
	}
	return ret
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

// 5602261504 was too small
// 89859464970 was too small

func indFirst(bs []byte, b byte) int {
	for i := 0; i < len(bs); i++ {
		if bs[i] == b {
			return i
		}
	}
	return -1
}

func indLast(bs []byte, b byte) int {
	ret := -1
	for i := 0; i < len(bs); i++ {
		if bs[i] == b {
			ret = i
		}
	}
	return ret
}

func indLastNon(bs []byte, b byte) int {
	ret := -1
	for i := 0; i < len(bs); i++ {
		if bs[i] != b {
			ret = i
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

// func indLastInt(vs []int, v int) int {
// 	ret := -1
// 	for i := 0; i < len(vs); i++ {
// 		if vs[i] == v {
// 			ret = i
// 		}
// 	}
// 	return ret
// }

func indLastNonInt(vs []int, v int) int {
	ret := -1
	for i := 0; i < len(vs); i++ {
		if vs[i] != v {
			ret = i
		}
	}
	return ret
}

//	func moveLastBlock(s string) bool {
//		bs := []byte(s)
//
// returns: modified
func moveLastBlock(bs *([]byte)) bool {
	// fmt.Println("first .", indFirst(*bs, '.'))
	// fmt.Println("last non-.", indLastNon(*bs, '.'))
	// if indLastNon(*bs, '.')+1 == indFirst(*bs, '.')
	indln := indLastNon(*bs, '.')
	indf := indFirst(*bs, '.')
	if indln < indf {
		// fmt.Println("nothing to change")
		return false
	}
	(*bs)[indf] = (*bs)[indln]
	(*bs)[indln] = '.'
	return true
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

// func filesystemChecksum(bs []byte) int {
func filesystemChecksum(s string) int {
	ret := 0
	for i, r := range s {
		ret += i * digitRuneAsInt(r)
	}
	return ret
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
	// dmtb := diskMapToBlocks(lines[0])
	// fmt.Println(dmtb)
	// bs := []byte(dmtb)
	// fmt.Println(string(bs))
	// for i := 0; i < 1e6; i++ {
	// 	if !moveLastBlock(&bs) {
	// 		break
	// 	}
	// 	fmt.Println("moving block", i, "among", len(bs), "bytes")
	// 	// fmt.Println(string(bs))
	// }
	// fmt.Println("final", string(bs))
	// var ret = filesystemChecksum(string(bs))

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
