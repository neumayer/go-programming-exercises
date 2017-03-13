package datastructures

type linkedlist struct {
	head *element
}

type element struct {
	value int
	next  *element
}

func newLinkedListFromSlice(values []int) linkedlist {
	ll := &linkedlist{}
	for _, value := range values {
		ll.insert(value)
	}
	return *ll
}

func (l *linkedlist) insertInFront(value int) {
	if l.head == nil {
		newElement := &element{value, nil}
		l.head = newElement
		return
	}
	originalFirstElement := l.head
	newElement := &element{value, originalFirstElement}
	l.head = newElement
}

func (l *linkedlist) insert(value int) {
	newElement := &element{value, nil}
	if l.head == nil {
		l.head = newElement
		return
	}
	var lastElement *element
	for element := l.head; element != nil; element = element.next {
		lastElement = element
	}
	lastElement.next = newElement
}

func (l linkedlist) size() int {
	size := 0
	for element := l.head; element != nil; element = element.next {
		size++
	}
	return size
}

func (l linkedlist) contains(value int) bool {
	for element := l.head; element != nil; element = element.next {
		if element.value == value {
			return true
		}
	}
	return false
}
