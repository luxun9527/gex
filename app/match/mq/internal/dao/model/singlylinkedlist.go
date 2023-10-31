package model

import (
	"github.com/pkg/errors"
)

// List holds the elements, where each element points to the next element
type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	Value *MatchData
	Next  *element
}

// NewSinglyList instantiates a new list and adds the passed values, if any, to the list
func NewSinglyList() *List {
	return &List{}
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (list *List) Add(values ...*MatchData) {
	for _, value := range values {
		newElement := &element{Value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.Next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// ResetHead 重置头指针到指定的位置。
func (list *List) ResetHead(index int) error {

	if e, ok := list.Get(index); ok {
		list.size -= index
		list.first = e
		return nil
	}
	return errors.New("ResetHead fail over range")
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (*element, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	element := list.first
	for e := 0; e != index; e, element = e+1, element.Next {
	}

	return element, true
}

// Empty returns true if list does not contain any elements.
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *List) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Check that the index is within bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
