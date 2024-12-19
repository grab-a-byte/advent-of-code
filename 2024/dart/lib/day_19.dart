
import 'dart:convert';
import 'dart:io';

void solution() {
  var file = File("inputs/day_19.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 19 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 19 - Part 2 : $res2");
}

int partOneSolution(String input){
  final lines = LineSplitter().convert(input);
  final patterns = lines[0].split(", ").map((e) => e.trim()).toList();
  final designs = lines.skip(2).toList();

  final chars = patterns.expand((x) => x.split("")).toSet();
  final useableDesigns = designs.where((x) => x.split("").toSet().difference(chars).isEmpty).toList();


  return getMakeableDesigns(useableDesigns, patterns).length;
}

int partTwoSolution(String input){
  final lines = LineSplitter().convert(input);
  final patterns = lines[0].split(", ").map((e) => e.trim()).toList();
  final designs = lines.skip(2).toList();

  final chars = patterns.expand((x) => x.split("")).toSet();
  final useableDesigns = designs.where((x) => x.split("").toSet().difference(chars).isEmpty).toList();

  final makeableDesigns = getMakeableDesigns(useableDesigns, patterns);
  int total = 0;
  for(var design in makeableDesigns){
    total += howManyWays(design, "", patterns);
  }

  return total;
}

int howManyWays(String design, String current, List<String> patterns){
  var newDesign = design.split("").skip(current.length).join();
  final usablePatterns = patterns.where(newDesign.startsWith);
  var total = 0;
  for (var pattern in usablePatterns.toSet()){
    if(current + pattern == design){
      total++;
      continue;
    }
    total += usablePatterns.where((x) => x == pattern).length * (howManyWays(design, current+pattern, patterns));
  }

  return total;
}

List<String> getMakeableDesigns(List<String> designs, List<String> patterns){
  final List<String> makeableDesigns = [];
  for (var design in designs){
    var toCheck = patterns.where((x) => design.startsWith(x)).toList();
    while(toCheck.isNotEmpty){
      final checking = toCheck.removeAt(0);
      if (checking == design){
        makeableDesigns.add(design);
        break;
      }
      final left = design.split("").skip(checking.length).join();
      final nextPatterns = patterns.where((x) => left.startsWith(x)).toList();
      toCheck.addAll(nextPatterns.map((e) => checking + e));
      toCheck = toCheck.toSet().toList();
    }
  }
  return makeableDesigns;
}