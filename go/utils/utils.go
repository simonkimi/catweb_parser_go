package utils

import "strings"

func Map[T any, E any](input []T, function func(item T) E) []E {
	var list []E
	for _, item := range input {
		list = append(list, function(item))
	}
	return list
}

func IsEmptyOrWhiteSpace(str string) bool {
	return strings.TrimSpace(str) == ""
}
