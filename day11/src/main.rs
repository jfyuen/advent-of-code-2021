fn flash(vals: &mut Vec<Vec<i32>>) -> usize {
    use std::collections::HashSet;
    let mut flashed = HashSet::new();
    let range: [i32; 3] = [-1, 0, 1];
    let mut has_new_flash = true;
    while has_new_flash {
        has_new_flash = false;
        for i in 0..vals.len() {
            for j in 0..vals[0].len() {
                if vals[i][j] <= 9 {
                    continue;
                }

                if flashed.contains(&(i, j)) {
                    continue;
                }
                flashed.insert((i, j));
                has_new_flash = true;
                for k in range {
                    for l in range {
                        if k + (i as i32) < 0 || k + (i as i32) >= vals.len() as i32 {
                            continue;
                        }
                        if l + (j as i32) < 0 || l + (j as i32) >= vals[0].len() as i32 {
                            continue;
                        }
                        vals[(i as i32 + k) as usize][(j as i32 + l) as usize] += 1
                    }
                }
            }
        }
    }
    flashed.len()
}

fn reset(vals: &mut Vec<Vec<i32>>) {
    for i in 0..vals.len() {
        for j in 0..vals[0].len() {
            if vals[i][j] > 9 {
                vals[i][j] = 0;
            }
        }
    }
}

fn increase(vals: &Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    let mut step_vals = vals.clone();

    for i in 0..vals.len() {
        for j in 0..vals[0].len() {
            step_vals[i][j] += 1;
        }
    }
    step_vals
}

fn main() {
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let mut vals: Vec<Vec<i32>> = contents
        .split("\n")
        .map(|s| s.chars().map(|c| c.to_digit(10).unwrap() as i32).collect())
        .collect();

    // let mut flash_count = 0;
    let mut i = 1;
    loop {
        vals = increase(&vals);
        let current_flash_count = flash(&mut vals);
        if current_flash_count == vals[0].len() * vals.len() {
            println!("first flash at {}", i);
            break;
        }
        // flash_count += current_flash_count;
        reset(&mut vals);
        i += 1;
    }
    // println!("{:?}", flash_count);
}
