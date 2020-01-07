package _1019

type ListNode struct {
    Val  int
    Next *ListNode
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//单调栈解法 感觉在这道题上有点作弊，因为是变成了数组
func nextLargerNodes(head *ListNode) []int {
    var stack []int
    var result []int
    var array []int

    for current := head; current != nil; current = current.Next {
        array = append(array, current.Val)
    }
    result = make([]int, len(array))
    for i := 0; i < len(array); i++ {
        if len(stack) == 0 || array[i] <= array[stack[len(stack)-1]] {
            stack = append(stack, i)
        } else {
            for len(stack) > 0 && array[stack[len(stack)-1]] < array[i] {
                j := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                result[j] = array[i]
            }
            stack = append(stack, i)
        }
    }
    return result
}
