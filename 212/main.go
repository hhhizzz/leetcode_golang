package _212

type Trie struct {
    data   map[byte]*Trie
    word   string
    isWord bool
}

func construct(words *[]string) *Trie {
    root := &Trie{data: map[byte]*Trie{}}
    for _, word := range *words {
        wordArray := []byte(word)
        current := root
        for _, c := range wordArray {
            if next, ok := current.data[c]; ok {
                current = next
            } else {
                current.data[c] = &Trie{data: map[byte]*Trie{}}
                current = current.data[c]
            }
        }
        current.isWord = true
        current.word = word
    }
    return root
}

func search(board *[][]byte, visited *[][]bool, i, j int, current *Trie, m map[string]bool) {
    if i < 0 || j < 0 || i >= len(*board) || j >= len((*board)[i]) {
        return
    }
    if (*visited)[i][j] {
        return
    }

    if next, ok := current.data[(*board)[i][j]]; ok {
        (*visited)[i][j] = true
        if next.isWord {
            m[next.word] = true
        }
        nextI := i
        nextJ := j
        search(board, visited, nextI-1, nextJ, next, m)
        search(board, visited, nextI, nextJ+1, next, m)
        search(board, visited, nextI+1, nextJ, next, m)
        search(board, visited, nextI, nextJ-1, next, m)
    }
    (*visited)[i][j] = false
}

func findWords(board [][]byte, words []string) []string {
    trie := construct(&words)
    visited := make([][]bool, len(board))
    for i := range visited {
        visited[i] = make([]bool, len(board[i]))
    }
    result := make(map[string]bool)
    for i := range board {
        for j := range board[i] {
            for v := range visited {
                for u := range visited[v] {
                    visited[v][u] = false
                }
            }
            search(&board, &visited, i, j, trie, result)
        }
    }
    var resultArray []string
    for k := range result {
        resultArray = append(resultArray, k)
    }
    return resultArray
}
