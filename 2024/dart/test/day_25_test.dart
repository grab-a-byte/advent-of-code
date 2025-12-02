import 'package:test/test.dart';
import 'package:dart/day_25.dart' as day_25;

const input = """#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####""";

void main(){
  test("Day 25 - Part One", (){
    final int result = day_25.partOneSolution(input);
    expect(result, 3);
  });
}
