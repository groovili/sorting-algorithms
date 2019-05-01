package sorting_algorithms

import (
	"math"
)

func MergeSort(input []int, asc bool) []int {
	if len(input) <= 1 {
		return input
	}

	middle := int(math.Floor(float64(len(input)) / 2))
	left := MergeSort(input[:middle], asc)
	right := MergeSort(input[middle:], asc)

	return merge(left, right, asc)
}

func merge(left, right []int, asc bool) []int {
	result := make([]int, 0)
	lIdx, rIdx := 0, 0

	var check bool
	for lIdx < len(left) && rIdx < len(right) {
		if asc {
			check = left[lIdx] <= right[rIdx]
		} else {
			check = right[rIdx] <= left[lIdx]
		}

		if check {
			result = append(result, left[lIdx])
			lIdx++
		} else {
			result = append(result, right[rIdx])
			rIdx++
		}
	}

	result = append(result, left[lIdx:]...)
	result = append(result, right[rIdx:]...)

	return result
}
