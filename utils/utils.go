package utils

func Map[T any, E any](input []T, function func(item T) E) []E {
	var list []E
	for _, item := range input {
		list = append(list, function(item))
	}
	return list
}
