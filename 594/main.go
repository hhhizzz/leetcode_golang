package _594

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func findLHS(nums []int) int {
    m := map[int]int{}
    result := 0
    for _, num := range nums {
        m[num] += 1
        if m[num+1] != 0 {
            result = max(result, m[num]+m[num+1])
        }
        if m[num-1] != 0 {
            result = max(result, m[num]+m[num-1])
        }
    }
    return result
}
