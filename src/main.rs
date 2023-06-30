mod day_01;
mod util;

// Suggested pre-commit actions:
// * cargo build --verbose
// * cargo clippy --all-targets -- -D warnings
// * cargo test --verbose
//   * from ./ (src/) or ../
// * cargo run
//   * from ./ (src/)
fn main() {
    println!("Hello, world!");
    day_01::solve_for_file("../data/01.txt");
}
