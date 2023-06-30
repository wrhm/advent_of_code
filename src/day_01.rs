use crate::util;

fn solve(file_contents: &str) -> (i32, i32) {
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

pub(crate) fn solve_for_file(filename: &str) {
    let file_contents = util::get_file_contents(filename);
    let (ans1, ans2) = solve(&file_contents);
    println!("Day 01: {:?}, {:?}", ans1, ans2);
}

#[test]
fn unit_test() {
    assert_eq!(solve("(())"), (0, -1));
    assert_eq!(solve(")"), (-1, 1));
    assert_eq!(solve("()())"), (-1, 5));
}
