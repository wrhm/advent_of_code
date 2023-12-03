import day01
import day02
import day03

assert day01.summed_calibrations(day01.example1) == 142
assert day01.summed_calibrations_v2(day01.example2) == 281

assert sum([day02.is_possible(line) for line in day02.example]) == 8
assert sum([day02.find_power(line) for line in day02.example]) == 2286

assert day03.solve_pt1(day03.example) == 4361
assert day03.solve_pt2(day03.example) == 467835
