use std::time::Instant;
use aoc2019::read_line;
use aoc2019::intcode::{parse_code, IntProgram};

fn main() {
    let now = Instant::now();

    let code = read_line();
    println!("Solution One: {}", solution_one(&code, 12, 2));
    println!("Solution Two: {}", solution_two(&code).unwrap());

    let took = now.elapsed();
    println!("took {}.{}", took.as_secs(), took.subsec_nanos());
}

fn solution_one(code: &String, a: i64, b: i64) -> i64 {
    let mut program = IntProgram::new(parse_code(&code));
    program.mem[1] = a;
    program.mem[2] = b;
    program.execute();
    return program.mem[0];
}

fn solution_two(code: &String) -> Option<i64> {
    let bytecode = parse_code(&code);
    for i in 0..100 {
        for j in 0..100 {
            let mut program = IntProgram::new(bytecode.clone());
            program.mem[1] = i;
            program.mem[2] = j;
            program.execute();
            if program.mem[0] == 19690720 {
                return Some(100 * i + j);
            }
        }
    }
    return None;
}

#[cfg(test)]
mod tests {
    use super::*;
    use aoc2019::read_line_from_file;

    #[test]
    fn solution_one_test() {
        let code = read_line_from_file(String::from("./inputs/02.input")).unwrap();
        assert_eq!(solution_one(&code, 12, 2), 3765464);
    }

    #[test]
    fn solution_two_test() {
        let code = read_line_from_file(String::from("./inputs/02.input")).unwrap();
        assert_eq!(solution_two(&code).unwrap(), 7610);
    }

}