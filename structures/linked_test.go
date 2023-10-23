package structures

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		value    []int
		wantSize int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9}, // add 9 items
		{[]int{1}, 1},                         // add 1 item
		{[]int{}, 0},                          // add 0 items
		{[]int{1, 2, 1}, 3},                   // add duplicate
	}

	for _, test := range tests {
		l := &LinkedList{}
		for _, v := range test.value {
			l.Add(v)
		}

		if l.Size() != test.wantSize {
			t.Errorf("got %v, want %v", l.Size(), test.wantSize)
		}
	}
}

func TestDoublyAdd(t *testing.T) {
	var tests = []struct {
		value    []int
		wantSize int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 6}, // add 6 items
		{[]int{1}, 1},                // add 1 item
		{[]int{}, 0},                 // add 0 items
		{[]int{1, 1, 1}, 3},          // add duplicate
	}

	for _, test := range tests {
		l := &DoublyLinkedList{}
		for _, v := range test.value {
			l.Add(v)
		}

		if l.Size() != test.wantSize {
			t.Errorf("got %v, want %v", l.Size(), test.wantSize)
		}
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		value     []int
		deleteVal int
		wantError bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, false}, // delete from middle
		{[]int{1}, 1, false},             // delete from root
		{[]int{1, 2}, 2, false},          // delete from end
		{[]int{}, 1, true},               // delete from empty list
		{[]int{1, 2}, 3, true},           // delete non-existent value
	}

	for _, test := range tests {
		var err error
		l := &LinkedList{}
		for _, v := range test.value {
			l.Add(v)
		}
		err = l.Delete(test.deleteVal)
		if err != nil && !test.wantError {
			t.Errorf("got error %v, want no error", err)
		}
		if err == nil && test.wantError {
			t.Error("got no error, want error")
		}
	}
}

func TestDoublyDelete(t *testing.T) {
	var tests = []struct {
		value     []int
		deleteVal int
		wantError bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, false}, // delete from middle
		{[]int{1}, 1, false},             // delete from root
		{[]int{1, 2}, 2, false},          // delete from end
		{[]int{}, 1, true},               // delete from empty list
		{[]int{1, 2}, 3, true},           // delete non-existent value
	}

	for _, test := range tests {
		var err error
		l := &DoublyLinkedList{}
		for _, v := range test.value {
			l.Add(v)
		}
		err = l.Delete(test.deleteVal)
		if err != nil && !test.wantError {
			t.Errorf("got error %v, want no error", err)
		}
		if err == nil && test.wantError {
			t.Error("got no error, want error")
		}
	}
}

func TestWriteString(t *testing.T) {
	var tests = []struct {
		value        []int
		wantText     string
		wantStrError bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, "123456789", false}, // add 9 items
		{[]int{1}, "1", false},                                 // add 1 item
		{[]int{}, "", true},                                    // add 0 items
	}

	for _, test := range tests {
		var err error
		l := &LinkedList{}
		for _, v := range test.value {
			l.Add(v)
		}

		buf := new(bytes.Buffer)
		err = l.WriteString(buf)
		if err != nil && !test.wantStrError {
			t.Errorf("got error %v, want no error", err)
		}
		if err == nil && test.wantStrError {
			t.Error("got no error, want error")
		}
		if buf.String() != test.wantText {
			t.Errorf("got %v, want %v", buf.String(), test.wantText)
		}
	}
}

func TestDoublyWriteString(t *testing.T) {
	var tests = []struct {
		value        []int
		wantText     string
		wantStrError bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, "123456789", false}, // add 9 items
		{[]int{1}, "1", false},                                 // add 1 item
		{[]int{}, "", true},                                    // add 0 items
	}

	for _, test := range tests {
		var err error
		l := &DoublyLinkedList{}
		for _, v := range test.value {
			l.Add(v)
		}

		buf := new(bytes.Buffer)
		err = l.WriteString(buf)
		if err != nil && !test.wantStrError {
			t.Errorf("got error %v, want no error", err)
		}
		if err == nil && test.wantStrError {
			t.Error("got no error, want error")
		}
		if buf.String() != test.wantText {
			t.Errorf("got %v, want %v", buf.String(), test.wantText)
		}
	}
}

func TestDoublyValues(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{"Empty list", []int{}, []int{}},
		{"Single element", []int{42}, []int{42}},
		{"Multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := &DoublyLinkedList{}
			for _, v := range test.values {
				l.Add(v)
			}

			got := l.Values()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
