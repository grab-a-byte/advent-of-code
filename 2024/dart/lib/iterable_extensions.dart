
extension Sliding<T> on Iterable<T> {
  Iterable<Iterable<T>> sliding(int window) sync* {
    for (int i = 0; i < length - (window - 1); i++) {
      yield skip(i).take(window);
    }
  }

  Iterable<Iterable<T>> slidingWithSkip(int window) sync* {
    for (int i = 0; i < length - (window - 1); i+=window) {
      yield skip(i).take(window);
    }
  }
}