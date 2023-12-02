# day01.py

import re

# To import ../util.py
if True:
    import sys
    import os
    current = os.path.dirname(os.path.realpath(__file__))
    parent = os.path.dirname(current)
    sys.path.append(parent)
    import util

lines = util.file_as_lines('a.txt')


def is_digit(c):
    return '0' <= c <= '9'


def digits(s):
    return [int(c) for c in s if is_digit(c)]


def calibration(s):
    d = digits(s)
    return 10*d[0]+d[-1]


def summed_calibrations(lines):
    return sum([calibration(x) for x in lines])


print(summed_calibrations(lines))


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


def summed_calibrations_v2(lines):
    return sum([calibration_v2(x) for x in lines])


print(summed_calibrations_v2(lines))
