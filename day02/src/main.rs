fn main() {
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let lines = contents.split("\n").filter(|s| !s.is_empty());

    let mut depth = 0;
    let mut pos = 0;
    let mut aim = 0;

    for line in lines {
        let split: Vec<&str> = line.split(" ").collect();
        let command = split[0];
        let value: i32 = split[1].parse().unwrap();
        match command {
            "down" => aim += value,
            "up" => aim -= value,
            "forward" => {
                pos += value;
                depth += value * aim
            }
            _ => (),
        }
    }
    println!("{}", depth * pos);
}
