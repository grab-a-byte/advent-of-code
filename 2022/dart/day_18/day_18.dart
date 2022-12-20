import 'dart:collection';
import 'dart:io';
import 'dart:math';

abstract class DayEighteen {
  static Future Solution() async {
    File file = File("day_18${Platform.pathSeparator}day_18_input.txt");
    List<String> input = await file.readAsLines();
    print("Part 1 :-  ${PartOne(input)}");
    print("Part 2 :-  ${PartTwo(input)}");
  }

  static int PartOne(List<String> input) {
    List<List<int>> parsed = _parseInput(input);
    int result = 0;
    for (var x in parsed) {
      bool containsLeft = _containsByValue(parsed, [x[0] - 1, x[1], x[2]]);
      bool containsRight = _containsByValue(parsed, [x[0] + 1, x[1], x[2]]);
      bool containsDown = _containsByValue(parsed, [x[0], x[1] - 1, x[2]]);
      bool containsUp = _containsByValue(parsed, [x[0], x[1] + 1, x[2]]);
      bool containsForward = _containsByValue(parsed, [x[0], x[1], x[2] - 1]);
      bool containsBack = _containsByValue(parsed, [x[0], x[1], x[2] + 1]);

      result += [
        containsLeft,
        containsRight,
        containsUp,
        containsDown,
        containsForward,
        containsBack
      ].map((e) => e ? 0 : 1).reduce((value, element) => value + element);
    }

    return result;
  }

  //test passes, unable to get tree walk to not re-check itself
  //meaning it takes to long to run for full solution
  static int PartTwo(List<String> input) {
    List<List<int>> parsed = _parseInput(input);
    int result = 0;

    var xMin = parsed.map((e) => e[0]).reduce(min);
    var xMax = parsed.map((e) => e[0]).reduce(max);
    var yMin = parsed.map((e) => e[1]).reduce(min);
    var yMax = parsed.map((e) => e[1]).reduce(max);
    var zMin = parsed.map((e) => e[2]).reduce(min);
    var zMax = parsed.map((e) => e[2]).reduce(max);

    var mins = [xMin, yMin, zMin];
    var maxes = [xMax, yMax, zMax];

    for (var x in parsed) {
      print("Checking $x");
      List<int> left = [x[0] - 1, x[1], x[2]];
      List<int> right = [x[0] + 1, x[1], x[2]];
      List<int> down = [x[0], x[1] - 1, x[2]];
      List<int> up = [x[0], x[1] + 1, x[2]];
      List<int> forward = [x[0], x[1], x[2] - 1];
      List<int> back = [x[0], x[1], x[2] + 1];
      bool containsLeft = _containsByValue(parsed, left);
      bool containsRight = _containsByValue(parsed, right);
      bool containsDown = _containsByValue(parsed, up);
      bool containsUp = _containsByValue(parsed, down);
      bool containsForward = _containsByValue(parsed, forward);
      bool containsBack = _containsByValue(parsed, back);

      containsLeft =
          containsLeft ? containsLeft : _treeWalk(parsed, left, mins, maxes);
      containsRight =
          containsRight ? containsRight : _treeWalk(parsed, right, mins, maxes);
      containsDown =
          containsDown ? containsDown : _treeWalk(parsed, down, mins, maxes);
      containsUp = containsUp ? containsUp : _treeWalk(parsed, left, up, maxes);
      containsForward = containsForward
          ? containsForward
          : _treeWalk(parsed, forward, mins, maxes);
      containsBack =
          containsBack ? containsBack : _treeWalk(parsed, back, mins, maxes);

      result += [
        containsLeft,
        containsRight,
        containsUp,
        containsDown,
        containsForward,
        containsBack
      ].map((e) => e ? 0 : 1).reduce((value, element) => value + element);
    }

    return result;
  }

  static bool _treeWalk(List<List<int>> values, List<int> intialCheck,
      List<int> minVal, List<int> maxVal) {
    List<List<int>> visitedNodes = [];
    Queue queue = Queue<List<int>>()..add(intialCheck);

    while (queue.isNotEmpty) {
      var item = queue.removeFirst();
      print("Processing Item from Queue $item");
      visitedNodes.add(item);
      print("VisitedNodes = $visitedNodes");
      if ((item[0] < minVal[0] || item[0] > maxVal[0]) ||
          (item[1] < minVal[1] || item[1] > maxVal[1]) ||
          (item[2] < minVal[2] || item[2] > maxVal[2])) {
        return false;
      }
      List<int> left = [item[0] - 1, item[1], item[2]];
      List<int> right = [item[0] + 1, item[1], item[2]];
      List<int> down = [item[0], item[1] - 1, item[2]];
      List<int> up = [item[0], item[1] + 1, item[2]];
      List<int> forward = [item[0], item[1], item[2] - 1];
      List<int> back = [item[0], item[1], item[2] + 1];
      if (!_containsByValue(values, left) &&
          !_containsByValue(visitedNodes, left)) {
        queue.add(left);
      }
      if (!_containsByValue(values, right) &&
          !_containsByValue(visitedNodes, right)) {
        queue.add(right);
      }
      if (!_containsByValue(values, down) &&
          !_containsByValue(visitedNodes, down)) {
        queue.add(down);
      }
      if (!_containsByValue(values, up) &&
          !_containsByValue(visitedNodes, up)) {
        queue.add(up);
      }
      if (!_containsByValue(values, forward) &&
          !_containsByValue(visitedNodes, forward)) {
        queue.add(forward);
      }
      if (!_containsByValue(values, back) &&
          !_containsByValue(visitedNodes, back)) {
        queue.add(back);
      }

      print("BREAKING");
    }
    return true;
  }

  static bool _containsByValue(List<List<int>> values, List<int> check) =>
      values.any((element) =>
          element[0] == check[0] &&
          element[1] == check[1] &&
          element[2] == check[2]);

  static List<List<int>> _parseInput(List<String> input) {
    return input
        .map((e) => e.split(",").map((e) => int.parse(e)).toList())
        .toList();
  }
}
