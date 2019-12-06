use std::collections::HashMap;
use std::io;
use std::io::prelude::*;

fn main() {
    let stdin = io::stdin();
    let mut edges = Vec::new();
    let mut reverse_edges = Vec::new();
    for line in stdin.lock().lines() {
        let unwrapped = line.unwrap();
        let mut t = unwrapped.split(")");
        let a = t.next().unwrap().to_string();
        let b = t.next().unwrap().to_string();
        edges.push((a.to_string(), b.to_string()));
        reverse_edges.push((b.to_string(), a.to_string()));
    }

    let tree = build_tree(edges);
    let mut res = vec![];
    traverse(String::from("COM"), &mut res, -1, &tree);
    println!("Soulution 1: {:?}", res.iter().sum::<i64>());

    let reverse_tree = build_tree(reverse_edges);
    let mut mine = vec![];
    find_ancestors(String::from("YOU"), &mut mine, -1, &reverse_tree);
    let mut santas = vec![];
    find_ancestors(String::from("SAN"), &mut santas, -1, &reverse_tree);

    let mut common_ancestors_dist = vec![];
    for m in &mine {
        for s in &santas {
            if m.0 == s.0 {
                common_ancestors_dist.push(m.1 + s.1 - 2);
            }
        }
    }

    println!("Soulution 2: {:?}", common_ancestors_dist.iter().min().unwrap());
}

fn build_tree(edges: Vec<(String,String)>) -> HashMap<String, Vec<String>> {
    let mut tree: HashMap<String, Vec<String>> = HashMap::new();
    for edge in edges {
        match tree.contains_key(&edge.0) {
            true => (*tree.get_mut(&edge.0).unwrap()).push(edge.1),
            false => { tree.insert(edge.0, vec![edge.1]); ()}
        }
    }
    tree
}

fn traverse(p: String, depths: &mut Vec<i64>, progress: i64, tree: &HashMap<String, Vec<String>>) {
    match tree.contains_key(&p) {
        false => {
            depths.push(progress + 1)
        },
        true => {
            depths.push(progress + 1);
            for n in tree.get(&p).unwrap() {
                traverse(n.to_string(), depths, progress + 1, tree);
            }
        }
    }
}

fn find_ancestors(p: String, ancestors: &mut Vec<(String, i64)>, progress: i64, tree: &HashMap<String, Vec<String>>) {
    match tree.get(&p) {
        Some(x) => {
            match x.len() {
                0 => (),
                _ => {
                    ancestors.push((p.to_string(), progress + 1));
                    for n in tree.get(&p).unwrap() {
                        find_ancestors(n.to_string(), ancestors, progress + 1, tree);
                    }
                }
            }
        },
        _ => ()
    }
}