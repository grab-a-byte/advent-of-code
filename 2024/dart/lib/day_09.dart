import 'package:dart/iterable_extensions.dart';
import 'dart:io';

void solution() {
  var file = File("inputs/day_09.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 2 - Part 1 : $res1");
  //Part 2 failed
  // var res2 = partTwoSolution(input);
  print("Day 2 - Part 2 : Wah Wah Wah");
}

int partOneSolution(String input) {
  final blocks = List<int>.empty(growable: true);
  int index = 0;
  for (int i = 0; i < input.length; i += 2) {
    blocks.addAll(List.filled(int.parse(input[i]), index));
    if (i + 1 < input.length) {
      blocks.addAll(List.filled(int.parse(input[i + 1]), -1));
    }
    index++;
  }

  for (int i = 0; i < blocks.length; i++) {
    if (blocks[i] == -1) {
      for (int x = blocks.length - 1; x > i; x--) {
        if (blocks[x] != -1) {
          blocks[i] = blocks[x];
          blocks[x] = -1;
          break;
        }
      }
    }
  }

  int total = 0;
  for (int i = 0; i < blocks.length; i++) {
    if (blocks[i] == -1) break;
    total += blocks[i] * i;
  }
  return total;
}

int partTwoSolution(String input) {
  var index = 0;
  if (input.length % 2 != 0) input += "0";
  var fileMap = input
      .split('')
      .slidingWithSkip(2)
      .map((v) => PuzFile(int.parse(v.first), int.parse(v.last), index++))
      .toList();

  for (int r = fileMap.length - 1; r >= 0; r--) {
    var right = fileMap[r];
    if (right.moved) continue;
    for (int l = 0; l < fileMap.length; l++) {
      var left = fileMap[l];
      if (left.freeSpace < right.size) continue;
      if (l >= r) break;
      //Add size and free space to one left of where we are
      fileMap[r - 1].freeSpace += right.freeSpace + right.size;
      //Capture remaining free space
      var remainingSpace = left.freeSpace - right.size;
      //Remove remainingSpace from left
      left.freeSpace = 0;
      //Take start to left
      var newList = fileMap.take(l + 1).toList();
      //Add new node that is marked a moved
      newList.add(PuzFile(right.size, remainingSpace, right.index, true));
      //Take the rest of the nodes up to left
      newList.addAll(fileMap.take(r).skip(l + 1));
      //Add everything after left
      newList.addAll(fileMap.take(fileMap.length).skip(r + 1));
      //Assign new list to continue
      fileMap = newList;
      //Go back to far right to try and move
      r = fileMap.length -1;
      break;
    }
    //Didnt make any of it, mark as moved
    right.moved = true;
  }

  final blocks = List<int>.empty(growable: true);
  for (var f in fileMap) {
    blocks.addAll(List.filled(f.size, f.index));
    blocks.addAll(List.filled(f.freeSpace, -1));
  }

  int total = 0;
  for (int i = 0; i < blocks.length; i++) {
    if (blocks[i] == -1) continue;
    total += blocks[i] * i;
  }

  //6378445102056 Too High
  return total;
}

void printMap(List<PuzFile> input) {
  final blocks = List<int>.empty(growable: true);
  for (var f in input) {
    blocks.addAll(List.filled(f.size, f.index));
    blocks.addAll(List.filled(f.freeSpace, -1));
  }
  print(blocks.map((x) {
    if (x == -1) {
      return ".";
    } else {
      return x.toString();
    }
  }).join(""));
}

class PuzFile {
  int size;
  int freeSpace;
  int index;
  bool moved = false;

  PuzFile(this.size, this.freeSpace, this.index, [this.moved = false]);
}