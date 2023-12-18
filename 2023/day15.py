# day15.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/15/a.txt')

example = '''rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7'''.split('\n')


def hash(x):
    ret = 0
    for c in x:
        ret = (17*(ret+ord(c))) % 256
    return ret


def solve_pt1(data):
    ret = 0
    for x in data[0].split(','):
        ret += hash(x)
    return ret


def solve_pt2(data):
    # k: box
    # v:
    #   k: label
    #   v: (focal_len, time)
    d = dict()
    for i, x in enumerate(data[0].split(',')):
        if x[-1] == '-':
            label = x[:-1]
            box = hash(label)
            if box not in d:
                continue
            hm = d[box]
            if label in hm:
                hm[label] = None
        else:
            [label, focal_len] = x.split('=')
            h = hash(label)
            if h not in d:
                d[h] = dict()
            hm = d[h]
            if label in hm and hm[label] is not None:
                (fl, ind) = hm[label]
                hm[label] = (focal_len, ind)
            else:
                hm[label] = (focal_len, i)

    ret = 0
    for box in range(256):
        if box not in d:
            continue
        labels = []
        for k in d[box]:
            v = d[box][k]
            if v is not None:
                labels.append((k, v[0], v[1]))
        labels = sorted(labels, key=lambda x: x[2])
        labels = [(x[0], int(x[1])) for x in labels]
        for j, (_, fl) in enumerate(labels):
            fp = (1+box)*(1+j)*fl
            ret += fp

    return ret


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
