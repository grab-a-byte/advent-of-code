import 'package:test/test.dart';
import 'package:dart/day_18.dart' as day_18;

const testInput = """5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0""";

void main() {
  test("Day 18- Part One",  (){
    var result = day_18.partOneSolution(testInput, 7, 12);
    expect(result, 22);
  });

  test("Day 18 - Part Two", (){
    var result = day_18.partTwoSolution(testInput, 7, 12);
    expect(result, "6,1");
  });
}