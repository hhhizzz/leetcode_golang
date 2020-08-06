package _131

func dfs(s string, palin [][]bool, pos int, current []string, result *[][]string) {
	if pos == len(s) {
		temp := make([]string, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	for i := pos; i < len(s); i++ {
		if palin[pos][i] {
			dfs(s, palin, i+1, append(current, s[pos:i+1]), result)
		}
	}
}

func partition(s string) [][]string {
	// palin[i][j] 表示 s[i:j+1]是回文串
	palin := make([][]bool, len(s))
	for i := 0; i < len(palin); i++ {
		palin[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		for dis := 0; i+dis < len(s) && i-dis >= 0; dis++ {
			if dis == 0 {
				palin[i][i+dis] = true
			} else {
				if s[i+dis] == s[i-dis] {
					palin[i-dis][i+dis] = true
				} else {
					break
				}
			}
		}
		j := i + 1
		for dis := 0; i-dis >= 0 && j+dis < len(s); dis++ {
			if s[i-dis] == s[j+dis] {
				palin[i-dis][j+dis] = true
			} else {
				break
			}
		}
	}
	var result [][]string
	dfs(s, palin, 0, []string{}, &result)
	return result
}
