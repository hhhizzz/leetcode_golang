package _144

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

//题目说了要用非递归，要对得起中等难度
func preorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var result []int
    if root == nil {
        return result
    }
    stack = append(stack, root)

    for len(stack) != 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.Val)
        if current.Right != nil {
            stack = append(stack, current.Right)
        }
        if current.Left != nil {
            stack = append(stack, current.Left)
        }
    }

    return result
}
