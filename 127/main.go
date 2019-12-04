package _127

func changed(word1, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    diff := 0
    for i := 0; i < len(word1); i++ {
        if word1[i] != word2[i] {
            diff++
            if diff > 1 {
                return false
            }
        }
    }
    return diff == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    g := make([][]bool, len(wordList))
    for i := range g {
        g[i] = make([]bool, len(wordList))
    }
    end := -1
    for i := 0; i < len(wordList); i++ {
        if wordList[i] == endWord {
            end = i
        }
        for j := i + 1; j < len(wordList); j++ {
            if changed(wordList[i], wordList[j]) {
                g[i][j] = true
                g[j][i] = true
            }
        }
    }
    if end == -1 {
        return 0
    }
    var queue []int
    dist := make([]int, len(wordList))
    for i := 0; i < len(wordList); i++ {
        dist[i] = -1
    }
    for i := 0; i < len(wordList); i++ {
        if beginWord == wordList[i] {
            dist[i] = 1
        }
        if changed(beginWord, wordList[i]) {
            queue = append(queue, i)
            dist[i] = 2
        }
    }
    for len(queue) != 0 {
        current := queue[0]
        queue = queue[1:]
        if current == end {
            return dist[current]
        }
        for i := 0; i < len(g[current]); i++ {
            if g[current][i] && dist[i] == -1 {
                dist[i] = dist[current] + 1
                queue = append(queue, i)
            }
        }
    }
    return 0
}
