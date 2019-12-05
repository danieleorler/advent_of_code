use std::cmp::{min, max};

pub mod intcode;

#[derive(Copy, Clone, Debug)]
pub struct Point {
    pub x: i64,
    pub y: i64,
}
 
impl Point {
    pub fn new(x: i64, y: i64) -> Self {
        Point { x, y }
    }

    pub fn is_origin(self) -> bool {
        return self.x == 0 && self.y == 0;
    }
}

#[derive(Copy, Clone, Debug)]
pub struct Line {
    pub from: Point,
    pub to: Point
}
 
impl Line {
    pub fn new(from: Point, to: Point) -> Self {
        Line { from, to }
    }
    pub fn intersect(self, other: Self) -> Option<Point> {
        let x = match self.from.x == self.to.x {
            true => min(other.from.x, other.to.x) <= self.from.x && self.from.x <= max(other.from.x, other.to.x),
            false => min(self.from.x, self.to.x) <= other.from.x && other.from.x <= max(self.from.x, self.to.x)
        };

        let y = match self.from.y == self.to.y {
            true => min(other.from.y, other.to.y) <= self.from.y && self.from.y <= max(other.from.y, other.to.y),
            false => min(self.from.y, self.to.y) <= other.from.y && other.from.y <= max(self.from.y, self.to.y)
        };
        
        return match x && y {
            true => Some(self.find_intersection(other)),
            false => None
        }
    }

    fn find_intersection(self, other: Self) -> Point {
        let x = match self.from.x == self.to.x {
            true => self.from.x,
            false => other.from.x
        };

        let y = match self.from.y == self.to.y {
            true => self.from.y,
            false => other.from.y
        };

        return Point::new(x, y);

    }
}

pub fn manhattan_distance(from: Point, to: Point) -> i64 {
    return (from.x - to.x).abs() + (from.y - to.y).abs(); 
}