use aoc2019::read_line;
use aoc2019::intcode::{parse_code, IntProgram};

fn main() {
    let code = read_line();
    println!("Solution 1: {}", solution(&code, 1));
    println!("Solution 2: {}", solution(&code, 5));
}

fn solution(code: &String, input: i64) -> i64 {
    let mut program = IntProgram::new(parse_code(&code));
    program.input.push(input);
    program.execute();
    return program.output.pop().unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;
    use aoc2019::read_line_from_file;

    #[test]
    fn solution_one() {
        let code = read_line_from_file(String::from("./inputs/05.input")).unwrap();
        assert_eq!(solution(&code, 1), 6069343);
    }

    #[test]
    fn solution_two() {
        let code = read_line_from_file(String::from("./inputs/05.input")).unwrap();
        assert_eq!(solution(&code, 5), 3188550);
    }


}