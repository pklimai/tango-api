package tool

func Map[T1 any, T2 any](input []T1, mapping func(T1) T2) []T2 {
	result := make([]T2, 0, len(input))

	for _, value := range input {
		result = append(result, mapping(value))
	}

	return result
}

func MapSafe[T1 any, T2 any](input []T1, mapping func(T1) (T2, error)) ([]T2, error) {
	result := make([]T2, 0, len(input))

	for _, value := range input {
		mapped, err := mapping(value)
		if err != nil {
			return nil, err
		}

		result = append(result, mapped)
	}

	return result, nil
}
