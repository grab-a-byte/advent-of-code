import 'package:dart/day_01.dart' as day_01;
import 'package:test/test.dart';

void main() {
  var testInput = """3   4
4   3
2   5
1   3
3   9
3   3
""";

  test('Day One Part One', () {
      var result = day_01.partOneSolution(testInput);
      expect(result, 11);
  });

  test('Day One Part Two', () {
      var result = day_01.partTwoSolution(testInput);
      expect(result, 31);
  });
}
