# day11.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/11/a.txt')

example = '''...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....'''.split('\n')


def em_rows(g):
    ret = []
    for r in range(len(g)):
        if '#' not in g[r]:
            ret.append(r)
    return ret


def em_cols(g):
    w, h = len(g[0]), len(g)
    ret = []
    for c in range(w):
        empty = True
        for r in range(h):
            if g[r][c] == '#':
                empty = False
                break
        if empty:
            ret.append(c)
    return ret


def expand(g, ers, ecs):
    gw, gh = len(g[0]), len(g)
    eg = []
    for r in range(gh):
        er = []
        for c in range(gw):
            er.append(g[r][c])
            if c in ecs:
                er.append('.')
        eg.append(er)
        if r in ers:
            eg.append(er)
    return eg


def add_dists(eg):
    w, h = len(eg[0]), len(eg)
    locs = []
    for r in range(h):
        for c in range(w):
            if eg[r][c] == '#':
                locs.append((r, c))
    ret = 0
    for i in range(len(locs)):
        (a, b) = locs[i]
        for j in range(i+1, len(locs)):
            (c, d) = locs[j]
            dist = abs(a-c)+abs(b-d)
            ret += dist
    return ret


def solve_pt1(data):
    g = [[c for c in r] for r in data]
    eg = expand(g, em_rows(g), em_cols(g))
    return add_dists(eg)


def add_big_dists(og, emr, emc, m):
    w, h = len(og[0]), len(og)
    locs = []
    for r in range(h):
        for c in range(w):
            if og[r][c] == '#':
                locs.append((r, c))
    ret = 0
    for i in range(len(locs)):
        (a, b) = locs[i]
        for j in range(i+1, len(locs)):
            (c, d) = locs[j]
            dist = abs(a-c)+abs(b-d)
            for r in range(min(a, c), max(a, c)+1):
                if r in emr:
                    dist += m-1
            for c in range(min(b, d), max(b, d)+1):
                if c in emc:
                    dist += m-1
            ret += dist
    return ret


def solve_pt2(data, m=1000000):
    g = [[c for c in r] for r in data]
    emr, emc = em_rows(g), em_cols(g)
    return add_big_dists(g, emr, emc, m)


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
