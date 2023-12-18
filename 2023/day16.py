# day16.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su


lines = su.file_as_lines('inputs/16/a.txt')

# example = '''.|...\....
# |.-.\.....
# .....|-...
# ........|.
# ..........
# .........\
# ..../.\\..
# .-.-/..|..
# .|....-|.\
# ..//.|....'''.split('\n')

example = su.file_as_lines('inputs/16/ex1.txt')


def energize_until_done(data, first_beam=(0, 0, 1, 0), max_steps=1000):
    data = [[x for x in y] for y in data]
    w, h = len(data[0]), len(data)
    energized = set()
    # (c,r,dc,dr)
    # beam = (0, 0, 1, 0)
    # c, r, dc, dr = 0, 0, 1, 0
    # beams = [(0, 0, 1, 0)]
    beams = set()
    beams.add(first_beam)
    time_since_change = 0
    max_steps = 1000
    for step in range(max_steps):
        time_since_change += 1
        next_beams = set()
        old_energized = len(energized)
        for beam in beams:
            (c, r, dc, dr) = beam
            if not (0 <= r < h and 0 <= c < w):
                continue
            # print(c, r, dc, dr, 'sees', data[r][c])
            # print('time since change', time_since_change)

            # if (c, r) not in energized:
            #     print('newly energized', c, r)
            #     time_since_change = 0
            # elif time_since_change >= 500:
            #     print('probably done')
            #     step = max_steps
            #     break
            energized.add((c, r))
            ch = data[r][c]
            if ch == '.':
                # print('motion unchanged')
                next_beams.add((c+dc, r+dr, dc, dr))
            elif ch == '\\':
                '''
                right (1,0) -> (0,1) down
                down (0,1) -> (1,0) right
                left (-1,0) -> (0,-1) up
                up (0,-1) -> (-1,0) left
                '''
                # print('reflecting off \\')
                dc, dr = dr, dc
                next_beams.add((c+dc, r+dr, dc, dr))
            elif ch == '/':
                '''
                right (1,0) -> (0,-1) up
                up (0,-1) -> (1,0) right
                left (-1,0) -> (0,1) down
                down (0,1) -> (-1,0) left
                '''
                # print('reflecting off /')
                dc, dr = -dr, -dc
                next_beams.add((c+dc, r+dr, dc, dr))
            elif ch == '-':
                if dr == 0:
                    # print('aligned with -')
                    next_beams.add((c+dc, r+dr, dc, dr))
                else:
                    # print('splitting at -')
                    # next_beams.append((c, r, 1, 0))
                    # next_beams.append((c, r, -1, 0))
                    next_beams.add((c, r, 1, 0))
                    next_beams.add((c, r, -1, 0))
            elif ch == '|':
                if dc == 0:
                    # print('aligned with |')
                    next_beams.add((c+dc, r+dr, dc, dr))
                else:
                    # print('splitting at |')
                    # next_beams.append((c, r, -1, 0))
                    # next_beams.append((c, r, 1, 0))
                    next_beams.add((c, r, 0, -1))
                    next_beams.add((c, r, 0, 1))
            # print('dc', dc, 'dr', dr)
            # c += dc
            # r += dr
            # if 0 <= c < w and 0 <= r < h:
            #     print('adding', c, r, dc, dr, 'to set')
            #     next_beams.add((c, r, dc, dr))
        # beams = [x for x in next_beams]
        # if len(next_beams) == len(beams):
        #     time_since_change += 1
        #     if time_since_change == 100:
        #         print('probably done')
        #         break
        beams = set()
        for x in next_beams:
            beams.add(x)
        # print(beams)
        if len(energized) == old_energized:
            time_since_change += 1
            if time_since_change > 100:
                print('probably done')
                break
        else:
            time_since_change = 0
        nrg = [['.' for y in x] for x in data]
        for (c, r) in energized:
            nrg[r][c] = '#'
            # if (c, r) in beams:
            #     nrg[r][c] = 'B'
        # print('=')
        # su.dispj(nrg)
        # print('after step', step, ': beams', len(
        #     beams), 'energized', len(energized), 't_s_c', time_since_change)

    return len(energized)


def solve_pt1(data):
    # data = [[x for x in y] for y in data]
    # w, h = len(data[0]), len(data)
    # energized = set()
    # # (c,r,dc,dr)
    # # beam = (0, 0, 1, 0)
    # # c, r, dc, dr = 0, 0, 1, 0
    # # beams = [(0, 0, 1, 0)]
    # beams = set()
    # beams.add((0, 0, 1, 0))
    # time_since_change = 0
    # max_steps = 1000
    # for step in range(max_steps):
    #     time_since_change += 1
    #     next_beams = set()
    #     old_energized = len(energized)
    #     for beam in beams:
    #         (c, r, dc, dr) = beam
    #         if not (0 <= r < h and 0 <= c < w):
    #             continue
    #         # print(c, r, dc, dr, 'sees', data[r][c])
    #         # print('time since change', time_since_change)

    #         # if (c, r) not in energized:
    #         #     print('newly energized', c, r)
    #         #     time_since_change = 0
    #         # elif time_since_change >= 500:
    #         #     print('probably done')
    #         #     step = max_steps
    #         #     break
    #         energized.add((c, r))
    #         ch = data[r][c]
    #         if ch == '.':
    #             # print('motion unchanged')
    #             next_beams.add((c+dc, r+dr, dc, dr))
    #         elif ch == '\\':
    #             '''
    #             right (1,0) -> (0,1) down
    #             down (0,1) -> (1,0) right
    #             left (-1,0) -> (0,-1) up
    #             up (0,-1) -> (-1,0) left
    #             '''
    #             # print('reflecting off \\')
    #             dc, dr = dr, dc
    #             next_beams.add((c+dc, r+dr, dc, dr))
    #         elif ch == '/':
    #             '''
    #             right (1,0) -> (0,-1) up
    #             up (0,-1) -> (1,0) right
    #             left (-1,0) -> (0,1) down
    #             down (0,1) -> (-1,0) left
    #             '''
    #             # print('reflecting off /')
    #             dc, dr = -dr, -dc
    #             next_beams.add((c+dc, r+dr, dc, dr))
    #         elif ch == '-':
    #             if dr == 0:
    #                 # print('aligned with -')
    #                 next_beams.add((c+dc, r+dr, dc, dr))
    #             else:
    #                 # print('splitting at -')
    #                 # next_beams.append((c, r, 1, 0))
    #                 # next_beams.append((c, r, -1, 0))
    #                 next_beams.add((c, r, 1, 0))
    #                 next_beams.add((c, r, -1, 0))
    #         elif ch == '|':
    #             if dc == 0:
    #                 # print('aligned with |')
    #                 next_beams.add((c+dc, r+dr, dc, dr))
    #             else:
    #                 # print('splitting at |')
    #                 # next_beams.append((c, r, -1, 0))
    #                 # next_beams.append((c, r, 1, 0))
    #                 next_beams.add((c, r, 0, -1))
    #                 next_beams.add((c, r, 0, 1))
    #         # print('dc', dc, 'dr', dr)
    #         # c += dc
    #         # r += dr
    #         # if 0 <= c < w and 0 <= r < h:
    #         #     print('adding', c, r, dc, dr, 'to set')
    #         #     next_beams.add((c, r, dc, dr))
    #     # beams = [x for x in next_beams]
    #     # if len(next_beams) == len(beams):
    #     #     time_since_change += 1
    #     #     if time_since_change == 100:
    #     #         print('probably done')
    #     #         break
    #     beams = set()
    #     for x in next_beams:
    #         beams.add(x)
    #     # print(beams)
    #     if len(energized) == old_energized:
    #         time_since_change += 1
    #         if time_since_change > 100:
    #             print('probably done')
    #             break
    #     else:
    #         time_since_change = 0
    #     nrg = [['.' for y in x] for x in data]
    #     for (c, r) in energized:
    #         nrg[r][c] = '#'
    #         # if (c, r) in beams:
    #         #     nrg[r][c] = 'B'
    #     # print('=')
    #     # su.dispj(nrg)
    #     print('after step', step, ': beams', len(
    #         beams), 'energized', len(energized), 't_s_c', time_since_change)

    # return len(energized)
    return energize_until_done(data, first_beam=(0, 0, 1, 0), max_steps=1000)


def solve_pt2(data):
    w, h = len(data[0]), len(data)
    best = 0
    for r in range(0, h):
        best = max(best, energize_until_done(data, (0, r, 1, 0)))
        print('row', r, 'of', h, ' (from left): ', best)
    for r in range(0, h):
        best = max(best, energize_until_done(data, (w-1, r, -1, 0)))
        print('row', r, 'of', h, '(from right): ', best)
    for c in range(0, w):
        best = max(best, energize_until_done(data, (c, 0, 0, 1)))
        print('col', c, 'of', w, ' (from top): ', best)
    for c in range(0, w):
        best = max(best, energize_until_done(data, (c, h-1, 0, -1)))
        print('col', c, 'of', w, ' (from bottom): ', best)
    return best


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    print('ex2', solve_pt2(example))
    # slow, takes about 5 min.
    print('part2', solve_pt2(lines))
