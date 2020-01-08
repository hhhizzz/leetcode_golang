package _820

import (
    "sort"
)

func minimumLengthEncoding(words []string) int {
    m := map[string]bool{}
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            if _, ok := m[word[i:]]; ok {
                delete(m, word[i:])
            }
        }
        m[word] = true
    }
    result := 0
    for word := range m {
        result += len(word) + 1
    }
    return result
}
