package trie

type Trie struct {
	next [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	for _, ch := range word {
		w := ch - 'a'
		if node.next[w] == nil {
			node.next[w] = &Trie{}
		}
		node = node.next[w]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		w := ch - 'a'
		if node.next[w] == nil {
			return false
		}

		node = node.next[w]
	}

	return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		w := ch - 'a'
		if node.next[w] == nil {
			return false
		}
		node = node.next[w]
	}

	return true
}

