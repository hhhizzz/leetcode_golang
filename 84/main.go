package _84

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func largestRectangleArea(heights []int) int {
    result := 0
    heights = append([]int{0}, heights...)
    heights = append(heights, 0)
    var stack []int
    for i := 0; i < len(heights); i++ {
        if len(stack) != 0 {
            topItem := stack[len(stack)-1]
            for heights[i] < heights[topItem] {
                topHeight := heights[topItem]
                stack = stack[:len(stack)-1]
                topItem = stack[len(stack)-1]
                result = max(result, (i-topItem-1)*topHeight)
            }
        }
        stack = append(stack, i)
    }
    return result
}
