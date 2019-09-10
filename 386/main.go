package _386

func lexicalOrder(n int) []int {
    if n < 1 {
        return []int{}
    }

    var stack []int
    var result []int
    for i := 9; i > 0; i-- {
        stack = append(stack, i)
    }
    for len(stack) != 0 {
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pop <= n {
            result = append(result, pop)
            pop *= 10
            if pop <= n {
                for i := 9; i >= 0; i-- {
                    stack = append(stack, pop+i)
                }
            }
        }
    }
    return result
}
