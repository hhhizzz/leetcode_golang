package _186

func reverseWords(s []byte) {
	var words [][]byte
	start := 0
	end := 0
	for end < len(s) {
		for end < len(s) && s[end] != ' ' {
			end++
		}
		temp := make([]byte, end-start)
		copy(temp, s[start:end])
		words = append(words, temp)
		end++
		start = end
	}
	pos := 0
	for i := len(words) - 1; i >= 0; i-- {
		for j := 0; j < len(words[i]); j++ {
			s[pos] = words[i][j]
			pos++
		}
		if pos < len(s) {
			s[pos] = ' '
			pos++
		}
	}
}
