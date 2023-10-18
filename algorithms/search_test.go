package algorithms

import "testing"

func Test_BinarySearch(t *testing.T) {
	var tests = []struct {
		sortedData    []int
		itm           int
		expectedIndex int
		expectedFound bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4, true},    // found
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, -1, false}, // not found
		{[]int{}, 5, -1, false},                           // empty slice
		{[]int{1}, 1, 0, true},                            // one item
	}

	for _, test := range tests {
		wantedIndex, wantedFound := BinarySearch(test.sortedData, test.itm)

		if wantedIndex != test.expectedIndex || wantedFound != test.expectedFound {
			t.Errorf("got %v %v, want %v %v", wantedIndex, wantedFound, test.expectedIndex, test.expectedFound)
		}
	}
}
