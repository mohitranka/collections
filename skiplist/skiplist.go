package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

// SkipList : Structure to hold the skiplist datastructure
type SkipList struct {
	MaxLevel   int
	P          float64
	Level      int
	Header     *node
	Randomizer *rand.Rand
}

// NewSkipList : Constructor for the `SkipList` datastructure
func NewSkipList(maxLevel int, p float64) *SkipList {

	return &SkipList{
		MaxLevel:   maxLevel,
		P:          p,
		Level:      0,
		Header:     newNode(-1, maxLevel),
		Randomizer: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (sl *SkipList) getRandomLevel() int {
	level := 0
	for sl.Randomizer.Float64() < sl.P && level < sl.MaxLevel {
		level++
	}
	return level
}

func (sl *SkipList) getUpdateNCurrent(key int) ([]*node, *node) {
	update := make([]*node, sl.MaxLevel+1)
	current := sl.Header
	for level := sl.Level; level >= 0; level-- {
		for current.Forward[level] != nil && current.Forward[level].Key < key {
			current = current.Forward[level]
		}
		update[level] = current
	}
	current = current.Forward[0]
	return update, current
}

// Display : Return a string representation of the `SkipList`
func (sl *SkipList) Display() string {
	s := ""
	for level := sl.Level; level >= 0; level-- {
		s += fmt.Sprintf("Level %d: ", level)
		node := sl.Header.Forward[level]
		for node != nil {
			s += fmt.Sprintf("%d ", node.Key)
		}
		s += "\n"
	}
	return s
}

// Search : search `key` in the `SkipList`, return a boolean to indicate whether the `key` was found in the `SkipList`
func (sl *SkipList) Search(key int) bool {
	_, current := sl.getUpdateNCurrent(key)
	if current != nil && current.Key == key {
		return true
	}
	return false
}

// Insert : insert `key` in the `SkipList`, return a boolean indicating if a new entry was added in the `SkipList`
func (sl *SkipList) Insert(key int) bool {
	update, current := sl.getUpdateNCurrent(key)

	if current != nil && current.Key != key {
		randomLevel := sl.getRandomLevel()
		if randomLevel > sl.Level {
			for level := sl.Level + 1; level < randomLevel+1; level++ {
				update[level] = sl.Header
			}
			sl.Level = randomLevel
		}

		node := newNode(key, randomLevel)
		for level := 0; level < randomLevel+1; level++ {
			node.Forward[level] = update[level].Forward[level]
			update[level].Forward[level] = node
		}
		return true
	}
	return false
}

// Delete : delete the `key` from the `SkipList` if present. Return a boolean indicating the actual removal.
func (sl *SkipList) Delete(key int) bool {
	update, current := sl.getUpdateNCurrent(key)
	if current != nil && current.Key == key {
		// Found key, we can delete it.
		for level := 0; level < sl.Level+1; level++ {
			if update[level].Forward[level] != current {
				break
			}
			update[level].Forward[level] = current.Forward[level]
		}

		for sl.Level > 0 && sl.Header.Forward[sl.Level] == nil {
			sl.Level--
		}

		return true
	}
	return false
}
