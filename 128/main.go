package _128

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func longestConsecutive(nums []int) int {
    m := map[int]bool{}
    result := 0
    for _, num := range nums {
        m[num] = true
    }
    for _, num := range nums {
        if _, ok := m[num-1]; !ok {
            current := 1
            for {
                if _, ok = m[num+1]; ok {
                    num += 1
                    current += 1
                } else {
                    break
                }
            }
            result = max(current, result)
        }
    }
    return result
}
