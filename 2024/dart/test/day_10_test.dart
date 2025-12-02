import 'dart:convert';

import 'package:test/test.dart';
import 'package:dart/day_10.dart' as day_10;

const testInput = """89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732""";

void main(){
  test('Day 10 - Part 1', () {
    final int result =  day_10.partOneSolution(testInput);
    expect(result, 36);
  });

  test('DepthFirstSearchCount', (){
  final grid = LineSplitter().convert(testInput)
    .map((x) => x.split('')
      .map((i) => int.parse(i)).toList()).toList();

    final int checkOne = day_10.depthFirstSearchCount((x: 2, y: 0), grid);
    expect(checkOne, 5);
    final int checkTwo = day_10.depthFirstSearchCount((x: 4, y: 0), grid);
    expect(checkTwo, 6);
    final int checkThree = day_10.depthFirstSearchCount((x: 4, y: 2), grid);
    expect(checkThree, 5);
    final int checkFour = day_10.depthFirstSearchCount((x: 6, y: 4), grid);
    expect(checkFour, 3);
    final int checkSix = day_10.depthFirstSearchCount((x: 2, y: 5), grid);
    expect(checkSix, 1);
    final int checkSeven = day_10.depthFirstSearchCount((x: 5, y: 5), grid);
    expect(checkSeven, 3);
    final int checkEight = day_10.depthFirstSearchCount((x: 0, y: 6), grid);
    expect(checkEight, 5);
    final int checkNine = day_10.depthFirstSearchCount((x: 6, y: 6), grid);
    expect(checkNine, 3);
    final int checkTen = day_10.depthFirstSearchCount((x: 1, y: 7), grid);
    expect(checkTen, 5);
  });


  const partTwoTestInput = """89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732""";

test("Day 10 - Part Two", (){
  var output = day_10.partTwoSolution(partTwoTestInput);
  expect(output, 81);
});
}