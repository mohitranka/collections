package ring

import (
	"fmt"
)

// Node : Struct to hold the information about a server node
type Node struct {
	name    string
	address string
}

// NewNode : Node constructor
func NewNode(name string, address string) *Node {
	return &Node{name: name, address: address}
}

// Print : Print Node
func (n *Node) Print() string {
	return fmt.Sprintf("<Node name: %s address: %s>", n.name, n.address)
}
