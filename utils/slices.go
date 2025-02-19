package utils

import "slices"

func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	if slice == nil {
		return nil
	}

	result := make([]T2, len(slice))
	for i, e := range slice {
		result[i] = f(e)
	}

	return result
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, e := range slice {
		if f(e) {
			result = append(result, e)
		}
	}

	return result
}

// All returns true if all elements match the given predicate
func All[T any](slice []T, f func(T) bool) bool {
	return slices.IndexFunc(slice, func(e T) bool { return !f(e) }) == -1
}

func AnyOf[T comparable](e T, values ...T) bool {
	return slices.Contains(values, e)
}
