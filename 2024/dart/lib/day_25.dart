import 'dart:convert';
import 'dart:io';

void solution() {
  var file = File("inputs/day_25.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 25 - Part 1 : $res1");
  // var res2 = partTwoSolution(input);
  // print("Day 22 - Part 2 : $res2");
}

int partOneSolution(String input){
  final parts = input.replaceAll("\r", "").split("\n\n");
  final locks = parts.where((x) => LineSplitter().convert(x).toList()[0][0] == "#").toList();
  final keys = parts.where((x) => LineSplitter().convert(x).toList()[0][0] == ".").toList();

   var lockHeights = locks.map((l){
    final grid = LineSplitter.split(l).map((s) => s.split("").toList()).toList();
    List<int> values = [];
    for(int i = 0; i < grid[0].length; i++){
      int val = 0;
      for (int y = 0; y < grid.length; y++){
        if(grid[y][i] == "#") val++;
      }
      values.add(val -1 );
    }
    return values;
   }).toList();


   var keyHeights = keys.map((k){
    final grid = LineSplitter.split(k).map((s) => s.split("").toList()).toList();
    List<int> values = [];
    for(int i = 0; i < grid[0].length; i++){
      int val = 0;
      for (int y = 0; y < grid.length; y++){
        if(grid[y][i] == "#") val++;
      }
      values.add(val -1 );
    }
    return values;
   }).toList();


  int count = 0;
  for(var key in keyHeights){
    for (var lock in lockHeights){
      final inc = zip(key, lock).map((e) => e.$1 + e.$2).every((x) => x < 6);
      if(inc) count++;
    }
  }
   return count;
}

List<(T,J)> zip<T,J>(Iterable<T> a, Iterable<J> b){
  List<(T,J)> items = [];
  for(int i = 0; i < a.length; i++){
    items.add((a.elementAt(i), b.elementAt(i)));
  }
  return items;
}