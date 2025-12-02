import 'package:test/expect.dart';
import 'package:test/scaffolding.dart';
import './day_18.dart';

void main() {
  test("Given sample input, prodeces sample solution result for part 1", () {
    var result = DayEighteen.PartOne([
      "2,2,2",
      "1,2,2",
      "3,2,2",
      "2,1,2",
      "2,3,2",
      "2,2,1",
      "2,2,3",
      "2,2,4",
      "2,2,6",
      "1,2,5",
      "3,2,5",
      "2,1,5",
      "2,3,5"
    ]);
    expect(result, 64);
  });

  test("Given sample input, prodeces sample solution result for part 2", () {
    var result = DayEighteen.PartTwo([
      "2,2,2",
      "1,2,2",
      "3,2,2",
      "2,1,2",
      "2,3,2",
      "2,2,1",
      "2,2,3",
      "2,2,4",
      "2,2,6",
      "1,2,5",
      "3,2,5",
      "2,1,5",
      "2,3,5"
    ]);
    expect(result, 58);
  });
}
