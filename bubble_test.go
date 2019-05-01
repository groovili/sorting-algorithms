package sorting_algorithms

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortAsc(t *testing.T) {
	sorted := Bubble([]int{10, 23423, 444, -77, 1231231, 11, 0}, true)
	assert.Equal(t, []int{-77, 0, 10, 11, 444, 23423, 1231231}, sorted)
}

func TestBubbleSortDesc(t *testing.T) {
	sorted := Bubble([]int{-1, 10, math.MaxInt32, 0, 7, math.MinInt32, math.MinInt8, 1231231, 11, 0, 11, 1, 100, 13, 1, 5}, false)
	assert.Equal(t, []int{math.MaxInt32, 1231231, 100, 13, 11, 11, 10, 7, 5, 1, 1, 0, 0, -1, math.MinInt8, math.MinInt32}, sorted)
}
