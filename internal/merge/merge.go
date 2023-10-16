package merge

func MergeSort(input []int) []int {
	if len(input) > 1 {
		mid := len(input) / 2
		return Merge(MergeSort(input[:mid]), MergeSort(input[mid:]))
	}
	return input
}

func Merge(left []int, right []int) []int {
	res := make([]int, len(left)+len(right))

	// merge
	i, j := 0, 0
	for k := range res {
		if i >= len(left) {
			res[k] = right[j]
			j++
		} else if j >= len(right) {
			res[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
	}

	return res
}
