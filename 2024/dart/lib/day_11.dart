import 'dart:io';

void solution(){
  var file = File("inputs/day_11.txt");
  var input = file.readAsStringSync();
  var res1 = partOneSolution(input);
  print("Day 11 - Part 1 : $res1");
  var res2 = partTwoSolution(input);
  print("Day 11 - Part 2 : $res2");
}

int partOneSolution(String input) => solve2(input, 25);
int partTwoSolution(String input) => solve2(input, 75);

int solve2(String input, int steps) {
  var inputNumbers = input.trim().split(' ').map((x) => int.parse(x));
  var numbers = {for (var element in inputNumbers) element: 1};

  for (int i = 0; i < steps; i++) {
    Map<int, int> newMap = {};
    numbers.forEach((key, value) {
      final nStr = key.toString();
      if (value == 0) {
        return;
      } else if (key == 0) {
        newMap[1] = (newMap[1] ?? 0) + value;
      } else if (nStr.length % 2 == 0) {
        var firstNum = int.parse(nStr.substring(0, (nStr.length / 2).floor()));
        var secondNum = int.parse(nStr.substring((nStr.length / 2).floor()));
        newMap[firstNum] = (newMap[firstNum] ?? 0) + value;
        newMap[secondNum] = (newMap[secondNum] ?? 0) + value;
      } else {
        newMap[key * 2024] = (newMap[key * 2024] ?? 0) + value;
      }
      numbers[key] = 0;
    });

    newMap.forEach((key, value) {
      numbers.update(key, (v) => v + value, ifAbsent: () => value);
    });
  }

  return numbers.values.fold(0, (a, b) => a + b);
}

/////// Initial implementation, worked for part 1, kept for historical reasons ///////
// int solve(String input, int steps){
//   var numbers = input.trim().split(' ').map((x) => int.parse(x));
//   for(int i = 0; i < steps; i++){
//     numbers = numbers.map((n) {
//       if(n == 0) return [1];
//       final nStr = n.toString();
//       if (nStr.length % 2 ==0 ){
//         var firstNum = int.parse(nStr.substring(0, (nStr.length / 2).floor()));
//         var secondNum  = int.parse(nStr.substring((nStr.length / 2).floor()));
//         return [firstNum, secondNum];
//       }
//       return [n*2024];
//     }).expand((a) => a).toList();
//   }

//   return numbers.length;
// }