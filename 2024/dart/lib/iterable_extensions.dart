extension Sliding<T> on Iterable<T> {
  Iterable<Iterable<T>> sliding(int window) sync* {
    for (int i = 0; i < length - (window - 1); i++) {
      yield skip(i).take(window);
    }
  }

  Iterable<Iterable<T>> slidingWithSkip(int window) sync* {
    for (int i = 0; i < length - (window - 1); i += window) {
      yield skip(i).take(window);
    }
  }

  bool isEqual(Iterable<T> other) {
    if (identical(this, other)) return true;
    var length = this.length;
    if (length != other.length) return false;
    for (var i = 0; i < length; i++) {
      if (elementAt(i) != other.elementAt(i)) return false;
    }
    return true;
  }
}
