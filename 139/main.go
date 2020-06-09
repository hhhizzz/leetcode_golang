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
