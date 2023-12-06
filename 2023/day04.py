# day04.py

import os
import sys

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/04/a.txt')

example = '''Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11'''.split('\n')


def id_and_winners(line):
    [cid, n] = line.split(':')
    cn = int(cid.split()[1])
    [w, h] = n.split('|')
    w = [int(x) for x in w.split()]
    h = [int(x) for x in h.split()]
    wn = [x for x in h if x in w]
    nwn = len(wn)
    return (cn, nwn)


def score(line):
    (_, nw) = id_and_winners(line)
    return 0 if nw == 0 else int(2**(nw-1))


def solve_pt1(data):
    sc = 0
    for line in data:
        sc += score(line)
    return sc


def solve_pt2(data):
    t = []
    copies = dict()
    for line in data:
        (cn, nw) = id_and_winners(line)
        t.append((cn, nw))
        copies[cn] = 1

    for (cn, nw) in t:
        for i in range(nw):
            copies[cn+i+1] += copies[cn]

    total = 0
    for k in copies:
        total += copies[k]

    return total


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
