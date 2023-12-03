import sys
import os
parent_dir = os.path.dirname(os.path.realpath(__file__))
sys.path.append(parent_dir)
sys.path.append(parent_dir+'/..')
import common.shared_utils as su


def file_as_lines(filename):
    lines = []
    with open(filename) as fd:
        lines = fd.read().split('\n')

    lines = [x for x in lines if x != '']
    return lines


def is_digit(c):
    return '0' <= c <= '9'
