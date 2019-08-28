package skiplist

type node struct {
	Key     int
	Forward []*node
}

func newNode(key int, level int) *node {
	return &node{
		Key:     key,
		Forward: make([]*node, level+1),
	}
}
