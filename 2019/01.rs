use std::io;
use std::io::prelude::*;

fn main() {
    run()
}

fn run() {
    let stdin = io::stdin();
    let mut sum = 0;
    for line in stdin.lock().lines() {
        let i = line.unwrap().parse::<i32>().unwrap();
        sum = sum + calculate_fuel(i);
    }
    println!("result: {}", sum);
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