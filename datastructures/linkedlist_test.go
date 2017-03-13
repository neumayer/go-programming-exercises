package datastructures

import (
	"reflect"
	"testing"
)

func toSlice(list linkedlist) []int {
	slice := make([]int, 0)
	for element := list.head; element != nil; element = element.next {
		slice = append(slice, element.value)
	}
	return slice
}

func TestInsert(t *testing.T) {
	var insertTests = []struct {
		n            []int
		expectedSize int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 2}, 3},
		{[]int{1, 2, 2}, 3},
		{[]int{1, 2, 3}, 3},
	}
	for _, tt := range insertTests {
		ll := newLinkedListFromSlice(tt.n)
		if !reflect.DeepEqual(tt.n, toSlice(ll)) {
			t.Errorf("Linkedlist (%d): expected %d, actual %d", ll, tt.n, toSlice(ll))
		}
		if len(toSlice(ll)) != ll.size() {
			t.Errorf("Linkedlist (%d): expected size %d, actual %d", ll, tt.expectedSize, ll.size())
		}
	}
}

func TestInsertNilList(t *testing.T) {
	ll := *(&linkedlist{})
	ll.insert(0)
	if !reflect.DeepEqual([]int{0}, toSlice(ll)) {
		t.Errorf("Linkedlist (%d): expected %d, actual %d", ll, []int{0}, toSlice(ll))
	}

	ll = *(&linkedlist{})
	ll.insertInFront(0)
	if !reflect.DeepEqual([]int{0}, toSlice(ll)) {
		t.Errorf("Linkedlist (%d): expected %d, actual %d", ll, []int{0}, toSlice(ll))
	}
}

func TestInsertInFront(t *testing.T) {
	var insertTests = []struct {
		n        []int
		element  int
		expected []int
	}{
		{[]int{1}, 0, []int{0, 1}},
		{[]int{1, 2}, 3, []int{3, 1, 2}},
	}
	for _, tt := range insertTests {
		ll := newLinkedListFromSlice(tt.n)
		ll.insertInFront(tt.element)
		if !reflect.DeepEqual(tt.expected, toSlice(ll)) {
			t.Errorf("Linkedlist (%d): expected %d, actual %d", tt.n, tt.expected, toSlice(ll))
		}
	}
}

func TestContains(t *testing.T) {
	var insertTests = []struct {
		n        []int
		value    int
		expected bool
	}{
		{[]int{1}, 1, true},
		{[]int{1, 2}, 2, true},
		{[]int{1, 2, 2}, 3, false},
		{[]int{1, 2, 3}, 3, true},
	}
	for _, tt := range insertTests {
		ll := newLinkedListFromSlice(tt.n)
		result := ll.contains(tt.value)
		if result != tt.expected {
			t.Errorf("Contains (%d): value %d expected %d, actual %d", tt.n, tt.value, tt.expected, result)
		}
	}
}
