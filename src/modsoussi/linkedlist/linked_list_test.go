package linkedlist

import "testing"

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList()
	if l.Head.Val != nil || l.Head.Next != nil || l.Size > 0 {
		t.Fatal("Initialization failed.")
	}
}

func TestAdd(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	if l.HeadValue() != 1 {
		t.Fatal("Getting head value failed")
	}
	if l.GetSize() != 2 {
		t.Fatal("Getting size of list failed")
	}
}

func TestGet(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	n := l.Get(1)
	if n != 2 {
		t.Fatalf("Expected %d, got %d", 2, n)
	}

	n = l.Get(4)
	if n != 5 {
		t.Fatalf("Expected %d, got %d", 5, n)
	}
}

func TestClear(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	l.Clear()
	n := l.HeadValue()
	if n != nil {
		t.Fatal("Failed to clear. HeadValue != nil")
	}

	s := l.GetSize()
	if s > 0 {
		t.Fatal("Failed to clear.Size > 0")
	}
}
