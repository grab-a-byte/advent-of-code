import 'dart:convert';
import 'dart:io';
import 'dart:math';

void solution() {
  var file = File("inputs/day_13.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 13 - Part 1 : $res1");
  // var res2 = partTwoSolution(input);
  // print("Day 13 - Part 2 : $res2");
}

//1388 - Too low
//71717 - Too High
//55144 - Just incorrect
//55193 - Also wrong
int partOneSolution(String input) {
  final inputLines =
      LineSplitter().convert(input).where((x) => x.isNotEmpty).toList();
  final List<Machine> machines = [];
  for (int i = 0; i < inputLines.length; i += 3) {
    final aParts = inputLines[i]
        .replaceAll("Button A: X+", "")
        .replaceAll("Y+", "")
        .split(", ");
    final aButton = Button(int.parse(aParts[0]), int.parse(aParts[1]));

    final bParts = inputLines[i + 1]
        .replaceAll("Button B: X+", "")
        .replaceAll("Y+", "")
        .split(", ");
    final bButton = Button(int.parse(bParts[0]), int.parse(bParts[1]));

    final prizeParts = inputLines[i + 2]
        .replaceAll("Prize: X=", "")
        .replaceAll("Y=", "")
        .split(", ");

    machines.add(Machine(
        aButton, bButton, int.parse(prizeParts[0]), int.parse(prizeParts[1])));
  }

  final minValues = machines
      .map((m) => m.getPressesForPrize(100))
      .where((x) => x != null)
      .map((m) => (m!.aPresses * 3) + m.bPresses);

  return minValues.reduce((a, b) => a + b);
}

class Button {
  int x, y;
  Button(this.x, this.y);
}

class Machine {
  Button a, b;
  int prizeX, prizeY;
  Machine(this.a, this.b, this.prizeX, this.prizeY);

  ({int aPresses, int bPresses})? getPressesForPrize(int maxPresses) {
    final maxVal = max((prizeX / b.x).floor(), (prizeY / b.y).floor());
    for (int i = min(maxVal, maxPresses); i > 0; i--) {
      final leftoverX = prizeX - (b.x * i);
      final leftoverY = prizeY - (b.y * i);
      if (leftoverX % a.x == 0 && leftoverY % a.y == 0) {
        final aPresses = (leftoverY / a.y).floor();
        if (aPresses < 0) break;
        return (aPresses: aPresses, bPresses: i);
      }
    }
    return null;
  }
}
