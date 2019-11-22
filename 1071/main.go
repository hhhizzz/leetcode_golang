package _1071

func devide(p, s []byte) bool {
    if len(p) == 0 {
        return true
    }
    if len(p) < len(s) {
        return false
    }
    for i := 0; i < len(s); i++ {
        if p[i] != s[i] {
            return false
        }
    }
    return devide(p[len(s):], s)
}

// 直接想到的办法
func gcdOfStrings1(str1 string, str2 string) string {
    if len(str2) > len(str1) {
        str1, str2 = str2, str1
    }
    s1 := []byte(str1)
    s2 := []byte(str2)

    d := []byte{}
    var result []byte
    for i := 0; i < len(s2); i++ {
        d = append(d, s2[i])
        if devide(s1, d) && devide(s2, d) {
            result = make([]byte, len(d))
            copy(result, d)
        }
    }
    return string(result)
}

//辗转相除法
func gcdOfStrings(str1 string, str2 string) string {
    if str1+str2 != str2+str1 {
        return ""
    }
    return string(helper([]byte(str1), []byte(str2)))
}

func helper(str1, str2 []byte) []byte {
    diff := len(str1) - len(str2)
    if diff == 0 {
        return str1
    } else if diff > 0 {
        str1 = str1[len(str2):]
        return helper(str1, str2)
    } else {
        str2 = str2[len(str1):]
        return helper(str1, str2)
    }
}
