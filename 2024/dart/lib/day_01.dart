import 'dart:convert';
import 'dart:io';

void solution() {
  var file = File("inputs/day_01.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 1 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 1 - Part 2 : $res2");
}

int partOneSolution(String input) {
  List<int> left = List.empty(growable: true);
  List<int> right = List.empty(growable: true);
  var lines = LineSplitter().convert(input);

  for (var line in lines) {
    final parts = line.split("   ");
    final leftNum = int.parse(parts[0]);
    final rightNum = int.parse(parts[1]);
    left.add(leftNum);
    right.add(rightNum);
  }

  left.sort();
  right.sort();

  var result = 0;
  for (int i = 0; i < left.length; i++) {
    result += (left[i] - right[i]).abs();
  }

  return result;
}

int partTwoSolution(String input) {
  List<int> left = List.empty(growable: true);
  List<int> right = List.empty(growable: true);
  var lines = LineSplitter().convert(input);

  for (var line in lines) {
    final parts = line.split("   ");
    final leftNum = int.parse(parts[0]);
    final rightNum = int.parse(parts[1]);
    left.add(leftNum);
    right.add(rightNum);
  }

  Map<int,int> counts = {};
  for(var val in right){
    if(counts.containsKey(val)){
      counts[val] = counts[val]! + 1;
    } else {
      counts[val] = 1;
    }
  }

  var result = 0;
  for (int i = 0; i < left.length; i++) {
    if(counts.containsKey(left[i])){
      result += left[i] * counts[left[i]]!;
    }
  }

  return result;
}
