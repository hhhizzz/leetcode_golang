package _14

func longestCommonPrefix(strs []string) string {
    prefix := ""
    if len(strs) == 0 {
        return prefix
    }
    l := len(strs[0])
    if l == 0 {
        return prefix
    }
    for i := 0; i < l; i++ {
        w := strs[0][i]
        for _, s := range strs {
            if i >= len(s) || s[i] != w {
                return prefix
            }
        }
        prefix += string(w)
    }
    return prefix
}
