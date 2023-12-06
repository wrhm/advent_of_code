# day03.py

import os
import sys

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/03/a.txt')

example = '''467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..'''.split('\n')

'''
parts (dict): { (start_row, start_col) : part_num }

can use length of part_num to calculate occupied cells.

symbols (set): { (row, col) }
'''


def locate_parts_and_symbols(grid):
    w, h = len(grid[0]), len(grid)
    parts = dict()
    symbols = set()
    gear_candidates = set()
    start_r, start_c = -1, -1
    val = 0
    for r in range(h):
        for c in range(w):
            ch = grid[r][c]
            if su.is_digit(ch):
                if start_r == -1:
                    start_r, start_c = r, c
                val = 10*val + int(ch)
            else:
                if ch != '.':
                    symbols.add((r, c))
                    if ch == '*':
                        gear_candidates.add((r, c))
                if start_r != -1:
                    parts[(start_r, start_c)] = val
                start_r, start_c = -1, -1
                val = 0
    return (parts, symbols, gear_candidates)


def touches(a, b):
    (r, c) = a
    (x, y) = b
    return abs(r-x) <= 1 and abs(c-y) <= 1 and not (r == x and c == y)


def touches_any(pos, locs):
    for x in locs:
        if touches(pos, x):
            return True
    return False


def parts_touching(parts, symbols):
    touching = []
    for (r, c) in parts:
        v = parts[(r, c)]
        nc = len(str(v))
        for i in range(nc):
            if touches_any((r, c+i), symbols):
                touching.append(v)
                break
    return touching


def total_gear_ratio(parts, gear_candidates):
    ratio = 0
    touching_parts = dict()
    for gc in gear_candidates:
        touching_parts[gc] = []
        for (r, c) in parts:
            v = parts[(r, c)]
            nc = len(str(v))
            for i in range(nc):
                if touches((r, c+i), gc):
                    touching_parts[gc].append(v)
                    break
    for k in touching_parts:
        vals = touching_parts[k]
        if len(vals) == 2:
            ratio += vals[0]*vals[1]
    return ratio


def solve_pt1(grid):
    (p, s, _) = locate_parts_and_symbols(grid)
    pt = parts_touching(p, s)
    return sum(pt)


def solve_pt2(grid):
    (p, _, gc) = locate_parts_and_symbols(grid)
    return total_gear_ratio(p, gc)


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
