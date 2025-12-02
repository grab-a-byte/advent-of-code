//This deffo needs re-writing 100%
pub fn day_one_part_one(input: String) -> u32 {
    let sanitized_input = input.replace("\r\n", "\n");
    let lines = sanitized_input.split("\n");
    let mut dial = 50;
    let mut counter = 0;
    for line in lines {
        let dir = if line.chars().next().unwrap() == 'L' {
            -1
        } else {
            1
        };
        let amount_string: String = line.chars().skip(1).collect();
        let amount: i32 = amount_string.parse().unwrap();
        dial += amount * dir;
        if (dial % 100) == 0 {
            counter += 1;
        }
    }

    return counter;
}

pub fn day_one_part_two(input: String) -> u32 {
    let sanitized_input = input.replace("\r\n", "\n");
    let lines = sanitized_input.split("\n");
    let moves = lines.flat_map(|line| {
        let dir = if line.chars().nth(0).unwrap() == 'L' {
            -1
        } else {
            1
        };
        let amount_string: String = line.chars().skip(1).collect();
        let amount: i32 = amount_string.parse().unwrap();
        return [dir].repeat(amount as usize);
    });

    let folded = moves.fold((50, 0), |acc, item| {
        let (dial, count) = acc;
        let new_dial = (dial + item) % 100;
        if new_dial == 0 {
            return (0, count + 1);
        }
        return (new_dial, count);
    });

    return folded.1;
}

#[cfg(test)]
mod test {
    use super::*;

    const TEST_INPUT: &str = "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82";

    #[test]
    fn test_part_one() {
        let answer = day_one_part_one(TEST_INPUT.to_string());
        assert_eq!(answer, 3);
    }

    #[test]
    fn test_part_two() {
        let answer = day_one_part_two(TEST_INPUT.to_string());
        assert_eq!(answer, 6)
    }
}
