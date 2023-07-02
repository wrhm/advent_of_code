use crate::util;

use md5::compute;
use std::collections::HashSet;

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
    assert_eq!(solve_day00(""), (0, 0));
}

*/

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
        let slack = std::cmp::min(l * w, std::cmp::min(l * h, w * h));
        total_area += area + slack;

        let ribbon = std::cmp::min(2 * (l + w), std::cmp::min(2 * (l + h), 2 * (w + h)));
        let bow = l * w * h;
        total_ribbon += ribbon + bow;
        // println!("{:?}", dims[0]);
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
    let mut ans2;
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
            println!("{}, {}", i, h);
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
