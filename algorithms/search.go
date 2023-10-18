package algorithms

// BinarySearch searches for an item in a sorted slice of integers.
// Given a sorted slice of integers and an item to search for,
// it returns the index of the item and true if the item is found.
func BinarySearch(sortedData []int, itm int) (int, bool) {
	if len(sortedData) == 0 {
		return -1, false
	}

	lower := 0
	upper := len(sortedData) - 1
	for lower <= upper {
		middle := (lower + upper) / 2
		found := sortedData[middle]

		switch {
		case found == itm:
			return middle, true
		case found < itm:
			lower = middle + 1
		case found > itm:
			upper = middle - 1
		}
	}

	return -1, false
}
