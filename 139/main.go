package _139

func wordBreak(s string, wordDict []string) bool {
	canCut := make([]bool, len(s)+1)
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	canCut[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if canCut[j] && wordMap[s[j:i]] {
				canCut[i] = true
			}
		}
	}
	return canCut[len(s)]
}

type Trie struct {
	isWord bool
	next   map[byte]*Trie
}

func constructTrie(wordDict []string) *Trie {
	root := &Trie{next: map[byte]*Trie{}}
	for _, word := range wordDict {
		current := root
		for i := 0; i < len(word); i++ {
			letter := word[i]
			if next, ok := current.next[letter]; !ok {
				current.next[letter] = &Trie{next: map[byte]*Trie{}}
				current = current.next[letter]
			} else {
				current = next
			}
		}
		current.isWord = true
	}
	return root
}

// 方法2：使用trie优化，可以使复杂度减小到O(MN) M为wordDict中最长的一个单词
func wordBreak2(s string, wordDict []string) bool {
	root := constructTrie(wordDict)
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		current := root
		if i != 0 && !dp[i-1] {
			continue
		} else {
			for j := i; j < len(s); j++ {
				if next, ok := current.next[s[j]]; ok {
					current = next
					if current.isWord {
						dp[j] = true
					}
				} else {
					break
				}
			}
		}
	}
	return dp[len(s)-1]
}
