fn compute_step(start: i32, end: i32) -> i32 {
    (end - start).abs() * ((end - start).abs() + 1) / 2
}

fn main() {
    use std::cmp;
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let mut vals: Vec<i32> = contents
        .split(",")
        .filter(|s| !s.is_empty())
        .map(|s| s.parse().unwrap())
        .collect();

    let mut min_fuel = 0;

    for i in 0..(vals.len() - 1) {
        let mut current_min_fuel = 0;
        for v in vals.iter() {
            current_min_fuel += compute_step(i as i32, *v);
        }
        if min_fuel == 0 {
            min_fuel = current_min_fuel;
        }
        min_fuel = cmp::min(min_fuel, current_min_fuel);
    }

    println!("{}", min_fuel);
}
