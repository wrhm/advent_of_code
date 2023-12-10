# day10.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/10/a.txt')

example = '''..F7.
.FJ|.
SJ.L7
|F--J
LJ...'''.split('\n')


def find_s(grid):
    w, h = len(grid[0]), len(grid)
    for r in range(h):
        for c in range(w):
            if grid[r][c] == 'S':
                return (r, c)
    return None


def connected_neighbors(grid, r, c):
    w, h = len(grid[0]), len(grid)
    # ch = grid[r][c]
    ret = []
    if r > 0 and grid[r-1][c] in '|7FS' and grid[r][c] in '|LJS':
        ret.append((r-1, c))
    if r < h-1 and grid[r+1][c] in '|LJS' and grid[r][c] in '|7FS':
        ret.append((r+1, c))
    if c > 0 and grid[r][c-1] in '-LFS' and grid[r][c] in '-J7S':
        ret.append((r, c-1))
    if c < w-1 and grid[r][c+1] in '-J7S' and grid[r][c] in '-LFS':
        ret.append((r, c+1))
    return ret


def viz_route(loop):
    route = '.'
    for i in range(1, len(loop)):
        fr, fc = loop[i-1]
        tr, tc = loop[i]
        if tr > fr:
            route += 'v'
        elif tr < fr:
            route += '^'
        elif tc > fc:
            route += '>'
        else:
            route += '<'
    return route


def disp(w, h, loop, grid, use_nums=False):
    m = [['.' for _ in range(w)] for _ in range(h)]
    for i in range(len(loop)):
        (r, c) = loop[i]
        if use_nums:
            if i == 0:
                m[r][c] = 'S'
            else:
                m[r][c] = '%d' % (min(i, len(loop)-i) % 10)
        else:
            if i == 0:
                m[r][c] = 'S'
            else:
                m[r][c] = grid[r][c]
    for row in m:
        print(''.join(row))


def solve_pt1(data):
    w, h = len(data[0]), len(data)
    print(find_s(data))
    (sr, sc) = find_s(data)
    print('S is at', sr, sc)
    print('S connections', connected_neighbors(data, sr, sc))
    loop = [(sr, sc)]
    for i in range(100000):
        # while True:
        (r, c) = loop[-1]
        if len(loop) > 1 and (r, c) == (sr, sc):
            break
        cn = connected_neighbors(data, r, c)
        # print(r, c, cn)
        print(i, len(loop), '%s at %d %d has %s' %
              (data[r][r], r, c, [(a, b, data[a][b]) for (a, b) in cn]))
        for n in cn:
            # if n != loop[0]:
            if n not in loop:
                loop.append(n)
                break
        # if (sr, sc) in cn[1:]:
        #     break
        # print(loop)
    print('\nfull loop', loop)
    print(''.join([data[i][j] for (i, j) in loop]))
    print(viz_route(loop))
    print(len(loop), len(loop)//2)
    disp(w, h, loop, data, False)
    print('\n')
    disp(w, h, loop, data, True)
    return len(loop)//2


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
