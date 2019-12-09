use aoc2019::read_line;
use aoc2019::intcode::{parse_code, IntProgram};

fn main() {
    let code = read_line();
    let bytecode = parse_code(&code);

    println!("Solution one: {}", solution(bytecode.clone(), 1));
    println!("Solution two: {}", solution(bytecode.clone(), 2));
}

fn solution(bytecode: Vec<i64>, mode: i64) -> i64 {
    let mut p = IntProgram::new(bytecode);
    p.input.push(mode);
    p.execute();
    return p.output.pop().unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;
    use aoc2019::intcode::parse_code;
    use aoc2019::read_line_from_file;

    #[test]
    fn solution_one() {
        let code = read_line_from_file(String::from("./inputs/09.input")).unwrap();
        assert_eq!(solution(parse_code(&code), 1), 3345854957);
    }

    #[test]
    fn solution_two() {
        let code = read_line_from_file(String::from("./inputs/09.input")).unwrap();
        assert_eq!(solution(parse_code(&code), 2), 68938);
    }

}