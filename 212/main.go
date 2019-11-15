package _212

type Trie struct {
    word string
    char rune
    next map[rune]*Trie
}

func dfs(board [][]byte, visited [][]bool, i, j int, result *[]string, current *Trie) {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
        return
    }
    if visited[i][j] {
        return
    }
    if next, ok := current.next[rune(board[i][j])]; !ok {
        return
    } else {
        if next.word != "" {
            *result = append(*result, next.word)
            next.word = ""
        }
        visited[i][j] = true
        dfs(board, visited, i+1, j, result, next)
        dfs(board, visited, i, j+1, result, next)
        dfs(board, visited, i-1, j, result, next)
        dfs(board, visited, i, j-1, result, next)
        visited[i][j] = false
    }
}

func findWords(board [][]byte, words []string) []string {
    root := &Trie{"", 0, map[rune]*Trie{}}
    for _, word := range words {
        current := root
        for _, c := range word {
            if next, ok := current.next[c]; ok {
                current = next
            } else {
                next = &Trie{"", c, map[rune]*Trie{}}
                current.next[c] = next
                current = next
            }
        }
        current.word = word
    }
    visited := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        visited[i] = make([]bool, len(board[i]))
    }
    var result []string
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            dfs(board, visited, i, j, &result, root)
        }
    }
    return result
}
