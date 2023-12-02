'''
day01.py

On each line, the calibration value can be found by combining the first
digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines
are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
'''

import re
lines = []
with open('a.txt') as fd:
    lines = fd.read().split('\n')

lines = [x for x in lines if x != '']

example_lines = '''1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet'''.split('\n')

# print(len(lines))
# print(lines[0])
# print(lines[-1])


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

'''
--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are
actually spelled out with letters: one, two, three, four, five, six,
seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first
and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14,
and 76. Adding these together produces 281.

What is the sum of all of the calibration values?
'''

example_lines2 = '''two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen'''.split('\n')


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
