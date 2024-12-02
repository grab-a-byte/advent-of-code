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
  var lines = LineSplitter().convert(input).map((l) => l.split("   "));
  var left = lines.map((p) => int.parse(p[0])).toList();
  var right = lines.map((p) => int.parse(p[1])).toList();

  left.sort();
  right.sort();

  var result = 0;
  for (int i = 0; i < left.length; i++) {
    result += (left[i] - right[i]).abs();
  }

  return result;
}

int partTwoSolution(String input) {
  var lines = LineSplitter().convert(input).map((l) => l.split("   "));
  var left = lines.map((p) => int.parse(p[0])).toList();
  var right = lines.map((p) => int.parse(p[1])).toList();

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
