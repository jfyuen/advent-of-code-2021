fn is9(digits: &[&str], digit: &str) -> bool {
    if digit.len() != 6 {
        return false;
    }
    let filtered: Vec<char> = digit
        .chars()
        .filter(|c| {
            let c_str = c.to_string();
            !digits[7].contains(&c_str) && !digits[4].contains(&c_str)
        })
        .collect();
    filtered.len() == 1
}

fn is0(digits: &[&str], digit: &str) -> bool {
    if digit.len() != 6 {
        return false;
    }
    let filtered: Vec<char> = digit
        .chars()
        .filter(|c| {
            let c_str = c.to_string();
            !digits[1].contains(&c_str)
        })
        .collect();
    filtered.len() == 4
}

fn is2(digits: &[&str], digit: &str) -> bool {
    if digit.len() != 5 {
        return false;
    }
    let filtered: Vec<char> = digit
        .chars()
        .filter(|c| {
            let c_str = c.to_string();
            !digits[9].contains(&c_str) && !digits[1].contains(&c_str)
        })
        .collect();
    filtered.len() == 1
}

fn is5(digits: &[&str], digit: &str) -> bool {
    if digit.len() != 5 {
        return false;
    }
    let filtered: Vec<char> = digit
        .chars()
        .filter(|c| {
            let c_str = c.to_string();
            !digits[2].contains(&c_str) && !digits[1].contains(&c_str)
        })
        .collect();
    filtered.len() == 1
}

fn main() {
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let mut outputs = vec![""; 10];
    let mut total = 0;

    for line in contents.split("\n") {
        let split: Vec<&str> = line.split(" | ").collect();
        let patterns: Vec<&str> = split[0].split(" ").collect();
        let digits: Vec<&str> = split[1].split(" ").collect();

        let mut not_founds = Vec::<&str>::new();
        for digit in &patterns {
            match digit.len() {
                2 => outputs[1] = digit, //1
                3 => outputs[7] = digit, //7
                4 => outputs[4] = digit, //4
                // 5 => 2, 3, 5
                // 6 => 0, 6, 9
                7 => outputs[8] = digit,     //8
                _ => not_founds.push(digit), // -1
            };
        }
        for not_found in &not_founds {
            if is9(&outputs, &not_found) {
                outputs[9] = not_found;
            }
        }
        not_founds = not_founds
            .into_iter()
            .filter(|e| e != &outputs[9])
            .collect();
        for not_found in &not_founds {
            if is2(&outputs, &not_found) {
                outputs[2] = not_found;
            }
        }
        not_founds = not_founds
            .into_iter()
            .filter(|e| e != &outputs[2])
            .collect();

        for not_found in &not_founds {
            if is5(&outputs, &not_found) {
                outputs[5] = not_found;
            }
        }
        not_founds = not_founds
            .into_iter()
            .filter(|e| e != &outputs[5])
            .collect();

        outputs[3] = not_founds.iter().filter(|e| e.len() == 5).next().unwrap(); // 3
        not_founds = not_founds
            .into_iter()
            .filter(|e| e != &outputs[3])
            .collect();

        for not_found in &not_founds {
            if is0(&outputs, &not_found) {
                outputs[0] = not_found;
            }
        }
        not_founds = not_founds
            .into_iter()
            .filter(|e| e != &outputs[0])
            .collect();

        outputs[6] = not_founds[0];

        for (i, digit) in digits.iter().enumerate() {
            for (j, output) in outputs.iter().enumerate() {
                if output.len() == digit.len()
                    && digit.chars().all(|c| output.contains(&c.to_string()))
                {
                    total += 10_i32.pow((digits.len() - i - 1) as u32) * (j as i32);
                    break;
                }
            }
        }
    }
    println!("{}", total);
}
