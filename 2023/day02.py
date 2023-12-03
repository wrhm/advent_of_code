# day02.py

import util

lines = util.file_as_lines('inputs/02/a.txt')

example = '''Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green'''.split('\n')


def gn_rgb(line):
    [lhs, rhs] = line.split(':')
    gn = lhs.split(' ')[1]
    grabs = rhs.split(';')
    d = {'red': 0, 'green': 0, 'blue': 0}
    for pg in grabs:
        vks = [x.split() for x in pg.split(',')]
        for [v, k] in vks:
            d[k] = max(d[k], int(v))
    r, g, b = d['red'], d['green'], d['blue']
    return (gn, r, g, b)


def is_possible(line):
    (gn, r, g, b) = gn_rgb(line)
    possible = r <= 12 and g <= 13 and b <= 14
    return int(gn) if possible else 0


def solve_pt1(data):
    return sum([is_possible(line) for line in data])


def find_power(line):
    (gn, r, g, b) = gn_rgb(line)
    return r*g*b


def solve_pt2(data):
    return sum([find_power(line) for line in data])


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    print('part2', solve_pt2(lines))
