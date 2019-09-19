package _94

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var result []int
    current := root
    for true {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        if len(stack) == 0 {
            break
        }
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.Val)
        current = current.Right
    }
    return result
}
