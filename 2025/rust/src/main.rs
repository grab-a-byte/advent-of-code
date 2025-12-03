use std::fs;

pub mod day_01;
pub mod day_02;

fn main() {
    // println!("{}", std::env::current_dir().unwrap().display());
    // let input1 = fs::read_to_string("./inputs/day_01.txt").unwrap();
    // let answer1 = day_01::day_one_part_one(input1.clone());
    // let answer1_2 = day_01::day_one_part_two(input1.clone());
    // println!("{}", answer1);
    // println!("{}", answer1_2);

    let input2 = fs::read_to_string("./inputs/day_02.txt").unwrap();
    let answer1 = day_02::day_two_part_one(input2.as_str());
    // let answer1_2 = day_01::day_two_part_two(input1.clone());
    println!("{}", answer1);
    // println!("{}", answer1_2);
}
