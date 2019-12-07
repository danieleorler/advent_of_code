pub struct IntProgram {
    pub mem: Vec<i64>,
    pub input: Vec<i64>,
    pub output: Vec<i64>,
    pub pointer: usize,
    pub status: Status
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
    EXIT
}

#[derive(Debug)]
pub enum Mode {
    POSITION,
    IMMEDIATE
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
        99 => Instruction::EXIT,
        _ => panic!("{} is not a valid Instruction", i)
    }
}
fn to_statement(raw: i64) -> Statement {
    let log = (raw as f64).log10().floor() as i64;
    match log {
        0 => {
            return Statement {
                inst: to_instruction(raw),
                first: Mode::POSITION,
                second: Mode::POSITION,
                third: Mode::POSITION
            }
        },
        1 => {
            return Statement {
                inst: to_instruction(raw),
                first: Mode::POSITION,
                second: Mode::POSITION,
                third: Mode::POSITION
            }
        },
        2 => {
            return Statement {
                inst: to_instruction(raw%10),
                first: match raw / 100 { 0 => Mode::POSITION, _ => Mode::IMMEDIATE },
                second: Mode::POSITION,
                third: Mode::POSITION
            }
        },
        3 => {
            return Statement {
                inst: to_instruction(raw%10),
                first: match (raw / 100) % 10 { 0 => Mode::POSITION, _ => Mode::IMMEDIATE },
                second: match raw / 1000 { 0 => Mode::POSITION, _ => Mode::IMMEDIATE },
                third: Mode::POSITION
            }
        },
        4 => {
            return Statement {
                inst: to_instruction(raw%10),
                first: match (raw / 100) % 10 { 0 => Mode::POSITION, _ => Mode::IMMEDIATE },
                second: match (raw / 1000) % 10 { 0 => Mode::POSITION, _ => Mode::IMMEDIATE },
                third: match raw / 10000 { 0 => Mode::POSITION, _ => Mode::IMMEDIATE }
            }
        },
        _ => panic!("{} is not parable to a Statement", raw)
    }
}

impl IntProgram {
    pub fn new(bytecode: Vec<i64>) -> Self {
        return IntProgram { 
            mem: bytecode,
            input: Vec::new(),
            output: Vec::new(),
            pointer: 0,
            status: Status::READY
         };
    }

    pub fn get_param(&self, i: usize, mode: &Mode) -> i64 {
        match mode {
            Mode::IMMEDIATE => self.mem[i],
            _ => self.mem[self.mem[i] as usize]
        }
    }

    pub fn is_running(&self) -> bool {
        match self.status {
            Status::EXECUTING => true,
            _ => false
        }
    }

    pub fn execute(&mut self) {
        let mut op_code = self.mem[self.pointer];
        self.status = Status::EXECUTING;
        while self.is_running() {
            let stm = to_statement(op_code);
            match stm.inst {
                Instruction::ADD => {
                    let dest = self.mem[self.pointer+3] as usize;
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    self.mem[dest] = first + second;
                    self.pointer += 4;
                },
                Instruction::MULT => {
                    let dest = self.mem[self.pointer+3] as usize;
                    let first = Self::get_param(self, self.pointer+1, &stm.first);
                    let second = Self::get_param(self, self.pointer+2, &stm.second);
                    self.mem[dest] = first * second;
                    self.pointer += 4;
                },
                Instruction::SAVE => {
                    if self.input.len() < 1 {
                        self.status = Status::WAITINGFORINPUT;
                    } else {
                        let dest = self.mem[self.pointer+1] as usize;
                        self.mem[dest] = self.input.pop().unwrap();
                        self.pointer += 2;
                    }
                    
                },
                Instruction::PRINT => {
                    let out = Self::get_param(self, self.pointer+1, &stm.first);
                    self.output.push(out);
                    self.pointer += 2;
                },
                Instruction::JUMPT => {
                    if Self::get_param(self, self.pointer+1, &stm.first) != 0 {
                        self.pointer = Self::get_param(self, self.pointer+2, &stm.second) as usize;
                    } else {
                        self.pointer += 3;
                    }
                },
                Instruction::JUMPF => {
                    if Self::get_param(self, self.pointer+1, &stm.first) == 0 {
                        self.pointer = Self::get_param(self, self.pointer+2, &stm.second) as usize;
                    } else {
                        self.pointer += 3;
                    }
                },
                Instruction::LT => {
                    let dest = self.mem[self.pointer+3] as usize;
                    if Self::get_param(self, self.pointer+1, &stm.first) < Self::get_param(self, self.pointer+2, &stm.second) {
                        self.mem[dest] = 1;
                    } else {
                        self.mem[dest] = 0;
                    }
                    self.pointer += 4;
                },
                Instruction::EQ => {
                    let dest = self.mem[self.pointer+3] as usize;
                    if Self::get_param(self, self.pointer+1, &stm.first) == Self::get_param(self, self.pointer+2, &stm.second) {
                        self.mem[dest] = 1;
                    } else {
                        self.mem[dest] = 0;
                    }
                    self.pointer += 4;
                },
                Instruction::EXIT => {
                    self.status = Status::DONE;
                }
            }
            op_code = self.mem[self.pointer];
        }
    }
}
