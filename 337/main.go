package _337

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func dfs(root *TreeNode, result *int) (maxMe, maxNoMe int) {
    maxMe = 0
    maxNoMe = 0
    if root == nil {
        return
    }
    left1, left2 := dfs(root.Left, result)
    right1, right2 := dfs(root.Right, result)
    maxMe = root.Val + left2 + right2
    maxNoMe = max(left1, left2) + max(right1, right2)

    *result = max(*result, maxMe)
    *result = max(*result, maxNoMe)

    return
}

func rob(root *TreeNode) int {
    result := 0
    dfs(root, &result)
    return result
}
