use std::io;
use std::io::prelude::*;
use aoc2019::intcode::{parse_code, IntProgram};

fn main() {
    let mut code = String::new();
    let stdin = io::stdin();
    stdin.lock().read_line(&mut code).unwrap();

    let mut program = IntProgram::new(parse_code(&code));
    program.execute();
}