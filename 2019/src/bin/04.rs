use std::collections::HashMap;

fn main() {
    let res1 = (359282..820401)
        .filter(|p| is_valid_one(reverse(*p)))
        .count();

    println!("Solution1: {}", res1);

    let res2 = (359282..820401)
        .map(|p| reverse(p))
        .filter(|p| is_valid_one(p.to_vec()))
        .filter(|p| is_valid_two(p.to_vec()))
        .count();

    println!("Solution2: {}", res2);
}

fn reverse(p: i32) -> Vec<i32> {
    let mut t = p;
    let mut digits = Vec::new();
    while t > 0 {
        digits.push(t%10);
        t /= 10;
    }
    return digits;
}

fn is_valid_one(p: Vec<i32>) -> bool {
    let mut monotone = true;
    let mut has_repetitions = false;
    for i in 0..p.len()-1 {
        monotone &= p[i] >= p [i+1];
        has_repetitions |= p[i] == p[i+1];
    }
    return monotone && has_repetitions;
}

fn is_valid_two(p: Vec<i32>) -> bool {
    let mut f: HashMap<i32, i32> = HashMap::new();
    for i in 0..p.len() {
        if !f.contains_key(&p[i]) {
            f.insert(p[i], 1);
        } else {
            f.insert(p[i], f.get(&p[i]).unwrap() + 1);
        }
    }
    return f.iter().filter(|x| x.1 == &2).count() > 0;
}