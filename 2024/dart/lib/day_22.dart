import 'dart:convert';
import 'dart:io';
import 'dart:math';
import 'package:dart/iterable_extensions.dart';

void solution() {
  var file = File("inputs/day_22.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 22 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 22 - Part 2 : $res2");
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
  List<int> totals = [];
  for (int i = 0; i < changes.length; i++) {
    final windows = changes[i].sliding(4);
    for (final window in windows) {
      var total = 0;
      for (int i = 0; i < changes.length; i++) {
        final pos = changes[i].sliding(4).indexed.firstWhere(
              (iter) => window.isEqual(iter.$2),
              orElse: () => (-1, []),
            );
        if (pos.$1 == -1) continue;
        total = total + prices[i][pos.$1 + 4];
      }
      totals.add(total);
    }
  }
  return totals.reduce(max);
}
extension IntExensions on int {
  int prune() => this % 16777216;
  int mix(int mixing) => this ^ mixing;
}