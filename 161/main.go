package _161

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func isOneEditDistance(s string, t string) bool {
    if abs(len(s)-len(t)) > 1 {
        return false
    }
    if len(s) == len(t) {
        diff := 0
        for i := 0; i < len(s); i++ {
            if s[i] != t[i] {
                diff++
                if diff > 1 {
                    return false
                }
            }
        }
        return diff == 1
    } else {
        if len(s) < len(t) {
            s, t = t, s
        }
        flag := false
        for i, j := 0, 0; j < len(t); {
            if s[i] != t[j] {
                if flag {
                    return false
                } else {
                    j--
                    flag = true
                }
            }
            j++
            i++
        }
        return true
    }
}
