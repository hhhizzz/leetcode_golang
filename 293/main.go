package _293

func generatePossibleNextMoves(s string) []string {
    var result []string
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '+' && s[i+1] == '+' {
            temp := []byte(s)
            temp[i] = '-'
            temp[i+1] = '-'
            result = append(result, string(temp))
        }
    }
    return result
}
