use std::collections::HashMap;

pub struct IntProgram {
    pub mem: HashMap<i64, i64>,
    pub input: Vec<i64>,
    pub output: Vec<i64>,
    pub pointer: i64,
    pub status: Status,
    pub relative_base: i64
}

#[derive(Debug)]
pub enum Instruction {
    ADD,
    MULT,
    SAVE,
    PRINT,
    JUMPT,
    JUMPF,
    LT,
    EQ,
    RBOFFSET,
    EXIT
}

#[derive(Debug)]
pub enum Mode {
    POSITION,
    IMMEDIATE,
    RELATIVE
}

#[derive(Debug)]
pub enum Status {
    READY,
    EXECUTING,
    WAITINGFORINPUT,
    DONE
}

#[derive(Debug)]
pub struct Statement {
    pub inst: Instruction,
    pub first: Mode,
    pub second: Mode,
    pub third: Mode
}

pub fn parse_code(code: &String) -> Vec<i64> {
    return code.split(",")
        .map(|x| x.parse::<i64>().unwrap())
        .collect();
}

pub fn to_mode(i: i64) -> Mode {
    return match i {
        0 => Mode::POSITION,
        1 => Mode::IMMEDIATE,
        2 => Mode::RELATIVE,
        _ => panic!("mode {} not recognized", i)
    }
}

fn to_instruction(i: i64) -> Instruction {
    match i {
        1 => Instruction::ADD,
        2 => Instruction::MULT,
        3 => Instruction::SAVE,
        4 => Instruction::PRINT,
        5 => Instruction::JUMPT,
        6 => Instruction::JUMPF,
        7 => Instruction::LT,
        8 => Instruction::EQ,
        9 => Instruction::RBOFFSET,
        99 => Instruction::EXIT,
        _ => panic!("{} is not a valid Instruction", i)
    }
}
fn to_statement(raw: i64) -> Statement {
    let log = (raw as f64).log10().floor() as i64;
    match log {
        0 => {
            Statement {
                inst: to_instruction(raw),
                first: Mode::POSITION,
                second: Mode::POSITION,
                third: Mode::POSITION
            }
        },
        1 => {
            Statement {
                inst: to_instruction(raw),
                first: Mode::POSITION,
                second: Mode::POSITION,
                third: Mode::POSITION
            }
        },
        2 => {
            Statement {
                inst: to_instruction(raw%10),
                first: to_mode(raw / 100),
                second: Mode::POSITION,
                third: Mode::POSITION
            }
        },
        3 => {
            Statement {
                inst: to_instruction(raw%10),
                first: to_mode((raw / 100) % 10),
                second: to_mode(raw / 1000),
                third: Mode::POSITION
            }
        },
        4 => {
            Statement {
                inst: to_instruction(raw%10),
                first: to_mode((raw / 100) % 10),
                second: to_mode((raw / 1000) % 10),
                third: to_mode(raw / 10000)
            }
        },
        _ => panic!("{} is not parable to a Statement", raw)
    }
}

impl IntProgram {
    pub fn new(bytecode: Vec<i64>) -> Self {
        let mut map: HashMap<i64, i64> = HashMap::new();
        for (i, e) in bytecode.iter().enumerate() {
            map.insert(i as i64, *e);
        }
        return IntProgram { 
            mem: map,
            input: Vec::new(),
            output: Vec::new(),
            pointer: 0,
            status: Status::READY,
            relative_base: 0
         };
    }

    pub fn get_param(&mut self, i: i64, mode: &Mode) -> i64 {
        let target = match mode {
            Mode::IMMEDIATE => i,
            Mode::POSITION => self.mem.get(&i).unwrap() + 0i64,
            Mode::RELATIVE => self.relative_base + self.mem.get(&i).unwrap()
        };

        if !self.mem.contains_key(&target) {
            self.mem.insert(target, 0);
        }

        return target;
    }

    pub fn store_value(&mut self, dest: i64, value: i64) {
        self.mem.insert(dest, value);
    }

    pub fn is_running(&self) -> bool {
        match self.status {
            Status::EXECUTING => true,
            _ => false
        }
    }

    pub fn get(&self, offset: i64) -> i64 {
        return *self.mem.get(&(self.pointer + offset)).unwrap()
    }

    pub fn get_at(&self, addr: i64) -> i64 {
        return *self.mem.get(&addr).unwrap()
    }

    pub fn execute(&mut self) {
        let mut op_code = Self::get(self, 0);
        self.status = Status::EXECUTING;
        while self.is_running() {
            let stm = to_statement(op_code);
            match stm.inst {
                Instruction::ADD => {
                    let dest = Self::get_param(self, self.pointer+3, &stm.third);
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    Self::store_value(self, dest, Self::get_at(self, first) + Self::get_at(self, second));
                    self.pointer += 4;
                },
                Instruction::MULT => {
                    let dest = Self::get_param(self, self.pointer+3, &stm.third);
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    Self::store_value(self, dest, Self::get_at(self, first) * Self::get_at(self, second));
                    self.pointer += 4;
                },
                Instruction::SAVE => {
                    if self.input.len() < 1 {
                        self.status = Status::WAITINGFORINPUT;
                    } else {
                        let dest = Self::get_param(self, self.pointer+1, &stm.first);
                        let value = self.input.pop().unwrap();
                        Self::store_value(self, dest, value);
                        self.pointer += 2;
                    }
                    
                },
                Instruction::PRINT => {
                    let out = Self::get_param(self, self.pointer+1, &stm.first);
                    self.output.push(Self::get_at(self, out));
                    self.pointer += 2;
                },
                Instruction::JUMPT => {
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    if Self::get_at(self, first) != 0 {
                        self.pointer = Self::get_at(self, second);
                    } else {
                        self.pointer += 3;
                    }
                },
                Instruction::JUMPF => {
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    if Self::get_at(self, first) == 0 {
                        self.pointer = Self::get_at(self, second);
                    } else {
                        self.pointer += 3;
                    }
                },
                Instruction::LT => {
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    let dest = Self::get_param(self, self.pointer+3, &stm.third);
                    if Self::get_at(self, first) < Self::get_at(self, second) {
                        Self::store_value(self, dest, 1);
                    } else {
                        Self::store_value(self, dest, 0);
                    }
                    self.pointer += 4;
                },
                Instruction::EQ => {
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    let dest = Self::get_param(self, self.pointer+3, &stm.third);
                    if Self::get_at(self, first) == Self::get_at(self, second) {
                        Self::store_value(self, dest, 1);
                    } else {
                        Self::store_value(self, dest, 0);
                    }
                    self.pointer += 4;
                },
                Instruction::RBOFFSET => {
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    self.relative_base += Self::get_at(self, first);
                    self.pointer += 2;
                },
                Instruction::EXIT => {
                    self.status = Status::DONE;
                }
            }
            op_code = Self::get(self, 0);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn quine() {
        let code = String::from("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99");
        let mut p = IntProgram::new(parse_code(&code));
        p.execute();
        let result = p.output.iter().map(|x| x.to_string()).collect::<Vec<String>>().join(",");
        assert_eq!(code, result);
    }

    #[test]
    fn sixteen_digits() {
        let mut p = IntProgram::new(parse_code(&String::from("1102,34915192,34915192,7,4,7,99,0")));
        p.execute();
        let result = p.output[0].to_string().len();
        assert_eq!(16, result);
    }

    #[test]
    fn the_one_in_the_middle() {
        let mut p = IntProgram::new(parse_code(&String::from("104,1125899906842624,99")));
        p.execute();
        let result = p.output[0];
        assert_eq!(1125899906842624, result);
    }
}