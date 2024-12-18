import os
import sys

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

import day00
import day01
import day02
import day03
import day04
import day05
import day06
import day07
import day08
import day09
import day10
import day11
import day12
import day13
import day14
import day15
import day16
import day19
import day21

assert day00.solve_pt1(day00.example) == ''
assert day00.solve_pt2(day00.example) == ''

assert day01.solve_pt1(day01.example1) == 142
assert day01.solve_pt2(day01.example2) == 281

assert day02.solve_pt1(day02.example) == 8
assert day02.solve_pt2(day02.example) == 2286

assert day03.solve_pt1(day03.example) == 4361
assert day03.solve_pt2(day03.example) == 467835

assert day04.solve_pt1(day04.example) == 13
assert day04.solve_pt2(day04.example) == 30

assert day05.solve_pt1(day05.example) == 35
assert day05.solve_pt2(day05.example) == 46

assert day06.solve_pt1(day06.example) == 288
assert day06.solve_pt2(day06.example) == 71503

assert day07.solve_pt1(day07.example) == 6440
assert day07.solve_pt2(day07.example) == 5905

assert day08.solve_pt1(day08.example) == 6
assert day08.solve_pt2(day08.example2) == 6

assert day09.solve_pt1(day09.example) == 114
assert day09.solve_pt2(day09.example) == 2

assert day10.solve_pt1(day10.example) == 8
assert day10.solve_pt2(day10.example2) == 4
assert day10.solve_pt2(day10.example3) == 4
assert day10.solve_pt2(day10.example4) == 8
assert day10.solve_pt2(day10.example5) == 10

assert day11.solve_pt1(day11.example) == 374
assert day11.solve_pt2(day11.example, 10) == 1030
assert day11.solve_pt2(day11.example, 100) == 8410

assert day12.solve_pt1(day12.example) == 21
assert day12.solve_pt2(day12.example) == 525152

assert day13.solve_pt1(day13.example) == 405
assert day13.solve_pt2(day13.example) == 400

assert day14.solve_pt1(day14.example) == 136
assert day14.solve_pt2(day14.example) == 64

assert day15.solve_pt1(day15.example) == 1320
assert day15.solve_pt2(day15.example) == 145

assert day16.solve_pt1(day16.example) == 46
assert day16.solve_pt2(day16.example) == 51

# day17
# day17

# day18
# day18

assert day19.solve_pt1(day19.example) == 19114
# assert day19.solve_pt2(day19.example) == 0

# day20
# day20

assert day21.solve_pt1(day21.example, 6) == 16
# day21
