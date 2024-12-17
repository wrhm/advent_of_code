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

// const day15example = `#######
// #.....#
// #..O..#
// #..OO.#
// #.@O..#
// #.....#
// #######

// ^^>><vv>v>>^`

// const day15example = `#######
// #.....#
// #..O..#
// #..OO.#
// #.@OO.#
// #.....#
// #######

// >><^^>^>v`

// const day15example = `#######
// #.....#
// #..O..#
// #..OO.#
// #.@OO.#
// #.....#
// #######

// >v>^v>>^^<<<^>v>>>>>^^<<<v`

// const day15example = `#######
// #.....#
// #..O..#
// #..OO.#
// #.@O..#
// #.....#
// #######

// ^^>>^>vv`

// Move v:
// robot at 2 12
// before v:
// ####################
// ##[]..[]......[][]##
// ##[]........@..[].##
// ##..........[][][]##
// ##...........[][].##
// ##..##[]..[]......##
// ##...[]...[]..[]..##
// ##.....[]..[].[][]##
// ##........[]......##
// ####################
// after v:
// ####################
// ##[]..[]......[][]##
// ##[]........@..[].##
// ##..........[][][]##
// ##...........[][].##
// ##..##[]..[]......##
// ##...[]...[]..[]..##
// ##.....[]..[].[][]##
// ##........[]......##
// ####################

// const day15example = `#######
// #.....#
// #.....#
// #..O..#
// #.@O..#
// #.O...#
// #######

// >><^^>v`

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
	// fmt.Println("robot at", rr, rc)
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

// orig
// func pushBoxUp(bs [][]byte, r int, c int) {
// 	fmt.Println("pushing up at", r, c)
// 	if bs[r][c] != '[' && bs[r][c] != ']' {
// 		return
// 	}
// 	var left int
// 	if bs[r][c] == '[' {
// 		left = c
// 	} else {
// 		left = c - 1
// 	}
// 	pushBoxUp(bs, r-1, left)
// 	pushBoxUp(bs, r-1, left+1)
// 	// simple case: push into gap
// 	if bs[r-1][left] == '.' && bs[r-1][left+1] == '.' {
// 		bs[r-1][left] = '['
// 		bs[r-1][left+1] = ']'
// 		bs[r][left] = '.'
// 		bs[r][left+1] = '.'
// 		return
// 	}
// }

func canPushBoxUp(bs [][]byte, r int, c int) bool {
	if bs[r][c] == '#' {
		return false
	}
	var left int
	if bs[r][c] == '[' {
		left = c
	} else {
		left = c - 1
	}
	if bs[r-1][left] == '.' && bs[r-1][left+1] == '.' {
		return true
	}
	rec_l := canPushBoxUp(bs, r-1, left)
	rec_r := canPushBoxUp(bs, r-1, left+1)
	return rec_l && rec_r
}

func pushBoxUp(bs [][]byte, r int, c int) {
	fmt.Println("pushing up at", r, c)
	if bs[r][c] != '[' && bs[r][c] != ']' {
		return
	}
	if !canPushBoxUp(bs, r, c) {
		return
	}
	var left int
	if bs[r][c] == '[' {
		left = c
	} else {
		left = c - 1
	}

	pushBoxUp(bs, r-1, left)
	pushBoxUp(bs, r-1, left+1)
	// simple case: push into gap
	if bs[r-1][left] == '.' && bs[r-1][left+1] == '.' {
		bs[r-1][left] = '['
		bs[r-1][left+1] = ']'
		bs[r][left] = '.'
		bs[r][left+1] = '.'
	}
}

// orig
// func pushBoxDown(bs [][]byte, r int, c int) {
// 	fmt.Println("pushing down at", r, c)
// 	if bs[r][c] != '[' && bs[r][c] != ']' {
// 		return
// 	}
// 	var left int
// 	if bs[r][c] == '[' {
// 		left = c
// 	} else {
// 		left = c - 1
// 	}
// 	pushBoxDown(bs, r+1, left)
// 	pushBoxDown(bs, r+1, left+1)
// 	// simple case: push into gap
// 	if bs[r+1][left] == '.' && bs[r+1][left+1] == '.' {
// 		bs[r+1][left] = '['
// 		bs[r+1][left+1] = ']'
// 		bs[r][left] = '.'
// 		bs[r][left+1] = '.'
// 		return
// 	}
// }

func canPushBoxDown(bs [][]byte, r int, c int) bool {
	if bs[r][c] == '#' {
		return false
	}
	var left int
	if bs[r][c] == '[' {
		left = c
	} else {
		left = c - 1
	}
	if bs[r+1][left] == '.' && bs[r+1][left+1] == '.' {
		return true
	}
	rec_l := canPushBoxDown(bs, r+1, left)
	rec_r := canPushBoxDown(bs, r+1, left+1)
	return rec_l && rec_r
}

func pushBoxDown(bs [][]byte, r int, c int) {
	fmt.Println("pushing down at", r, c)
	if bs[r][c] != '[' && bs[r][c] != ']' {
		return
	}
	if !canPushBoxDown(bs, r, c) {
		return
	}
	var left int
	if bs[r][c] == '[' {
		left = c
	} else {
		left = c - 1
	}

	pushBoxDown(bs, r+1, left)
	pushBoxDown(bs, r+1, left+1)
	// simple case: push into gap
	if bs[r+1][left] == '.' && bs[r+1][left+1] == '.' {
		bs[r+1][left] = '['
		bs[r+1][left+1] = ']'
		bs[r][left] = '.'
		bs[r][left+1] = '.'
	}
}

// Move v:
// robot at 2 12
// before v:
// ####################
// ##[]..[]......[][]##
// ##[]........@..[].##
// ##..........[][][]##
// ##...........[][].##
// ##..##[]..[]......##
// ##...[]...[]..[]..##
// ##.....[]..[].[][]##
// ##........[]......##
// ####################
// after v:
// ####################
// ##[]..[]......[][]##
// ##[]........@..[].##
// ##..........[][][]##
// ##...........[][].##
// ##..##[]..[]......##
// ##...[]...[]..[]..##
// ##.....[]..[].[][]##
// ##........[]......##
// ####################

// Move #312: v:
// before v:
// ####################
// ##[]..[]......[][]##
// ##[]...........[].##
// ##...........@[][]##
// ##..........[].[].##
// ##..##[]..[].[]...##
// ##...[]...[]..[]..##
// ##.....[]..[].[][]##
// ##........[]......##
// ####################
// after v:
// ####################
// ##[]..[]......[][]##
// ##[]...........[].##
// ##...........@[][]##
// ##..........[].[].##
// ##..##[]..[].[]...##
// ##...[]...[]..[]..##
// ##.....[]..[].[][]##
// ##........[]......##
// ####################

func canPushBox(bs [][]byte, r int, c int, dr int) bool {
	// fmt.Println("can push", r, c, "?")
	if bs[r][c] == '#' {
		// fmt.Println("can't push", r, c)
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
		// fmt.Println("can push", r, c)
		return true
	}
	// rec_l := canPushBox(bs, r+dr, left, dr)
	// rec_r := canPushBox(bs, r+dr, left+1, dr)

	rec_l := bs[r+dr][left] == '.' || canPushBox(bs, r+dr, left, dr)
	rec_r := bs[r+dr][left+1] == '.' || canPushBox(bs, r+dr, left+1, dr)

	return rec_l && rec_r
}

func pushBox(bs [][]byte, r int, c int, dr int) {
	// fmt.Println("pushing down at", r, c)
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
	// if bs[r][left] == '.' && bs[r][left+1] == '.' {
	// 	return
	// }
	pushBox(bs, r+dr, left, dr)
	pushBox(bs, r+dr, left+1, dr)
	// simple case: push into gap
	if bs[r+dr][left] == '.' && bs[r+dr][left+1] == '.' {
		bs[r+dr][left] = '['
		bs[r+dr][left+1] = ']'
		bs[r][left] = '.'
		bs[r][left+1] = '.'
	}
}

func moveWideWarehouseBot(bs [][]byte, dir byte) {
	rr, rc := findByte(bs, '@')
	// fmt.Println("robot at", rr, rc)
	dr, dc := arrowToDeltas(dir)
	push_r := rr + dr
	push_c := rc + dc
	pushed := bs[push_r][push_c]
	// fmt.Printf("before %c:\n", dir)
	// 	print2dBytesList(bs)
	if pushed == '#' {
		// fmt.Println("pushing wall does nothing")
		return
	}
	if pushed == '.' {
		// fmt.Println("robot moves")
		bs[push_r][push_c] = '@'
		bs[rr][rc] = '.'
		return
	}

	// simpler case: horizontal pushing
	if dr == 0 {
		// fmt.Println("horizontal pushing")
		gapc := push_c
		for {
			if bs[rr][gapc] == '.' {
				// fmt.Printf("pushing into gap at %d,%d\n", rr, gapc)
				break
			}
			if bs[rr][gapc] == '#' {
				// fmt.Println("pushed stack stopped by wall")
				return
			}
			gapc += dc
		}
		for {
			// fmt.Printf("filling gap horizontally at %d,%d\n", rr, gapc)
			bs[rr][gapc] = bs[rr][gapc-dc]
			gapc -= dc
			if gapc == push_c {
				break
			}

		}
		bs[push_r][push_c] = '@'
		bs[rr][rc] = '.'
	} else {
		// fmt.Println("vertical pushing not implemented")
		fmt.Printf("before %c:\n", dir)
		print2dBytesList(bs)
		if dr == -1 {
			// pushBoxUp(bs, rr-1, rc)
			pushBox(bs, rr-1, rc, -1)
			if bs[rr-1][rc] == '.' {
				bs[rr-1][rc] = '@'
				bs[rr][rc] = '.'
			}
		} else {
			// pushBoxDown(bs, rr+1, rc)
			pushBox(bs, rr+1, rc, 1)
			if bs[rr+1][rc] == '.' {
				bs[rr+1][rc] = '@'
				bs[rr][rc] = '.'
			}
		}
		fmt.Printf("after %c:\n", dir)
		print2dBytesList(bs)
	}
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
		fmt.Printf("\nMove #%d: %c:\n", i, moves[i])
		moveWarehouseBot(bs, moves[i])
		// print2dBytesList(bs)
	}
	fmt.Println("final")
	print2dBytesList(bs)
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
	print2dBytesList(bs)
	for i := range moves {
		fmt.Printf("\nMove #%d: %c:\n", i, moves[i])
		moveWideWarehouseBot(bs, moves[i])
		// print2dBytesList(bs)
		// fmt.Println("remaining:", moves[i+1:])
		time.Sleep(time.Millisecond * 0) // 200
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
			// if bs[r][c] != '[' {
			// 	continue
			// }
			// dist_l := c
			// dist_r := w - 1 - c - 1
			// dist_t := r
			// dist_b := h - 1 - r
			// min_lr := dist_l
			// if dist_r < dist_l {
			// 	min_lr = dist_r
			// }
			// min_tb := dist_t
			// if dist_b < dist_t {
			// 	min_tb = dist_b
			// }
			// coord := 100*min_tb + min_lr
			// // fmt.Println(r, c, dist_l, dist_r, dist_t, dist_b, min_lr, min_tb, coord)
			// fmt.Printf("rc(%d,%d) lr(%d,%d) tb(%d,%d) mlr:%d mtb:%d co:%d\n", r, c, dist_l, dist_r, dist_t, dist_b, min_lr, min_tb, coord)
			// ret += coord
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
