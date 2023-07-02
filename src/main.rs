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
    year2015::solve2015(vec![1, 2, 3]);
}
