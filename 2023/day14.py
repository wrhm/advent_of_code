# day14.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/14/a.txt')

example = '''O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....'''.split('\n')


# def parse(data):
#     return [[x for x in y] for y in data]

def solve_pt1(data):
    g = [[x for x in y] for y in data]
    w, h = len(g[0]), len(g)
    for r in range(h):
        for c in range(w):
            if g[r][c] == 'O' and r > 0 and g[r-1][c] == '.':
                newr = r-1
                while g[newr][c] == '.' and newr > 0 and g[newr-1][c] == '.':
                    newr -= 1
                g[newr][c] = 'O'
                g[r][c] = '.'
    ret = 0
    for i, r in enumerate(g):
        print(''.join(r), h-i)
        for x in r:
            if x == 'O':
                ret += h-i
    return ret


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
