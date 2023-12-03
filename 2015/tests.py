import sys
import os
parent_dir = os.path.dirname(os.path.realpath(__file__))
sys.path.append(parent_dir)
sys.path.append(parent_dir+'/..')
import common.shared_utils as su

import day00
import day01


assert day00.solve_pt1(day00.example) == ''
assert day00.solve_pt2(day00.example) == ''

assert day01.solve_pt1(day01.example1) == 3
assert day01.solve_pt2(day01.example2) == 1

