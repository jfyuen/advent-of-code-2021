fn filter_values<'a>(vals: &Vec<&'a str>, c: char, pos: usize) -> Vec<&'a str> {
    let mut r: Vec<&str> = Vec::new();
    for v in vals
        .into_iter()
        .filter(|s| s.chars().nth(pos).unwrap() == c)
    {
        r.push(v)
    }
    r
}

fn count_zero_ones(vals: &[&str]) -> (Vec<i32>, Vec<i32>) {
    let size = vals[0].len();

    let mut ones = vec![0; size];
    let mut zeros = vec![0; size];

    for val in vals {
        for (i, c) in val.chars().enumerate() {
            match c {
                '0' => zeros[i] += 1,
                '1' => ones[i] += 1,
                _ => (),
            }
        }
    }
    (ones, zeros)
}

fn filter_vec<'a>(vals: &Vec<&'a str>, most_common: bool) -> Vec<&'a str> {
    let mut filtered = vals.to_vec();
    for i in 0..vals[0].len() {
        if filtered.len() == 1 {
            break;
        }
        let (ones, zeros) = count_zero_ones(&filtered);
        let mut check = ones[i] >= zeros[i];
        check = if most_common { check } else { !check };
        let char = if check { '1' } else { '0' };
        filtered = filter_values(&filtered, char, i);
    }
    filtered
}

fn to_int(s: &str) -> i32 {
    let mut r = 0;
    for (i, c) in s.chars().enumerate() {
        if c == '1' {
            r += 2_i32.pow((s.len() - i - 1) as u32);
        }
    }
    r
}

fn main() {
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let vals: Vec<&str> = contents.split("\n").filter(|s| !s.is_empty()).collect();

    println!(
        "{:?}",
        to_int(filter_vec(&vals, true)[0]) * to_int(filter_vec(&vals, false)[0])
    )
}
