import 'dart:collection';
import 'dart:convert';
import 'dart:io';


void solution() {
  var file = File("inputs/day_10.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 10 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 10 - Part 2 : $res2");
}

int partOneSolution(String input){
  final grid = LineSplitter().convert(input)
    .map((x) => x.split('')
      .map((i) => int.parse(i)).toList()).toList();

  List<({int x, int y})> startingPositions = [];
  for(int y = 0; y < grid.length; y++){
    for(int x = 0; x < grid[0].length; x++){
      if(grid[y][x] == 0) {
        startingPositions.add((x: x, y: y));
      }
    }

  }
  int totalCount =0;
  //Depth First search on each node
  for(final start in startingPositions){
    var count = depthFirstSearchCount(start, grid);
   totalCount += count;
  }

  return totalCount;
}


int partTwoSolution(String input){
  final grid = LineSplitter().convert(input)
    .map((x) => x.split('')
      .map((i) => int.parse(i)).toList()).toList();

  List<({int x, int y})> startingPositions = [];
  for(int y = 0; y < grid.length; y++){
    for(int x = 0; x < grid[0].length; x++){
      if(grid[y][x] == 0) {
        startingPositions.add((x: x, y: y));
      }
    }

  }
  int totalCount =0;
  //Depth First search on each node
  for(final start in startingPositions){
    var count = depthFirstSearchCount(start, grid, false);
   totalCount += count;
  }

  return totalCount;
}

int depthFirstSearchCount(({int x, int y}) start, List<List<int>> grid, [bool distinct = true]){
    int count =0;
   Queue<({int x, int y})> toMove = Queue.from([start]);
   while(toMove.isNotEmpty){
    //Add all nodes next to it and pop this one
    final checkFrom = toMove.removeFirst();
    if(grid[checkFrom.y][checkFrom.x] == 9) count++;
    //Check Left
    if((checkFrom.x - 1) >= 0 && grid[checkFrom.y][checkFrom.x -1] == grid[checkFrom.y][checkFrom.x] +1){
      toMove.add((x: checkFrom.x-1, y:checkFrom.y));
    }
    //Check Right
    if((checkFrom.x + 1) < grid[0].length && grid[checkFrom.y][checkFrom.x +1] == grid[checkFrom.y][checkFrom.x] +1){
      toMove.add((x: checkFrom.x+1, y:checkFrom.y));
    }
    //Check Up
    if((checkFrom.y -1) >= 0 && grid[checkFrom.y -1][checkFrom.x] == grid[checkFrom.y][checkFrom.x] +1){
      toMove.add((x: checkFrom.x, y:checkFrom.y -1));
    }
    //Check Down
    if((checkFrom.y +1) < grid.length && grid[checkFrom.y +1][checkFrom.x] == grid[checkFrom.y][checkFrom.x] +1){
      toMove.add((x: checkFrom.x, y:checkFrom.y +1));
    }

    toMove = distinct ? Queue.from(toMove.toSet().toList()) : toMove;
   }

   return count;
}
