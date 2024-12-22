import 'dart:collection';
import 'dart:convert';
import 'dart:io';
import 'dart:math';

import 'package:dart/iterable_extensions.dart';

void solution() {
  var file = File("inputs/day_21.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 21 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 21 - Part 2 : $res2");
}

int partOneSolution(String input) {
  final lines =
      LineSplitter().convert(input).map((x) => int.parse(x.trim())).toList();

  int total = 0;
  for (int l = 0; l < lines.length; l++) {
    int secretNumber = lines[l];
    for (int i = 0; i < 2000; i++) {
      //Stage One
      var a = secretNumber * 64;
      secretNumber = secretNumber.mix(a);
      secretNumber = secretNumber.prune();

      //Stage Two
      var b = (secretNumber / 32).floor();
      secretNumber = secretNumber.mix(b);
      secretNumber = secretNumber.prune();

      //Stage 3
      var c = secretNumber * 2048;
      secretNumber = secretNumber.mix(c);
      secretNumber = secretNumber.prune();
    }
    total += secretNumber;
  }

  return total;
}

int partTwoSolution(String input) {
  final lines =
      LineSplitter().convert(input).map((x) => int.parse(x.trim())).toList();

  List<List<int>> secretNumbers = [];
  for (int l = 0; l < lines.length; l++) {
    int secretNumber = lines[l];
    List<int> sequence = [secretNumber];
    for (int i = 0; i < 2000; i++) {
      //Stage One
      var a = secretNumber * 64;
      secretNumber = secretNumber.mix(a);
      secretNumber = secretNumber.prune();

      //Stage Two
      var b = (secretNumber / 32).floor();
      secretNumber = secretNumber.mix(b);
      secretNumber = secretNumber.prune();

      //Stage 3
      var c = secretNumber * 2048;
      secretNumber = secretNumber.mix(c);
      secretNumber = secretNumber.prune();

      sequence.add(secretNumber);
    }
    secretNumbers.add(sequence);
  }

  final prices = secretNumbers
      .map((x) => x.map((i) {
            final s = i.toString();
            return int.parse(s[s.length - 1]);
          }).toList())
      .toList();

  final changes = prices
      .map((x) => x.sliding(2).map((s) => s.elementAt(1) - s.first).toList())
      .toList();

  List<LinkedHashMap<List<int>, int>> parts = List.generate(
      changes.length,
      (_) => LinkedHashMap(
            equals: (p0, p1) => p0.isEqual(p1),
          ));

  for (int i = 0; i < changes.length; i++) {
    final windows = changes[i].sliding(4);
    for (final window in windows) {
      if (parts[i].containsKey(window.toList())) continue;
      final pos = changes[i].sliding(4).indexed.firstWhere(
            (iter) => window.isEqual(iter.$2),
            orElse: () => (-1, []),
          );
      if (pos.$1 == -1) continue;
      parts[i][window.toList()] = prices[i][pos.$1 + 4];
    }
  }

  LinkedHashMap<List<int>, int> totals = LinkedHashMap(
    equals: (p0, p1) => p0.isEqual(p1),
  );
  for (var map in parts) {
    for (var kv in [
      [-2, 1, 2, -3]
    ]) {
      //If kv exists, skip
      if (totals.containsKey(kv)) continue;
      var total = parts.map((x) => x[kv] ?? 0).reduce((a, b) => a + b);
      totals[kv] = total;
    }
  }

//find highest in values
  return totals.values.reduce(max);
  return 42;
}
// total = total + prices[i][pos.$1 + 4];

extension IntExensions on int {
  int prune() => this % 16777216;
  int mix(int mixing) => this ^ mixing;
}
