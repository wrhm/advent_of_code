use crate::util;

use regex::Regex;
use std::cmp::max;
use std::cmp::min;
use std::collections::HashMap;
use std::collections::HashSet;
use topological_sort::TopologicalSort;

/* New day template

fn solve_day00(file_contents: &str) -> (i32, i32) {
    let lines: Vec<&str> = file_contents.split('\n').collect();
    let line = lines.first().unwrap();
    (0, 0)
}

pub(crate) fn solve_day00_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve_day00(&file_contents);
    println!("Day 00: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test_day00() {
    assert_eq!(solve_day00(""), (-1, -1));
}

*/

pub(crate) fn solve2015(days: Vec<i32>) {
    for d in days {
        match d {
            1 => solve_day01_for_file("../data/2015/01.txt"),
            2 => solve_day02_for_file("../data/2015/02.txt"),
            3 => solve_day03_for_file("../data/2015/03.txt"),
            4 => solve_day04(),
            5 => solve_day05_for_file("../data/2015/05.txt"),
            6 => solve_day06_for_file("../data/2015/06.txt"),
            7 => solve_day07_for_file("../data/2015/07.txt"),
            _ => println!("Day {} not solved yet.", d),
        }
    }
}

fn solve_day01(file_contents: &str) -> (i32, i32) {
    let lines: Vec<&str> = file_contents.split('\n').collect();
    let line = lines.first().unwrap();
    let mut flr = 0;
    let mut ind = -1;
    for (i, c) in line.chars().enumerate() {
        if c == '(' {
            flr += 1;
        } else {
            flr -= 1;
            if ind < 0 && flr < 0 {
                ind = 1 + i as i32;
            }
        }
    }
    (flr, ind)
}

pub(crate) fn solve_day01_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve_day01(&file_contents);
    println!("Day 01: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test_day01() {
    assert_eq!(solve_day01("(())"), (0, -1));
    assert_eq!(solve_day01(")"), (-1, 1));
    assert_eq!(solve_day01("()())"), (-1, 5));
}

// 20x3x11
// 15x27x5
// 6x29x7
// 30x15x9
// 19x29x21
// 10x4x15
// 1x26x4

// For example:

// A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
// A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
// All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?

fn solve_day02(file_contents: &str) -> (i32, i32) {
    let lines: Vec<&str> = file_contents.split('\n').collect();
    // let line = lines.first().unwrap();
    let mut total_area = 0;
    let mut total_ribbon = 0;
    for line in lines {
        if line.is_empty() {
            continue;
        }
        let dims: Vec<i32> = line.split('x').map(|x| x.parse::<i32>().unwrap()).collect();
        let l = dims[0];
        let w = dims[1];
        let h = dims[2];
        let area = 2 * (l * w + l * h + w * h);
        let slack = min(l * w, min(l * h, w * h));
        total_area += area + slack;

        let ribbon = min(2 * (l + w), min(2 * (l + h), 2 * (w + h)));
        let bow = l * w * h;
        total_ribbon += ribbon + bow;
    }
    (total_area, total_ribbon)
}

pub(crate) fn solve_day02_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve_day02(&file_contents);
    println!("Day 02: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test_day02() {
    assert_eq!(solve_day02("2x3x4"), (58, 34));
}

fn solve_day03(file_contents: &str) -> (i32, i32) {
    let lines: Vec<&str> = file_contents.split('\n').collect();
    let line = lines.first().unwrap();
    let mut path: HashSet<(i32, i32)> = HashSet::new();
    let mut row = 0;
    let mut col = 0;
    path.insert((row, col));
    let mut s_path: HashSet<(i32, i32)> = HashSet::new();
    let mut s_row = 0;
    let mut s_col = 0;
    s_path.insert((s_row, s_col));
    let mut rs_path: HashSet<(i32, i32)> = HashSet::new();
    let mut rs_row = 0;
    let mut rs_col = 0;
    rs_path.insert((rs_row, rs_col));
    for (ci, c) in line.chars().enumerate() {
        let mut dr = 0;
        let mut dc = 0;
        match c {
            '>' => dc = 1,
            '<' => dc = -1,
            '^' => dr = 1,
            _ => dr = -1,
        }
        row += dr;
        col += dc;
        path.insert((row, col));
        if ci % 2 == 0 {
            s_row += dr;
            s_col += dc;
            s_path.insert((s_row, s_col));
        } else {
            rs_row += dr;
            rs_col += dc;
            rs_path.insert((rs_row, rs_col));
        }
    }
    (path.len() as i32, s_path.union(&rs_path).count() as i32)
}

pub(crate) fn solve_day03_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve_day03(&file_contents);
    println!("Day 03: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test_day03() {
    assert_eq!(solve_day03("^v"), (2, 3));
    assert_eq!(solve_day03("^>v<"), (4, 3));
    assert_eq!(solve_day03("^v^v^v^v^v"), (2, 11));
}

pub(crate) fn solve_day04() {
    let prefix = "yzbqklnj";
    let mut i = 0;
    let mut best = 0;
    let mut ans1 = 0;
    let ans2;
    loop {
        let s = format!("{}{}", prefix, i);
        let b = s.as_bytes();
        let h = format!("{:?}", md5::compute(b));

        let mut leading_zeros = 0;
        for c in h.chars() {
            if c == '0' {
                leading_zeros += 1;
            } else {
                break;
            }
        }
        if leading_zeros > best {
            best = leading_zeros
        }
        if ans1 == 0 && best == 5 {
            ans1 = i;
        }
        if best == 6 {
            ans2 = i;
            break;
        }
        i += 1;
    }
    println!("Day 04: {:?}, {:?}", ans1, ans2);
}

fn contains_three_vowels(s: &str) -> bool {
    let mut vowels: HashSet<char> = HashSet::new();
    for v in "aeiou".chars() {
        vowels.insert(v);
    }
    let mut n = 0;
    for c in s.chars() {
        if vowels.contains(&c) {
            n += 1;
        }
    }
    n >= 3
}

fn contains_adjacent_repeat(s: &str) -> bool {
    let mut prev: char = '?';
    for (i, c) in s.chars().enumerate() {
        if i > 0 && c == prev {
            return true;
        }
        prev = c;
    }
    false
}

fn contains_bad_pair(s: &str) -> bool {
    let mut prev: char;
    let mut curr: char = '?';
    for c in s.chars() {
        prev = curr;
        curr = c;
        match (prev, curr) {
            ('a', 'b') => return true,
            ('c', 'd') => return true,
            ('p', 'q') => return true,
            ('x', 'y') => return true,
            _ => (),
        }
    }
    false
}

fn is_nice_string_v1(s: &str) -> bool {
    contains_three_vowels(s) && contains_adjacent_repeat(s) && !contains_bad_pair(s)
}

fn has_repeated_pair(s: &str) -> bool {
    for i in 0..s.len() - 3 {
        for j in (i + 2)..s.len() - 1 {
            if s.chars().nth(i) == s.chars().nth(j) && s.chars().nth(i + 1) == s.chars().nth(j + 1)
            {
                return true;
            }
        }
    }
    false
}

fn has_sandwich(s: &str) -> bool {
    for i in 0..s.len() - 2 {
        if s.chars().nth(i) == s.chars().nth(i + 2) {
            return true;
        }
    }
    false
}

fn is_nice_string_v2(s: &str) -> bool {
    has_repeated_pair(s) && has_sandwich(s)
}

fn solve_day05(file_contents: &str) -> (i32, i32) {
    let lines: Vec<&str> = file_contents.split('\n').collect();
    let mut ans1 = 0;
    let mut ans2 = 0;
    for line in lines {
        if line.is_empty() {
            continue;
        }
        if is_nice_string_v1(line) {
            ans1 += 1;
        }
        if is_nice_string_v2(line) {
            ans2 += 1;
        }
    }
    (ans1, ans2)
}

pub(crate) fn solve_day05_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve_day05(&file_contents);
    println!("Day 05: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test_day05() {
    assert!(!has_repeated_pair("aaa"));
    assert!(!has_repeated_pair("aabb"));
    assert!(has_repeated_pair("abab"));
    assert_eq!(solve_day05("ugknbfddgicrmopn").0, 1);
    assert_eq!(solve_day05("jchzalrnumimnmhp").0, 0);
    assert_eq!(solve_day05("qjhvhtzxzqqjkmpb").1, 1);
}

/*
For example:

turn on 0,0 through 999,999 would turn on (or leave on) every light.
toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.
After following the instructions, how many lights are lit?
 */

/*
You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

turn on 0,0 through 0,0 would increase the total brightness by 1.
toggle 0,0 through 999,999 would increase the total brightness by 2000000.
 */

fn solve_day06(file_contents: &str) -> (i32, i32) {
    let lines: Vec<&str> = file_contents.split('\n').collect();

    let mut lights: Vec<Vec<bool>> = vec![vec![false; 1000]; 1000];
    let mut brightness: Vec<Vec<i32>> = vec![vec![0; 1000]; 1000];

    for line in lines {
        if line.is_empty() {
            continue;
        }
        let words: Vec<&str> = line.split_whitespace().collect();
        let mut turn_off = false;
        let mut turn_on = false;
        match (words[0], words[1]) {
            ("turn", "off") => turn_off = true,
            ("turn", "on") => turn_on = true,
            ("toggle", _) => (),
            (_, _) => panic!(),
        }
        let re = Regex::new(r"\d+").unwrap();
        let nums = re
            .find_iter(line)
            .map(|m| m.as_str().parse::<i32>().unwrap())
            .collect::<Vec<i32>>();
        let [r1, c1, r2, c2] = nums[..] else {panic!()};
        let rlo = min(r1, r2);
        let rhi = max(r1, r2);
        let clo = min(c1, c2);
        let chi = max(c1, c2);
        for r in rlo..rhi + 1 {
            for c in clo..chi + 1 {
                let old_val = lights[r as usize][c as usize];
                let old_brightness = brightness[r as usize][c as usize];
                let new_val: bool;
                let new_brightness: i32;
                //                 The phrase turn on actually means that you should increase the brightness of those lights by 1.

                // The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

                // The phrase toggle actually means that you should increase the brightness of those lights by 2.
                if turn_off {
                    new_val = false;
                    new_brightness = max(old_brightness - 1, 0);
                } else if turn_on {
                    new_val = true;
                    new_brightness = old_brightness + 1;
                } else {
                    new_val = !old_val;
                    new_brightness = old_brightness + 2;
                }
                lights[r as usize][c as usize] = new_val;
                brightness[r as usize][c as usize] = new_brightness;
            }
        }
    }
    let ans1 = lights
        .iter()
        .map(|row| row.iter().filter(|&&x| x).count())
        .reduce(|acc, e| acc + e)
        .unwrap();

    let ans2: i32 = brightness.iter().map(|x| x.iter().sum::<i32>()).sum();
    (ans1 as i32, ans2)
}

pub(crate) fn solve_day06_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve_day06(&file_contents);
    println!("Day 06: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test_day06() {
    assert_eq!(solve_day06("toggle 0,0 through 999,999").0, 1_000_000);
    assert_eq!(solve_day06("toggle 0,0 through 999,0").0, 1000);
    assert_eq!(solve_day06("turn on 0,0 through 9,9").0, 100);
    assert_eq!(
        solve_day06("turn on 0,0 through 9,9\nturn off 0,0 through 0,3").0,
        96
    );
    assert_eq!(
        solve_day06("turn on 0,0 through 9,9\nturn off 0,0 through 0,3\ntoggle 9,0 through 9,5").0,
        90
    );
    assert_eq!(solve_day06("turn on 499,499 through 500,501").0, 6);
    assert_eq!(solve_day06("turn on 0,0 through 0,0").1, 1);
    assert_eq!(solve_day06("toggle 0,0 through 999,999").1, 2_000_000);
}

// 123 -> x             [0]
// 456 -> y             [1]
// x AND y -> d         [2]
// x OR y -> e          [3]
// x LSHIFT 2 -> f      [4]
// y RSHIFT 2 -> g      [5]
// NOT x -> h           [6]
// NOT y -> i           [7]
/*
2 requires 0 and 1
3 requires 0 and 1
4 requires 0
5 requires 1
6 requires 0
7 requires 1
*/
fn simulate_wires(lines: &Vec<&str>, var: &str) -> i32 {
    // Stores variables values once they are known.
    let mut hm_var: HashMap<&str, u16> = HashMap::new();
    // Tracks variables' dependencies on other values.
    let mut ts = TopologicalSort::<&str>::new();
    // Tracks the needed inputs and operators from expressions.
    // (outname, (var/op, var/op, var/op))
    let mut hm_deps: HashMap<&str, (&str, &str, &str)> = HashMap::new();
    for line in lines {
        if line.is_empty() {
            continue;
        }
        let words: Vec<&str> = line.split_whitespace().collect();
        if words.len() == 3 {
            // x -> y
            let inname = words[0];
            let outname = words[2];
            ts.insert(outname);
            ts.add_dependency(inname, outname);
            hm_deps.insert(outname, (inname, "", ""));
        } else if words.len() == 4 {
            // NOT x -> y
            let inname = words[1];
            let outname = words[3];
            ts.add_dependency(inname, outname);
            hm_deps.insert(outname, ("NOT", inname, ""));
        } else {
            // x (AND / OR / LSHIFT / RSHIFT) y -> z
            let inname0 = words[0];
            let opname = words[1];
            let inname1 = words[2];
            let outname = words[4];
            ts.add_dependency(inname0, outname);
            ts.add_dependency(inname1, outname);
            hm_deps.insert(outname, (inname0, opname, inname1));
        }
    }
    while !ts.is_empty() {
        let elt = ts.pop().unwrap();
        match hm_deps.get(elt) {
            Some((x, y, z)) => {
                let result: u16 = match (x, y, z) {
                    (_, &"", &"") => *hm_var.get(x).unwrap(),
                    (&"NOT", _, _) => !hm_var.get(y).unwrap(),
                    (_, &"AND", _) => hm_var.get(x).unwrap() & hm_var.get(z).unwrap(),
                    (_, &"OR", _) => hm_var.get(x).unwrap() | hm_var.get(z).unwrap(),
                    (_, &"LSHIFT", _) => hm_var.get(x).unwrap() << hm_var.get(z).unwrap(),
                    (_, &"RSHIFT", _) => hm_var.get(x).unwrap() >> hm_var.get(z).unwrap(),
                    _ => panic!(),
                };
                hm_var.insert(elt, result);
            }
            None => {
                hm_var.insert(elt, elt.parse::<u16>().unwrap());
            }
        }
    }
    *hm_var.get(var).unwrap() as i32
}

fn solve_day07(file_contents: &str, var: &str) -> (i32, i32) {
    let mut lines: Vec<&str> = file_contents.split('\n').collect();
    let ans1 = simulate_wires(&lines, var);
    let new_rule = format!("{} -> b", ans1);
    lines.push(&new_rule);
    let ans2 = simulate_wires(&lines, var);
    (ans1, ans2)
}

pub(crate) fn solve_day07_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve_day07(&file_contents, "a");
    println!("Day 07: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test_day07() {
    assert_eq!(
        vec!["d", "e", "h", "i"]
            .iter()
            .map(|v| solve_day07(
                "123 -> x
456 -> y
y -> a
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i",
                v
            )
            .0)
            .collect::<Vec<i32>>(),
        [72, 507, 65412, 65079]
    );
}
