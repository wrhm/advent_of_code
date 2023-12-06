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
    for (map_name, ranges) in maps:
        new_locs = []
        for item in locs:
            position = lookup_in_ranges(item, ranges)
            new_locs.append(position)
        locs = new_locs

    return min(locs)


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
