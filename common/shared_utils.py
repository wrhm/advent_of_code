# shared_utils.py

def file_as_lines(filename):
    lines = []
    with open(filename) as fd:
        lines = fd.read().split('\n')
    if lines[-1] == '':
        return lines[::-1]
    return lines


def is_digit(c):
    return '0' <= c <= '9'
