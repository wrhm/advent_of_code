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


def disp(g):
    for r in g:
        print(r)


def dispj(g):
    for r in g:
        print(''.join(r))


def tilt_north(g):
    w, h = len(g[0]), len(g)
    for r in range(h):
        for c in range(w):
            if g[r][c] == 'O' and r > 0 and g[r-1][c] == '.':
                newr = r-1
                while g[newr][c] == '.' and newr > 0 and g[newr-1][c] == '.':
                    newr -= 1
                g[newr][c] = 'O'
                g[r][c] = '.'
    return g


def total_load(g):
    h = len(g)
    ret = 0
    for i, r in enumerate(g):
        # print(''.join(r), h-i)
        for x in r:
            if x == 'O':
                ret += h-i
    return ret


def solve_pt1(data):
    g = [[x for x in y] for y in data]
    w, h = len(g[0]), len(g)
    g = tilt_north(g)
    return total_load(g)


def rotate_cw(data):
    '''
    01    20
    23 -> 31

    012    30
    345 -> 41
        52

    012    630
    345 -> 741
    678    852
    '''
    w, h = len(data[0]), len(data)
    ret = []
    for c in range(w):
        col = [data[r][c] for r in range(h-1, -1, -1)]
        assert(len(col) == h)
        ret.append(col)
    return ret


def rotate_180(data):
    return rotate_cw(rotate_cw(data))


def rotate_ccw(data):
    return rotate_180(rotate_cw(data))


def tilt_east(g):
    return rotate_cw(tilt_north(rotate_ccw(g)))


def tilt_south(g):
    return rotate_180(tilt_north(rotate_180(g)))


def tilt_west(g):
    return rotate_ccw(tilt_north(rotate_cw(g)))


def spin_cycle(g):
    g = [[x for x in y] for y in g]
    return tilt_east(tilt_south(tilt_west(tilt_north(g))))


def double_join(data):
    return ''.join([''.join(x) for x in data])


def solve_pt2(data):
    g = [[x for x in y] for y in data]
    d = dict()
    d[double_join(g)] = (0, total_load(g))
    start = None
    gap = None
    for i in range(1, 10000):
        g = spin_cycle(g)
        # print('\n=%d=\n' % i)
        # dispj(g)
        dj = double_join(g)
        if dj in d:
            print(i, 'matches', d[dj])
            start = d[dj][0]
            gap = i-d[dj][0]
            break
        d[dj] = (i, total_load(g))

    n = 1000000000
    if gap is None:
        return None
    # print(n, gap, n % gap)
    print('start=%d, gap=%d' % (start, gap))
    # pinrt('N %% %d = %d'%())
    for x in d:
        # print(x, d[x])
        print("[len %d]" % len(x), d[x])
        if d[x][0] == ((n-start) % gap)+start:
            return d[x][1]
    return None


if __name__ == '__main__':
    # print('ex1', solve_pt1(example))
    # print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
