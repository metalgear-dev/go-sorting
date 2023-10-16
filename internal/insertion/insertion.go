package insertion

func InsertionSort(input []int) {
	if len(input) <= 1 {
		return
	}

	for i := 1; i < len(input); i += 1 {
		key := input[i]

		j := i - 1
		for j >= 0 && key < input[j] {
			input[j+1] = input[j]
			j -= 1
		}
		input[j+1] = key
	}
}
