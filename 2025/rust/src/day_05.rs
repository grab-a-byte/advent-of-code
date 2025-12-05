pub fn day_five_part_one_just_plain_daft_approach(input: &str) -> u128 {
    let fresh_ranges = input.lines().take_while(|l| !l.is_empty());
    let fresh_values: Vec<u128> = fresh_ranges
        .flat_map(|r| {
            let parts: Vec<&str> = r.split('-').collect();
            let min: u128 = parts[0].parse().unwrap();
            let max: u128 = parts[1].parse().unwrap();
            min..=max
        })
        .collect();

    let items: Vec<u128> = input
        .lines()
        .skip_while(|l| !l.is_empty())
        .skip(1)
        .map(|n| n.parse::<u128>().unwrap())
        .collect();

    let result = items.iter().filter(|i| fresh_values.contains(i)).count();
    result as u128
}

pub fn day_five_part_one(input: &str) -> u128 {
    let fresh_ranges: Vec<&str> = input.lines().take_while(|l| !l.is_empty()).collect();

    let items: Vec<u128> = input
        .lines()
        .skip_while(|l| !l.is_empty())
        .skip(1)
        .map(|n| n.parse::<u128>().unwrap())
        .collect();

    let mut count: u128 = 0;

    for i in items {
        for r in fresh_ranges.clone() {
            let parts: Vec<&str> = r.split('-').collect();
            let min: u128 = parts[0].parse().unwrap();
            let max: u128 = parts[1].parse().unwrap();
            let full_range = min..=max;
            if full_range.contains(&i) {
                count += 1;
                break;
            }
        }
    }

    count
}

pub fn day_five_part_two(input: &str) -> u128 {
    let mut fresh_ranges: Vec<(u128, u128)> = input
        .lines()
        .take_while(|l| !l.is_empty())
        .map(|l| {
            let parts: Vec<&str> = l.split('-').collect();
            let min: u128 = parts[0].parse().unwrap();
            let max: u128 = parts[1].parse().unwrap();
            (min, max)
        })
        .collect();

    fresh_ranges.sort_by(|a, b| a.1.partial_cmp(&b.1).unwrap());

    let (initial_range_start, initial_range_end) = fresh_ranges[0];
    let mut result: u128 = (initial_range_start..=initial_range_end).count() as u128;

    for i in 1..fresh_ranges.len() {
        let (_, previous_end) = fresh_ranges[i - 1];
        let (current_start, current_end) = fresh_ranges[i];
        if current_start < previous_end {
            result += (previous_end + 1..=current_end).count() as u128;
        } else {
            result += (current_start..=current_end).count() as u128;
        }
    }

    result as u128
}

#[cfg(test)]
mod test {

    use super::*;

    const TEST_INPUT: &str = "3-5
10-14
16-20
12-18

1
5
8
11
17
32";

    #[test]
    pub fn part_one_test() {
        let result = day_five_part_one(TEST_INPUT);
        assert_eq!(3, result);
    }

    #[ignore = "Passes on test input but currently fails on main input"]
    #[test]
    pub fn part_two_test() {
        let result = day_five_part_two(TEST_INPUT);
        assert_eq!(14, result);
    }
}
