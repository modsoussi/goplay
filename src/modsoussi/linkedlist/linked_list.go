// A singly linked list implementation
// (c) 2017. modsoussi.

package linkedlist

// Node ...
type Node struct {
	Val  interface{}
	Next *Node
}

// LinkedList ...
type LinkedList struct {
	Head *Node
	Size int
}

// NewLinkedList ...
func NewLinkedList() LinkedList {
	return LinkedList{Head: &Node{Val: nil, Next: nil}, Size: 0}
}

// Add to linked list
func (l *LinkedList) Add(val interface{}) {
	if l.Head.Val == nil {
		l.Head.Val = val
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &Node{Val: val, Next: nil}
	}
	l.Size++
}

// HeadValue returns head value
func (l *LinkedList) HeadValue() interface{} {
	return l.Head.Val
}

// GetSize returns current size of singly linked list
func (l *LinkedList) GetSize() int {
	return l.Size
}

// Get returns element at given index
func (l *LinkedList) Get(index int) interface{} {
	if index == 0 {
		return l.Head.Val
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Val
}

// Clear clears LinkedList
func (l *LinkedList) Clear() {
	l.Size = 0
	l.Head = &Node{Val: nil, Next: nil}
}
