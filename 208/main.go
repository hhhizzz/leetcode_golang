package _208

type Trie struct {
    data map[int32]*Trie
    cut  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    trie := Trie{data: map[int32]*Trie{}}
    return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
    current := this
    for _, c := range word {
        if next, ok := current.data[c]; ok {
            current = next
        } else {
            current.data[c] = &Trie{data: map[int32]*Trie{}}
            current = current.data[c]
        }
    }
    current.cut = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    current := this
    for _, c := range word {
        if next, ok := current.data[c]; ok {
            current = next
        } else {
            return false
        }
    }
    return current.cut
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    current := this
    for _, c := range prefix {
        if next, ok := current.data[c]; ok {
            current = next
        } else {
            return false
        }
    }
    return true
}
