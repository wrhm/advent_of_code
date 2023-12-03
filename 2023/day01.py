# day01.py

import sys
import os
parent_dir = os.path.dirname(os.path.realpath(__file__))
sys.path.append(parent_dir)
sys.path.append(parent_dir+'/..')
import common.shared_utils as su

import re
import util

lines = su.file_as_lines('inputs/01/a.txt')

example1 = '''1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet'''.split('\n')

example2 = '''two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen'''.split('\n')


def digits(s):
    return [int(c) for c in s if su.is_digit(c)]


def calibration(s):
    d = digits(s)
    return 10*d[0]+d[-1]


def summed_calibrations(data):
    return sum([calibration(x) for x in data])


def solve_pt1(data):
    return summed_calibrations(data)


def find_digits_and_number_words(s):
    matches = re.finditer(
        r'(?=(one|two|three|four|five|six|seven|eight|nine|[0-9]))', s)
    return [m.group(1) for m in matches]


def a2i(s):
    nw = 'one,two,three,four,five,six,seven,eight,nine'.split(',')
    if s in nw:
        return nw.index(s)+1
    return int(s)


def calibration_v2(s):
    dnw = find_digits_and_number_words(s)
    d = list(map(a2i, dnw))
    return 10*d[0]+d[-1]


def summed_calibrations_v2(data):
    return sum([calibration_v2(x) for x in data])


def solve_pt2(data):
    return summed_calibrations_v2(data)


if __name__ == '__main__':
    print('ex1', solve_pt1(example1))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example2))
    print('part2', solve_pt2(lines))
