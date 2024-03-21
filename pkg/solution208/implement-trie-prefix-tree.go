package solution208

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

// Constructor Initialize your data structure here
func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert Inserts a word into the trie
func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isEndOfWord = true // 标记单词结束
}

// Search Returns if the word is in the trie
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false // 如果字符不存在，返回false
		}
		node = node.children[c]
	}
	return node.isEndOfWord // 如果完整走完word且isEndOfWord为真，则说明找到了完整的单词
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false // 如果字符不存在，返回false
		}
		node = node.children[c]
	}
	return true // 走完prefix，只要找到前缀就可返回true，无需是完整的单词
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
