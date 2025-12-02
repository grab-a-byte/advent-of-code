import 'dart:convert';
import 'dart:io';

void solution(){
  var file = File("inputs/day_18.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input, 71, 1024);
  print("Day 13 - Part 1 : $res1");

  var res2 = partTwoSolution(input, 71, 1024);
  print("Day 13 - Part 2 : $res2");
}

const defaultValue = 2000000;

int partOneSolution(String input, int length, int byteCount){
  var bytes = LineSplitter().convert(input).map((el){
    final parts = el.split(',');
    return (int.parse(parts[0]), int.parse(parts[1]));
  }).take(byteCount);

  final arr = aStar(length, bytes.toList());

  // printArr(arr);
  return arr[length-1][length-1];
}

String partTwoSolution(String input, int length, int byteCount){
  var bytes = LineSplitter().convert(input).map((el){
    final parts = el.split(',');
    return (int.parse(parts[0]), int.parse(parts[1]));
  }).toList();

  for (int i = byteCount; i < bytes.length; i++){
    final arr = aStar(length, bytes.take(i+1).toList());
    if(arr[length-1][length-1] == defaultValue) {
      return "${bytes[i].$1},${bytes[i].$2}";
    }
  }

  return "INVALID";
}

List<List<int>> aStar(int length, List<(int,int)> objects){

  var arr = List.generate(length, (_) => List<int>.filled(length, defaultValue));
  for(var(x,y) in objects){
    arr[y][x] = -1;
  }
  arr[0][0] = 0;
  List<(int, int)> toCheck = [(0,0)];
  while(toCheck.isNotEmpty){
    var (x, y) = toCheck.removeAt(0);
    //Check left and add
    if (x > 0 && arr[y][x] +1 < arr[y][x-1]){
      arr[y][x-1] = arr[y][x] +1;
      toCheck.add((x-1, y));
    }
    //Check Right and add
    if(x < arr[0].length-1 && arr[y][x] +1 < arr[y][x+1]){
      arr[y][x+1] = arr[y][x] +1;
      toCheck.add((x+1, y));
    }
    //Check up and add
    if(y > 0 && arr[y][x] +1 < arr[y-1][x]){
      arr[y-1][x] = arr[y][x] +1;
      toCheck.add((x, y-1));
    }

    //Check down and add
    if(y < arr.length-1 && arr[y][x] +1 < arr[y+1][x]){
      arr[y+1][x] = arr[y][x] +1;
      toCheck.add((x, y+1));
    }
  }
  return arr;
}

void printArr(List<List<int>> arr){
  for(int y =0; y < arr.length; y++){
    print(arr[y].map((e) => e.toString()).join(","));
  }
}