# day09.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/09/a.txt')

example = '''0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45'''.split('\n')


def parse(data):
    seqs = []
    for line in data:
        seqs.append([int(x) for x in line.split()])
    return seqs


def az(seq):
    for x in seq:
        if x != 0:
            return False
    return True


def extrapolate(seq):
    iters = [seq]
    while not az(iters[-1]):
        prev = iters[-1]
        next = []
        for i in range(1, len(prev)):
            next.append(prev[i]-prev[i-1])
        iters.append(next)
    iters.reverse()
    iters[0].append(0)
    for i in range(1, len(iters)):
        iters[i].append(iters[i][-1]+iters[i-1][-1])
    return iters[-1][-1]


def solve_pt1(data):
    seqs = parse(data)
    ret = 0
    for s in seqs:
        ex = extrapolate(s)
        ret += ex
    return ret


def extrapolate2(seq):
    iters = [seq]
    while not az(iters[-1]):
        prev = iters[-1]
        next = []
        for i in range(1, len(prev)):
            next.append(prev[i]-prev[i-1])
        iters.append(next)
    iters.reverse()
    iters[0].append(0)
    for i in range(1, len(iters)):
        newv = iters[i][0]-iters[i-1][0]
        iters[i] = [newv]+iters[i]
    return iters[-1][0]


def solve_pt2(data):
    seqs = parse(data)
    ret = 0
    for s in seqs:
        ex = extrapolate2(s)
        ret += ex
    return ret


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
