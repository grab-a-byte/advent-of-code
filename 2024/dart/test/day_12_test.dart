import 'package:test/test.dart';
import 'package:dart/day_12.dart' as day_12;

void main(){
  test("Day Twelve - Part 1 - Example 1", (){
    const input = """AAAA
BBCD
BBCC
EEEC""";
    final int result = day_12.partOneSolution(input);
    expect(result, 140);
  });

  test("Day Twelve - Part 1 - Example 2", (){
    const input = """OOOOO
OXOXO
OOOOO
OXOXO
OOOOO""";
    final int result = day_12.partOneSolution(input);
    expect(result, 772);
  });

  test("Day Twelve - Part 1 - Example 3", (){
    const input = """RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE""";
    final int result = day_12.partOneSolution(input);
    expect(result, 1930);
  });
}
