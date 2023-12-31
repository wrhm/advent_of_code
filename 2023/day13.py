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


def dispj(g):
    for r in g:
        print(''.join(r))


def vert_line_of_sym(g):
    w, h = len(g[0]), len(g)
    d = dict()
    for r in g:
        syms = []
        for i in range(w-1):
            if row_sym(r, i, i+1):
                syms.append((i, i+1))
                if (i, i+1) in d:
                    d[(i, i+1)] += 1
                else:
                    d[(i, i+1)] = 1
    ret = None
    for k in d:
        (l, r) = k
        if d[k] == h:
            if ret is None or l > ret:
                ret = l
    return ret


def vert_lines_of_sym(g):
    w, h = len(g[0]), len(g)
    d = dict()
    for r in g:
        syms = []
        for i in range(w-1):
            if row_sym(r, i, i+1):
                syms.append((i, i+1))
                if (i, i+1) in d:
                    d[(i, i+1)] += 1
                else:
                    d[(i, i+1)] = 1
    ret = []
    for k in d:
        (l, r) = k
        if d[k] == h:
            ret.append(l)
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
        if vs_or is not None:
            vsum += 1+vs_or
        if hs_or is not None:
            hsum += 1+hs_or
    return vsum+100*hsum


def opp(c):
    if c == '.':
        return '#'
    return '.'


def solve_pt2(data):
    vsum = 0
    hsum = 0
    for g in parse(data):
        w, h = len(g[0]), len(g)
        g = [[x for x in y] for y in g]
        gt = transpose(g)
        vs_or = vert_lines_of_sym(g)
        hs_or = vert_lines_of_sym(gt)
        found = False
        for r in range(h):
            for c in range(w):
                g[r][c] = opp(g[r][c])
                gt = transpose(g)
                vs_or2 = vert_lines_of_sym(g)
                hs_or2 = vert_lines_of_sym(gt)
                g[r][c] = opp(g[r][c])
                if (vs_or2, hs_or2) != ([], []):
                    vsd = set(vs_or2)-set(vs_or)
                    hsd = set(hs_or2)-set(hs_or)
                    if len(vsd) > 0:
                        newv = vsd.pop()
                        vsum += 1+newv
                        found = True
                    if len(hsd) > 0:
                        newh = hsd.pop()
                        hsum += 1+newh
                        found = True
                if found:
                    break
            if found:
                break
        assert(found)
    return vsum+100*hsum


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
