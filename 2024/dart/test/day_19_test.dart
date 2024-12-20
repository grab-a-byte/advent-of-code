import 'package:dart/day_19.dart' as day_19;
import 'package:test/test.dart';

const testInput = """r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb""";

void main() {
  test("Day 19 - Part 1", () {
    final result = day_19.partOneSolution(testInput);
    expect(result, 6);
  });

  test("Day 19 - Part 2", (){
    final result = day_19.partTwoSolution(testInput);
    expect(result, 16);
  });
}
