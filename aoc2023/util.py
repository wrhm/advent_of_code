def file_as_lines(filename):
    lines = []
    with open('a.txt') as fd:
        lines = fd.read().split('\n')

    lines = [x for x in lines if x != '']
    return lines
