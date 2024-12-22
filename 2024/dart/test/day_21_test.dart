
import 'package:test/test.dart';
import 'package:dart/day_21.dart' as day_21;


const testInput = """1
10
100
2024""";

void main(){
  test("Day 21- Part 1", (){
    var res = day_21.partOneSolution(testInput);
    expect(res, 37327623);
  });

  test("Day 21- Part 2", (){
    var res = day_21.partTwoSolution("""1
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