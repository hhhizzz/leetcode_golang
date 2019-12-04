package _129

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func dfs(root *TreeNode, sum int, result *int) {
    if root == nil {
        return
    }
    current := sum*10 + root.Val
    if root.Left == nil && root.Right == nil {
        *result += current
    } else {
        dfs(root.Left, current, result)
        dfs(root.Right, current, result)
    }
}

func sumNumbers(root *TreeNode) int {
    result := 0
    dfs(root, 0, &result)
    return result
}
