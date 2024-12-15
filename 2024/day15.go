package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// const day15example = `########
// #..O.O.#
// ##@.O..#
// #...O..#
// #.#.O..#
// #...O..#
// #......#
// ########

// <^^>>>vv<v>>v<<`

const day15example = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

func arrowToDeltas(dir byte) (int, int) {
	if dir == '>' {
		return 0, 1
	}
	if dir == '<' {
		return 0, -1
	}
	if dir == '^' {
		return -1, 0
	}
	if dir == 'v' {
		return 1, 0
	}
	return 0, 0
}

func moveWarehouseBot(bs [][]byte, dir byte) {
	rr, rc := findByte(bs, '@')
	fmt.Println("robot at", rr, rc)
	dr, dc := arrowToDeltas(dir)
	push_r := rr + dr
	push_c := rc + dc
	pushed := bs[push_r][push_c]
	if pushed == '#' {
		fmt.Println("pushing wall does nothing")
		return
	}
	if pushed == '.' {
		fmt.Println("robot moves")
		bs[push_r][push_c] = '@'
		bs[rr][rc] = '.'
		return
	}
	gapr := push_r
	gapc := push_c
	for {
		if bs[gapr][gapc] == '.' {
			fmt.Printf("pushing into gap at %d,%d\n", gapr, gapc)
			break
		}
		if bs[gapr][gapc] == '#' {
			fmt.Println("pushed stack stopped by wall")
			return
		}
		gapr += dr
		gapc += dc
	}
	for {
		fmt.Printf("filling gap at %d,%d\n", gapr, gapc)
		bs[gapr][gapc] = 'O'
		gapr -= dr
		gapc -= dc
		if gapr == push_r {
			break
		}

	}
	bs[push_r][push_c] = '@'
	bs[rr][rc] = '.'
}

func day15partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var grid []string
	var moves string
	for _, line := range lines {
		if strings.Contains(line, "#") {
			grid = append(grid, line)
		} else if len(line) > 0 {
			moves += line
		}
	}
	fmt.Println(moves)
	bs := strListAs2dBytes(grid)
	print2dBytesList(bs)
	for i := range moves {
		fmt.Printf("\nMove %c:\n", moves[i])
		moveWarehouseBot(bs, moves[i])
		// print2dBytesList(bs)
	}
	h := len(bs)
	w := len(bs[0])
	var ret = 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if bs[r][c] != 'O' {
				continue
			}
			ret += 100*r + c
		}
	}
	LogPartOneResult(ret, start)
}

func day15partTwo(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	fmt.Printf("lines has size %d\n", len(lines))
	var ret = 0
	LogPartTwoResult(ret, start)
}

func day15main() time.Duration {
	start := time.Now()
	fmt.Println("Example:")
	day15partOne(day15example)
	day15partTwo(day15example)
	data, _ := os.ReadFile("inputs/day15.txt")
	content := string(data)
	fmt.Println("\nFrom file:")
	day15partOne(content)
	day15partTwo(content)
	elapsed := time.Since(start)
	return elapsed
}
