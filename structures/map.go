package structures

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type keyVal struct {
	key   string
	value int
	next  *keyVal
}

const (
	FillFactorThreshold = 0.75
	InitialSize         = 3
)

type HashTable struct {
	size       int
	loadfactor float64
	entries    []*keyVal
}

func (h *HashTable) Add(key string, value int) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("key cannot be empty")
	}
	if h.hasKey(key) {
		return errors.New("key already exists")
	}
	if h.entries == nil {
		h.entries = make([]*keyVal, InitialSize)
	}
	kv := &keyVal{key: key, value: value}
	addEntry(h.entries, kv, len(h.entries))
	h.size = h.size + 1
	h.loadfactor = math.Trunc(float64(h.size)/float64(len(h.entries))*100) / 100.0
	if h.loadfactor > FillFactorThreshold {
		h.entries = resize(h.entries)
	}
	return nil
}

func addEntry(entries []*keyVal, kv *keyVal, size int) {
	index := hashStringFold(kv.key, size)
	root := entries[index]
	if root == nil {
		entries[index] = kv
		return
	}
	addNext(root, kv)
}

func addNext(curr *keyVal, kv *keyVal) {
	if curr.next == nil {
		curr.next = kv
		return
	}
	addNext(curr.next, kv)
}

func resize(entries []*keyVal) []*keyVal {
	newEntries := make([]*keyVal, len(entries)*2)
	for _, v := range entries {
		if v != nil {
			addEntry(newEntries, v, len(newEntries))
		}
	}
	return newEntries
}

func (h *HashTable) Get(key string) (int, error) {
	if h.entries == nil {
		return 0, errors.New("hash table is empty")
	}
	kv, _ := findKey(h.entries, key)
	if kv == nil {
		return 0, fmt.Errorf("key %s does not exist", key)
	}
	return kv.value, nil
}

func (h *HashTable) hasKey(key string) bool {
	kv, _ := findKey(h.entries, key)
	return kv != nil
}

func findKey(entries []*keyVal, key string) (*keyVal, *keyVal) {
	if entries == nil {
		return nil, nil
	}
	index := hashStringFold(key, len(entries))
	root := entries[index]
	if root == nil {
		return nil, nil
	}
	var next = root
	var prev *keyVal
	for next != nil {
		if next.key == key {
			return next, prev
		}
		next = next.next
		prev = next
	}
	return nil, nil
}

func (h *HashTable) Remove(key string) error {
	if h.entries == nil {
		return errors.New("hash table is empty")
	}
	kv, prevKv := findKey(h.entries, key)
	if kv == nil {
		return fmt.Errorf("key %s does not exist", key)
	}
	if prevKv != nil && kv.next != nil {
		prevKv.next = kv.next
		kv = nil
		return nil
	}
	if kv.next != nil {
		kv = kv.next
	}
	kv = nil
	return nil
}

func (h *HashTable) Size() int {
	return h.size
}

// WriteString writes the linked list items to a writer.
// It returns an error if the list is empty.
func (h *HashTable) WriteString(w io.Writer) error {
	if h.entries == nil {
		return errors.New("hash table is empty")
	}
	return writeKeyValString(h.entries, w)
}

func writeKeyValString(kv []*keyVal, w io.Writer) error {
	var kvStr []string
	for _, v := range kv {
		if v == nil {
			continue
		}
		var next = v
		for next != nil {
			str := fmt.Sprintf("%s=%s", next.key, strconv.Itoa(next.value))
			kvStr = append(kvStr, str)
			next = next.next
		}
	}
	out := strings.Join(kvStr, ", ")
	if _, err := w.Write([]byte(out)); err != nil {
		return err
	}
	return nil
}

func hashStringFold(key string, size int) int {
	if size == 0 {
		size = 1
	}
	var sum int64 = 0
	var mul int64 = 1
	for i := 0; i < len(key); i++ {
		rem := i % 4
		switch {
		case rem == 0:
			mul = 1
		default:
			mul *= 256
		}
		sum += int64(key[i]) * mul
	}
	return int(sum % int64(size))
}
