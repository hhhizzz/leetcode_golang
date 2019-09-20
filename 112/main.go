package _112

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

//挑战一下非递归前序和不改变原树
func hasPathSum(root *TreeNode, sum int) bool {
    var stack []*TreeNode
    if root == nil {
        return false
    }
    current := &TreeNode{root.Val, root.Left, root.Right}
    for true {
        for current != nil {
            if current.Left == nil && current.Right == nil && current.Val == sum {
                return true
            }
            if current.Right != nil {
                newRight := &TreeNode{current.Right.Val + current.Val, current.Right.Left, current.Right.Right}
                stack = append(stack, newRight)
            }
            if current.Left != nil {
                newLeft := &TreeNode{current.Left.Val + current.Val, current.Left.Left, current.Left.Right}
                current = newLeft
            } else {
                current = current.Left
            }
        }
        if len(stack) == 0 {
            break
        }
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if current.Left == nil && current.Right == nil && current.Val == sum {
            return true
        }
    }
    return false
}
