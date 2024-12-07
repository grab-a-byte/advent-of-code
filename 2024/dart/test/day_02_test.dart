import 'package:dart/day_02.dart' as day_02;
import 'package:test/test.dart';

var sampleInput = """7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9""";

void main() {
  test("Day Two - Part One", () {
    var result = day_02.partOneSolution(sampleInput);
    expect(result, 2);
  });

  test("Day Two - Part Two", () {
    var result = day_02.partTwoSolution(sampleInput);
    expect(result, 4);
  });

  test("Remove One, returns as expected", () {
    final input = [1, 2, 3];
    var results = input.removeOne();
    final expected = [
      [1, 2],
      [2, 3],
      [1, 3]
    ];

    for (var res in results) {
      final r = expected.any((x) => _listEquals(x, res.toList()));
      assert(r, true);
    }
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
