package structures

import (
	"errors"
	"io"
	"strconv"
)

type node struct {
	prev  *node
	value int
	next  *node
}

type LinkedList struct {
	root *node
}

// Add adds a value to the linked list.
func (l *LinkedList) Add(value int) {
	if l.root == nil {
		l.root = &node{value: value}
		return
	}
	addNode(l.root, value, l)
}

// addNode adds a value to the linked list.
func addNode(elm *node, value int, caller interface{}) {
	if elm.next == nil {
		switch caller.(type) {
		case *LinkedList:
			elm.next = &node{value: value}
		case *DoublyLinkedList:
			elm.next = &node{value: value, prev: elm}
		default:
			panic("unknown caller")
		}
		return
	}
	addNode(elm.next, value, caller)
}

// Delete deletes a value from the linked list.
// It returns an error if the list is empty or the value does not exist in the list.
func (l *LinkedList) Delete(value int) error {
	if l.root == nil {
		return errors.New("list is empty")
	}
	if l.root.value == value {
		l.root = l.root.next
		return nil
	}
	found := deleteNode(l.root, l.root.next, value)
	if !found {
		return errors.New("value does not exists in list")
	}
	return nil
}

// deleteNode deletes a value from the linked list.
// It returns true if the value was found and deleted.
func deleteNode(prev, curr *node, value int) bool {
	if curr == nil {
		return false
	}
	if curr.value == value {
		prev.next = curr.next
		curr = nil
		return true
	}
	return deleteNode(curr, curr.next, value)
}

// WriteString writes the linked list items to a writer.
// It returns an error if the list is empty.
func (l *LinkedList) WriteString(w io.Writer) error {
	if l.root == nil {
		return errors.New("list is empty")
	}
	return writeNodeString(l.root, w)
}

func writeNodeString(elm *node, w io.Writer) error {
	var next = elm
	for next != nil {
		str := strconv.Itoa(next.value)
		if _, err := w.Write([]byte(str)); err != nil {
			return err
		}
		next = next.next
	}
	return nil
}

// Size returns the number of items in the linked list.
func (l *LinkedList) Size() int {
	if l.root == nil {
		return 0
	}
	return size(l.root)
}

func size(elm *node) int {
	size := 0
	var next = elm
	for next != nil {
		size++
		next = next.next
	}
	return size
}

type DoublyLinkedList struct {
	root *node
}

// Add adds a value to the linked list.
func (d *DoublyLinkedList) Add(value int) {
	if d.root == nil {
		d.root = &node{value: value}
		return
	}
	addNode(d.root, value, d)
}

// Delete deletes a value from the linked list.
// It returns an error if the list is empty or the value does not exist in the list.
func (d *DoublyLinkedList) Delete(value int) error {
	if d.root == nil {
		return errors.New("list is empty")
	}
	if d.root.value == value {
		d.root = d.root.next
		return nil
	}
	found := deleteNode(d.root, d.root.next, value)
	if !found {
		return errors.New("value does not exists in list")
	}
	return nil
}

// WriteString writes the linked list items to a writer.
// It returns an error if the list is empty.
func (d *DoublyLinkedList) WriteString(w io.Writer) error {
	if d.root == nil {
		return errors.New("list is empty")
	}
	return writeNodeString(d.root, w)
}

// Size returns the number of items in the linked list.
func (d *DoublyLinkedList) Size() int {
	if d.root == nil {
		return 0
	}
	return size(d.root)
}
