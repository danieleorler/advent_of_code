use std::io::{self, prelude::*};
use aoc2019::geometry::{Point,Coordinates};
use std::collections::HashSet;
use std::collections::HashMap;

fn main () {
    let stdin = io::stdin();
    let mut row = 0;
    let mut asteroids = vec![];
    for line in stdin.lock().lines() {
        let l = line.unwrap();
        let mut column = 0;
        for c in l.chars() {
            match c {
                '#' => asteroids.push(Point::new(column as i64,row as i64)),
                _ => ()
            };
            column +=1;
        }
        row += 1;
    }

    let res1 = solution_one(&asteroids);
    println!("Solution one: {:?}, {:?}", res1.0, res1.1);
    println!("Solution two: {}", solution_two(Point::new(25,31), &asteroids));
}

fn solution_one(asteroids: &Vec<Point>) -> (Point, usize) {
    let mut rank = HashMap::new();
    for (i,a) in asteroids.iter().enumerate() {
        let mut angles = HashSet::new();
        for (j,o) in asteroids.iter().enumerate() {
            if i != j {
                let coord = a.coordinates_to(*o);
                angles.insert((coord.direction, coord.angle));
            }
        }
        rank.insert(a, angles.len());
    }

    let mut best_rank: usize = 0;
    let mut best_point: Point = Point::new(0,0);
    for (p, v) in rank.iter() {
        if *v > best_rank {
            best_rank = *v;
            best_point = **p;
        }
    }

    return (best_point, best_rank);
}

fn solution_two(center: Point, asteroids: &Vec<Point>) -> i64 {
    let mut other: Vec<Coordinates> = asteroids.iter()
     .filter(|a| **a != center)
     .map(|a| center.coordinates_to(*a))
     .collect();

    other.sort();
    other.reverse();

    let mut d = HashSet::new();
    let mut count = 0;
    let mut last_destroyed = other[0];
    other.retain(|c| { 
        let destroyed = d.insert((c.direction, c.angle));
        
        if count >= 200 {
            return true;
        }
        if destroyed {
            last_destroyed = *c;
            count +=1;
        }
        
        !destroyed
    });

    return last_destroyed.dest.x*100 + last_destroyed.dest.y;
}
