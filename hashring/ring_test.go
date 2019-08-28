package ring

import (
	"testing"
)

func TestNewHashRing(t *testing.T) {
	hr := NewHashRing([]*Node{}, 3)
	switch interface{}(hr).(type) {
	case *HashRing:
		{
		}
	default:
		{
			t.Errorf("Expected *HashRing, got %T\n", hr)
		}
	}
}

func TestAddNewNode(t *testing.T) {
	hr := NewHashRing([]*Node{}, 3)
	node := &Node{name: "first", address: "127.0.0.1"}
	hr = hr.AddNode(node)
	if len(hr.nodes) != 1 {
		t.Errorf("Expected 1 node in the node, found %v\n", len(hr.nodes))
	}
}

func TestAddAndRemoveNewNode(t *testing.T) {
	hr := NewHashRing([]*Node{}, 3)
	nodeToBeAdded := &Node{name: "first", address: "127.0.0.1"}
	nodeToBeRemoved := &Node{name: "first", address: "127.0.0.1"}

	hr = hr.AddNode(nodeToBeAdded)
	if len(hr.nodes) != 1 {
		t.Errorf("Expected 1 node in the ring, found %d\n", len(hr.nodes))
	}

	hr = hr.RemoveNode(nodeToBeRemoved)
	if len(hr.nodes) != 0 {
		t.Errorf("Expected 0 node in the ring, found %d\n", len(hr.nodes))
	}
}

func TestGetTheNodeForTheKey(t *testing.T) {
	testString := "testString"
	var node *Node
	nodes := []*Node{
		{name: "First", address: "First"},
		{name: "Second", address: "Second"},
		{name: "Third", address: "Third"},
	}
	hr := NewHashRing(nodes, 3)
	node = hr.GetNode(testString)

	if node.name != "First" {
		t.Errorf("Expected 'First', got %s\n", node.name)
	}

	hr = hr.RemoveNode(&Node{name: "Second", address: "Second"})
	if len(hr.nodes) != 2 {
		t.Errorf("Expected 2 node in the ring, found %d\n", len(hr.nodes))
	}

	node = hr.GetNode(testString)
	if node.name != "Second" {
		t.Errorf("Expected 'Second', got %s\n", node.name)
	}
}

func TestRingPrint(t *testing.T) {
	nodes := []*Node{
		{name: "First", address: "First"},
		{name: "Second", address: "Second"},
		{name: "Third", address: "Third"},
	}
	hr := NewHashRing(nodes, 3)
	got := len(hr.Print()) //Because go sucks at multiline strings
	expected := 426
	if expected != got {
		t.Errorf("HashRing Print:: Expected '%d', got '%d'", expected, got)
	}
}
