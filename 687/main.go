package _487

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

func longestPathFromRoot(root *TreeNode, result *int) int {
    if root == nil {
        return 0
    }
    current := 1
    left := longestPathFromRoot(root.Left, result)
    right := longestPathFromRoot(root.Right, result)

    if root.Left != nil && root.Left.Val == root.Val {
        current += left
        left += 1
    } else {
        left = 1
    }
    if root.Right != nil && root.Right.Val == root.Val {
        current += right
        right += 1
    } else {
        right = 1
    }
    *result = max(*result, current)
    return max(left, right)
}

func longestUnivaluePath(root *TreeNode) int {
    result := 0
    longestPathFromRoot(root, &result)
    if result <= 0 {
        return result
    } else {
        return result - 1
    }
}
