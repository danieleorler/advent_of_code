use aoc2019::read_line;
use aoc2019::intcode::{parse_code, IntProgram, Status};

fn main() {
    let code = read_line();
    let bytecode = parse_code(&code);

    println!("Solution one: {}", solution(bytecode.clone(), 0, 4));
    println!("Solution two: {}", solution(bytecode.clone(), 5, 9));
}

fn solution(bytecode: Vec<i64>, low: i64, high: i64) -> i64 {
    let phases = generate_phases(low, high);
    return phases.iter()
        .map(|p| simulation(&bytecode, *p))
        .max()
        .unwrap();
}

fn simulation(bytecode: &Vec<i64>, phases: (i64, i64, i64, i64, i64)) -> i64 {
    let amp1 = &mut IntProgram::new(bytecode.clone());
    let amp2 = &mut IntProgram::new(bytecode.clone());
    let amp3 = &mut IntProgram::new(bytecode.clone());
    let amp4 = &mut IntProgram::new(bytecode.clone());
    let amp5 = &mut IntProgram::new(bytecode.clone());

    start_amplifier(amp1, phases.0, 0);
    start_amplifier(amp2, phases.1, amp1.output.pop().unwrap());
    start_amplifier(amp3, phases.2, amp2.output.pop().unwrap());
    start_amplifier(amp4, phases.3, amp3.output.pop().unwrap());
    start_amplifier(amp5, phases.4, amp4.output.pop().unwrap());

    while !are_all_done(vec![amp1, amp2, amp3, amp4, amp5]) {
        resume_amplifier(amp1, amp5.output.pop().unwrap());
        resume_amplifier(amp2, amp1.output.pop().unwrap());
        resume_amplifier(amp3, amp2.output.pop().unwrap());
        resume_amplifier(amp4, amp3.output.pop().unwrap());
        resume_amplifier(amp5, amp4.output.pop().unwrap());
    }

    return amp5.output.pop().unwrap();
}

fn start_amplifier(program: &mut IntProgram, phase: i64, input: i64) {
    program.input.push(input);
    program.input.push(phase);
    program.execute();
}

fn are_all_done(amps: Vec<&mut IntProgram>) -> bool {
    let mut done = true;
    for amp in amps {
        match amp.status {
            Status::DONE => done &= true,
            _ => done &= false
        }
    }
    return done
}

fn resume_amplifier(program: &mut IntProgram, input: i64) {
    program.input.push(input);
    program.execute();
}

fn generate_phases(low: i64, high: i64) -> Vec<(i64, i64, i64, i64, i64)> {
    let mut all = vec![];
    for a in low..=high {
        for b in low..=high {
            for c in low..=high {
                for d in low..=high {
                    for e in low..=high {
                        all.push(vec![a,b,c,d,e]);
                    }
                }
            }
        }
    }
    
    let mut no_repetition = vec![];
    for c in all {
        if !has_repetitions(&mut c.clone()) {
            no_repetition.push((c[0],c[1],c[2],c[3],c[4]));
        }
    }
    return no_repetition;
}

fn has_repetitions(l: &mut Vec<i64>) -> bool {
    let original = l.len();
    l.sort();
    l.dedup();
    return original != l.len();
}

#[cfg(test)]
mod tests {
    use super::*;
    use aoc2019::intcode::parse_code;
    use aoc2019::read_line_from_file;

    #[test]
    fn simulation_test() {
        assert_eq!(simulation(&parse_code(&String::from("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")), (4,3,2,1,0)), 43210);
        assert_eq!(simulation(&parse_code(&String::from("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")), (0,1,2,3,4)), 54321);
        assert_eq!(simulation(&parse_code(&String::from("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")), (1,0,4,3,2)), 65210);
        assert_eq!(simulation(&parse_code(&String::from("3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5")), (9,8,7,6,5)), 139629729);
        assert_eq!(simulation(&parse_code(&String::from("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10")), (9,7,8,5,6)), 18216);
    }

    #[test]
    fn solution_one_test() {
        assert_eq!(solution(parse_code(&String::from("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")), 0, 4), 43210);
        assert_eq!(solution(parse_code(&String::from("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")), 0, 4), 54321);
        assert_eq!(solution(parse_code(&String::from("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")), 0, 4), 65210);
        let code = read_line_from_file(String::from("./inputs/07.input")).unwrap();
        assert_eq!(solution(parse_code(&code), 0, 4), 206580);
        assert_eq!(solution(parse_code(&code), 5, 9), 2299406);
    }

    #[test]
    fn has_repetitions_test() {
        assert_eq!(has_repetitions(&mut vec![0,1,2,3,4]), false);
        assert_eq!(has_repetitions(&mut vec![1,2,2,3,4]), true);
        assert_eq!(has_repetitions(&mut vec![0,0,0,0,0]), true);
    }
}