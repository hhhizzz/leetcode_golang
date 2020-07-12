package _68

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var lineLen = 0
	var wordsLen = 0
	var wordArray []string
	for i := 0; i < len(words); i++ {
		wordArray = append(wordArray, words[i])
		wordsLen += len(words[i])
		if len(wordArray) == 1 {
			lineLen += len(words[i])
		} else {
			lineLen += len(words[i]) + 1
		}

		if i == len(words)-1 || lineLen+len(words[i+1])+1 > maxWidth {
			spaceLen := maxWidth - wordsLen
			spaceArray := make([]string, len(wordArray)-1)
			if len(spaceArray) != 0 {
				if i != len(words)-1 {
					for i := 0; i < spaceLen; i++ {
						spaceArray[i%len(spaceArray)] += " "
					}
					spaceLen = 0
				} else {
					for i := 0; i < len(spaceArray); i++ {
						spaceArray[i] = " "
						spaceLen -= 1
					}
				}
			}
			var current string
			for i := 0; i < len(wordArray)+len(spaceArray); i++ {
				if i%2 == 0 {
					current += wordArray[i>>1]
				} else {
					current += spaceArray[i>>1]
				}
			}
			for spaceLen != 0 {
				current += " "
				spaceLen -= 1
			}
			res = append(res, current)
			lineLen = 0
			wordsLen = 0
			wordArray = make([]string, 0)
		}

	}
	return res
}
