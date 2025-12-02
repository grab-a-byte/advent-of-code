
import 'dart:collection';
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

  return designs.map((d) => getPermutationsCount(d, patterns, {})).fold(0, (a,b) => a + b);
}

int getPermutationsCount(String design, List<String> patterns, Map<String, int> cache){
  if(cache.containsKey(design)) return cache[design]!;
  var count = 0;
  for(final pattern in patterns){
    if(pattern == design) {
      count++;
    } else if (design.startsWith(pattern)){
      count += getPermutationsCount(design.substring(pattern.length), patterns, cache);
    }
  }
  cache[design] = count;
  return count;
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