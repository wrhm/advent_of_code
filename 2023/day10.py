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


def find_loop(data):
    (sr, sc) = find_s(data)
    loop = [(sr, sc)]
    for i in range(100000):
        (r, c) = loop[-1]
        if len(loop) > 1 and (r, c) == (sr, sc):
            break
        cn = connected_neighbors(data, r, c)
        for n in cn:
            if n not in loop:
                loop.append(n)
                break
        if len(loop) > 3 and (sr, sc) in cn:
            break
    return loop


def solve_pt1(data):
    w, h = len(data[0]), len(data)
    loop = find_loop(data)
    return len(loop)//2


example2 = '''...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........'''.split('\n')

example3 = '''..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........'''.split('\n')

example4 = '''.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...'''.split('\n')

example5 = '''FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L'''.split('\n')


def zoom2(data):
    w, h = len(data[0]), len(data)
    z = [['.' for _ in range(w*2)] for _ in range(h*2)]
    for r in range(h):
        for c in range(w):
            z[r*2][c*2] = data[r][c]

    # |-LJ7F.S
    for r in range(h):
        for c in range(w):
            ch = z[r*2][c*2]
            if ch in 'LF-' and c+1 < w and z[r*2][(c+1)*2] in '-J7S':
                z[r*2][c*2+1] = '-'
            if ch in 'J7-' and c-1 >= 0 and z[r*2][(c-1)*2] in '-LFS':
                z[r*2][c*2-1] = '-'
            if ch in 'F7|' and r+1 < h and z[(r+1)*2][c*2] in '|JLS':
                z[r*2+1][c*2] = '|'
            if ch in 'JL|' and r-1 >= 0 and z[(r-1)*2][c*2] in '|F7S':
                z[r*2-1][c*2] = '|'

    return [''.join(x) for x in z]


def fill(data, r, c, ch='*'):
    w, h = len(data[0]), len(data)
    g = [[data[r][c] for c in range(w)] for r in range(h)]
    touched_edge = False
    q = [(r, c)]
    while len(q) > 0:
        (fr, fc), q = q[0], q[1:]
        if fr in [0, h-1] or fc in [0, w-1]:
            touched_edge = True
        g[fr][fc] = ch
        for (dr, dc) in [(0, 1), (0, -1), (-1, 0), (1, 0)]:
            nr, nc = fr+dr, fc+dc
            if 0 <= nr < h and 0 <= nc < w and g[nr][nc] == '.' and (nr, nc) not in q:
                q.append((nr, nc))
    return ([''.join(x) for x in g], touched_edge)


def count_3x3s(data):
    w, h = len(data[0]), len(data)
    boxes = 0
    for r in range(0, h, 2):
        for c in range(0, w, 2):
            dots = 0
            for (dr, dc) in [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 0), (0, 1), (1, -1), (1, 0), (1, 1)]:
                if not (0 <= r+dr < h and 0 <= c+dc < w):
                    continue
                if data[r+dr][c+dc] == '*':
                    dots += 1
            if dots == 9:
                boxes += 1
    return boxes


def solve_pt2(data):
    w, h = len(data[0]), len(data)
    if __name__ == '__main__':
        print('finding loop')
    loop = find_loop(data)

    if __name__ == '__main__':
        print('erasing off-loop cells')
    data = [[c for c in row] for row in data]
    for r in range(h):
        for c in range(w):
            if (r, c) not in loop:
                data[r][c] = '.'

    if __name__ == '__main__':
        print('computing z2')
    z2 = zoom2(data)

    # fill starting from all 8 squares around S if they are '.'
    # ignore any that BFS to include a grid edge.
    #
    # then count 3x3 squares of '.' centered at even indices in z2
    if __name__ == '__main__':
        print('trying all fills')
    ret = None
    (sr, sc) = find_s(data)
    for (dr, dc) in [(-1, -1), (-1, 0), (-1, 1), (-1, 0), (1, 0), (1, -1), (1, 0), (1, 1)]:
        fr, fc = 2*sr+dr, 2*sc+dc
        if not (0 <= fr < 2*h and 0 <= fc < 2*w):
            continue
        if z2[fr][fc] != '.':
            continue
        dt = [[c for c in row] for row in z2]
        (f2, te) = fill(dt, fr, fc)
        c3x3 = count_3x3s(f2)
        if not te:
            ret = c3x3
            break
    return ret


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example2))
    print('ex3', solve_pt2(example3))
    print('ex4', solve_pt2(example4))
    print('ex4', solve_pt2(example5))
    print('part2', solve_pt2(lines))
