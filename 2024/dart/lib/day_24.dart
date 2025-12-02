import 'dart:convert';
import 'dart:io';

void solution() {
  var file = File("inputs/day_24.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 24 - Part 1 : $res1");
  // var res2 = partTwoSolution(input);
  // print("Day 22 - Part 2 : $res2");
}

int partOneSolution(String input){
  final lines = LineSplitter.split(input);
  final starting = lines.takeWhile((s) => s.isNotEmpty);
  final gates = lines.skipWhile((s) => s.isNotEmpty).skip(1);

  final startParsed = <String, bool>{};
  for(final s in starting){
    final parts = s.split(": ");
    startParsed[parts[0]] = int.parse(parts[1]) == 1 ;
  }

  final mappedGates = gates.map((g) {
    final parts = g.split(" -> ").toList();
    final p = parts[0].split(" ").toList();
    final output = parts[1];
    final type = switch (p[1]){
      "XOR" => GateType.xor,
      "OR" => GateType.or,
      "AND" => GateType.and,
      _ => throw Exception("Unknown Type")
    };

    return Gate(p[0], p[2], type, output);
  }).toList();

  final values = mappedGates
    .where((x) => x.outputName.startsWith("z"))
    .map((g) => (g.outputName, getValue(g, startParsed, mappedGates))).toList();


  values.sort((a,b) => b.$1.compareTo(a.$1));
  final outputStr = values.map((e) => e.$2 ? "1" : "0").join();
  return int.parse(outputStr, radix: 2);
}

bool getValue(Gate gate, Map<String, bool> cache, List<Gate> gates){
  if(cache.containsKey(gate.outputName)) return cache[gate.outputName]!;
  final a = gates.firstWhere((g) => g.outputName == gate.a, orElse: () => Gate("", "", GateType.and, gate.a),);
  final b = gates.firstWhere((g) => g.outputName == gate.b, orElse: () => Gate("", "", GateType.and, gate.b),);
  final aVal = getValue(a, cache, gates);
  final bVal = getValue(b, cache, gates);
  return switch(gate.type){
    GateType.and => aVal && bVal,
    GateType.xor => aVal ^ bVal,
    GateType.or => aVal || bVal,
  };
}

enum GateType{
  and,
  xor,
  or,
}

class Gate {
  String a,b;
  GateType type;
  String outputName;

  Gate(this.a, this.b, this.type, this.outputName);
}