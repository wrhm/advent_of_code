# day12.py

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

# def agree()


# def ways_to_place(total, nums):
#     # print('\nwtp', total, nums)
#     if len(nums) == 0:
#         # return 1 if total == 0 else 0
#         return 0
#     x = nums[0]
#     # print('x', x)
#     rest_need = sum(nums[1:])+len(nums)-1
#     # print('rest_need', rest_need)
#     if rest_need > total-x:
#         # print('no room')
#         return 0
#     # ret = 1
#     ret = 0
#     print('block of size', x, 'fits', total-x+1, 'ways in gap of size', total)
#     for i in range(total-x+1):
#         # ret += 1
#         print('recursing for ', total-x-i, nums[1:])
#         # rec = ways_to_place(total-x-i, nums[1:])
#         # print('rec for ', total-x-i, nums[1:], 'is', rec)
#         rec = ways_to_place(total-x-i, nums[1:])
#         print('rec for ', total-x-i, nums[1:], 'is', rec)
#         ret += rec
#     print('returning', ret, 'for', total, nums)
#     return ret

def ways_to_place_one(gap_size, block_size):
    return max(0, gap_size-block_size+1)


'''
3, [1,1]

X _ X [0,2]

5 [1,2]

X _ X X _ [0,2] because 3 [2] is possible
X _ _ X X [0,3] because 2 [2] is possible
_ X _ X X [1,3] because 2 [2] is possible
'''


def ways_to_place(gap_size, blocks, pos=0):
    if blocks == []:
        return []
    # if len(blocks) == 1:
    #     if ways_to_place_one(gap_size, blocks[0]) == 0:
    #         return []
    #     return [pos]
    # if len(blocks)==1
    b = blocks[0]
    ret = []
    for i in range(ways_to_place_one(gap_size, b)):
        rec = ways_to_place(gap_size-b-1, blocks[1:], pos+b+i+1)
        # print('block size', b, 'at', 'i=', i, 'possible with', rec)
        if rec != []:
            ret.append([i]+rec)
    return ret


# def can_place_blocks(gap_size, blocks):
#     if blocks == []:
#         return True
#     if gap_size == 0:
#         return False
#     if sum(blocks)+len(blocks)-1 > gap_size:
#         return False
#     b = blocks[0]
#     return can_place_blocks(gap_size-b-1, blocks[1:])


# def ways_to_place_blocks(gap_size, blocks, pos=0):
#     print('called wtpb', gap_size, blocks, pos)
#     if blocks == []:
#         print('possible to place zero blocks anywhere')
#         return [pos]
#     if gap_size == 0:
#         print('nothing fits in empty gap')
#         return []
#     if sum(blocks)+len(blocks)-1 > gap_size:
#         print('not enough room for remaining blocks and gaps')
#         return []
#     b = blocks[0]
#     # return [pos]+ways_to_place_blocks(gap_size-b-1, blocks[1:], pos+b+1)
#     # rec = ways_to_place_blocks(gap_size-b-1, blocks[1:], pos+b+1)
#     # return [[pos]+x for x in rec] if rec is not None else None
#     ret = []
#     for i in range(max(0, gap_size-b)):
#         print('could place', b, 'at', pos+i)
#         rec = ways_to_place_blocks(gap_size-b-1, blocks[1:], pos+b+1)
#         print('combined:', i, rec)
#         if rec != []:
#             ret.append(i)
#     return ret


def solve_pt1(data):
    return ''


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    # print('ex1', solve_pt1(example))
    # print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
    # print('ans', ways_to_place(3, [1, 1]))
    pass
