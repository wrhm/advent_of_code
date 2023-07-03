mod util;
mod year2015;

// To build iteratively on save:
// * cargo watch -x build
//
// Suggested pre-commit actions:
// * cargo build --verbose
// * cargo clippy --all-targets -- -D warnings
// * cargo test --verbose
//   * from ./ (src/) or ../
// * cargo run
//   * from ./ (src/)
fn main() {
    // Day 4 involves md5 hashing and is very slow.
    let mut days = vec![1, 2, 3];
    let mut later_days: Vec<i32> = (5..26).collect();
    days.append(&mut later_days);
    year2015::solve2015(days);
}
