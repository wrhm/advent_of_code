use std::fs::File;
use std::io::prelude::*;
use std::path::Path;

pub(crate) fn get_file_contents(filename: &str) -> String {
    let path = Path::new(filename);

    let mut file = match File::open(path) {
        Err(why) => panic!("Failed to open file at {}: {}", filename, why),
        Ok(file) => file,
    };
    let mut s = String::new();
    if let Err(why) = file.read_to_string(&mut s) {
        panic!("{}", why)
    }
    s
}
