package _3

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func lengthOfLongestSubstring(s string) int {
    m := map[byte]int{}
    result := 0
    for i, j := 0, 0; j < len(s); j++ {
        if _, ok := m[s[j]]; ok {
            i = max(i, m[s[j]]+1)
        }
        result = max(result, j-i+1)
        m[s[j]] = j
    }
    return result
}
