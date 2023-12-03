import day00
import day01
import day02
import day03

import sys
import os
parent_dir = os.path.dirname(os.path.realpath(__file__))
sys.path.append(parent_dir)
sys.path.append(parent_dir+'/..')
import common.shared_utils as su


assert day00.solve_pt1(day00.example) == ''
assert day00.solve_pt2(day00.example) == ''

assert day01.solve_pt1(day01.example1) == 142
assert day01.solve_pt2(day01.example2) == 281

assert day02.solve_pt1(day02.example) == 8
assert day02.solve_pt2(day02.example) == 2286

assert day03.solve_pt1(day03.example) == 4361
assert day03.solve_pt2(day03.example) == 467835
