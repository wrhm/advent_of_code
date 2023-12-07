# day00.py

import functools
import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/07/a.txt')

example = '''32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483'''.split('\n')


def parse(data):
    resp = []
    for line in data:
        [h, b] = line.split()
        resp.append((h, int(b)))
    return resp


def counts(h):
    d = dict()
    for c in h:
        if c in d:
            d[c] += 1
        else:
            d[c] = 1
    return d


def vs(d):
    return sorted([d[k] for k in d])


def hand_type(h):
    v = vs(counts(h))
    r = {(5,): '5oak', (1, 4): '4oak', (2, 3): 'fh', (1, 1, 3): '3oak',
         (1, 2, 2): '2p', (1, 1, 1, 2): '1p', (1, 1, 1, 1, 1): 'hc'}
    return r[tuple(v)]


def cmp_card(a, b):
    if a == b:
        return 0
    cds = '23456789TJQKA'
    return 1 if cds.index(a) > cds.index(b) else -1


def cmp_hand(a, b):
    htypes = 'hc 1p 2p 3oak fh 4oak 5oak'.split()
    aty = htypes.index(hand_type(a))
    bty = htypes.index(hand_type(b))
    if aty > bty:
        return 1
    if aty < bty:
        return -1
    for i in range(5):
        ac, bc = a[i], b[i]
        cc = cmp_card(ac, bc)
        if cc != 0:
            return cc
    return 0


def solve_pt1(data):
    hbs = parse(data)
    bids = dict()
    for (h, b) in hbs:
        bids[h] = b
    srtd = sorted([h for (h, _) in hbs], key=functools.cmp_to_key(cmp_hand))
    score = 0
    for i in range(len(srtd)):
        score += (i+1)*bids[srtd[i]]
    return score


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
