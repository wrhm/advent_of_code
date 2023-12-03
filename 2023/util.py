def file_as_lines(filename):
    lines = []
    with open(filename) as fd:
        lines = fd.read().split('\n')

    lines = [x for x in lines if x != '']
    return lines


def is_digit(c):
    return '0' <= c <= '9'
