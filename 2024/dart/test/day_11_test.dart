import 'package:test/test.dart';
import 'package:dart/day_11.dart' as day_11;

const String testInput = '125 17';
void main(){
  test("Day 11 - Part One", (){
    final result = day_11.partOneSolution(testInput);
    expect(result, 55312);
  });

  test("Day 11 - Part 2", (){
    final result = day_11.solve2(testInput, 25);
    expect(result, 55312);
  });
}