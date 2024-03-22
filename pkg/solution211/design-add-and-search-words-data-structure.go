package solution211

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word)
}

func (this *WordDictionary) search(node *TrieNode, word string) bool {
	for i, c := range word {
		if c == '.' {
			for _, child := range node.children {
				if this.search(child, word[i+1:]) {
					return true
				}
			}
			return false
		}
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return node.isEndOfWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
