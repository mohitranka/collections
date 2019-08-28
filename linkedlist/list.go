package linkedlist

import (
	"fmt"
)

// List : the basic data structure to operate on the linked list
type List struct {
	head *Node
}

// NewList : Constructor
func NewList(head *Node) *List {
	return &List{head: head}
}

func (l *List) Search(value interface{}) *Node {
	for curr := l.head; curr != nil; curr = curr.Next {
		if curr == nil || curr.Value == value {
			return curr
		}
	}
	return nil
}

func (l *List) Print() string {
	s := ""
	for curr := l.head; curr != nil; curr = curr.Next {
		s += fmt.Sprintf("%v", curr.Value)
		if curr.Next != nil {
			s += "->"
		}
	}
	return s // should not be reachable
}

func (l *List) AddNode(value interface{}) {

	// If the list is empty
	if l.head == nil {
		l.head = NewNode(value)
		return
	}
	// List is not empty, go to the end
	curr := l.head
	for ; curr.Next != nil; curr = curr.Next {
	}
	curr.Next = NewNode(value)
}

func (l *List) RemoveNode(value interface{}) {
	fmt.Printf("Trying to remove %v from the list: %s \n", value, l.Print())
	for curr := l.head; curr != nil; curr = curr.Next {
		// Looped through the entire list,
		// and have not found the value

		if curr.Value == value {
			// 1 -> 5 -> 7
			// Target is 1
			l.head = curr.Next
			return
		}
		if curr.Next == nil {
			return
		}
		if curr.Next.Value == value {
			// 1 -> 5-> 7
			// target is 5
			// curr is at the node with Value 1
			curr.Next = curr.Next.Next
		}
	}
}

func (l *List) Reverse() {
	curr := l.head
	var prev, next *Node
	for curr != nil {
		// 1->2->3->4->nil
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *List) PairWiseSwap() {
	// 1 => 1
	// 1 -> 2 -> 3 => 2 -> 1 -> 3
	// 1 -> 2 -> 3 -> 4 => 2 -> 1 -> 3 -> 4
	curr := l.head
	for curr != nil && curr.Next != nil {
		nextValue := curr.Next.Value
		curr.Next.Value = curr.Value
		curr.Value = nextValue
		curr = curr.Next.Next
	}
}
