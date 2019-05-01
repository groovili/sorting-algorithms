package sorting_algorithms

func InsertionSort(input []int, asc bool) []int {
	for i := 1; i < len(input); i++ {
		val := input[i]
		prevIdx := i - 1

		for prevIdx > -1 {
			if asc && input[prevIdx] < val {
				break
			}

			if !asc && input[prevIdx] > val {
				break
			}

			input[prevIdx+1] = input[prevIdx]
			prevIdx--
		}

		input[prevIdx+1] = val
	}

	return input
}
