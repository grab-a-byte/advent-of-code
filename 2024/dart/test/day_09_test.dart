import 'package:test/test.dart';
import 'package:dart/day_09.dart' as day_09;

const _testInput = "2333133121414131402";

main(){
  test("Day 9 - Part 1", (){
    final result = day_09.partOneSolution(_testInput);
    expect(result, 1928);
  });

  test("Day 9 - Part 2", (){
    final result = day_09.partTwoSolution(_testInput);
    expect(result, 2858);
  });
}