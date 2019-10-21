package _140

func wordBreak(s string, wordDict []string) []string {
    dp := map[int][]string{}
    return wordBreaks(s, wordDict, 0, dp)
}

func wordBreaks(s string, wordDict []string, start int, dp map[int][]string) []string {
    var result []string
    if result, ok := dp[start]; ok {
        return result
    }
    for _, word := range wordDict {
        if start+len(word) <= len(s) && s[start:len(word)+start] == word {
            if start+len(word) == len(s) {
                result = append(result, word)
            } else {
                nextResult := wordBreaks(s, wordDict, start+len(word), dp)
                for _, nextR := range nextResult {
                    nextR = word + " " + nextR
                    result = append(result, nextR)
                }
            }
        }
    }
    dp[start] = result
    return result
}
