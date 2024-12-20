import 'package:test/test.dart';
import 'package:dart/day_20.dart' as day_20;


const testInput = """###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############""";

void main(){
  test("Day 20- Part 1", (){
    var res = day_20.partOneSolution(testInput);
    print(res);
  });
}