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
    year2015::solve_day01_for_file("../data/2015_01.txt");
}
