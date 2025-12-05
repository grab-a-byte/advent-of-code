pub fn day_three_part_one(input: &str) -> u32 {
    let sanitized_input = input.replace("\r\n", "\n");
    let lines = sanitized_input.lines().filter(|l| !l.is_empty());
    let mut numbers: Vec<u32> = Vec::new();
    for line in lines {
        let start = &line[0..line.len() - 1];
        let (first, first_index) = get_highest_digit_from_string(start);
        let rest = &line[first_index + 1..line.len()];
        let (second, _second_index) = get_highest_digit_from_string(rest);

        let full = first.to_string() + &second.to_string();
        let num: u32 = full.parse().unwrap();
        numbers.push(num);
    }
    numbers.iter().fold(0, |acc, el| acc + el)
}

pub fn day_three_part_two(input: &str) -> u128 {
    let sanitized_input = input.replace("\r\n", "\n");
    let lines = sanitized_input.lines().filter(|l| !l.is_empty());
    let mut numbers: Vec<u128> = Vec::new();
    for line in lines {
        let mut current_str: String = String::from("");
        let mut start = 0;
        let mut count = 11;

        for _ in 0..12 {
            let segment = &line[start..line.len() - count];
            let (value, idx) = get_highest_digit_from_string(segment);
            current_str = current_str.to_owned() + &value.to_string();
            start += idx + 1;
            count = if count == 0 { 0 } else { count - 1 }
        }

        let num: u128 = current_str.parse().unwrap();

        numbers.push(num);
    }
    numbers.iter().fold(0, |acc, el| acc + el)
}

fn get_highest_digit_from_string(input: &str) -> (char, usize) {
    let nums: Vec<u32> = input
        .chars()
        .map(|c| c.to_string().parse().unwrap())
        .collect();

    let max = char::from_digit(nums.iter().max().unwrap().clone(), 10).unwrap();
    let idx = input.chars().position(|el| el == max).unwrap();

    (max, idx)
}

#[cfg(test)]
mod test {

    use super::*;

    const TEST_INPUT: &str = "987654321111111
811111111111119
234234234234278
818181911112111";

    #[test]
    pub fn part_one_test() {
        let result = day_three_part_one(TEST_INPUT);
        assert_eq!(357, result);
    }

    #[test]
    pub fn part_two_test() {
        let result = day_three_part_two(TEST_INPUT);
        assert_eq!(3121910778619, result);
    }
}
