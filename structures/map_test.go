package structures

import (
	"bytes"
	"testing"
)

func TestMapAdd(t *testing.T) {
	tests := []struct {
		keys      []string
		values    []int
		wantSize  int
		wantError bool
	}{
		{[]string{"a", "b", "c", "d", "e"}, []int{1, 2, 3, 4, 5}, 5, false}, // add 5 items
		{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 11, false}, // add 11 items
		{[]string{"a"}, []int{1}, 1, false},        // add 1 item
		{[]string{"a", "a"}, []int{1, 2}, 1, true}, // add duplicate item - error
		{[]string{""}, []int{0}, 0, true},          // add 0 items - error
	}

	for _, test := range tests {
		var err error
		h := &HashTable{}
		for i, v := range test.keys {
			err = h.Add(v, test.values[i])
		}
		if h.Size() != test.wantSize {
			t.Errorf("got %v, want %v", h.Size(), test.wantSize)
		}
		if err != nil && !test.wantError {
			t.Errorf("got error %v, want no error", err)
		}
		if err == nil && test.wantError {
			t.Error("got no error, want error")
		}
	}
}

func TestMapGet(t *testing.T) {
	tests := []struct {
		keys      []string
		values    []int
		invKey    string
		wantError bool
	}{
		{[]string{"apple", "banana", "orange", "kiwi"}, []int{1080, 3002, 1200, 2000}, "", false},
		{[]string{"bread", "milk", "eggs", "cheese"}, []int{100, 200, 300, 400}, "tomato", true},
	}

	for _, test := range tests {
		h := &HashTable{}
		for i, v := range test.keys {
			err := h.Add(v, test.values[i])
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
		for i, v := range test.keys {
			mv, err := h.Get(v)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if test.values[i] != mv {
				t.Errorf("got %v, want %v", mv, test.values[i])
			}
		}
		if test.invKey == "" {
			return
		}
		_, err := h.Get(test.invKey)
		if err == nil && test.wantError {
			t.Error("got no error, want error")
		}
	}
}

func TestMapRemove(t *testing.T) {
	tests := []struct {
		keys   []string
		values []int
	}{
		{[]string{"apple", "banana", "orange", "kiwi"}, []int{1080, 3002, 1200, 2000}}, // remove 4 items
		{[]string{"bread"}, []int{100}}, // remove from root
	}
	for _, test := range tests {
		h := &HashTable{}
		for i, v := range test.keys {
			err := h.Add(v, test.values[i])
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
		for _, v := range test.keys {
			err := h.Remove(v)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
	}
	// test error conditions
	// remove from empty list
	h := &HashTable{}
	err := h.Remove("tomato") // remove from empty list
	if err == nil {
		t.Error("got no error, want error")
	}
	// remove non-existent key
	err = h.Add("tomato", 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = h.Remove("peanut") // remove non-existent key
	if err == nil {
		t.Error("got no error, want error")
	}
}

func TestMapWriteString(t *testing.T) {
	var tests = []struct {
		keys         []string
		values       []int
		wantText     string
		wantStrError bool
	}{
		{[]string{"a", "b", "c", "d", "e"}, []int{1, 2, 3, 4, 5}, "a=1, b=2, c=3, d=4, e=5", false}, // add 5 items
		{[]string{"apple"}, []int{100}, "apple=100", false},                                         // add 1 item
		{[]string{}, []int{}, "", true},                                                             // add 0 items - error
	}

	for _, test := range tests {
		var err error
		h := &HashTable{}
		for i, v := range test.keys {
			err = h.Add(v, test.values[i])
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
		buf := new(bytes.Buffer)
		err = h.WriteString(buf)
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
