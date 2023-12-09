# day05.py

import os
import sys

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/05/a.txt')

example = '''seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4'''.split('\n')


def parse_data(data):
    seeds = None
    maps = []
    map_name = None
    map_triples = []
    for line in data:
        if "seeds:" in line:
            seeds = [int(x) for x in line.split(':')[1].split()]
        elif "map:" in line:
            map_name = line.split()[0]
            map_triples = []
        elif line == "":
            if map_name is not None:
                maps.append((map_name, map_triples))
        else:
            [a, b, c] = [int(x) for x in line.split()]
            map_triples.append([a, b, c])
    if maps[-1][0] != map_name:
        maps.append((map_name, map_triples))
    return (seeds, maps)


def lookup_in_ranges(item, ranges):
    result = item
    for [dst, src, rlen] in ranges:
        if item in range(src, src+rlen):
            result = item-src+dst
            break
    return result


def solve_pt1(data):
    (seeds, maps) = parse_data(data)

    locs = seeds
    new_locs = []
    for (_, ranges) in maps:
        new_locs = []
        for item in locs:
            position = lookup_in_ranges(item, ranges)
            new_locs.append(position)
        locs = new_locs

    return min(locs)


def solve_pt2_brute(data):
    (seeds, maps) = parse_data(data)
    best = 1e20
    ns = len(seeds)
    for i in range(0, ns, 2):
        start, run = seeds[i], seeds[i+1]
        print(start, run)
        loc = None
        for j in range(start, start+run):
            if j % 10000 == 0:
                progress_pct = 100.*(j-start)/run
                print('range #%d of %d, progress: %.3f%%' %
                      (i, ns//2, progress_pct))
            loc = j
            for (_, ranges) in maps:
                loc = lookup_in_ranges(loc, ranges)
            best = min(best, loc)
    return best


'''
Collapsing mappings

Consider again the example seed-to-soil map:

50 98 2
52 50 48

The first line has a destination range start of 50,
a source range start of 98, and a range length of 2.

98 -> 50 start
99 -> 51 end

50 -> 52 start
51 -> 53
...
97 -> 99 end

seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51

seed -> soil
[0,49] -> [0,49] (self)
[50,97] -> [52,99]
[98,99] -> [50,51]
[100,inf] -> [100,inf] (self)

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

15 -> 0 start
15+37-1 = 51 -> 36 end

52 -> 37 start
53 -> 38 end

0 -> 39 start
0+15-1 = 14 -> 53 end


soil fertilizer
0    39
1    40
...
13   52
14   53
=
15   0
16   1
...
51   36
=
52   37
53   38

soil -> fertilizer
[0,14] -> [39,53] (+39)
[15,51] -> [0,36] (-15)
[52,53] -> [37,38] (-15)
[54,inf] -> [54,inf] (+0, self)

seed -> soil
[0,49] -> [0,49] (+0, self)
[50,97] -> [52,99] (+2)
[98,99] -> [50,51] (-48)
[100,inf] -> [100,inf] (+0, self)


Combining

seed     soil    fertilizer   net_delta   result
[0,14]   +0      +39          +39         [39,52]
[15,49]  +0      -15          -15         [0,34]
[50,51]  +2      -15          -13         [37,38]
[...]

'''


def range_w_delta(trio):
    [dst, src, run] = trio
    return ([src, src+run-1], dst-src)


# left and right are both lists, whose elements are each
# results from range_w_delta.
# DO NOT SUBMIT: actually want to apply deltas to left
# and somehow combine with right and its deltas.
def combine_rwds(left, right):
    intvs = []
    for (iv, d) in left:
        intvs.append(iv)
    for (iv, d) in right:
        intvs.append(iv)
    intvs.sort(key=lambda x: x[0])
    print(intvs)


def solve_pt2v2(data):
    # return solve_pt2_brute(data)
    (seeds, maps) = parse_data(data)

    # for (nm, ranges) in maps[:2]:
    #     print(nm, ranges)
    #     srtd_ranges = sorted(ranges, key=lambda x: x[1])
    #     # print('sorted', srtd_ranges)
    #     for trio in srtd_ranges:
    #         print('\t%s: %s' % (trio, range_w_delta(trio)))

    # rwd_list_1 = []
    # rwd_list_2 = []
    rwds0 = []
    (_, r0) = maps[0]
    for trio in sorted(r0, key=lambda x: x[1]):
        rwds0.append(range_w_delta(trio))
    rwds1 = []
    (_, r1) = maps[1]
    for trio in sorted(r1, key=lambda x: x[1]):
        rwds1.append(range_w_delta(trio))
    print(rwds0)
    print(rwds1)
    combine_rwds(rwds0, rwds1)

    return ''


'''
add [0,x] if missing

 0  1  2  3  4  5
+0 +0 +0 +0 +0 +0

[dest,src,range_len] --> [start,end,delta]

composing lists of [[start,end,delta], [start,end,delta], ...]

 0  1  2  3  4  5
+0 +0 +0 +0 +0 +0

add [start=1,end=3,delta=+2]

 0  1  2  3  4  5
+0 +0 +0 +0 +0 +0
   +2 +2 +2

[[0,5,0]] becomes [[0,0,0],[1,3,2],[4,5,0]]

add [start=2,end=4,delta=-1]

 0  1  2  3  4  5
+0 +0 +0 +0 +0 +0
   +2 +2 +2
      -1 -1 -1

[[0,0,0],[1,3,2],[4,5,0]] becomes
[[0,0,0],[1,1,2],[2,3,1],[4,4,-1],[5,5,0]]

'''


def dsr_to_sed(dsr):
    [dest, src, range_len] = dsr
    start = src
    end = src+range_len-1
    delta = dest-src
    return [start, end, delta]


'''
 0  1  2  3  4  5
+0 +0 +0 +0 +0 +0
   +2 +2 +2

[[1,3,2]] with [[0,5,0]] becomes
[[0,0,0],[1,3,2],[4,5,0]]

[  ]    []    [ ]     [ ]
 []    [  ]    [ ]   [ ]
'''


def combine_seds(a, b):
    ret = []
    # a.reverse()
    # b.reverse()
    while len(a) > 0 and len(b) > 0:
        print('a', a)
        print('b', b)
        print('ret', ret)
        fa, a = a[0], a[1:]
        fb, b = b[0], b[1:]
        print('fa', fa)
        print('fb', fb)
        [sa, ea, da] = fa
        [sb, eb, db] = fb
        if sa == sb and ea == eb:
            print('case1')
            ret.append([sa, ea, da+db])
        elif sa < sb and ea > eb:
            print('case2')
            ret.append([sa, sb-1, da])
            a = [[sb, ea, da]]+a
            b = [fb]+b
        elif sb < sa and eb > ea:
            print('case3')
            ret.append([sb, sa-1, db])
            b = [[sa, eb, db]]+b
            a = [fa]+a
        elif sa < sb and ea < eb:
            print('case4')
            ret.append([sa, sb-1, da])
            a = [[sb, ea, da]]+a
            b = [fb]+b
        elif sa > sb and ea > eb:
            print('case5')
            ret.append([sb, sa-1, db])
            b = [[sa, eb, db]]+b
            a = [fa]+a
        else:
            print('intervals share 1 boundary: ', fa, fb)
            return None
    print(ret)
    return ret


def solve_pt2(data):
    (seeds, maps) = parse_data(data)
    print(seeds)
    print(maps)

    for (label, maps_lst) in maps:
        print(label)
        seds = sorted([dsr_to_sed(x) for x in maps_lst])
        print(seds)
        if seds[0][0] != 0:
            seds = [[0, seds[0][0]-1, 0]]+seds
        print(seds)
    pass


if __name__ == '__main__':
    # print('ex1', solve_pt1(example))
    # print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))

    combine_seds([[0, 5, 0]], [[1, 3, 2]])
