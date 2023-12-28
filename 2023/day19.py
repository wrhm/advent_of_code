# day19.py

import sys
import os

if True:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    sys.path.append(parent_dir)
    sys.path.append(parent_dir+'/..')
    import common.shared_utils as su

lines = su.file_as_lines('inputs/19/a.txt')

example = '''px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}'''.split('\n')


def parse_wf(line):
    default = None
    [name, rest] = line.split('{')
    rules = []
    for part in rest.split(','):
        if ':' not in part:
            default = part[:-1]
            continue
        [cmp, next] = part.split(':')
        rules.append((cmp, next))
    return (name, rules, default)


def parse_part(line):
    d = dict()
    ps = line.split(',')
    for p in ps:
        [k, v] = p.split('=')
        if '{' in k:
            k = k[1:]
        if '}' in v:
            v = v[:-1]
        d[k] = int(v)
    return d


def parse_lines(data):
    parsing_wfs = True
    workflows = dict()
    parts = []
    for line in data:
        if line == '':
            parsing_wfs = False
            continue
        if parsing_wfs:
            # [wfn,rest] = line.split('{')
            # workflows.append(parse_wf(line))
            (name, rules, default) = parse_wf(line)
            workflows[name] = (rules, default)
        else:
            parts.append(parse_part(line))
    return (workflows, parts)


# def iterate_part(p, wfs):
#     k = 'in'
#     while True:
#         if k not in wfs:
#             return None

def applies(cmp, p):
    if '<' in cmp:
        field, thresh = cmp.split('<')
        return p[field] < int(thresh)
    else:
        field, thresh = cmp.split('>')
        thresh = int(thresh)
        return p[field] > int(thresh)


def lookup(wf, p):
    (cmps, default) = wf
    for (cmp, nxt) in cmps:
        if applies(cmp, p):
            return nxt
    return default


def solve_pt1(data):
    (wfs, ps) = parse_lines(data)
    for x in ps:
        print(x)
    print('===')
    for k in wfs:
        print(k, wfs[k])
    print('===')
    result = 0
    for p in ps:
        k = 'in'
        # print(p, lookup(wfs[k], p))
        while k not in ['A', 'R']:
            k = lookup(wfs[k], p)
        print(p, 'finishes at', k)
        if k == 'A':
            for c in 'xmas':
                result += p[c]

    return result


def solve_pt2(data):
    return ''


if __name__ == '__main__':
    print('ex1', solve_pt1(example))
    print('part1', solve_pt1(lines))
    # print('ex2', solve_pt2(example))
    # print('part2', solve_pt2(lines))
