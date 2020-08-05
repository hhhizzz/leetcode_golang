package _127

func compare(a, b string) bool {
	num := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			num++
		}
	}
	return num == 1
}

func find(word string, wordList []string) bool {
	for _, w := range wordList {
		if compare(w, word) {
			return true
		}
	}
	return false
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	currentList := []string{beginWord}
	result := 0
	for len(wordList) != 0 {
		var nextList []string
		var leftList []string
		result++
		for _, word := range wordList {
			if find(word, currentList) {
				if word == endWord {
					return result + 1
				}
				nextList = append(nextList, word)
			} else {
				leftList = append(leftList, word)
			}
		}
		currentList = nextList
		if len(leftList) == len(wordList) {
			break
		}
		wordList = leftList
	}
	return 0
}
