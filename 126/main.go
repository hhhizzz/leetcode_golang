package _126

func dfs(current, endWord string, path []string, g map[string][]string, result *[][]string) {
    path = append(path, current)
    if current == endWord {
        newPath := make([]string, len(path))
        copy(newPath, path)
        *result = append(*result, newPath)
    } else {
        for _, next := range g[current] {
            dfs(next, endWord, path, g, result)
        }
    }
}

//可以通过广搜的方法来构建图，然后再深搜找出路径
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    g := map[string][]string{}
    //dis表示这个字符离beginWord的距离
    dis := map[string]int{beginWord: 0}
    dict := map[string]bool{}
    for _, word := range wordList {
        dict[word] = true
    }
    wordLen := len(beginWord)

    q := []string{beginWord}
    step := 0
    for len(q) != 0 {
        level := make([]string, len(q))
        copy(level, q)
        q = []string{}
        var nextLevel []string
        for len(level) != 0 {
            current := level[0]
            level = level[1:]
            for l := 0; l < wordLen; l++ {
                word := []byte(current)
                for c := 'a'; c <= 'z'; c++ {
                    word[l] = byte(c)
                    if string(word) == current {
                        continue
                    }
                    if _, ok := dict[string(word)]; ok {
                        if wordDis, ok := dis[string(word)]; !ok {
                            dis[string(word)] = step
                            g[current] = append(g[current], string(word))
                            nextLevel = append(nextLevel, string(word))
                        } else if wordDis == step {
                            g[current] = append(g[current], string(word))
                        }
                    }
                }
            }

        }
        q = nextLevel
        step++
    }

    var result [][]string
    dfs(beginWord, endWord, []string{}, g, &result)

    return result
}
