package mapx

// Keys returns all keys of provided map
func Keys[K comparable, V any](provide map[K]V) []K {
	if provide == nil || len(provide) == 0 {
		return []K{}
	}

	keys := make([]K, 0, len(provide))
	for key := range provide {
		keys = append(keys, key)
	}
	return keys
}

// Values return all values of provided map
func Values[K comparable, V any](provide map[K]V) []V {
	if provide == nil || len(provide) == 0 {
		return []V{}
	}

	values := make([]V, 0, len(provide))
	for _, value := range provide {
		values = append(values, value)
	}
	return values
}

// Copy return copy of provided map
func Copy[K comparable, V any](provide map[K]V) map[K]V {
	if provide == nil || len(provide) == 0 {
		return map[K]V{}
	}

	copied := make(map[K]V, len(provide))
	for key, value := range provide {
		copied[key] = value
	}

	return copied
}
