import 'dart:convert';

int partOneSolution(String input) {
  final Iterable<Iterable<int>> grid =
      LineSplitter().convert(input).map((x) => x.split(" ").map(int.parse));

  grid.map((Iterable<int> row){
    row.getSliding(2).every((x) {
      return (x.first - x.last).abs() < 3;
      //Check direction
    });
  });

  return 42;
}

extension Sliding<T> on Iterable<T> {
  Iterable<Iterable<T>> getSliding(int window) sync* {
    for (int i = 0; i < length - (window - 1); i++) {
      yield skip(i).take(window);
    }
  }
}
