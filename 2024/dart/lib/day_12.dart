import 'dart:collection';
import 'dart:convert';
import 'dart:io';

void solution(){
  var file = File("inputs/day_12.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 12 - Part 1 : $res1");
  // var res2 = partTwoSolution(input);
  // print("Day 11 - Part 2 : $res2");
}

int partOneSolution(String input) {
  List<_Position> positions = [];
  var lines = LineSplitter().convert(input);
  for (int x = 0; x < lines[0].length; x++) {
    for (int y = 0; y < lines.length; y++) {
      positions.add(_Position(x, y, lines[y][x]));
    }
  }
  print(positions.length);

  List<_Region> regions = [];

  for (int i = 0; i < positions.length; i++) {
    final start = positions[i];
    if (start.checked) continue;
    start.checked = true;
    List<_Position> regionPositions = [];
    var toCheck = Queue<_Position>.from([start]);
    while (toCheck.isNotEmpty) {
      final checking = toCheck.removeFirst();
      checking.checked = true;
      regionPositions.add(checking);
      final toAdd = positions.where((p) => p.isAdjacent(checking) && p.value == checking.value && !p.checked).toList();
      toCheck.addAll(toAdd);
    }
    regions.add(_Region(regionPositions));
  }

  return regions.map((x) => x.cost()).reduce((a, b) => a + b);
}

class _Position {
  int x, y;
  String value;
  bool checked = false;

  _Position(this.x, this.y, this.value);

  bool isAdjacent(_Position check) {
    return (x == check.x - 1 && y == check.y) ||
        (x == check.x + 1 && y == check.y) ||
        (x == check.x && y == check.y + 1) ||
        (x == check.x && y == check.y - 1);
  }
}

class _Region {
  List<_Position> positions;
  int? _perimeter;

  _Region(List<_Position> positions) : positions = positions.toSet().toList();

  int area() => positions.length;

  int perimeter() {
    if (_perimeter != null) return _perimeter!;

    _perimeter = positions
        .map((p) => 4 - positions.where((c) => c.isAdjacent(p)).length)
        .reduce((a, b) => a + b);

    return _perimeter!;
  }

  int cost() => area() * perimeter();
}
