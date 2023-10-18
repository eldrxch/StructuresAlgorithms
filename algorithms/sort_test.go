package algorithms

import "testing"

func TestSelectionSort(t *testing.T) {
	var tests = []struct {
		items []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, // already sorted
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}}, // reverse sorted
		{[]int{1, 3, 2, 5, 4}, []int{1, 2, 3, 4, 5}}, // unsorted
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}}, // all equal
		{[]int{1}, []int{1}},                         // single element
		{[]int{}, []int{}},                           // empty slice
	}

	for _, test := range tests {
		result := SelectionSort(test.items)
		if !equal(result, test.want) {
			t.Errorf("got %v, want %v", result, test.want)
		}
	}
}

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		items []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, // already sorted
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}}, // reverse sorted
		{[]int{1, 3, 2, 5, 4}, []int{1, 2, 3, 4, 5}}, // unsorted
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}}, // all the same
		{[]int{4, 2, 3, 2, 4}, []int{2, 2, 3, 4, 4}}, // some duplicates
		{[]int{1}, []int{1}},                         // single element
		{[]int{}, []int{}},                           // empty slice
	}

	for _, test := range tests {
		result := QuickSort(test.items)
		if !equal(result, test.want) {
			t.Errorf("got %v, want %v", result, test.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range b {
		if a[i] != v {
			return false
		}
	}

	return true
}
