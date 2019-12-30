package _135

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func candy(ratings []int) int {
    result := 0
    left := make([]int, len(ratings))
    right := make([]int, len(ratings))

    left[0] = 1
    for i := 1; i < len(ratings); i++ {
        if ratings[i] > ratings[i-1] {
            left[i] = left[i-1] + 1
        } else {
            left[i] = 1
        }
    }
    right[len(right)-1] = 1
    for i := len(right) - 2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] {
            right[i] = right[i+1] + 1
        } else {
            right[i] = 1
        }
    }
    for i := 0; i < len(left); i++ {
        result += max(left[i], right[i])
    }
    return result
}
