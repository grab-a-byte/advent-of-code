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
}
