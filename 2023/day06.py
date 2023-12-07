# day00.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/06/a.txt')

example = '''Time:      7  15   30
Distance:  9  40  200'''.split('\n')


def races(data):
    times = [int(x) for x in data[0].split(':')[1].split()]
    dists = [int(x) for x in data[1].split(':')[1].split()]
    return [x for x in zip(times, dists)]


def dists_with_delays(time):
    return [(i, i*(time-i)) for i in range(time+1)]


def ways_to_win(time, dist_to_beat):
    w = 0
    for i in range(time+1):
        if i*(time-i) > dist_to_beat:
            w += 1
    return w


def solve_pt1(data):
    p = 1
    r = races(data)
    for (t, d) in r:
        p *= ways_to_win(t, d)
    return p


def parse_single_race(data):
    time = int(''.join(data[0].split(':')[1].split()))
    dist = int(''.join(data[1].split(':')[1].split()))
    return (time, dist)


def solve_pt2(data):
    (t, d) = parse_single_race(data)
    return ways_to_win(t, d)


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
