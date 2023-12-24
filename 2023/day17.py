# day17.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/17/a.txt')

example = '''2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533'''.split('\n')


def bfs(grid):
    w, h = len(grid[0]), len(grid)
    q = [((0, 0), 0)]
    first_hl = None
    pred = dict()
    visited = set()
    for i in range(100000):
        print(i, len(q))
        if len(q) == 0:
            break
        ((r, c), heatloss), q = q[0], q[1:]
        if (r, c) == (h-1, w-1):
            first_hl = heatloss
            break
        nexts = []
        if c-1 >= 0:
            nexts.append((r, c-1))
        if c+1 < w:
            nexts.append((r, c+1))
        if r-1 >= 0:
            nexts.append((r-1, c))
        if r+1 < h:
            nexts.append((r+1, c))
        for (nc, nr) in nexts:
            pred[(nc, nr)] = (c, r)
            if (c, r) in pred and pred[(c, r)] == (nc, nr):
                # don't reverse direction
                continue
            # if (c,r) in pred
            candidate = ((nc, nr), heatloss+int(grid[r][c]))
            if candidate not in visited:
                q.append(candidate)
            visited.add(candidate)
    print(first_hl)


def solve_pt1(data):

    return ''


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
