package sorting_algorithms

import (
	"math/rand"
)

func Quick(input []int, asc bool) []int {
	if len(input) <= 1 {
		return input
	}

	pivot := input[rand.Intn(len(input)-1)]
	equal := make([]int, 0)
	less := make([]int, 0)
	more := make([]int, 0)

	for _, v := range input {
		if v > pivot {
			more = append(more, v)
		} else if v < pivot {
			less = append(less, v)
		} else {
			equal = append(equal, v)
		}
	}

	less = Quick(less, asc)
	more = Quick(more, asc)

	if asc {
		return append(append(less, equal...), more...)
	}

	return append(append(more, equal...), less...)
}
