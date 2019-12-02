use std::io;
use std::io::prelude::*;
use std::sync::mpsc;
use threadpool::ThreadPool;
use std::time::Instant;

fn main() {
    let mut line = String::new();
    let stdin = io::stdin();
    stdin.lock().read_line(&mut line).unwrap();

    let int_code: Vec<usize> = line.split(",")
        .map(|x| x.parse::<usize>().unwrap())
        .collect();

    println!("solution one: {}", execute_program(&mut int_code.to_vec()));

    let n_workers = 4;
    let pool = ThreadPool::new(n_workers);

    let now = Instant::now();

    let (tx, rx) = mpsc::channel();
    for i in 0..100 {
        for j in 0..100 {
            let tx = tx.clone();
            let mut program = int_code.to_vec();
            program[1] = i;
            program[2] = j;
            // producer
            pool.execute(move|| {
                let candidate = execute_program(&mut program);
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

fn execute_program(int_code: &mut [usize]) -> usize {
    let mut i = 0;
    let mut op_code = int_code[0];
    while op_code != 99 {
        let dest = int_code[i+3];
        let a = int_code[i+2];
        let b = int_code[i+1];
        if op_code == 1 {
            int_code[dest] = int_code[a] + int_code[b];
        }
        if op_code == 2 {
            int_code[dest] = int_code[a] * int_code[b];
        }
        i = i + 4;
        op_code = int_code[i];
    }

    int_code[0]
}
