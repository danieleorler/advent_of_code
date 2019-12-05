use std::io;
use std::io::prelude::*;
use std::sync::mpsc;
use threadpool::ThreadPool;
use std::time::Instant;
use aoc2019::intcode::{parse_code, IntProgram};

fn main() {
    let mut code = String::new();
    let stdin = io::stdin();
    stdin.lock().read_line(&mut code).unwrap();

    let mut program = IntProgram::new(parse_code(&code));

    println!("solution one: {}", program.execute());

    let n_workers = 4;
    let pool = ThreadPool::new(n_workers);

    let now = Instant::now();

    let (tx, rx) = mpsc::channel();
    let bytecode = parse_code(&code);
    for i in 0..100 {
        for j in 0..100 {
            let tx = tx.clone();
            let mut program = IntProgram::new(bytecode.clone());
            program.mem[1] = i;
            program.mem[2] = j;
            // producer
            pool.execute(move|| {
                let candidate = program.execute();
                tx.send((candidate, i, j)).unwrap();
            });
        }
    }
    drop(tx);

    // consumer
    for received in rx {
        if received.0 == 19690720 {
            println!("Solution 2: {}", 100 * received.1 + received.2);
        }
    }

    let took = now.elapsed();
    println!("took {}.{}", took.as_secs(), took.subsec_nanos());
}
