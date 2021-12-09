fn is_local_min(vals: &Vec<Vec<i32>>, i: i32, j: i32) -> bool {
    let range: [(i32, i32); 4] = [(-1, 0), (0, -1), (1, 0), (0, 1)];
    for (k, l) in &range {
        if k + i < 0 || k + i >= vals.len() as i32 {
            continue;
        }
        if l + j < 0 || l + j >= vals[0].len() as i32 {
            continue;
        }
        if vals[(i + k) as usize][(j + l) as usize] <= vals[i as usize][j as usize] {
            return false;
        }
    }
    true
}

fn get_basin_size(vals: &mut Vec<Vec<i32>>, i: i32, j: i32) -> i32 {
    let mut total = 1;
    vals[i as usize][j as usize] = 9;

    let range: [(i32, i32); 4] = [(-1, 0), (0, -1), (1, 0), (0, 1)];
    for (k, l) in &range {
        if k + i < 0 || k + i >= vals.len() as i32 {
            continue;
        }
        if l + j < 0 || l + j >= vals[0].len() as i32 {
            continue;
        }
        if vals[(i + k) as usize][(j + l) as usize] != 9 {
            total += get_basin_size(vals, i + k, j + l);
        }
    }
    total
}

fn main() {
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let vals: Vec<Vec<i32>> = contents
        .split("\n")
        .map(|s| s.chars().map(|c| c.to_digit(10).unwrap() as i32).collect())
        .collect();

    let mut basins: Vec<i32> = Vec::<i32>::new();
    for i in 0..vals.len() {
        for j in 0..vals[0].len() {
            if is_local_min(&vals, i as i32, j as i32) {
                basins.push(get_basin_size(&mut vals.clone(), i as i32, j as i32));
            }
        }
    }
    basins.sort();
    let mut total = 1;
    for i in 0..3 {
        println!("{:?}", basins[basins.len() - i - 1]);
        total *= basins[basins.len() - i - 1];
    }
    println!("{}", total);
}
