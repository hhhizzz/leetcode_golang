package algorithm

func kmp(pattern string, text string) bool {
    if len(pattern) > len(text) {
        return false
    }
    next := make([]int, len(pattern))
    next[0] = -1
    for j, t := 0, -1; j < len(pattern)-1; {
        if t == -1 || pattern[j] == pattern[t] {
            j++
            t++
            if pattern[j] == pattern[t] {
                next[j] = next[t]
            } else {
                next[j] = t
            }
        } else {
            t = next[t]
        }
    }
    for i, j := 0, 0; i < len(text) && j < len(pattern); {
        if j == -1 || text[i] == pattern[j] {
            i++
            j++
            if j == len(pattern) {
                return true
            }
        } else {
            j = next[j]
        }

    }
    return false
}
