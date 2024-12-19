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

  test("How many ways", (){
    final patterns = [
      "r",
      "wr",
      "b",
      "g",
      "bwu",
      "rb",
      "gb",
      "br"
    ];
    var check1 = day_19.howManyWays("brwrr", "", patterns);
    expect(check1, 2);

    var check2 = day_19.howManyWays("bggr", "", patterns);
    expect(check2, 1);

    var check3 = day_19.howManyWays("gbbr", "", patterns);
    expect(check3, 4);

    var check4 = day_19.howManyWays("rrbgbr", "", patterns);
    expect(check4, 6);
  });

  test("TEMP", (){
    String input = """a, b, aa, bb

aab
bba
""";
  var b = day_19.partTwoSolution(input);
  print(b);
  });
}
