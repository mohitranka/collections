package trie

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	current := t.Root
	var node *TrieNode
	var ok bool
	for _, ch := range word {
		if node, ok = current.Children[ch]; !ok {
			node = NewTrieNode()
			current.Children[ch] = node
		}
		current = node
	}
	current.IsWordEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t.Root

	var node *TrieNode
	var ok bool

	for _, ch := range word {
		if node, ok = current.Children[ch]; !ok {
			return false
		}
		current = node
	}
	return current.IsWordEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.Root

	var node *TrieNode
	var ok bool

	for _, ch := range prefix {
		if node, ok = current.Children[ch]; !ok {
			return false
		}
		current = node
	}
	return true
}

func (t *Trie) Delete(word string) bool {
	return t.deleteWord(t.Root, word, 0)
}

func (t *Trie) deleteWord(current *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !current.IsWordEnd {
			return false
		}
		current.IsWordEnd = false
		return len(current.Children) == 0
	}

	ch := []rune(word)[index]

	var node *TrieNode
	var ok bool

	if node, ok = current.Children[ch]; !ok {
		return false
	}

	shouldDeleteCurrentNode := t.deleteWord(node, word, index+1)

	if shouldDeleteCurrentNode {
		delete(current.Children, ch)
		return len(current.Children) == 0
	}
	return true
}
