use crate::util;

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
