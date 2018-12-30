package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []struct {
		array    []int
		expected []int
	}{
		{
			array:    []int{2, 3, 5, 4, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			array:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			array:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			array:    []int{0, 3, 6, 8, -1},
			expected: []int{-1, 0, 3, 6, 8},
		},
	}

	for _, d := range data {
		QuickSort(d.array, 0, len(d.array)-1)

		if fmt.Sprintf("%v", d.array) != fmt.Sprintf("%v", d.expected) {
			t.Errorf("Expected: %v, actually: %v", d.expected, d.array)
		}
	}
}

func TestSwap(t *testing.T) {
	array1 := []int{
		2, 3, 5, 4, 1,
	}
	array2 := []int{
		2, 3, 5, 4, 1,
	}

	swap(array1, 1, 2)
	// fmt.Println(array1)

	swap(array2, 0, len(array2)-1)
	// fmt.Println(array2)
}
