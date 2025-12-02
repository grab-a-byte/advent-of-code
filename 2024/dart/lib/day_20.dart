import 'dart:convert';
import 'dart:io';

void solution() {
  var file = File("inputs/day_20.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 20 - Part 1 : $res1");
  // var res2 = partTwoSolution(input);
  // print("Day 20 - Part 2 : $res2");
}

int partOneSolution(String input) {
  final List<List<String>> grid = LineSplitter()
      .convert(input)
      .map((l) => l.split("").skip(1).take(l.length - 2).toList())
      .toList();

  grid.removeLast();
  grid.removeAt(0);

  late (int, int) start, end;
  List<(int, int)> objects = [];
  for (int y = 0; y < grid.length; y++) {
    for (int x = 0; x < grid[0].length; x++) {
      var place = grid[y][x];
      if (place == "#") {
        objects.add((x, y));
      } else if (place == "S") {
        start = (x, y);
      } else if (place == "E") {
        end = (x, y);
      }
    }
  }

  var baseValue =
      aStar(grid.length, grid[0].length, objects, start)[end.$2][end.$1];

  var total = 0;
  // Map<int, int> debugMap = {};
  for (int i = 0; i < objects.length; i++) {
    var newObjects = objects.take(i).toList()..addAll(objects.skip(i + 1));

    var newValue =
        aStar(grid.length, grid[0].length, newObjects, start)[end.$2][end.$1];

    // debugMap.update(baseValue - newValue, (v) => v+1, ifAbsent: () => 1);
    if (baseValue - newValue >= 100) {
      total++;
    }
  }
  // for(var m in debugMap.entries){
  //   print("${m.key}:${m.value}");
  // }
  return total;
}

// int partTwoSolution(String input) {
//   final List<List<String>> grid = LineSplitter()
//       .convert(input)
//       .map((l) => l.split("").skip(1).take(l.length - 2).toList())
//       .toList();

//   grid.removeLast();
//   grid.removeAt(0);

//   late (int, int) start, end;
//   List<(int, int)> objects = [];
//   for (int y = 0; y < grid.length; y++) {
//     for (int x = 0; x < grid[0].length; x++) {
//       var place = grid[y][x];
//       if (place == "#") {
//         objects.add((x, y));
//       } else if (place == "S") {
//         start = (x, y);
//       } else if (place == "E") {
//         end = (x, y);
//       }
//     }
//   }

//   var baseValue =
//       aStar(grid.length, grid[0].length, objects, start)[end.$2][end.$1];

//   int total = 0;

//   for (int i = 0; i < objects.length; i++) {
//     for (int skip = 1; skip < 50 && i + 50 < objects.length; skip++) {
//       var newObjects = objects.take(i).toList()..addAll(objects.skip(i + skip));

//       var newValue =
//           aStar(grid.length, grid[0].length, newObjects, start)[end.$2][end.$1];

//       if (baseValue - newValue >= 100) {
//         total++;
//       }
//     }
//   }
//   return total;
// }

const int maxInteger = 0x7FFFFFFFFFFFFFFF;

List<List<int>> aStar(
    int xLen, int yLen, List<(int, int)> objects, (int, int) start) {
  var arr = List.generate(yLen, (_) => List<int>.filled(xLen, maxInteger));
  for (var (x, y) in objects) {
    arr[y][x] = -1;
  }
  arr[start.$2][start.$1] = 0;
  List<(int, int)> toCheck = [start];
  while (toCheck.isNotEmpty) {
    var (x, y) = toCheck.removeAt(0);

    //Check left and add
    if (x > 0 && arr[y][x] + 1 < arr[y][x - 1]) {
      arr[y][x - 1] = arr[y][x] + 1;
      toCheck.add((x - 1, y));
    }

    //Check Right and add
    if (x < arr[0].length - 1 && arr[y][x] + 1 < arr[y][x + 1]) {
      arr[y][x + 1] = arr[y][x] + 1;
      toCheck.add((x + 1, y));
    }
    //Check up and add
    if (y > 0 && arr[y][x] + 1 < arr[y - 1][x]) {
      arr[y - 1][x] = arr[y][x] + 1;
      toCheck.add((x, y - 1));
    }

    //Check down and add
    if (y < arr.length - 1 && arr[y][x] + 1 < arr[y + 1][x]) {
      arr[y + 1][x] = arr[y][x] + 1;
      toCheck.add((x, y + 1));
    }
  }
  return arr;
}
