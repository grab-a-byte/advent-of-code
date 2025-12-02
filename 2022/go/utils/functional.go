package utils

func Map[T interface{}, U interface{}](input []T, mapper func(*T) U) []U {
	var output []U
	for _, val := range input {
		output = append(output, mapper(&val))
	}
	return output
}

func Max[T float64 | int](input []T) T {
	var output T
	for _, val := range input {
		if val > output {
			output = val
		}
	}
	return output
}

func Any(check int, items []int, predicate func(int, int) bool) bool {
	for _, item := range items {
		if predicate(check, item) {
			return true
		}
	}
	return false
}
