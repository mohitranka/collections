package linkedlist

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value, Next: nil}
}

func (n *Node) Print() string {
	return fmt.Sprintf("%v", n.Value)
}
