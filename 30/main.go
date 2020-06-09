package _30

func findSubstring(s string, words []string) []int {
	N := len(s)
	M := len(words)
	if M == 0 || N == 0 {
		return []int{}
	}
	wordLength := len(words[0])
	if N < M*wordLength {
		return []int{}
	}

	count := make(map[string]int)
	var result []int
	for i := 0; i < M; i++ {
		count[words[i]]++
	}
	for i := 0; i < wordLength; i++ {
		start := i
		num := 0
		occur := make(map[string]int)
		for j := i; j <= N-wordLength; j += wordLength {
			temp := s[j : j+wordLength]
			if count[temp] > 0 {
				num++
				occur[temp]++
				for occur[temp] > count[temp] {
					occur[s[start:start+wordLength]]--
					start += wordLength
					num--
				}
				if num == M {
					result = append(result, start)
				}
			} else {
				start = j + wordLength
				num = 0
				occur = make(map[string]int)
			}
		}
	}
	return result
}
