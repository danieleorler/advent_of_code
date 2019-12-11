use aoc2019::read_line;
use aoc2019::intcode::{parse_code, IntProgram, Status};
use aoc2019::geometry::{Point, Direction};
use std::collections::HashMap;

fn main() {
    
    let code = read_line();
    let bytecode = parse_code(&code);

    println!("Solution one: {}", solution(&bytecode, 0).len());
    let res2 = solution(&bytecode, 1);
    println!("Solution two:");
    render(&res2);
}

fn solution(bytecode: &Vec<i64>, color: i64) -> HashMap<Point, i64> {
    let mut p = IntProgram::new(bytecode.clone());

    let mut tails = HashMap::new();
    tails.insert(Point::new(0,0), color);
    let mut loc = Point::new(0,0);
    let mut direction = Direction::U;
    loop {
        match p.status {
            Status::DONE => break,
            _ => {
                p.input.push(*tails.get(&loc).unwrap());
                p.execute();
                direction = turn(direction, p.output.pop().unwrap());
                tails.insert(loc, p.output.pop().unwrap());
                let next = next(&loc, &direction);
                if !tails.contains_key(&next) {
                    tails.insert(next, 0);
                }
                loc = next;
            }
        }
    }
    return tails;
}

fn render(map: &HashMap<Point, i64>) {
    let min_x = map.keys().map(|p| p.x).min().unwrap();
    let max_x = map.keys().map(|p| p.x).max().unwrap();
    let min_y = map.keys().map(|p| p.y).min().unwrap();
    let max_y = map.keys().map(|p| p.y).max().unwrap();

    for x in min_x..=max_x {
        for y in min_y..=max_y {
            let point = Point::new((x as i64) as i64, (y as i64) as i64);
            match map.contains_key(&point) {
                true => match map.get(&point).unwrap() {
                    0 => print!(" "),
                    _ => print!("*"),
                },
                false => print!(" ")
            }
        }
        println!("");
    }
}

fn turn(original: Direction, m: i64) -> Direction {
    match original {
        Direction::U => match m {
            0 => Direction::L,
            _ => Direction::R
        },
        Direction::D => match m {
            0 => Direction::R,
            _ => Direction::L
        },
        Direction::L => match m {
            0 => Direction::D,
            _ => Direction::U
        },
        Direction::R => match m {
            0 => Direction::U,
            _ => Direction::D
        },
        _ => panic!("not a valid direction")
    }
}

fn next(cur: &Point, dir: &Direction) -> Point {
    match dir {
        Direction::U => Point::new(cur.x, cur.y + 1),
        Direction::D => Point::new(cur.x, cur.y - 1),
        Direction::L => Point::new(cur.x - 1 , cur.y),
        Direction::R => Point::new(cur.x + 1, cur.y),
        _ => panic!("not a valid direction")
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use aoc2019::intcode::parse_code;
    use aoc2019::read_line_from_file;

    #[test]
    fn solution_one() {
        let code = read_line_from_file(String::from("./inputs/11.input")).unwrap();
        assert_eq!(solution(&parse_code(&code), 0).len(), 1885);
    }
}