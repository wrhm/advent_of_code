# day13.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/13/a.txt')

example = '''#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#'''.split('\n')


def parse(lines):
    gs = []
    g = []
    for line in lines:
        if len(line) == 0:
            gs.append(g)
            g = []
        else:
            g.append(line)
    if len(g) > 0:
        gs.append(g)
    return gs


def row_sym(row, left, right):
    n = len(row)
    while left >= 0 and right < n:
        if row[left] != row[right]:
            return False
        left -= 1
        right += 1
    return True


def disp(g):
    for r in g:
        print(r)


def vert_line_of_sym(g):
    w, h = len(g[0]), len(g)
    print('\n')
    disp(g)
    print('\n')
    d = dict()
    for r in g:
        # for i in range(w):
        #     for j in range(i+1, w):
        #         if row_sym(r, i, j):
        #             # print(i, j, r)
        #             if (i, j) in d:
        #                 d[(i, j)] += 1
        #             else:
        #                 d[(i, j)] = 1
        syms = []
        for i in range(w-1):
            if row_sym(r, i, i+1):
                # print(i, i+1, r)
                syms.append((i, i+1))
                if (i, i+1) in d:
                    d[(i, i+1)] += 1
                else:
                    d[(i, i+1)] = 1
        print(r, syms)
    print(d)
    print('want v=', h)
    ret = None
    for k in d:
        (l, r) = k
        if d[k] == h:
            print(k, d[k])
            if ret is None or l > ret:
                ret = l
    return ret


def transpose(g):
    w, h = len(g[0]), len(g)
    return [''.join([g[i][j] for i in range(h)]) for j in range(w)]


def solve_pt1(data):
    vsum = 0
    hsum = 0
    for g in parse(data):
        gt = transpose(g)
        vs_or = vert_line_of_sym(g)
        hs_or = vert_line_of_sym(gt)
        print(vs_or, hs_or)
        if vs_or is not None:
            vsum += 1+vs_or
        if hs_or is not None:
            hsum += 1+hs_or
        # v = vert_line_of_sym(g)+1
        # h = vert_line_of_sym(gt)+1
        # print(v, h, 100*h+v)

        # break
    return vsum+100*hsum


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
