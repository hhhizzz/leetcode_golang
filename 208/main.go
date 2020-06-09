package _208

type Trie struct {
	pre *letter
}

type letter struct {
	c        int32
	isLetter bool
	next     map[int32]*letter
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{pre: &letter{0, false, map[int32]*letter{}}}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	current := this.pre
	for _, c := range word {
		if nextLetter, ok := current.next[c]; ok {
			current = nextLetter
		} else {
			current.next[c] = &letter{c, false, map[int32]*letter{}}
			current = current.next[c]
		}
	}
	current.isLetter = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	current := this.pre
	for _, c := range word {
		if nextLetter, ok := current.next[c]; ok {
			current = nextLetter
		} else {
			return false
		}
	}
	return current.isLetter
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	current := this.pre
	for _, c := range prefix {
		if nextLetter, ok := current.next[c]; ok {
			current = nextLetter
		} else {
			return false
		}
	}
	return true
}
