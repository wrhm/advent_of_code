# day01.py

import sys
import os
parent_dir = os.path.dirname(os.path.realpath(__file__))
sys.path.append(parent_dir)
sys.path.append(parent_dir+'/..')
import common.shared_utils as su

lines = su.file_as_lines('inputs/01/a.txt')

example1 = '(()(()('
example2 = ')'

def to_floor(ps):
    f = 0
    for c in ps:
        if c=='(':
            f += 1
        else:
            f -= 1
    return f

def solve_pt1(data):
    return to_floor(data)

def btime(ps):
    f = 0
    t = 1
    for c in ps:
        if c=='(':
            f += 1
        else:
            f -= 1
        if f<0:
            return t
        t += 1
    return t

def solve_pt2(data):
    return btime(data)


if __name__ == '__main__':
    print('ex1', solve_pt1(example1))
    print('part1', solve_pt1(lines[0]))
    print('ex2', solve_pt2(example2))
    print('part2', solve_pt2(lines[0]))
