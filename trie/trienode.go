package trie

type TrieNode struct {
	Children  map[rune]*TrieNode
	IsWordEnd bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children:  make(map[rune]*TrieNode),
		IsWordEnd: false,
	}
}
