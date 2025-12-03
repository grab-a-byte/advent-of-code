pub fn day_two_part_one(input: &str) -> u128 {
    let parts: Vec<&str> = input.split(",").collect();
    let ranges: Vec<(u64, u64)> = parts
        .into_iter()
        .map(|line| {
            let mut line_parts = line.split("-");
            let min: u64 = line_parts.next().unwrap().trim().parse().unwrap();
            let max: u64 = line_parts.next().unwrap().trim().parse().unwrap();
            (min, max)
        })
        .collect();

    let mut numbers: Vec<u128> = Vec::new();
    for range in ranges {
        let (min, max) = range;
        let r = min..=max;
        for num in r {
            let str = num.to_string();
            if str.len() % 2 != 0 {
                continue;
            };
            let mid_point = str.len() / 2;
            let left = &str[0..mid_point];
            let right = &str[mid_point..];

            if left.starts_with(right) {
                numbers.push(num as u128);
            }
        }
    }

    numbers.iter().fold(0, |acc, elem| acc + elem)
}

pub fn day_two_part_two(input: &str) -> u32 {
    42
}

#[cfg(test)]
mod test {
    use super::*;

    const TEST_INPUT: &str = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

    #[test]
    pub fn part_one_test() {
        let answer = day_two_part_one(TEST_INPUT);
        assert_eq!(answer, 1227775554);
    }

    //Expect to fail as unimplemented
    #[ignore = "Not yet implemented"]
    #[test]
    pub fn part_two_test() {
        let answer = day_two_part_two(TEST_INPUT);
        assert_eq!(answer, 4174379265);
    }
}
