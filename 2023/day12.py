# day12.py

# big help from
# https://www.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/

import functools

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/12/a.txt')

example = '''???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1'''.split('\n')


def parse_line(line):
    [pat, nums] = line.split()
    nums = [int(x) for x in nums.split(',')]
    return (pat, nums)


@functools.cache
def ways(pat, nums):
    if not nums:
        return '#' not in pat
    if not pat:
        return 0
    ch = pat[0]
    g = nums[0]

    def if_hash():
        grp = pat[:g]
        grp = grp.replace("?", "#")

        if grp != g * "#":
            return 0

        if len(pat) == g:
            return len(nums) == 1

        if pat[g] in ".?":
            return ways(pat[g+1:], nums[1:])
        return 0

    def if_dot():
        return ways(pat[1:], nums)

    ret = None
    if ch == '#':
        ret = if_hash()
    elif ch == '.':
        ret = if_dot()
    else:
        ret = if_dot()+if_hash()
    return ret


def solve_pt1(data):
    ret = 0
    for line in data:
        (p, n) = parse_line(line)
        w = ways(p, tuple(n))
        ret += w
    return ret


def solve_pt2(data):
    ret = 0
    for line in data:
        (p, n) = parse_line(line)
        p5 = '?'.join([p, p, p, p, p])
        w = ways(p5, tuple(n*5))
        ret += w
    return ret


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
