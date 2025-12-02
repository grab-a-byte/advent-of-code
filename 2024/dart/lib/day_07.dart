import 'dart:convert';
import 'dart:io';

void solution() {
  var file = File("inputs/day_07.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 7 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 7 - Part 2 : $res2");
}

int partOneSolution(String input) {
  var c = parseInput(input);
  var total = 0;
  for (int i = 0; i < c.length; i++) {
    final checking = c[i];
    final List<List<String>> toCheck =
        getPermutations("+*", checking.calibration.length - 1);
    if (checkLines(toCheck, checking.calibration, checking.target)) {
      total += checking.target;
    }
  }

  return total;
}

int partTwoSolution(String input) {
  var c = parseInput(input);
  var total = 0;
  for (int i = 0; i < c.length; i++) {
    final checking = c[i];
    final List<List<String>> toCheck =
        getPermutations("+*|", checking.calibration.length - 1);
    if (checkLines(toCheck, checking.calibration, checking.target)) {
      total += checking.target;
    }
  }

  return total;
}

bool checkLines(List<List<String>> operators, List<int> value, int target) {
  for (int i = 0; i < operators.length; i++) {
    int total = value[0];
    for (int o = 0; o < operators[i].length; o++) {
      switch (operators[i][o]) {
        case "*":
          total = total * value[o + 1];
          break;
        case "+":
          total = total + value[o + 1];
          break;
        case "|":
          total = int.parse(total.toString() + value[o + 1].toString());
          break;
      }
    }
    if (total == target) return true;
  }
  return false;
}

List<List<String>> getPermutations(String validChars, int length) {
  final chars = validChars.split("").toSet().toList();
  return getPermutationsB(chars, "", chars.length, length);
}

List<List<String>> getPermutationsB(
    List<String> chars, String prefix, int n, int k) {
  List<List<String>> possibilities = [];

  if (k == 0) {
    return [prefix.split("")];
  }

  for (int i = 0; i < n; ++i) {
    String newPrefix = prefix + chars[i];

    possibilities.addAll(getPermutationsB(chars, newPrefix, n, k - 1));
  }
  return possibilities;
}

List<({List<int> calibration, int target})> parseInput(String input) {
  return LineSplitter().convert(input).map((l) {
    final parts = l.split(": ");
    final target = int.parse(parts[0]);
    final calibration = parts[1].split(" ").map(int.parse).toList();
    return (target: target, calibration: calibration);
  }).toList();
}
