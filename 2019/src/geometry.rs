use std::cmp::{min, max};
use std::cmp;

#[derive(Copy, Clone, Debug, PartialEq, Eq, Hash)]
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

    pub fn coordinates_to(self, other: Point) -> Coordinates {
        let distance = manhattan_distance(self, other);
        let mut direction = Direction::U;
        
        if self.x == other.x {
            match self.y < other.y {
                true => direction = Direction::D,
                false => direction = Direction::U
            }
        }

        if self.y == other.y {
            match self.x < other.x {
                true => direction = Direction::R,
                false => direction = Direction::L
            }
        }

        if self.x < other.x && self.y < other.y {
            direction = Direction::DR;
        } else if self.x < other.x && self.y > other.y {
            direction = Direction::UR;
        } else if self.x > other.x && self.y < other.y {
            direction = Direction::DL;
        } else if self.x > other.x && self.y > other.y {
            direction = Direction::UL;
        }

        let opposite = (self.x - other.x).abs() as f64;
        let adjacent = (self.y - other.y).abs() as f64;
        let angle = (opposite/adjacent).atan();
        return Coordinates {
            direction: direction,
            angle: (angle * 1000000f64) as i64,
            distance: distance,
            dest: other
        }
    }
}

#[derive(Copy, Clone, Debug, PartialEq, Eq, Hash)]
pub enum Direction {
    U,
    D,
    L,
    R,
    UL,
    UR,
    DL,
    DR
} 

pub fn direction_order(d: Direction) -> i32 {
    match d {
        Direction::U => 0,
        Direction::UR => 1,
        Direction::R => 2,
        Direction::DR => 3,
        Direction::D => 4,
        Direction::DL => 5,
        Direction::L => 6,
        Direction::UL => 7
    }
}

#[derive(Copy, Clone, Debug, PartialEq, Eq, Hash)]
pub struct Coordinates {
    pub direction: Direction,
    pub angle: i64,
    pub distance: i64,
    pub dest: Point
}

impl PartialOrd for Coordinates {
    fn partial_cmp(&self, other: &Coordinates) -> Option<cmp::Ordering> {
        Some(other.cmp(self))
    }
}

impl Ord for Coordinates {
    fn cmp(&self, other: &Coordinates) -> cmp::Ordering {
        match self.direction == other.direction {
            false => direction_order(self.direction).cmp(&direction_order(other.direction)),
            true => match self.angle == other.angle {
                true => self.distance.cmp(&other.distance),
                false => match self.direction {
                    Direction::UL => other.angle.cmp(&self.angle),
                    Direction::DR => other.angle.cmp(&self.angle),
                    _ => self.angle.cmp(&other.angle)
                }
            }
        }
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