package _168

func convertToTitle(n int) string {
    var result string
    for n > 0 {
        n -= 1
        current := n % 26
        result = string([]byte{'A' + byte(current)}) + result
        n /= 26
    }
    return result
}
