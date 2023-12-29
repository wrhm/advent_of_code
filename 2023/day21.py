# day21.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/21/a.txt')

example = '''...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........'''.split('\n')


def solve_pt1(data, steps):
    srow, scol = 0, 0
    w, h = len(data[0]), len(data)
    for r in range(h):
        for c in range(w):
            if data[r][c] == 'S':
                srow, scol = r, c
                r, c = h, w
                break
    q = {(srow, scol)}
    for _ in range(steps):
        next = set()
        for (r, c) in q:
            for (dr, dc) in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                mr, mc = r+dr, c+dc
                if 0 <= mr < h and 0 <= mc < w and data[mr][mc] != '#':
                    next.add((mr, mc))
        q = next
    return len(q)


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example, 6))
    print('part1', solve_pt1(lines, 64))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
