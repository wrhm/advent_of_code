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


def disp(g):
    '''
    Print all rows of a grid, where each row is one string.
    '''
    for r in g:
        print(r)


def dispj(g):
    '''
    Print all rows of a grid, where each row is multiple strings.
    '''
    for r in g:
        print(''.join(r))
