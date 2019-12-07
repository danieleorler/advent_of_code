use std::io;
use std::io::prelude::*;
use aoc2019::geometry::{Point, Line, manhattan_distance};

fn main() {
    let stdin = io::stdin();

    let mut a = String::new();
    stdin.lock().read_line(&mut a).unwrap();
    let first = to_lines(to_points(to_instructions(String::from(a.trim()))));
    let mut b = String::new();
    stdin.lock().read_line(&mut b).unwrap();
    let second = to_lines(to_points(to_instructions(String::from(b.trim()))));

    let mut straight_distances: Vec<i64> = Vec::new();
    let mut manhattan_distances: Vec<i64> = Vec::new();
    let mut manhattan_distance_a = 0;

    for la in first {
        manhattan_distance_a += manhattan_distance(la.from, la.to);
        let mut manhattan_distance_b = 0;
        for lb in &second {
            let intersect = la.intersect(lb.clone());
            if intersect.is_some() && !intersect.unwrap().is_origin() {
                manhattan_distance_b += manhattan_distance(lb.from, intersect.unwrap());
                straight_distances
                    .push(manhattan_distance(Point::new(0,0), intersect.unwrap()));
                manhattan_distances
                    .push(manhattan_distance_a + manhattan_distance_b - manhattan_distance(intersect.unwrap(), la.to));
            } else {
                manhattan_distance_b += manhattan_distance(lb.from, lb.to);
            }
        }
    }

    let res1 = straight_distances.iter().min();
    println!("Solution 1: {}", res1.unwrap());

    let res2 = manhattan_distances.iter().min();
    println!("Solution 2: {}", res2.unwrap());
}

fn to_instructions(input: String) -> Vec<(char, i64)> {
    return input
        .split(",")
        .map(|x| (x.chars().next().unwrap(), (&x[1..]).parse::<i64>().unwrap()))
        .collect();
}

fn to_points(input: Vec<(char, i64)>) -> Vec<Point> {
    let mut points = vec!(Point::new(0,0));
    for instruction in input {
        points.push(follow_instruction(&points.last().unwrap(), instruction));
    }

    return points;
}

fn to_lines(input: Vec<Point>) -> Vec<Line> {
    let mut start = input[0];
    let mut lines: Vec<Line> = Vec::new();
    for i in 1..input.len() {
        lines.push(Line::new(start, input[i]));
        start = input[i];
    }

    return lines;
}

fn follow_instruction(start: &Point, instruction: (char, i64)) -> Point {
    match instruction.0 {
        'U' =>  Point::new(start.x, start.y + instruction.1),
        'D' =>  Point::new(start.x, start.y - instruction.1),
        'R' =>  Point::new(start.x + instruction.1, start.y),
        'L' =>  Point::new(start.x - instruction.1, start.y),
        _ => panic!("{}", instruction.0)
    }
}