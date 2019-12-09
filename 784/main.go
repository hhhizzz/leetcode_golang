package _784

func dfs(S []byte, word []int, result *[]string) {
    if len(word) == 0 {
        *result = append(*result, string(S))
        return
    }
    dfs(S, word[1:], result)
    if S[word[0]] >= 'a' && S[word[0]] <= 'z' {
        S[word[0]] -= 'a' - 'A'
        dfs(S, word[1:], result)
        S[word[0]] += 'a' - 'A'
    } else {
        S[word[0]] += 'a' - 'A'
        dfs(S, word[1:], result)
        S[word[0]] -= 'a' - 'A'
    }
}

func letterCasePermutation(S string) []string {
    word := []int{}
    for i := 0; i < len(S); i++ {
        if S[i] >= 'a' && S[i] <= 'z' || S[i] >= 'A' && S[i] <= 'Z' {
            word = append(word, i)
        }
    }
    result := []string{}
    dfs([]byte(S), word, &result)
    return result
}
