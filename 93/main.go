package _93

import "strconv"

func checkIp(bytes []byte) bool {
    if len(bytes) > 1 && bytes[0] == '0' {
        return false
    }
    number, _ := strconv.Atoi(string(bytes))
    if number > 255 {
        return false
    }
    return true
}

func dfs(bytes []byte, left int, result *[]string, current string) {
    if left == 1 {
        if !checkIp(bytes) {
            return
        } else {
            current += string(bytes)
            *result = append(*result, current)
            return
        }
    } else {
        for i := 1; i <= 3 && i < len(bytes); i++ {
            if checkIp(bytes[:i]) {
                next := current
                number, _ := strconv.Atoi(string(bytes[:i]))
                if number > 255 {
                    break
                }
                next += string(bytes[:i])
                next += "."
                dfs(bytes[i:], left-1, result, next)
            }
        }
    }
}

func restoreIpAddresses(s string) []string {
    var result []string
    dfs([]byte(s), 4, &result, "")
    return result
}
