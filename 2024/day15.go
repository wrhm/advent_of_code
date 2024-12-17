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

// const day15example = `#######
// #...#.#
// #.....#
// #..OO@#
// #..O..#
// #.....#
// #######

// <vv<<^^<<^^`

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
	dr, dc := arrowToDeltas(dir)
	push_r := rr + dr
	push_c := rc + dc
	pushed := bs[push_r][push_c]
	if pushed == '#' {
		return
	}
	if pushed == '.' {
		bs[push_r][push_c] = '@'
		bs[rr][rc] = '.'
		return
	}
	gapr := push_r
	gapc := push_c
	for {
		if bs[gapr][gapc] == '.' {
			break
		}
		if bs[gapr][gapc] == '#' {
			return
		}
		gapr += dr
		gapc += dc
	}
	for {
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

func canPushBox(bs [][]byte, r int, c int, dr int) bool {
	if bs[r][c] == '#' {
		return false
	}
	var left int
	if bs[r][c] == '[' {
		left = c
	} else {
		left = c - 1
	}
	if bs[r][left] == '.' && bs[r][left+1] == '.' {
		return true
	}
	if bs[r+dr][left] == '.' && bs[r+dr][left+1] == '.' {
		return true
	}
	rec_l := bs[r+dr][left] == '.' || canPushBox(bs, r+dr, left, dr)
	rec_r := bs[r+dr][left+1] == '.' || canPushBox(bs, r+dr, left+1, dr)

	return rec_l && rec_r
}

func pushBox(bs [][]byte, r int, c int, dr int) {
	if bs[r][c] != '[' && bs[r][c] != ']' {
		return
	}
	if !canPushBox(bs, r, c, dr) {
		return
	}
	var left int
	if bs[r][c] == '[' {
		left = c
	} else {
		left = c - 1
	}
	pushBox(bs, r+dr, left, dr)
	pushBox(bs, r+dr, left+1, dr)
	if bs[r+dr][left] == '.' && bs[r+dr][left+1] == '.' {
		bs[r+dr][left] = '['
		bs[r+dr][left+1] = ']'
		bs[r][left] = '.'
		bs[r][left+1] = '.'
	}
}

func moveWideWarehouseBot(bs [][]byte, dir byte) {
	rr, rc := findByte(bs, '@')
	dr, dc := arrowToDeltas(dir)
	push_r := rr + dr
	push_c := rc + dc
	pushed := bs[push_r][push_c]
	if pushed == '#' {
		return
	}
	if pushed == '.' {
		bs[push_r][push_c] = '@'
		bs[rr][rc] = '.'
		return
	}

	// simpler case: horizontal pushing
	if dr == 0 {
		gapc := push_c
		for {
			if bs[rr][gapc] == '.' {
				break
			}
			if bs[rr][gapc] == '#' {
				return
			}
			gapc += dc
		}
		for {
			bs[rr][gapc] = bs[rr][gapc-dc]
			gapc -= dc
			if gapc == push_c {
				break
			}

		}
		bs[push_r][push_c] = '@'
		bs[rr][rc] = '.'
	} else {
		if dr == -1 {
			pushBox(bs, rr-1, rc, -1)
			if bs[rr-1][rc] == '.' {
				bs[rr-1][rc] = '@'
				bs[rr][rc] = '.'
			}
		} else {
			pushBox(bs, rr+1, rc, 1)
			if bs[rr+1][rc] == '.' {
				bs[rr+1][rc] = '@'
				bs[rr][rc] = '.'
			}
		}
	}
}

func day15partOne(contents string) {
	start := time.Now()
	lines := strings.Split(contents, "\n")
	var grid []string
	var moves string
	for _, line := range lines {
		if strings.Contains(line, "#") {
			grid = append(grid, line)
		} else if len(line) > 0 {
			moves += line
		}
	}
	bs := strListAs2dBytes(grid)
	for i := range moves {
		moveWarehouseBot(bs, moves[i])
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
	var grid []string
	var moves string
	for _, line := range lines {
		if strings.Contains(line, "#") {
			grid = append(grid, line)
		} else if len(line) > 0 {
			moves += line
		}
	}
	var stretched []string
	for _, line := range grid {
		st := ""
		for _, c := range line {
			if c == 'O' {
				st += "[]"
			} else if c == '@' {
				st += "@."
			} else if c == '#' {
				st += "##"
			} else {
				st += ".."
			}
		}
		stretched = append(stretched, st)
	}
	bs := strListAs2dBytes(stretched)
	for i := range moves {
		moveWideWarehouseBot(bs, moves[i])
	}
	h := len(bs)
	w := len(bs[0])
	var ret = 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if bs[r][c] != '[' {
				continue
			}
			ret += 100*r + c
		}
	}
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
