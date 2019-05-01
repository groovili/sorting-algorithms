package sorting_algorithms

func Bubble(input []int, asc bool) []int {
	i := len(input) - 1
	for i >= 0 {
		for j := 0; j < i; j++ {
			if asc && input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}

			if !asc && input[j] < input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
		i--
	}

	return input
}
