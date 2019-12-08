use aoc2019::read_line;

fn main() {
    let input = read_line();

    let mut pixel = 0;
    let mut layers = vec![];
    while pixel < input.len() {
        layers.push(&input[pixel..pixel+150]);
        pixel += 150;
    }

    println!("Solution one: {}", check(layers.clone()));
    println!("Solution two: ");
    render(&mut layers.clone());
}

fn render(layers: &mut Vec<&str>) {
    let mut rendered = vec![];
    for c in layers[0].chars() {
        rendered.push(match c { '0' => ' ', _ => c });
    }

    for layer in layers {
        for (i, c) in layer.chars().enumerate() {
            if rendered[i] == '2' {
                rendered[i] = match c { '0' => ' ', _ => c };
            }
        }
    }

    let mut pixel = 0;
    while pixel < rendered.len() {
        println!("{:?}", &rendered[pixel..pixel+25].to_vec().into_iter().collect::<String>());
        pixel += 25;
    } 
}

fn check(layers: Vec<&str>) -> usize {
    let mut min_i = 0;
    let mut min_zeros = 150;
    for (i, layer) in layers.iter().enumerate() {
        let n_zeros = layer.matches("0").count();
        if n_zeros < min_zeros {
            min_zeros = n_zeros;
            min_i = i;
        }
    }

    return layers[min_i].matches("1").count() * layers[min_i].matches("2").count();
}