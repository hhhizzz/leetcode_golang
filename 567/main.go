package _567

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    m := map[byte]int{}
    for i := 0; i < len(s1); i++ {
        m[s1[i]] += 1
    }
    for i, j := 0, 0; j < len(s2); {
        if n, ok := m[s2[j]]; ok {
            if n == 0 {
                for s2[i] != s2[j] {
                    m[s2[i]]++
                    i++
                }
                if i != j {
                    m[s2[i]]++
                    i++
                } else {
                    m[s2[i]]++
                    i++
                    j++
                }
            } else {
                m[s2[j]]--
                j++
                if j-i == len(s1) {
                    return true
                }
            }
        } else {
            for i < j {
                m[s2[i]]++
                i++
            }
            i++
            j = i
        }
    }
    return false
}
