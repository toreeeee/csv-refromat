package array

func Map[T any, O any](input []T, fn func(*T, int) O) []O {
	output := make([]O, len(input))

	for i := 0; i < len(input); i++ {
		output[i] = fn(&input[i], i)
	}

	return output
}

func Filter[T any](input []T, testFn func(*T) bool) (ret []T) {
	for _, s := range input {
		if testFn(&s) {
			ret = append(ret, s)
		}
	}
	return
}
