pub fn day_four_part_one(input: &str) -> u128 {
    let parsed: Vec<Vec<u8>> = input
        .lines()
        .filter(|l| !l.is_empty())
        .map(|l| l.chars().map(|c| if c == '@' { 1 } else { 0 }).collect())
        .collect();

    let max_x = parsed[0].len();
    let max_y = parsed.len();

    let mut result: u128 = 0;
    for x in 0..max_x {
        for y in 0..max_y {
            let count = check_surrounding_rolls(&parsed, y, x);
            if count < 4 {
                result += 1
            }
        }
    }

    result
}

pub fn day_four_part_two(input: &str) -> u128 {
    let parsed: Vec<Vec<u8>> = input
        .lines()
        .filter(|l| !l.is_empty())
        .map(|l| l.chars().map(|c| if c == '@' { 1 } else { 0 }).collect())
        .collect();

    let max_x = parsed[0].len();
    let max_y = parsed.len();

    let mut result: u128 = 0;
    let mut keep_running = true;

    let mut initial = parsed.clone();
    let mut copy = initial.clone();

    while keep_running {
        keep_running = false;
        for x in 0..max_x {
            for y in 0..max_y {
                let count = check_surrounding_rolls(&initial, y, x);
                if count < 4 {
                    result += 1;
                    copy[y][x] = 0;
                    keep_running = true;
                }
            }
        }

        initial = copy;
        copy = initial.clone();
    }

    result
}

fn check_surrounding_rolls(grid: &Vec<Vec<u8>>, y: usize, x: usize) -> u32 {
    if grid[y][x] == 0 {
        return 420; //ensure not counted
    }

    fn get_grid_position(grid: &Vec<Vec<u8>>, y: i32, x: i32) -> u8 {
        if x < 0 || y < 0 {
            return 0;
        }
        grid.get(y as usize)
            .and_then(|l| l.get(x as usize))
            .copied()
            .unwrap_or(0)
    }

    let y: i32 = y as i32;
    let x: i32 = x as i32;

    (get_grid_position(grid, y - 1, x)
        + get_grid_position(grid, y + 1, x)
        + get_grid_position(grid, y - 1, x - 1)
        + get_grid_position(grid, y + 1, x + 1)
        + get_grid_position(grid, y - 1, x + 1)
        + get_grid_position(grid, y + 1, x - 1)
        + get_grid_position(grid, y, x + 1)
        + get_grid_position(grid, y, x - 1))
    .into()
}

#[cfg(test)]
mod test {

    use super::*;

    const TEST_INPUT: &str = "..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
";

    #[test]
    pub fn part_one_test() {
        let result = day_four_part_one(TEST_INPUT);
        assert_eq!(13, result);
    }

    #[test]
    pub fn part_two_test() {
        let result = day_four_part_two(TEST_INPUT);
        assert_eq!(43, result);
    }
}
