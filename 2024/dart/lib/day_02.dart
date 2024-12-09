import 'dart:convert';
import 'dart:io';

import 'package:dart/iterable_extensions.dart';

void solution() {
  var file = File("inputs/day_02.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 2 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 2 - Part 2 : $res2");
}

int partOneSolution(String input) {
  final Iterable<Iterable<int>> grid =
      LineSplitter().convert(input).map((x) => x.split(" ").map(int.parse));

  return grid.map(isSafe).where((x) => x).length;
}

int partTwoSolution(String input) {
  final Iterable<Iterable<int>> grid =
      LineSplitter().convert(input).map((x) => x.split(" ").map(int.parse));

  final safeLines = grid.where(isSafe);
  final fixedLines = grid
    .where((x) =>!isSafe(x))
    .map((x)=> x.removeOne())
    .where((x) => x.any(isSafe)).length;

  return safeLines.length + fixedLines;
}

bool isSafe(Iterable<int> row) => row.sliding(2).every((x) {
      final num = x.first - x.last;
      final rowIsNegative = (row.first - row.skip(1).first) < 0;
      return num.abs() < 4 &&
          num.abs() > 0 &&
          (rowIsNegative ? num < 0 : num > 0);
    });

extension Sliding<T> on Iterable<T> {
  Iterable<Iterable<T>> removeOne() sync* {
    final list = toList();
    for (int i = 0; i < length; i++) {
      yield list.sublist(0, i) + list.sublist(i+1, length);
    }
  }
}
