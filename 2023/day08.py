# day08.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

import math

lines = su.file_as_lines('inputs/08/a.txt')

example = '''LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)'''.split('\n')


def parse(data):
    path = data[0]
    d = dict()
    for line in data[2:]:
        lhs = line[:3]
        next = [line[7:10], line[12:15]]
        d[lhs] = next
    return (path, d)


def solve_pt1(data):
    (path, d) = parse(data)
    pos = 'AAA'
    for i in range(1000000):
        inst = path[i % len(path)]
        ind = 0 if inst == 'L' else 1
        pos = d[pos][ind]
        if pos == 'ZZZ':
            return i+1
    return None


example2 = '''LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)'''.split('\n')


def all_zs(xs):
    for x in xs:
        if x[-1] != 'Z':
            return False
    return True


def solve_pt2_bf(data):
    (path, d) = parse(data)
    locs = [k for k in d if k[2] == 'A']
    for i in range(1000000):
        inst = path[i % len(path)]
        ind = 0 if inst == 'L' else 1
        for j in range(len(locs)):
            locs[j] = d[locs[j]][ind]
        if all_zs(locs):
            return i+1

    return None


def lcm_of_list(xs):
    if len(xs) == 0:
        return None
    ret = 1
    for x in xs:
        ret = math.lcm(ret, x)
    return ret


def solve_pt2(data):
    (path, d) = parse(data)
    locs = [k for k in d if k[2] == 'A']
    to_lcm = []
    for x in locs:
        steps = None
        for i in range(1000000):
            inst = path[i % len(path)]
            ind = 0 if inst == 'L' else 1
            x = d[x][ind]
            if x[-1] == 'Z':
                steps = i+1
                to_lcm.append(steps)
                break
        if steps is None:
            return None

    return lcm_of_list(to_lcm)


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example2))
    print('part2', solve_pt2(lines))
