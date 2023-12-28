# day18.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/18/a.txt')

example = '''R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)'''.split('\n')


def parse_line(line):
    [dir, dist, _] = line.split()
    return (dir, int(dist))


def solve_pt1(data):
    legs = [parse_line(x) for x in data]
    # print(legs)
    # (r,c)
    (r, c) = (0, 0)
    dpos = {'R': (0, 1), 'L': (0, -1), 'U': (-1, 0), 'D': (1, 0)}
    loop = [(0, 0)]
    for (dir, dist) in legs:
        (dr, dc) = dpos[dir]
        r += dist*dr
        c += dist*dc
        # print(r, c)
        loop.append((r, c))
    # print(loop)
    trench = 0
    walls = set()
    for i in range(1, len(loop)):
        (sr, sc) = loop[i-1]
        (er, ec) = loop[i]
        # trench += abs(loop[i][0]-loop[i-1][0])+abs(loop[i][1]-loop[i-1][1])
        trench += abs(er-sr)+abs(ec-sc)
        for r in range(min(sr, er), max(sr, er)+1):
            for c in range(min(sc, ec), max(sc, ec)+1):
                walls.add((r, c))
    # print('nwalls', len(walls))
    # print('walls', walls)
    # ws = [x for x in walls]
    # ws = sorted(ws, key=lambda x:x[0]*1000+x[1])
    # print('walls',ws)

    print('trench', trench)

    minr, minc, maxr, maxc = int(1e6), int(1e6), int(-1e6), int(-1e6)
    for (r, c) in loop:
        minr = min(minr, r)
        maxr = max(maxr, r)
        minc = min(minc, c)
        maxc = max(maxc, c)
    print(minr, maxr, minc, maxc)

    interior = 0
    for r in range(minr, maxr+1):
        c = minc
        inside = False
        was_wall = False
        wall_crossings = 0
        while c <= maxc:
            is_wall = (r, c) in walls
            # if was_wall and not is_wall:
            #     inside = not inside
            #     # if inside:
            #     #     print('inside', r, c)
            # if not is_wall:
            #     interior += 1
            #     print('interior', r, c)
            # was_wall = is_wall
            if was_wall and not is_wall:
                wall_crossings += 1
                inside = wall_crossings % 2 == 1
            if inside and not is_wall:
                # print('inside', r, c)
                interior += 1
            was_wall = is_wall
            c += 1
    print('interior', interior)
    return trench+interior


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
