use std::fs;

pub mod day_01;

fn main() {
    println!("{}", std::env::current_dir().unwrap().display());
    let input1 = fs::read_to_string("./inputs/day_01.txt").unwrap();
    let answer1 = day_01::day_one_part_one(input1.clone());
    let answer1_2 = day_01::day_one_part_two(input1.clone());
    println!("{}", answer1);
    println!("{}", answer1_2);
}
