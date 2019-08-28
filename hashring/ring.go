package ring

import (
	"crypto/md5"
	"fmt"
	"hash"
	"sort"
	"strconv"
)

// HashRing : Struct for the constient hash main entity
type HashRing struct {
	ring         map[int]*Node
	nodes        []*Node
	hasher       hash.Hash
	replicaCount int
}

// NewHashRing : HashRing constructor
func NewHashRing(nodes []*Node, replicaCount int) *HashRing {
	if replicaCount < 1 {
		fmt.Println("Need 1 or more replicas to create the ring")
		return nil
	}
	hr := &HashRing{nodes: nodes, hasher: md5.New(), ring: make(map[int]*Node), replicaCount: replicaCount}
	for _, node := range nodes {
		nodeKey := fmt.Sprintf("%s-%s-%s", node.name, node.address, strconv.FormatInt(int64(0), 10))
		bKey := hr.hasher.Sum([]byte(nodeKey))
		for i := 0; i < hr.replicaCount; i++ {
			key := hashVal(bKey[i*4 : i*4+4])
			hr.ring[key] = node
		}
	}
	return hr
}

// GetNode : Get the appropriate node to store the stringKey
func (hr *HashRing) GetNode(stringKey string) *Node {
	pos, ok := hr.GetNodePos(stringKey)
	if !ok {
		return nil
	}
	vnodes := hr.getSortedKeys()
	key := vnodes[pos]
	return hr.ring[key]
}

// GetNodePos : Get the position of the Node in the Ring, for the given stringKey
func (hr *HashRing) GetNodePos(stringKey string) (int, bool) {
	if len(hr.ring) == 0 {
		return int(0), false
	}

	key := hr.genKey(stringKey)
	vnodes := hr.getSortedKeys()
	pos := sort.Search(len(vnodes), func(i int) bool { return vnodes[i] > key })
	if pos == len(vnodes) {
		return int(0), true
	}
	return int(pos), true
}

// AddNode : Add a new Node in the HashRing
func (hr *HashRing) AddNode(node *Node) *HashRing {
	nodes := make([]*Node, len(hr.ring), len(hr.ring)+1)
	nodes = append(nodes, node)
	return NewHashRing(nodes, hr.replicaCount)
}

// RemoveNode : Remove the node from the HashRing
func (hr *HashRing) RemoveNode(node *Node) *HashRing {
	nodes := make([]*Node, 0)
	for _, eNode := range hr.nodes {
		if !(eNode.name == node.name && eNode.address == node.address) {
			nodes = append(nodes, node)
		}
	}
	return NewHashRing(nodes, hr.replicaCount)
}

// Print : Print the hash ring
func (hr *HashRing) Print() string {
	keys := hr.getSortedKeys()
	s := ""
	for _, key := range keys {
		node := hr.ring[key]
		s += fmt.Sprintf("%s<Degree: %v>\n", node.Print(), strconv.Itoa(key))
	}
	return s
}

// Private Methods

func (hr *HashRing) genKey(stringKey string) int {
	bKey := hr.hasher.Sum([]byte(stringKey))
	return hashVal(bKey[0:4])
}

func (hr *HashRing) getSortedKeys() []int {
	i := 0
	keys := make([]int, len(hr.ring))
	for k := range hr.ring {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

// Helper methods

func hashVal(bKey []byte) int {
	return ((int(bKey[3]) << 24) |
		(int(bKey[2]) << 16) |
		(int(bKey[1]) << 8) |
		(int(bKey[0]))) % 360
}
