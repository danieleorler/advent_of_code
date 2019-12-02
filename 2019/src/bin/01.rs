use std::io;
use std::io::prelude::*;

fn main() {
    let stdin = io::stdin();
    let mut solution_one = 0;
    let mut solution_two = 0;
    for line in stdin.lock().lines() {
        let i = line.unwrap().parse::<i32>().unwrap();
        solution_one = solution_one + (i / 3 as i32) - 2;
        solution_two = solution_two + calculate_fuel(i);
    }
    println!("soulution one: {}", solution_one);
    println!("soulution two: {}", solution_two);
}

fn calculate_fuel(x: i32) -> i32 {
    let mut a = (x / 3 as i32) - 2;
    let mut sum = a;
    while a > 0 {
        a = (a / 3 as i32) - 2;
        if a > 0 {
            sum = sum + a
        }
    }
    sum
}