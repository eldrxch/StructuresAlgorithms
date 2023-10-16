package main

import "fmt"

func main() {
	sortedData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	itm := 5

	index, found := BinarySearch(sortedData, itm)

	if found {
		fmt.Println("Item found at index:", index)
		return
	}

	fmt.Println("Item not found")
}

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
