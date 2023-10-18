package algorithms

// SelectionSort sorts a slice of integers using the selection sort algorithm.
// It returns a new sorted slice in ascending order.
func SelectionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var sorted []int
	for len(arr) > 0 {
		smallest := arr[0]
		smallestIndex := 0
		for i, v := range arr {
			if v < smallest {
				smallest = v
				smallestIndex = i
			}
		}
		sorted = append(sorted, smallest)
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return sorted
}

// QuickSort sorts a slice of integers using the quick sort algorithm.
// It returns a new sorted slice in ascending order.
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0] // default pivot
	if len(arr) > 2 {
		pivot = arr[len(arr)/2] // approximate middle pivot
	}

	var less, greater, dups, sorted []int
	for _, v := range arr {
		switch {
		case v < pivot:
			less = append(less, v)
		case v > pivot:
			greater = append(greater, v)
		case v == pivot:
			dups = append(dups, v)
		}
	}
	sorted = append(sorted, QuickSort(less)...)
	sorted = append(sorted, pivot)
	if len(dups) > 1 {
		sorted = append(sorted, dups[1:]...)
	}
	sorted = append(sorted, QuickSort(greater)...)
	return sorted
}
