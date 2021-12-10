fn check_close(open: char, close: char) -> bool {
    match close {
        ')' => open == '(',
        ']' => open == '[',
        '}' => open == '{',
        '>' => open == '<',
        _ => false,
    }
}

fn check_corrupted(line: &str) -> Option<char> {
    let mut stack = Vec::<char>::new();
    for c in line.chars() {
        match c {
            '{' | '(' | '[' | '<' => stack.push(c),
            _ => {
                if !check_close(stack[stack.len() - 1], c) {
                    return Some(c);
                };
                stack.remove(stack.len() - 1);
            }
        }
    }
    None
}

fn get_corrupted_score(c: char) -> i32 {
    match c {
        ')' => 3,
        ']' => 57,
        '}' => 1197,
        '>' => 25137,
        _ => 0,
    }
}

fn get_closing(line: &str) -> Vec<char> {
    let mut stack = Vec::<char>::new();
    for c in line.chars() {
        match c {
            '{' | '(' | '[' | '<' => stack.push(c),
            _ => {
                stack.remove(stack.len() - 1);
            }
        }
    }
    let mut closings = Vec::<char>::new();
    for c in stack.iter().rev() {
        match c {
            '(' => closings.push(')'),
            '[' => closings.push(']'),
            '{' => closings.push('}'),
            '<' => closings.push('>'),
            _ => (),
        }
    }
    closings
}

fn get_autocomplete_score(chars: &Vec<char>) -> u64 {
    let mut score = 0;
    for c in chars {
        score *= 5;
        score += match c {
            ')' => 1,
            ']' => 2,
            '}' => 3,
            '>' => 4,
            _ => 0,
        }
    }
    score
}

fn main() {
    use std::env;
    use std::fs;
    let args: Vec<String> = env::args().collect();

    let contents = fs::read_to_string(&args[1]).unwrap();
    let lines: Vec<&str> = contents.split("\n").collect();

    // let mut score = 0;
    let mut incomplete_lines = Vec::<&str>::new();
    for line in lines {
        match check_corrupted(&line) {
            None => incomplete_lines.push(&line), //score += get_corrupted_score(c),
            _ => (),
        };
    }
    let mut scores: Vec<u64> = incomplete_lines
        .iter()
        .map(|line| get_autocomplete_score(&get_closing(&line)))
        .collect();
    scores.sort();
    println!("{}", scores[scores.len() / 2]);
}
