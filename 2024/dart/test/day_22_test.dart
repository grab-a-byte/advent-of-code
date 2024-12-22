
import 'package:test/test.dart';
import 'package:dart/day_22.dart' as day_22;


const testInput = """1
10
100
2024""";

void main(){
  test("Day 21- Part 1", (){
    var res = day_22.partOneSolution(testInput);
    expect(res, 37327623);
  });

  test("Day 21- Part 2", (){
    var res = day_22.partTwoSolution("""1
2
3
2024""");
    expect(res, 23);
  });

  test("Mix", (){
    var res = 42.mix(15);
    expect(res, 37);
  });

  test("Prune", (){
    var res = 100000000.prune();
    expect(res, 16113920);
  });
}