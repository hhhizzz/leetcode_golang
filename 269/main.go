package _269

//使用拓扑排序的方法
func alienOrder(words []string) string {
	//所有字母中最长的长度
	maxLength := 0

	//用于存储某个字母的下一个字母的图结构 例如letterNext[0][2]=1 表示c出现在a的后面
	letterNext := [26][26]int{}
	//存储出现过的字母
	letterMap := map[byte]bool{}

	for _, word := range words {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}
	//记录第一个单词出现的字母
	for i := 0; i < len(words[0]); i++ {
		letterMap[words[0][i]] = true
	}

	//position表示字母的位置
	for position := 0; position < maxLength; position++ {
		for i := 1; i < len(words); i++ {
			lastWord := words[i-1]
			word := words[i]
			if position < len(word) {
				//如果出现两个单词 ab* ac* 就说明c的位置在b后面，置letterNext['b'-'a']['c'-'a'] = 1
				if position < len(lastWord) && ((position == 0 || word[position-1] == lastWord[position-1]) && word[position] != lastWord[position]) {
					letterNext[lastWord[position]-'a'][word[position]-'a'] = 1
					//一个简单优化，如果出现b在c后面，c又在b后面，也就是出现了2个节点的环，可以直接判定无解
					if letterNext[word[position]-'a'][lastWord[position]-'a'] == 1 {
						return ""
					}
				}
				//记录这个单词出现过
				letterMap[word[position]] = true
			}
		}
	}
	var result []byte
	//进入拓扑排序流程
	for len(letterMap) != 0 {
		hasLoop := true
		for letter := range letterMap {
			noIn := true
			for i := 0; i < 26; i++ {
				if letterNext[i][letter-'a'] == 1 {
					noIn = false
					break
				}
			}
			if noIn {
				result = append(result, letter)
				//把这个节点的后继节点入度减少
				for i := 0; i < 26; i++ {
					letterNext[letter-'a'][i] = 0
				}
				delete(letterMap, letter)
				//如果存在入度为0的节点，说明此次循环没有环
				hasLoop = false
				break
			}
		}
		//有环就说明无解
		if hasLoop {
			return ""
		}
	}
	return string(result)
}
