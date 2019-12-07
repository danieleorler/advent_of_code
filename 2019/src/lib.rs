
use std::io::{self, prelude::*, BufReader};
use std::fs::File;

pub mod intcode;
pub mod geometry;

pub fn read_line() -> String {
    let mut line = String::new();
    let stdin = io::stdin();
    stdin.lock().read_line(&mut line).unwrap();
    return line.to_string();
}

pub fn read_line_from_file(file_name: String) -> io::Result<String> {
    let file = File::open(file_name)?;
    let reader = BufReader::new(file);

    return reader.lines().next().unwrap();
} 
