package _758

type trie struct {
	isWord bool
	c      byte
	next   map[byte]*trie
}

func boldWords(words []string, S string) string {
	root := &trie{next: map[byte]*trie{}}
	for _, word := range words {
		current := root
		for i := 0; i < len(word); i++ {
			c := word[i]
			if _, ok := current.next[c]; !ok {
				current.next[c] = &trie{isWord: false, c: c, next: map[byte]*trie{}}
			}
			current = current.next[c]
			if i == len(word)-1 {
				current.isWord = true
			}
		}
	}
	boldMark := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		current := root
		for j := i; j < len(S); j++ {
			if next, ok := current.next[S[j]]; ok {
				current = next
				if current.isWord {
					for begin := i; begin <= j; begin++ {
						boldMark[begin] = 1
					}
				}
			} else {
				break
			}
		}
	}
	result := ""
	startBold := false
	for i := 0; i < len(boldMark); i++ {
		if boldMark[i] == 1 {
			if !startBold {
				result += "<b>"
				startBold = true
			}
		} else {
			if startBold {
				result += "</b>"
				startBold = false
			}
		}
		result += string(S[i])
	}
	if startBold {
		result += "</b>"
		startBold = false
	}
	return result
}
