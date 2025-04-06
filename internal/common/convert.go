package common

func Convert[T, V any](data []T, converter func(T) V) []V {
	result := make([]V, 0, len(data))
	for _, element := range data {
		result = append(result, converter(element))
	}
	return result
}
