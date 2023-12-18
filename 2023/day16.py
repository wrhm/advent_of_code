# day16.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su


lines = su.file_as_lines('inputs/16/a.txt')

example = su.file_as_lines('inputs/16/ex1.txt')


def energize_until_done(data, first_beam=(0, 0, 1, 0), max_steps=1000):
    data = [[x for x in y] for y in data]
    w, h = len(data[0]), len(data)
    energized = set()
    beams = set()
    beams.add(first_beam)
    time_since_change = 0
    max_steps = 1000
    for step in range(max_steps):
        time_since_change += 1
        next_beams = set()
        old_energized = len(energized)
        for beam in beams:
            (c, r, dc, dr) = beam
            if not (0 <= r < h and 0 <= c < w):
                continue
            energized.add((c, r))
            ch = data[r][c]
            if ch == '.':
                next_beams.add((c+dc, r+dr, dc, dr))
            elif ch == '\\':
                '''
                right (1,0) -> (0,1) down
                down (0,1) -> (1,0) right
                left (-1,0) -> (0,-1) up
                up (0,-1) -> (-1,0) left
                '''
                dc, dr = dr, dc
                next_beams.add((c+dc, r+dr, dc, dr))
            elif ch == '/':
                '''
                right (1,0) -> (0,-1) up
                up (0,-1) -> (1,0) right
                left (-1,0) -> (0,1) down
                down (0,1) -> (-1,0) left
                '''
                dc, dr = -dr, -dc
                next_beams.add((c+dc, r+dr, dc, dr))
            elif ch == '-':
                if dr == 0:
                    next_beams.add((c+dc, r+dr, dc, dr))
                else:
                    next_beams.add((c, r, 1, 0))
                    next_beams.add((c, r, -1, 0))
            elif ch == '|':
                if dc == 0:
                    next_beams.add((c+dc, r+dr, dc, dr))
                else:
                    next_beams.add((c, r, 0, -1))
                    next_beams.add((c, r, 0, 1))
        beams = set()
        for x in next_beams:
            beams.add(x)
        if len(energized) == old_energized:
            time_since_change += 1
            if time_since_change > 100:
                # probably done energizing
                break
        else:
            time_since_change = 0
        nrg = [['.' for y in x] for x in data]
        for (c, r) in energized:
            nrg[r][c] = '#'
    return len(energized)


def solve_pt1(data):
    return energize_until_done(data, first_beam=(0, 0, 1, 0), max_steps=1000)


def solve_pt2(data):
    w, h = len(data[0]), len(data)
    best = 0
    for r in range(0, h):
        best = max(best, energize_until_done(data, (0, r, 1, 0)))
    for r in range(0, h):
        best = max(best, energize_until_done(data, (w-1, r, -1, 0)))
    for c in range(0, w):
        best = max(best, energize_until_done(data, (c, 0, 0, 1)))
    for c in range(0, w):
        best = max(best, energize_until_done(data, (c, h-1, 0, -1)))
    return best


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    # slow, takes about 5 min.
    print('part2', solve_pt2(lines))
