package _94

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    stack = append(stack, root)
    var result []int
    for len(stack) != 0 {
        current := stack[len(stack)-1]
        
    }
    return
}
