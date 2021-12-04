fn main() {
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let mut previous = 0;
    let mut count = -1;
    let vals: Vec<i32> = contents
        .split("\n")
        .filter(|s| !s.is_empty())
        .map(|s| s.parse().unwrap())
        .collect();

    for i in 0..(vals.len() - 2) {
        let next = vals[i] + vals[i + 1] + vals[i + 2];
        if next > previous {
            count += 1;
        }
        previous = next;
    }

    println!("{}", count);
}
