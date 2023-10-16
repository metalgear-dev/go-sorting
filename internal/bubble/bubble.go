package bubble

func BubbleSort(input []int) {
	n := len(input)

	for i := 0; i < n; i += 1 {
		swapped := false
		for j := 0; j < n-i-1; j += 1 {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
