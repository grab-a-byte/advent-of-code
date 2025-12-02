import 'package:dart/day_07.dart';
import 'package:test/test.dart';

void main() {
  const testInput = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20""";

  test("Day Seven, Part One", () {
    var result = partOneSolution(testInput);
    expect(result, 3749);
  });

  test("Day Seven, Part Two", () {
    var result = partTwoSolution(testInput);
    expect(result, 11387);
  });

  test("Check Lines", () {
    final res1 = checkLines([
      ["*"]
    ], [
      10,
      19
    ], 190);
    expect(res1, true);
    final res2 = checkLines([
      ["+", "*"]
    ], [
      81,
      40,
      27
    ], 3267);
    expect(res2, true);
    final res3 = checkLines([
      ["+", "*", "+"]
    ], [
      11,
      6,
      16,
      20
    ], 292);
    expect(res3, true);

    //Check for failure
    final res4 = checkLines([
      ["*", "+", "+"]
    ], [
      81,
      40,
      27,
      50
    ], 3267);
    expect(res4, false);

    final res5 = checkLines([
      ["+"],
      ["*"]
    ], [
      10,
      19
    ], 190);
    expect(res5, true);
  });

  test("Check lines with concat", (){
    final res1 = checkLines([["|"]], [15, 6], 156);
    expect(res1, true);
    final res2 = checkLines([["*", "|", "*"]], [6,8,6,15], 7290);
    expect(res2, true);
  });

  test("GetPermutations returns correct permutations", () {
    final result = getPermutations("ab", 3);
    var expected = [
      "aaa".split(""),
      "aab".split(""),
      "aba".split(""),
      "abb".split(""),
      "baa".split(""),
      "bab".split(""),
      "bba".split(""),
      "bbb".split("")
    ];
    var check = expected.every((x) => result.any((y) => _listEquals(x, y)));
    expect(check, true);
    expect(expected.length, result.length);
  });
}

bool _listEquals<E>(List<E>? list1, List<E>? list2) {
  if (identical(list1, list2)) return true;
  if (list1 == null || list2 == null) return false;
  var length = list1.length;
  if (length != list2.length) return false;
  for (var i = 0; i < length; i++) {
    if (list1[i] != list2[i]) return false;
  }
  return true;
}
