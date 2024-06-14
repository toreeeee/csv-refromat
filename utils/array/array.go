package array

func Map[T any, O any](input []T, fn func(T, int) O) []O {
	output := make([]O, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = fn(input[i], i)
	}

	return output
}
