package _140

func wordBreak(s string, wordDict []string) []string {
	dp := map[int][]string{}
	return wordBreaks(s, wordDict, 0, dp)
}

func wordBreaks(s string, wordDict []string, start int, dp map[int][]string) []string {
	var result []string
	if result, ok := dp[start]; ok {
		return result
	}
	for _, word := range wordDict {
		if start+len(word) <= len(s) && s[start:len(word)+start] == word {
			if start+len(word) == len(s) {
				result = append(result, word)
			} else {
				nextResult := wordBreaks(s, wordDict, start+len(word), dp)
				for _, nextR := range nextResult {
					nextR = word + " " + nextR
					result = append(result, nextR)
				}
			}
		}
	}
	dp[start] = result
	return result
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
			if next, ok := current.next[word[i]]; ok {
				current = next
			} else {
				current.next[word[i]] = &Trie{next: map[byte]*Trie{}}
				current = current.next[word[i]]
			}
		}
		current.isWord = true
	}
	return root
}

func dfs(s string, start int, root *Trie, cache map[int][]string) []string {
	if result, ok := cache[start]; ok {
		return result
	}
	current := root
	var result []string
	for i := start; i < len(s); i++ {
		if next, ok := current.next[s[i]]; ok {
			current = next
			if current.isWord {
				if i == len(s)-1 {
					result = append(result, s[start:i+1])
				} else {
					nextResults := dfs(s, i+1, root, cache)
					for _, nextResult := range nextResults {
						result = append(result, s[start:i+1]+" "+nextResult)
					}
				}
			}
		} else {
			break
		}
	}
	cache[start] = result
	return result
}

// 方法2 用trie树优化
func wordBreak2(s string, wordDict []string) []string {
	root := constructTrie(wordDict)
	return dfs(s, 0, root, map[int][]string{})
}
